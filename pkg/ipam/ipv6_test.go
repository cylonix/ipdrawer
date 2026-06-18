package ipam

import (
	"fmt"
	"net"
	"testing"

	"github.com/stretchr/testify/assert"
	"golang.org/x/net/context"

	"github.com/hatena/ipdrawer/gen/go/model"
	"github.com/hatena/ipdrawer/pkg/storage"
)

// TestIPv6KeyRoundTrip verifies that the Redis key parsers correctly handle
// IPv6 literals (which contain colons that collide with the key delimiter).
func TestIPv6KeyRoundTrip(t *testing.T) {
	ip := net.ParseIP("2001:db8::1")
	s := net.ParseIP("2001:db8::1")
	e := net.ParseIP("2001:db8::ff")

	gotIP, err := parseIPDetailsKey(makeIPDetailsKey(testNS, ip))
	assert.NoError(t, err)
	assert.True(t, gotIP.Equal(ip), "parseIPDetailsKey round trip: got %v", gotIP)

	gotIP, err = parseTempReservedIPKey(makeIPTempReserved(testNS, ip))
	assert.NoError(t, err)
	assert.True(t, gotIP.Equal(ip), "parseTempReservedIPKey round trip: got %v", gotIP)

	ps, pe, pip, uuid, err := parsePoolUuidIPKey(makeIPPoolUuid(testNS, s, e, ip, "dev-uuid"))
	assert.NoError(t, err)
	assert.True(t, ps.Equal(s) && pe.Equal(e) && pip.Equal(ip), "parsePoolUuidIPKey IPs")
	assert.Equal(t, "dev-uuid", uuid)

	ds, de, err := ParsePoolDetailsKey(makePoolDetailsKey(testNS, s, e))
	assert.NoError(t, err)
	assert.True(t, ds.Equal(s) && de.Equal(e), "ParsePoolDetailsKey round trip")
}

func TestIPv6Activation(t *testing.T) {
	r, deferFunc := storage.NewTestRedis()
	defer deferFunc()

	m := NewTestIPManager(r)
	ctx := context.Background()

	pool := &model.Pool{
		Start: "2001:db8::1",
		End:   "2001:db8::ff",
	}

	for _, ip := range []string{"2001:db8::1", "2001:db8::a"} {
		if err := m.CreateIP(ctx, []*model.Pool{pool}, &model.IPAddr{Namespace: testNS, Ip: ip}); err != nil {
			t.Fatalf("CreateIP(%s) error: %v", ip, err)
		}
	}

	zkey := makePoolUsedIPZSet(testNS, net.ParseIP(pool.Start), net.ParseIP(pool.End))
	cnt, err := r.Client.ZCard(zkey).Result()
	assert.NoError(t, err)
	assert.Equal(t, int64(2), cnt, "used zset should hold both activated IPv6 addresses")
}

func TestIPv6DrawSequential(t *testing.T) {
	r, deferFunc := storage.NewTestRedis()
	defer deferFunc()

	m := NewTestIPManager(r)
	ctx := context.Background()

	pool := &model.Pool{
		Start: "2001:db8::1",
		End:   "2001:db8::ff",
	}

	// Activate ::1 and ::3, leaving ::2 as the next free sequential address.
	for _, ip := range []string{"2001:db8::1", "2001:db8::3"} {
		assert.NoError(t, m.CreateIP(ctx, []*model.Pool{pool}, &model.IPAddr{
			Namespace: testNS, Ip: ip, Status: model.IPAddr_ACTIVE,
		}))
	}

	got, err := m.DrawIP(ctx, testNS, pool, testUUID, "", false /* sequential */, true, false, false, nil)
	assert.NoError(t, err)
	assert.Equal(t, net.ParseIP("2001:db8::2").String(), got.String())
}

func TestIPv6DrawRandom(t *testing.T) {
	r, deferFunc := storage.NewTestRedis()
	defer deferFunc()

	m := NewTestIPManager(r)
	ctx := context.Background()

	pool := &model.Pool{
		Start: "2001:db8::1",
		End:   "2001:db8::ff",
	}
	_, poolNet, _ := net.ParseCIDR("2001:db8::/120") // 2001:db8::/120 == ::0 - ::ff

	seen := make(map[string]bool)
	for i := 0; i < 20; i++ {
		uuid := fmt.Sprintf("v6-rand-%d", i)
		got, err := m.DrawIP(ctx, testNS, pool, uuid, "", true /* random */, true, false, false, nil)
		assert.NoError(t, err)
		assert.True(t, poolNet.Contains(got), "drawn IP %v should be within pool range", got)
		assert.False(t, seen[got.String()], "duplicate IP allocated: %v", got)
		seen[got.String()] = true
	}
}

func TestIPv6DrawWithExclude(t *testing.T) {
	r, deferFunc := storage.NewTestRedis()
	defer deferFunc()

	m := NewTestIPManager(r)
	ctx := context.Background()

	pool := &model.Pool{
		Start: "2001:db8::1",
		End:   "2001:db8::ff",
	}
	_, excludeNet, err := net.ParseCIDR("2001:db8::80/121") // ::80 - ::ff
	assert.NoError(t, err)

	for i := 0; i < 30; i++ {
		uuid := fmt.Sprintf("v6-excl-%d", i)
		got, err := m.DrawIP(ctx, testNS, pool, uuid, "", true /* random */, true, false, false, excludeNet)
		assert.NoError(t, err)
		assert.False(t, excludeNet.Contains(got), "drawn IP %v must not be in excluded range", got)
	}
}
