package ipam

import (
	"fmt"
	"math/rand"
	"net"
	"time"

	"github.com/go-redis/redis"
	tracing "github.com/opentracing/opentracing-go"
	"github.com/pkg/errors"
	"golang.org/x/net/context"

	"github.com/hatena/ipdrawer/gen/go/model"
	"github.com/hatena/ipdrawer/pkg/base"
	pm "github.com/hatena/ipdrawer/pkg/model"
	"github.com/hatena/ipdrawer/pkg/storage"
	nu "github.com/hatena/ipdrawer/pkg/utils/netutil"
	"github.com/sirupsen/logrus"
)

type IPManager struct {
	redis  *storage.Redis
	locker storage.Locker
}

// NewIPManager creates IPManager instance
func NewIPManager(cfg *base.Config) *IPManager {
	redis, err := storage.NewRedis(cfg)
	if err != nil {
		panic(fmt.Errorf("failed to connect to db: %w", err))
	}
	locker := storage.NewLocker(redis)
	return &IPManager{
		redis:  redis,
		locker: locker,
	}
}

// The IP addresses are organized as a z-range in redis:
// https://redis.io/commands/zrange/ whereas it starts with lower IP address to
// higher based on the uint order of the IP addresses.
//
// IP address can be marked used in following ways:
// 1. a temporary reserved status of 24 hours in 'ipTempReserved', or
// 2. a record in 'ipDetails'.
// The 1st method is to hold an IP address so that we don't have race conditions.
// Caller is expected to add a record to 'ip details' by activating the IP within
// 24 hours.
//
// We further use the following two set of keys to help the allocation process:
// 1. 'poolUsedIPZSet' is used to store a sorted (z-set) list of allocated IPs.
// 2. 'ipPoolUuid' to list the association of UUIDs to IPs. That way a device
//    can try to get an available IP that it used before on first come first
//    serve basis. Such association is not removed from the DB unless it is
//    explicitly done so through API.
//
// TODO:
// - DON"T USE KEYS command as it is very bad for performance. Use scan/sets:
//   https://redis.io/commands/keys/

// checkIPDetails return true if an IP is available based on 'ipDetails' key
func (m *IPManager) checkIPDetails(namespace, uuid string, ip net.IP, log *logrus.Entry) bool {
	addr, err := getIPAddr(m.redis, namespace, ip)
	if err != nil {
		if errors.Is(err, errUnmarshalAddr) {
			return false
		}
		// TODO: Other errors may indicate hard failures or no entry exists.
		// TODO: They need to be differentiated.
		// Fall through for now.
		log.WithError(err).Debugln("failed to get ip detail")
	} else {
		if addr != nil && addr.Status != model.IPAddr_UNKNOWN {
			if addr.Uuid != uuid {
				log.WithField("ip-details-uuid", addr.Uuid).Debugln("already assigned")
				return false
			}
			log.Infoln("already assigned to this device. Repeated request?")
		} else {
			log.Debugln("unknown ip status or nil entry")
		}
	}
	return true
}

// checkTempReserved returns true if an IP is available based on 'ipTempReserved'
func (m *IPManager) checkTempReserved(namespace, uuid string, ip net.IP, log *logrus.Entry) bool {
	check, err := m.redis.Client.Get(makeIPTempReserved(namespace, ip)).Result()
	if err == nil && check != uuid {
		log.WithField("temp-reserved-uuid", check).Debugln("already reserved")
		return false
	}
	return true
}

func (m *IPManager) isIPAvailable(namespace, uuid string, ip net.IP, log *logrus.Entry) bool {
	if !m.checkIPDetails(namespace, uuid, ip, log) {
		return false
	}
	if !m.checkTempReserved(namespace, uuid, ip, log) {
		return false
	}
	return true
}

