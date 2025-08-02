package ipam

import (
	"fmt"
	"net"
	"strconv"
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"golang.org/x/net/context"

	"github.com/hatena/ipdrawer/gen/go/model"
	"github.com/hatena/ipdrawer/pkg/storage"
)

func (m *IPManager) reserveTemporary(ip net.IP, uuid string) {
	_, _ = m.redis.Client.Set(makeIPTempReserved(testNS, ip), uuid, 24*time.Hour).Result()
}

func TestIPActivation(t *testing.T) {
	r, deferFunc := storage.NewTestRedis()
	defer deferFunc()

	m := NewTestIPManager(r)

	ctx := context.Background()

	pool := &model.Pool{
		Start: "10.0.0.1",
		End:   "10.0.0.254",
	}

	if err := m.CreateIP(ctx, []*model.Pool{pool}, &model.IPAddr{Namespace: testNS, Ip: "10.0.0.1"}); err != nil {
		t.Fatalf("Got error: %v", err)
	}
	if err := m.CreateIP(ctx, []*model.Pool{pool}, &model.IPAddr{Namespace: testNS, Ip: "10.0.0.4"}); err != nil {
		t.Fatalf("Got error: %v", err)
	}

	s := net.ParseIP(pool.Start)
	e := net.ParseIP(pool.End)
	zkey := makePoolUsedIPZSet(testNS, s, e)
	cnt, err := r.Client.ZCard(zkey).Result()
	if err != nil {
		t.Errorf("Got error: %v", err)
	}
	if cnt != 2 {
		t.Errorf("Expected %d, but got %d", 2, cnt)
	}
}

