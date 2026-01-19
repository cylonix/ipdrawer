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

		actual, err := m.DrawIP(ctx, testNS, c.pool, testUUID, c.wanted, false /* not random */, true, false, c.mustHaveWanted, nil /* no exclude */)
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

	actual, err := m.DrawIP(ctx, testNS, pools[1], testUUID, "", false /* not random */, true, false, false, nil /* no exclude */)
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
		ip, err := m.DrawIP(ctx, testNS, pool, uuid, "", true /* random */, true, false, false, nil /* no exclude */)
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
		ip, err := m.DrawIP(ctx, testNS, pool, uuid, "", true /* random */, true, false, false, nil /* no exclude */)
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

func TestDrawIPWithExclude(t *testing.T) {
	r, deferFunc := storage.NewTestRedis()
	defer deferFunc()

	m := NewTestIPManager(r)
	ctx := context.Background()

	// Create a pool similar to CGNAT range 100.64.0.0/10
	// This spans 100.64.0.0 - 100.127.255.255 (4M IPs)
	// We'll use a smaller subset for testing
	pool := &model.Pool{
		Start: "100.115.90.1",
		End:   "100.115.95.254",
	}

	// ChromeOS VM range to exclude: 100.115.92.0/23
	// This covers 100.115.92.0 - 100.115.93.255
	_, excludeNet, err := net.ParseCIDR("100.115.92.0/23")
	assert.NoError(t, err)

	// Track allocated IPs
	allocatedIPs := make(map[string]bool)
	excludedIPCount := 0
	validIPCount := 0

	// Allocate many IPs and verify none fall in the excluded range
	iterations := 100
	for i := 0; i < iterations; i++ {
		uuid := fmt.Sprintf("exclude-test-uuid-%d", i)
		ip, err := m.DrawIP(ctx, testNS, pool, uuid, "", true /* random */, true, false, false, excludeNet)
		assert.NoError(t, err)

		ipStr := ip.String()

		// Check IP is not in excluded range
		if excludeNet.Contains(ip) {
			excludedIPCount++
			t.Errorf("Allocated IP %s falls within excluded range %s", ipStr, excludeNet.String())
		} else {
			validIPCount++
		}

		// Check for duplicates
		if allocatedIPs[ipStr] {
			t.Errorf("Duplicate IP allocated: %s", ipStr)
		}
		allocatedIPs[ipStr] = true
	}

	t.Logf("Allocated %d IPs, %d valid (outside exclude), %d in excluded range",
		iterations, validIPCount, excludedIPCount)
	assert.Equal(t, 0, excludedIPCount, "No IPs should be allocated in excluded range")
	assert.Equal(t, iterations, validIPCount, "All IPs should be outside excluded range")

	// Verify allocated IPs are in valid ranges
	// Valid ranges: 100.115.90.1 - 100.115.91.255 and 100.115.94.0 - 100.115.95.254
	for ipStr := range allocatedIPs {
		ip := net.ParseIP(ipStr)
		parts := strings.Split(ipStr, ".")
		thirdOctet, _ := strconv.Atoi(parts[2])

		// Should be in 90-91 or 94-95 range
		validRange := (thirdOctet >= 90 && thirdOctet <= 91) || (thirdOctet >= 94 && thirdOctet <= 95)
		assert.True(t, validRange, "IP %s should be in valid range (90-91 or 94-95), got third octet %d", ipStr, thirdOctet)

		// Double-check not in excluded range
		assert.False(t, excludeNet.Contains(ip), "IP %s should not be in excluded range", ipStr)
	}
}

func TestDrawIPWithExcludeSequential(t *testing.T) {
	r, deferFunc := storage.NewTestRedis()
	defer deferFunc()

	m := NewTestIPManager(r)
	ctx := context.Background()

	// Pool that spans the excluded range with enough IPs
	// Range: 100.115.91.250 - 100.115.94.50
	// Usable: 91.250-91.255 (6 IPs) + 94.0-94.50 (51 IPs) = 57 IPs
	// Excluded: 92.0-93.255 (512 IPs)
	pool := &model.Pool{
		Start: "100.115.91.250",
		End:   "100.115.94.50",
	}

	// Exclude 100.115.92.0/23 (100.115.92.0 - 100.115.93.255)
	_, excludeNet, err := net.ParseCIDR("100.115.92.0/23")
	assert.NoError(t, err)

	// Sequential allocation should skip from 100.115.91.255 to 100.115.94.0
	allocatedIPs := []string{}

	// Allocate 15 IPs - enough to cross the excluded range
	for i := 0; i < 15; i++ {
		uuid := fmt.Sprintf("seq-exclude-test-%d", i)
		ip, err := m.DrawIP(ctx, testNS, pool, uuid, "", false /* sequential */, true, false, false, excludeNet)
		assert.NoError(t, err)

		ipStr := ip.String()
		allocatedIPs = append(allocatedIPs, ipStr)

		// Verify not in excluded range
		assert.False(t, excludeNet.Contains(ip), "Sequential IP %s should not be in excluded range", ipStr)
	}

	t.Logf("Sequential allocation with exclude: %v", allocatedIPs)

	// First allocations should be in 100.115.91.x range
	// Then jump to 100.115.94.x range
	foundJump := false
	for i := 1; i < len(allocatedIPs); i++ {
		prevParts := strings.Split(allocatedIPs[i-1], ".")
		currParts := strings.Split(allocatedIPs[i], ".")
		prevThird, _ := strconv.Atoi(prevParts[2])
		currThird, _ := strconv.Atoi(currParts[2])

		// Check for jump from 91 to 94 (skipping 92 and 93)
		if prevThird == 91 && currThird == 94 {
			foundJump = true
			t.Logf("Found expected jump from %s to %s (skipping excluded range)", allocatedIPs[i-1], allocatedIPs[i])
		}
	}

	assert.True(t, foundJump, "Sequential allocation should jump over excluded range 92-93")
}