// check IPs that were allocated to this UUID before
func (m *IPManager) checkUuidIP(namespace, uuid string, s, e net.IP, log *logrus.Entry) (net.IP, error) {
	uuidKeys, err := m.redis.Client.Keys(makePoolUuidIPListPattern(namespace, s, e, uuid)).Result()
	if err != nil {
		return nil, errors.Wrap(err, "Failed to fetch uuid IP list keys")
	}

	log.WithField("uuid-keys", uuidKeys).Infoln("try to find a uuid match")
	for _, key := range uuidKeys {
		_, _, avail, _, err := parsePoolUuidIPKey(key)
		if err != nil {
			// Skip bad entries
			log.WithField("key", key).WithError(err).Warnln("Parse failed")
			continue
		}

		logger := log.WithField("ip", avail.String())
		if !m.isIPAvailable(namespace, uuid, avail, logger) {
			continue
		}

		// It is free. Reserve it.
		k := makeIPTempReserved(namespace, avail)
		_, err = m.redis.Client.Set(k, uuid, 24*time.Hour).Result()
		if err == nil {
			logger.WithField("key", k).Infoln("is now reserved")
			return avail, nil
		}
		logger.WithError(err).Errorln("Failed to reserve")
	}
	return nil, nil
}

func isIPInUsedZSet(ip net.IP, i int, set []string) bool {
	if i < len(set) {
		usedIP := net.ParseIP(set[i])
		if usedIP != nil {
			if ip.Equal(usedIP) {
				return true
			}
		}
	}
	return false
}

// hasIPBeenUsedByAnyUuid checks if an IP has ever been used by any UUID. When
// an IP is freed back to the pool its association with an UUID is not cleared.
// That way a device can get its IP back next time.
func (m *IPManager) hasIPBeenUsedByAnyUuid(namespace string, ip net.IP, log *logrus.Entry) bool {
	key := makeIPUuidKey(namespace, "*", "*", ip.String())
	uuidKeys, err := m.redis.Client.Keys(key).Result()
	if err == nil && len(uuidKeys) > 0 {
		log.WithField("key", key).WithField("len-uuids", len(uuidKeys)).Infoln("used by")
		return true
	}
	return false
}

func (m *IPManager) temporaryReserveIP(namespace, uuid string, s, e, avail net.IP, log *logrus.Entry) (net.IP, error) {
	if _, err := m.redis.Client.Set(makeIPTempReserved(namespace, avail), uuid, 24*time.Hour).Result(); err != nil {
		log.WithError(err).Errorln("Failed to mark temp reserve")
		return nil, err
	}
	if uuid != "" {
		if _, err := m.redis.Client.Set(makeIPPoolUuid(namespace, s, e, avail, uuid), 1, 0).Result(); err != nil {
			log.WithError(err).Errorln("Failed to mark the uuid")
			return nil, err
		}
	}
	return avail, nil
}

func (m *IPManager) checkUsedPoolRandom(namespace, uuid string, s, e net.IP, log *logrus.Entry) (net.IP, error) {
	zkey := makePoolUsedIPZSet(namespace, s, e)
	zset, err := m.redis.Client.ZRange(zkey, 0, -1).Result()
	if err != nil {
		log.WithError(err).Errorln("failed to get used set")
		return nil, err
	}

	// Convert used IPs to a map for O(1) lookup
	usedIPs := make(map[string]bool)
	for _, ip := range zset {
		usedIPs[ip] = true
	}

	// Calculate total range size
	start := nu.IP2Uint(s)
	end := nu.IP2Uint(e)
	totalSize := end - start + 1

	// Define segment size (e.g., 256 IPs per segment)
	const segmentSize uint32 = 256
	numSegments := (totalSize + segmentSize - 1) / segmentSize // Round up division

	// Initialize random source
	rnd := rand.New(rand.NewSource(time.Now().UnixNano()))

	// Create random order of segments to try
	segmentOrder := rnd.Perm(int(numSegments))

	var usedByOtherUUID net.IP

	// Try each segment in random order
	for _, segmentIdx := range segmentOrder {
		// Calculate segment boundaries
		segmentStart := start + uint32(segmentIdx)*segmentSize
		segmentEnd := segmentStart + segmentSize - 1
		if segmentEnd > end {
			segmentEnd = end
		}

		// Create random order of IPs within this segment
		segmentSize := segmentEnd - segmentStart + 1
		ipOrder := rnd.Perm(int(segmentSize))

		// Try each IP in the segment in random order
		for _, offset := range ipOrder {
			curr := segmentStart + uint32(offset)
			avail := nu.Int2IP(curr)

			// Skip if IP is in used set
			if usedIPs[avail.String()] {
				continue
			}

			logger := log.WithField("ip", avail.String())
			if !m.checkTempReserved(namespace, uuid, avail, logger) {
				continue
			}

			if m.hasIPBeenUsedByAnyUuid(namespace, avail, logger) {
				usedByOtherUUID = avail
				continue
			}

			// Found a free IP
			return m.temporaryReserveIP(namespace, uuid, s, e, avail, logger)
		}
	}

	// Fall back to an IP that was used by other UUID if available
	if usedByOtherUUID != nil {
		logger := log.WithField("ip", usedByOtherUUID.String())
		return m.temporaryReserveIP(namespace, uuid, s, e, usedByOtherUUID, logger)
	}

	return nil, errNoIPAvailable
}