func TestDrawIPSeq(t *testing.T) {
	ctx := context.Background()

	testCases := []struct {
		name           string
		pool           *model.Pool
		ips            []*model.IPAddr
		expected       net.IP
		wanted         string
		mustHaveWanted bool
		err            error
	}{
		{
			name: "IP available in the middle of the range",
			pool: &model.Pool{
				Start: "10.0.0.1",
				End:   "10.0.0.254",
			},
			ips: []*model.IPAddr{
				{
					Namespace: testNS,
					Ip:        "10.0.0.1",
					Status:    model.IPAddr_ACTIVE,
				}, {
					Namespace: testNS,
					Ip:        "10.0.0.3",
					Status:    model.IPAddr_ACTIVE,
				},
			},
			expected: net.ParseIP("10.0.0.2"),
		},
		{
			name: "1st X IPs either temp reserved or active",
			pool: &model.Pool{
				Start: "10.0.0.1",
				End:   "10.0.0.254",
			},
			ips: []*model.IPAddr{
				{
					Namespace: testNS,
					Ip:        "10.0.0.1",
					Status:    model.IPAddr_ACTIVE,
				}, {
					Namespace: testNS,
					Ip:        "10.0.0.2",
					Status:    model.IPAddr_TEMPORARY_RESERVED,
				}, {
					Namespace: testNS,
					Ip:        "10.0.0.3",
					Status:    model.IPAddr_ACTIVE,
				}, {
					Namespace: testNS,
					Ip:        "10.0.0.4",
					Status:    model.IPAddr_ACTIVE,
				},
			},
			expected: net.ParseIP("10.0.0.5"),
		},
		{
			name: "1st X IPs all temp reserved",
			pool: &model.Pool{
				Start: "10.0.0.1",
				End:   "10.0.0.254",
			},
			ips: []*model.IPAddr{
				{
					Namespace: testNS,
					Ip:        "10.0.0.1",
					Status:    model.IPAddr_TEMPORARY_RESERVED,
				}, {
					Namespace: testNS,
					Ip:        "10.0.0.2",
					Status:    model.IPAddr_TEMPORARY_RESERVED,
				}, {
					Namespace: testNS,
					Ip:        "10.0.0.3",
					Status:    model.IPAddr_TEMPORARY_RESERVED,
				}, {
					Namespace: testNS,
					Ip:        "10.0.0.4",
					Status:    model.IPAddr_TEMPORARY_RESERVED,
				},
			},
			expected: net.ParseIP("10.0.0.5"),
		},
		{
			name: "1st X IPs temp reserved and last Y IPs active",
			pool: &model.Pool{
				Start: "10.0.0.1",
				End:   "10.0.0.254",
			},
			ips: []*model.IPAddr{
				{
					Namespace: testNS,
					Ip:        "10.0.0.1",
					Status:    model.IPAddr_TEMPORARY_RESERVED,
				}, {
					Namespace: testNS,
					Ip:        "10.0.0.2",
					Status:    model.IPAddr_TEMPORARY_RESERVED,
				}, {
					Namespace: testNS,
					Ip:        "10.0.0.3",
					Status:    model.IPAddr_TEMPORARY_RESERVED,
				}, {
					Namespace: testNS,
					Ip:        "10.0.0.4",
					Status:    model.IPAddr_ACTIVE,
				},
			},
			expected: net.ParseIP("10.0.0.5"),
		},
		{
			name: "No IP available",
			pool: &model.Pool{
				Start: "10.0.0.1",
				End:   "10.0.0.2",
			},
			ips: []*model.IPAddr{
				{
					Namespace: testNS,
					Ip:        "10.0.0.1",
					Status:    model.IPAddr_TEMPORARY_RESERVED,
				}, {
					Namespace: testNS,
					Ip:        "10.0.0.2",
					Status:    model.IPAddr_TEMPORARY_RESERVED,
				},
			},
			err: errNoIPAvailable,
		},
		{
			name: "Want IP available",
			pool: &model.Pool{
				Start: "10.0.0.1",
				End:   "10.0.0.5",
			},
			ips: []*model.IPAddr{
				{
					Namespace: testNS,
					Ip:        "10.0.0.1",
					Status:    model.IPAddr_TEMPORARY_RESERVED,
				}, {
					Namespace: testNS,
					Ip:        "10.0.0.2",
					Status:    model.IPAddr_TEMPORARY_RESERVED,
				},
			},
			wanted:   "10.0.0.3",
			expected: net.ParseIP("10.0.0.3"),
		},
		{
			name: "Want IP not available",
			pool: &model.Pool{
				Start: "10.0.0.1",
				End:   "10.0.0.5",
			},
			ips: []*model.IPAddr{
				{
					Namespace: testNS,
					Ip:        "10.0.0.1",
					Status:    model.IPAddr_TEMPORARY_RESERVED,
				}, {
					Namespace: testNS,
					Ip:        "10.0.0.2",
					Status:    model.IPAddr_TEMPORARY_RESERVED,
				},
			},
			wanted:   "10.0.0.1",
			expected: net.ParseIP("10.0.0.3"),
		},
		{
			name: "Must have want IP not available",
			pool: &model.Pool{
				Start: "10.0.0.1",
				End:   "10.0.0.5",
			},
			ips: []*model.IPAddr{
				{
					Namespace: testNS,
					Ip:        "10.0.0.1",
					Status:    model.IPAddr_TEMPORARY_RESERVED,
				}, {
					Namespace: testNS,
					Ip:        "10.0.0.2",
					Status:    model.IPAddr_TEMPORARY_RESERVED,
				},
			},
			wanted:         "10.0.0.1",
			mustHaveWanted: true,
			err:            errAddrNotAvailable,
		},
	}

	for _, c := range testCases {
		r, deferFunc := storage.NewTestRedis()
		m := NewTestIPManager(r)

		for _, ip := range c.ips {
			switch ip.Status {
			case model.IPAddr_ACTIVE:
				m.CreateIP(ctx, []*model.Pool{c.pool}, ip)
			case model.IPAddr_TEMPORARY_RESERVED:
				m.reserveTemporary(net.ParseIP(ip.Ip), ip.Uuid)
			case model.IPAddr_RESERVED:
				m.Reserve(c.pool, net.ParseIP(ip.Ip))
			}
		}

		actual, err := m.DrawIP(ctx, testNS, c.pool, testUUID, c.wanted, false /* not random */, true, false, c.mustHaveWanted)
		if c.err == nil {
			assert.NoError(t, err)
			assert.Equal(t, c.expected.String(), actual.String())
		} else {
			assert.ErrorIs(t, err, c.err)
		}

		deferFunc()
	}
}

func TestDeactivateAfterActivating(t *testing.T) {
	r, deferFunc := storage.NewTestRedis()
	defer deferFunc()

	m := NewTestIPManager(r)

	ctx := context.Background()

	pool := &model.Pool{
		Start: "10.0.0.1",
		End:   "10.0.0.254",
	}

	ip := &model.IPAddr{
		Ip: "10.0.0.1",
	}

	m.CreateIP(ctx, []*model.Pool{pool}, ip)

	if err := m.Deactivate(ctx, testNS, []*model.Pool{pool}, ip); err != nil {
		t.Errorf("Failed deactivating: %#+v", err)
	}

	keys, _ := r.Client.Keys(makeIPTempReserved(testNS, net.ParseIP(ip.Ip))).Result()
	if len(keys) != 0 {
		t.Errorf("Deactivation should remove temporary reserved key")
	}
}

