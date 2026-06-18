package ipam

import (
	"fmt"
	"net"
	"time"

	"github.com/bufbuild/protovalidate-go"
	"github.com/hatena/ipdrawer/gen/go/model"
	pm "github.com/hatena/ipdrawer/pkg/model"
	"github.com/hatena/ipdrawer/pkg/storage"
	"github.com/pkg/errors"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/timestamppb"
)

// getNetworks retrieves all networks in the given namespace from Redis.
// It returns a slice of model.Network pointers or an error if any occurs.
// No result is not an error. Empty slice will be returned if there is no result.
func getNetworks(r *storage.Redis, namespace string) ([]*model.Network, error) {
	lkey := makeNetworkListKey(namespace)
	ps, err := r.Client.SMembers(lkey).Result()
	if err != nil {
		return nil, err
	}

	keys := make([]string, len(ps))
	for i, p := range ps {
		_, ipnet, err := net.ParseCIDR(p)
		if err != nil {
			return nil, err
		}
		keys[i] = makeNetworkDetailsKey(namespace, ipnet)
	}

	if len(keys) == 0 {
		return []*model.Network{}, nil
	}

	data, err := r.Client.MGet(keys...).Result()
	if err != nil {
		return nil, err
	}

	networks := make([]*model.Network, len(keys))

	for i, d := range data {
		if s, ok := d.(string); ok {
			networks[i] = &model.Network{}
			if err := proto.Unmarshal([]byte(s), networks[i]); err != nil {
				return nil, err
			}
		}
	}
	return networks, nil
}

func setTSToNetwork(r *storage.Redis, namespace string, n *model.Network) error {
	if existsNetwork(r, n) {
		_, pre, _ := net.ParseCIDR(n.Prefix)
		stored, _ := getNetwork(r, namespace, pre)
		n.CreatedAt = stored.CreatedAt
	}

	now := timestamppb.New(time.Now())
	if n.CreatedAt == nil {
		n.CreatedAt = now
	} else {
		n.LastModifiedAt = now
	}

	return nil
}

func setNetwork(r *storage.Redis, namespace string, n *model.Network) error {
	if err := protovalidate.Validate(n); err != nil {
		return err
	}

	pipe := r.Client.TxPipeline()

	data, err := proto.Marshal(n)
	if err != nil {
		return err
	}
	_, pre, err := net.ParseCIDR(n.Prefix)
	if err != nil {
		return err
	}
	// Set details
	dkey := makeNetworkDetailsKey(namespace, pre)
	pipe.Set(dkey, string(data), 0)
	pipe.SAdd(makeNetworkListKey(namespace), pre.String())

	_, err = pipe.Exec()

	return err
}

func deleteNetwork(r *storage.Redis, namespace string, n *model.Network) error {
	if err := protovalidate.Validate(n); err != nil {
		return err
	}
	_, pre, _ := net.ParseCIDR(n.Prefix)
	pipe := r.Client.TxPipeline()
	pipe.Del(makeNetworkDetailsKey(n.Namespace, pre))
	pipe.SRem(makeNetworkListKey(namespace), pre.String())
	_, err := pipe.Exec()
	return err
}

func getNetwork(r *storage.Redis, namespace string, ipnet *net.IPNet) (*model.Network, error) {
	k := makeNetworkDetailsKey(namespace, ipnet)

	check, err := r.Client.Exists(k).Result()
	if err != nil {
		return nil, errors.Wrap(ErrNetworkNotFound, err.Error())
	}
	if check == 0 {
		return nil, ErrNetworkNotFound
	}

	data, err := r.Client.Get(k).Result()
	if err != nil {
		return nil, err
	}
	n := &model.Network{}
	if err := proto.Unmarshal([]byte(data), n); err != nil {
		return nil, err
	}

	return n, nil
}

func addPoolToNetwork(r *storage.Redis, namespace string, network *model.Network, pool *model.Pool) error {
	if err := protovalidate.Validate(pool); err != nil {
		return err
	}
	_, pre, err := net.ParseCIDR(network.Prefix)
	if err != nil {
		return err
	}
	poolKey := makeNetworkPoolKey(namespace, pre)
	_, err = r.Client.SAdd(poolKey, pm.PoolKey(pool)).Result()
	return err
}

// ParseMask32 parses a single-host address. It accepts a bare IP (v4 or v6) or
// a full-length prefix: /32 for IPv4 and /128 for IPv6.
func ParseMask32(s string) (net.IP, error) {
	ip, ipnet, err := net.ParseCIDR(s)
	if err != nil {
		ip := net.ParseIP(s)
		if ip == nil {
			return nil, errors.New("Failed parse IP")
		}
		return ip, nil
	}
	ones, bits := ipnet.Mask.Size()
	if ones == bits {
		return ip, nil
	}
	return nil, errors.New("Only accepts a full-length mask (/32 for IPv4, /128 for IPv6)")
}

func existsNetwork(r *storage.Redis, network *model.Network) bool {
	if err := protovalidate.Validate(network); err != nil {
		return false
	}
	_, pre, _ := net.ParseCIDR(network.Prefix)
	check, _ := r.Client.Exists(makeNetworkDetailsKey(network.Namespace, pre)).Result()
	return check != 0
}

func getNetworkIncludingPool(r *storage.Redis, namespace string, start net.IP, end net.IP) (*model.Network, error) {
	networks, err := getNetworks(r, namespace)
	if err != nil {
		return nil, err
	}

	for _, n := range networks {
		_, ipnet, _ := net.ParseCIDR(n.Prefix)
		if ipnet.Contains(start) && ipnet.Contains(end) {
			return n, nil
		}
	}

	return nil, fmt.Errorf("%w: no such a pool(start=%s, end=%s)", ErrNetworkNotFound, start, end)
}