// CheckUsedPoolSequential walks for all the addresses from s to e for an address that is
// not in used pool yet.
func (m *IPManager) CheckUsedPoolSequential(namespace, uuid string, s, e net.IP, log *logrus.Entry) (net.IP, error) {
	zkey := makePoolUsedIPZSet(namespace, s, e)
	zset, err := m.redis.Client.ZRange(zkey, 0, -1).Result()
	if err != nil {
		log.WithError(err).Errorln("failed to get used set")
		return nil, err
	}

	var usedByOtherUUID net.IP
	for i, avail := 0, s; !nu.PrevIP(avail).Equal(e); avail = nu.NextIP(avail) {
		if isIPInUsedZSet(avail, i, zset) {
			i += 1
			continue
		}
		logger := log.WithField("ip", avail.String())
		if !m.checkTempReserved(namespace, uuid, avail, logger) {
			continue
		}
		if m.hasIPBeenUsedByAnyUuid(namespace, avail, logger) {
			usedByOtherUUID = avail
			continue
		}

		// A true free IP
		return m.temporaryReserveIP(namespace, uuid, s, e, avail, logger)
	}
	// OK we failed to get a free IP if trying to avoid any IP ever used by
	// other UUIDs. Add this UUID to the list of lucky guys.
	if usedByOtherUUID != nil {
		avail := usedByOtherUUID
		logger := log.WithField("ip", avail.String())
		return m.temporaryReserveIP(namespace, uuid, s, e, avail, logger)
	}
	return nil, errNoIPAvailable
}

// DrawIP returns an available IP.
func (m *IPManager) DrawIP(ctx context.Context, namespace string, pool *model.Pool, uuid, wantIP string, random, reserve, ping, mustHaveWantIP bool) (net.IP, error) {
	span, ctx := tracing.StartSpanFromContext(ctx, "IPManager.DrawIP")
	span.SetTag("pool", pm.PoolKey(pool))
	defer span.Finish()

	log := logrus.WithFields(logrus.Fields{
		"namespace":         namespace,
		"handler":           "IPAM-draw-ip",
		"uuid":              uuid,
		"want-ip":           wantIP,
		"mast-have-want-ip": mustHaveWantIP,
	})
	log.Infoln("draw-ip")
	token, err := m.locker.Lock(ctx, makeGlobalLock())
	if err != nil {
		log.WithError(err).Errorln("failed to get lock")
		return nil, err
	}
	defer m.locker.Unlock(ctx, makeGlobalLock(), token)

	s := net.ParseIP(pool.Start)
	e := net.ParseIP(pool.End)

	// Check if the wanted IP is available or not
	if wantIP != "" {
		avail := net.ParseIP(wantIP)
		if avail == nil {
			return nil, fmt.Errorf("invalid ip '%v' wanted", wantIP)
		}
		if m.isIPAvailable(namespace, uuid, avail, log) {
			return m.temporaryReserveIP(namespace, uuid, s, e, avail, log)
		}
		if mustHaveWantIP {
			return nil, fmt.Errorf("%w: ip '%v' wanted but is not available", errAddrNotAvailable, wantIP)
		}
	}

	// Check if the last used address is free or not base on the uuid
	if uuid != "" {
		avail, err := m.checkUuidIP(namespace, uuid, s, e, log)
		if err != nil {
			return nil, err
		}
		if avail != nil {
			return avail, nil
		}
		// Fall through
	}
	// Look it up from the rest of the free entries
	if random {
		return m.checkUsedPoolRandom(namespace, uuid, s, e, log)
	}
	return m.CheckUsedPoolSequential(namespace, uuid, s, e, log)
}