func TestActivateIPInSeveralPools(t *testing.T) {
	r, deferFunc := storage.NewTestRedis()
	defer deferFunc()

	m := NewTestIPManager(r)

	ctx := context.Background()

	pools := []*model.Pool{
		{
			Start: "10.0.0.1",
			End:   "10.0.0.254",
		},
		{
			Start: "10.0.0.30",
			End:   "10.0.0.50",
		},
	}

	ip := &model.IPAddr{
		Ip: "10.0.0.40",
	}

	err := m.CreateIP(ctx, pools, ip)
	if err != nil {
		t.Errorf("Activate(%v, %v) returns %#+v; want success", pools, ip, err)
	}
}

func TestDeactivateIPInSeveralPools(t *testing.T) {
	r, deferFunc := storage.NewTestRedis()
	defer deferFunc()

	m := NewTestIPManager(r)

	ctx := context.Background()

	pools := []*model.Pool{
		{
			Start: "10.0.0.1",
			End:   "10.0.0.254",
		},
		{
			Start: "10.0.0.30",
			End:   "10.0.0.50",
		},
	}

	ip := &model.IPAddr{
		Ip: "10.0.0.40",
	}

	m.CreateIP(ctx, pools, ip)

	if err := m.Deactivate(ctx, testNS, pools, ip); err != nil {
		t.Errorf("Deactivate(%v, %v) returns %#+v; want success", pools, ip, err)
	}
}

func TestCorrectDrawIPFromInclusivePools(t *testing.T) {
	r, deferFunc := storage.NewTestRedis()
	defer deferFunc()

	m := NewTestIPManager(r)

	ctx := context.Background()

	pools := []*model.Pool{
		{
			Start: "10.0.0.1",
			End:   "10.0.0.254",
		},
		{
			Start: "10.0.0.1",
			End:   "10.0.0.10",
		},
	}

	ip := &model.IPAddr{
		Namespace: testNS,
		Ip:        "10.0.0.1",
	}

	m.CreateIP(ctx, pools, ip)

	actual, err := m.DrawIP(ctx, testNS, pools[1], testUUID, "", false /* not random */, true, false, false)
	if err != nil {
		t.Errorf("DrawIP returns err(%v); want success", err)
	}
	if !actual.Equal(net.ParseIP("10.0.0.2")) {
		t.Errorf("DrawIP returns incorrect IP(%v); want 10.0.0.2", actual.String())
	}
}