// CreateIP activates IP.
func (m *IPManager) CreateIP(ctx context.Context, ps []*model.Pool, addr *model.IPAddr) error {
	span, ctx := tracing.StartSpanFromContext(ctx, "IPManager.Activate")
	span.SetTag("ip", addr.Ip)
	defer span.Finish()
	log := logrus.WithFields(
		logrus.Fields{
			"ip":        addr.Ip,
			"uuid":      addr.Uuid,
			"namespace": addr.Namespace,
			"handler":   "CreateIP",
		},
	)
	namespace := addr.Namespace
	token, err := m.locker.Lock(ctx, makeGlobalLock())
	if err != nil {
		return err
	}
	defer m.locker.Unlock(ctx, makeGlobalLock(), token)

	ip := net.ParseIP(addr.Ip)

	if err := setIPAddrTimestamp(m.redis, namespace, addr); err != nil {
		return err
	}

	if err := setIPAddr(m.redis, namespace, addr); err != nil {
		return err
	}

	pipe := m.redis.Client.TxPipeline()
	// Remove temporary reserved key in any way
	pipe.Del(makeIPTempReserved(namespace, ip))
	// Add IP to used IP zset
	score := float64(nu.IP2Uint(ip))
	z := redis.Z{
		Score:  score,
		Member: ip.String(),
	}
	for _, p := range ps {
		if pm.PoolContains(p, ip) {
			s := net.ParseIP(p.Start)
			e := net.ParseIP(p.End)
			log.WithField("pool-s", s).WithField("pool-e", e).Infoln("created ip")
			pipe.ZAdd(makePoolUsedIPZSet(namespace, s, e), z)
		}
	}
	if _, err := pipe.Exec(); err != nil {
		return err
	}
	return nil
}

func (m *IPManager) removeKeys(pattern string) {
	keys, err := m.redis.Client.Keys(pattern).Result()
	if err == nil {
		for _, key := range keys {
			m.redis.Client.Del(key)
		}
	}
}

func (m *IPManager) RemoveIPUuids(namespace, ip string) {
	if ip != "" {
		m.removeKeys(makeIPUuidKey(namespace, "*", "*", ip))
	}
}

func (m *IPManager) RemoveUuidIPs(namespace, uuid string) {
	if uuid != "" {
		m.removeKeys(makeIPUuidKey(namespace, "*", "*", uuid))
	}
}

func (m *IPManager) Deactivate(ctx context.Context, namespace string, ps []*model.Pool, addr *model.IPAddr) error {
	span, ctx := tracing.StartSpanFromContext(ctx, "IPManager.Deactivate")
	span.SetTag("pool", pm.PoolKey(ps[0]))
	span.SetTag("ip", addr.Ip)
	defer span.Finish()
	log := logrus.WithField("ip", addr.Ip).WithField("handler", "Deactivate")
	token, err := m.locker.Lock(ctx, makeGlobalLock())
	if err != nil {
		return err
	}
	defer m.locker.Unlock(ctx, makeGlobalLock(), token)

	ip := net.ParseIP(addr.Ip)
	// m.RemoveUuidIPs(namespace, addr.Uuid)
	pipe := m.redis.Client.TxPipeline()
	pipe.Del(makeIPTempReserved(namespace, ip))
	pipe.Del(makeIPDetailsKey(namespace, ip))
	for _, p := range ps {
		if pm.PoolContains(p, ip) {
			s := net.ParseIP(p.Start)
			e := net.ParseIP(p.End)
			log.WithField("pool-s", s).WithField("pool-e", e).Infoln("removed ip")
			pipe.ZRem(makePoolUsedIPZSet(namespace, s, e), ip.String())
		}
	}
	if _, err := pipe.Exec(); err != nil {
		return err
	}
	return nil
}