func TestDrawIPRandom(t *testing.T) {
	r, deferFunc := storage.NewTestRedis()
	defer deferFunc()

	m := NewTestIPManager(r)
	ctx := context.Background()

	// Create a larger pool to test segment randomization
	pool := &model.Pool{
		Start: "10.0.0.1",
		End:   "10.0.3.254", // Spans multiple segments
	}

	// Pre-allocate some IPs
	allocatedIPs := []*model.IPAddr{
		{
			Namespace: testNS,
			Ip:        "10.0.0.1",
			Status:    model.IPAddr_ACTIVE,
		},
		{
			Namespace: testNS,
			Ip:        "10.0.1.100",
			Status:    model.IPAddr_ACTIVE,
		},
		{
			Namespace: testNS,
			Ip:        "10.0.2.200",
			Status:    model.IPAddr_ACTIVE,
		},
	}

	// Create the allocated IPs
	for _, ip := range allocatedIPs {
		err := m.CreateIP(ctx, []*model.Pool{pool}, ip)
		assert.NoError(t, err)
	}

	// Track allocated IPs by segment and UUID
	segmentResults := make(map[int]map[string]bool) // segment -> set of IPs
	uuidToIP := make(map[string]string)             // uuid -> ip mapping

	for i := 0; i < 4; i++ {
		segmentResults[i] = make(map[string]bool)
	}

	// Test 1: Verify random distribution across segments
	iterations := 50
	for i := 0; i < iterations; i++ {
		uuid := fmt.Sprintf("test-uuid-%d", i)
		ip, err := m.DrawIP(ctx, testNS, pool, uuid, "", true /* random */, true, false, false)
		assert.NoError(t, err)

		// Store UUID -> IP mapping
		uuidToIP[uuid] = ip.String()

		// Track segment distribution
		ipParts := strings.Split(ip.String(), ".")
		segment, err := strconv.Atoi(ipParts[2])
		assert.NoError(t, err)

		segmentResults[segment][ip.String()] = true
	}

	// Test 2: Verify same UUID gets same IP
	for uuid, expectedIP := range uuidToIP {
		ip, err := m.DrawIP(ctx, testNS, pool, uuid, "", true /* random */, true, false, false)
		assert.NoError(t, err)
		assert.Equal(t, expectedIP, ip.String(),
			"Same UUID %s should get same IP. Expected %s, got %s",
			uuid, expectedIP, ip.String())
	}

	// Test 3: Verify segment distribution
	segmentsWithMultipleIPs := 0
	for segment, ips := range segmentResults {
		if len(ips) > 1 {
			segmentsWithMultipleIPs++
		}
		t.Logf("Segment %d had %d different IPs allocated", segment, len(ips))
	}

	assert.GreaterOrEqual(t, segmentsWithMultipleIPs, 3,
		"Expected allocations spread across at least 3 segments, got %d segments",
		segmentsWithMultipleIPs)

	// Verify pre-allocated IPs were never returned
	for _, ip := range allocatedIPs {
		segment, _ := strconv.Atoi(strings.Split(ip.Ip, ".")[2])
		assert.False(t, segmentResults[segment][ip.Ip],
			"Pre-allocated IP %s should not be returned", ip.Ip)
	}

	// Clean up all IPs
	for _, ip := range uuidToIP {
		err := m.Deactivate(ctx, testNS, []*model.Pool{pool}, &model.IPAddr{
			Namespace: testNS,
			Ip:        ip,
		})
		assert.NoError(t, err)
	}
}
func TestCreatePoolWhenExistingActivatedIP(t *testing.T) {
	r, def := storage.NewTestRedis()
	defer def()

	m := NewTestIPManager(r)
	ctx := context.Background()

	ips := []*model.IPAddr{
		{
			Namespace: testNS,
			Ip:        "192.168.0.1",
			Status:    model.IPAddr_ACTIVE,
		},
		{
			Namespace: testNS,
			Ip:        "192.168.0.11",
			Status:    model.IPAddr_ACTIVE,
		},
	}
	for _, ip := range ips {
		if err := m.CreateIP(ctx, nil, ip); err != nil {
			t.Fatalf("CreateIP retusn error %v; want success", err)
		}
	}

	network := &model.Network{
		Prefix: "192.168.0.0/24",
	}
	pool := &model.Pool{
		Start: "192.168.0.1",
		End:   "192.168.0.10",
	}
	if err := m.CreatePool(ctx, testNS, network, pool); err != nil {
		t.Fatalf("CreatePool returns error %v; want success", err)
	}

	num, err := r.Client.ZCard(makePoolUsedIPZSet(testNS, net.ParseIP(pool.Start), net.ParseIP(pool.End))).Result()
	if err != nil {
		t.Errorf("ZCard returns error %v; want success", err)
	}
	assert.Equal(t, int64(1), num, "number of used zset should be one")
}

func TestDeletePool(t *testing.T) {
	r, def := storage.NewTestRedis()
	defer def()

	m := NewTestIPManager(r)
	ctx := context.Background()

	network := &model.Network{
		Prefix: "192.168.0.0/24",
	}
	pool := &model.Pool{
		Start: "192.168.0.1",
		End:   "192.168.0.10",
	}

	if err := m.CreateNetwork(ctx, testNS, network); err != nil {
		t.Fatalf("CreateNetwork returns error `%v`; want success", err)
	}

	if err := m.CreatePool(ctx, testNS, network, pool); err != nil {
		t.Fatalf("CreatePool returns error `%v`; want success", err)
	}

	if err := m.DeletePool(ctx, testNS, net.ParseIP(pool.Start), net.ParseIP(pool.End)); err != nil {
		t.Fatalf("DeletePool returns error `%v`; want success", err)
	}

	pools, err := m.GetPools(ctx, testNS)
	if err != nil {
		t.Fatalf("GetPools returns error `%v`; want success", err)
	}
	assert.Equal(t, 0, len(pools), "there should be no pools")

	nothing, err := m.GetPoolsInNetwork(ctx, testNS, network)
	if err != nil {
		t.Fatalf("GetPoolsInNetwork returns error `%v`; want success", err)
	}
	assert.Empty(t, nothing, "there should be no pool in the network")
}