func (m *IPManager) UpdateIP(ctx context.Context, addr *model.IPAddr) error {
	span, ctx := tracing.StartSpanFromContext(ctx, "IPManager.UpdateIP")
	span.SetTag("ip", addr.Ip)
	defer span.Finish()

	namespace := addr.Namespace
	token, err := m.locker.Lock(ctx, makeGlobalLock())
	if err != nil {
		return err
	}
	defer m.locker.Unlock(ctx, makeGlobalLock(), token)

	if !existsIP(m.redis, namespace, addr) {
		return errors.New("Not found IP")
	}

	if err := setIPAddrTimestamp(m.redis, namespace, addr); err != nil {
		return err
	}

	return setIPAddr(m.redis, namespace, addr)
}

// Reserve makes the status of given IP reserved.
func (m *IPManager) Reserve(p *model.Pool, ip net.IP) error {
	return nil
}

// Release makes the status of given IP available.
func (m *IPManager) Release(p *model.Pool, ip net.IP) error {
	return nil
}

// GetNetworkIncludingIP returns a network including given IP.
func (m *IPManager) GetNetworkIncludingIP(ctx context.Context, namespace string, ip net.IP) (*model.Network, error) {
	span, _ := tracing.StartSpanFromContext(ctx, "IPManager.GetNetworkIncludingIP")
	span.SetTag("ip", ip.String())
	defer span.Finish()

	ps, err := m.redis.Client.SMembers(makeNetworkListKey(namespace)).Result()
	if err != nil {
		return nil, err
	}
	for _, p := range ps {
		_, ipnet, err := net.ParseCIDR(p)
		if err != nil {
			continue
		}
		if ipnet.Contains(ip) {
			net, err := getNetwork(m.redis, namespace, ipnet)
			return net, err
		}
	}
	return nil, errors.New(fmt.Sprintf("Network including %s does not exist", ip.String()))
}

// GetPoolsInNetwork gets pools.
func (m *IPManager) GetPoolsInNetwork(ctx context.Context, namespace string, n *model.Network) ([]*model.Pool, error) {
	span, _ := tracing.StartSpanFromContext(ctx, "IPManager.GetPoolsInNetwork")
	span.SetTag("network", n.String())
	defer span.Finish()

	return getPoolsInNetwork(m.redis, namespace, n)
}

// GetNetworkByIP returns network by IP.
func (m *IPManager) GetNetworkByIP(ctx context.Context, namespace string, ipnet *net.IPNet) (*model.Network, error) {
	span, _ := tracing.StartSpanFromContext(ctx, "IPManager.GetNetworkByIP")
	span.SetTag("ip", ipnet.String())
	defer span.Finish()

	return getNetwork(m.redis, namespace, ipnet)
}

// GetNetworkByName returns network by name.
func (m *IPManager) GetNetworkByName(ctx context.Context, namespace string, name string) (*model.Network, error) {
	span, _ := tracing.StartSpanFromContext(ctx, "IPManager.GetNetworkByName")
	span.SetTag("name", name)
	defer span.Finish()

	networks, err := getNetworks(m.redis, namespace)
	if err != nil {
		return nil, err
	}

	target := &model.Tag{
		Key:   "Name",
		Value: name,
	}
	for _, n := range networks {
		if pm.NetworkHasTag(n, target) {
			return n, nil
		}
	}

	return nil, ErrNetworkNotFound
}

// CreateNetwork creates network.
func (m *IPManager) CreateNetwork(ctx context.Context, namespace string, n *model.Network) error {
	span, ctx := tracing.StartSpanFromContext(ctx, "IPManager.CreateNetwork")
	span.SetTag("network", n.String())
	defer span.Finish()

	token, err := m.locker.Lock(ctx, makeGlobalLock())
	if err != nil {
		return err
	}
	defer m.locker.Unlock(ctx, makeGlobalLock(), token)

	if err := setTSToNetwork(m.redis, namespace, n); err != nil {
		return err
	}

	return setNetwork(m.redis, namespace, n)
}

// DeleteNetwork deletes the network
func (m *IPManager) DeleteNetwork(ctx context.Context, namespace string, n *model.Network) error {
	span, ctx := tracing.StartSpanFromContext(ctx, "IPManager.DeleteNetwork")
	span.SetTag("network", n.String())
	defer span.Finish()

	token, err := m.locker.Lock(ctx, makeGlobalLock())
	if err != nil {
		return err
	}
	defer m.locker.Unlock(ctx, makeGlobalLock(), token)

	return deleteNetwork(m.redis, namespace, n)
}

// UpdateNetwork updates the network
func (m *IPManager) UpdateNetwork(ctx context.Context, n *model.Network) error {
	span, ctx := tracing.StartSpanFromContext(ctx, "IPManager.UpdateNetwork")
	span.SetTag("network", n.String())
	defer span.Finish()

	namespace := n.Namespace
	token, err := m.locker.Lock(ctx, makeGlobalLock())
	if err != nil {
		return err
	}
	defer m.locker.Unlock(ctx, makeGlobalLock(), token)

	if !existsNetwork(m.redis, n) {
		return ErrNetworkNotFound
	}

	if err := setTSToNetwork(m.redis, namespace, n); err != nil {
		return err
	}

	return setNetwork(m.redis, namespace, n)
}

// CreatePool creates pool
func (m *IPManager) CreatePool(ctx context.Context, namespace string, n *model.Network, pool *model.Pool) error {
	span, ctx := tracing.StartSpanFromContext(ctx, "IPManager.CreatePool")
	span.SetTag("network", n.String())
	span.SetTag("pool", pm.PoolKey(pool))
	defer span.Finish()

	if err := addPoolToNetwork(m.redis, namespace, n, pool); err != nil {
		return err
	}

	// Get IPs in this pool to add these to a used zset.
	addrs, err := m.ListIP(ctx, namespace)
	if err != nil {
		return err
	}

	s := net.ParseIP(pool.Start)
	e := net.ParseIP(pool.End)
	usedkey := makePoolUsedIPZSet(namespace, s, e)
	members := make([]redis.Z, 0, len(addrs))
	for _, _ip := range addrs {
		ip := net.ParseIP(_ip.Ip)
		if !pm.PoolContains(pool, ip) {
			continue
		}

		score := float64(nu.IP2Uint(ip))
		z := redis.Z{
			Score:  score,
			Member: ip.String(),
		}
		members = append(members, z)
	}
	if len(members) != 0 {
		_, err = m.redis.Client.ZAdd(usedkey, members...).Result()
		if err != nil {
			return err
		}
	}

	if err := setTSToPool(m.redis, namespace, pool); err != nil {
		return err
	}

	return setPool(m.redis, namespace, pool)
}

func (m *IPManager) ListIP(ctx context.Context, namespace string) ([]*model.IPAddr, error) {
	span, _ := tracing.StartSpanFromContext(ctx, "IPManager.ListIP")
	defer span.Finish()

	keys, err := m.redis.Client.Keys(makeIPListPattern(namespace)).Result()
	if err != nil {
		return nil, errors.Wrap(err, "failed to list keys")
	}

	ips := make([]net.IP, len(keys))
	for i, key := range keys {
		ip, err := parseIPDetailsKey(key)
		if err != nil {
			return nil, errors.Wrapf(err, "failed to parse key '%s'", key)
		}
		ips[i] = ip
	}

	return getIPAddrs(m.redis, namespace, ips)
}

// GetTemporaryReservedIPs returns all temporary reserved ips.
func (m *IPManager) GetTemporaryReservedIPs(ctx context.Context, namespace string) ([]*model.IPAddr, error) {
	span, _ := tracing.StartSpanFromContext(ctx, "IPManager.GetTemporaryReservedIPs")
	defer span.Finish()

	keys, err := m.redis.Client.Keys(makeTempReservedIPListPattern(namespace)).Result()
	if err != nil {
		return nil, errors.Wrap(err, "Failed to fetch IP list keys")
	}

	addrs := make([]*model.IPAddr, len(keys))
	for i, key := range keys {
		ip, err := parseTempReservedIPKey(key)
		if err != nil {
			return nil, errors.Wrapf(err, "Parse failed: %s", key)
		}
		addrs[i] = &model.IPAddr{
			Ip:     ip.String(),
			Status: model.IPAddr_TEMPORARY_RESERVED,
		}
	}
	return addrs, nil
}

// GetNetworks returns all network.
func (m *IPManager) GetNetworks(ctx context.Context, namespace string) ([]*model.Network, error) {
	span, _ := tracing.StartSpanFromContext(ctx, "IPManager.GetNetworks")
	defer span.Finish()

	return getNetworks(m.redis, namespace)
}

// GetPools returns all pools.
func (m *IPManager) GetPools(ctx context.Context, namespace string) ([]*model.Pool, error) {
	span, _ := tracing.StartSpanFromContext(ctx, "IPManager.GetPools")
	defer span.Finish()

	keys, err := m.redis.Client.Keys(makePoolListPattern(namespace)).Result()
	if err != nil {
		return nil, errors.Wrap(err, "Failed fetch Pool list keys")
	}

	return getPools(m.redis, keys)
}

func (m *IPManager) GetPool(ctx context.Context, namespace string, s net.IP, e net.IP) (*model.Pool, error) {
	span, _ := tracing.StartSpanFromContext(ctx, "IPManager.GetPool")
	defer span.Finish()

	return getPool(m.redis, namespace, s, e)
}

func (m *IPManager) UpdatePool(ctx context.Context, pool *model.Pool) error {
	span, ctx := tracing.StartSpanFromContext(ctx, "IPManager.UpdatePool")
	span.SetTag("pool", pm.PoolKey(pool))
	defer span.Finish()

	namespace := pool.Namespace
	token, err := m.locker.Lock(ctx, makeGlobalLock())
	if err != nil {
		return err
	}
	defer m.locker.Unlock(ctx, makeGlobalLock(), token)

	if !existsPool(m.redis, namespace, pool) {
		return errors.New("Not found Pool")
	}

	if err := setTSToPool(m.redis, namespace, pool); err != nil {
		return err
	}

	return setPool(m.redis, namespace, pool)
}

func (m *IPManager) DeletePool(ctx context.Context, namespace string, start, end net.IP) error {
	span, ctx := tracing.StartSpanFromContext(ctx, "IPManager.DeletePool")
	defer span.Finish()

	token, err := m.locker.Lock(ctx, makeGlobalLock())
	if err != nil {
		return err
	}
	defer m.locker.Unlock(ctx, makeGlobalLock(), token)

	network, err := getNetworkIncludingPool(m.redis, namespace, start, end)
	if err != nil {
		return err
	}
	_, ipnet, err := net.ParseCIDR(network.Prefix)
	if err != nil {
		return err
	}

	pipe := m.redis.Client.TxPipeline()
	pipe.Del(makePoolDetailsKey(namespace, start, end))
	pipe.Del(makePoolUsedIPZSet(namespace, start, end))
	pipe.SRem(makeNetworkPoolKey(namespace, ipnet), pm.PoolKey(&model.Pool{
		Start: start.String(),
		End:   end.String(),
	}))
	_, err = pipe.Exec()

	return err
}
