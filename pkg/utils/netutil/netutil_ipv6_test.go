package netutil

import (
	"net"
	"testing"
)

func TestNextPrevIPv6(t *testing.T) {
	testCases := []struct {
		in   string
		next string
		prev string
	}{
		{in: "2001:db8::1", next: "2001:db8::2", prev: "2001:db8::"},
		// Carry across a byte boundary.
		{in: "2001:db8::ff", next: "2001:db8::100", prev: "2001:db8::fe"},
	}

	for i, c := range testCases {
		ip := net.ParseIP(c.in)
		if got := NextIP(ip); !got.Equal(net.ParseIP(c.next)) {
			t.Errorf("#%d: NextIP(%s) = %v, want %s", i, c.in, got, c.next)
		}
		if got := PrevIP(ip); !got.Equal(net.ParseIP(c.prev)) {
			t.Errorf("#%d: PrevIP(%s) = %v, want %s", i, c.in, got, c.prev)
		}
		// NextIP must preserve the IPv6 (16-byte) family.
		if got := NextIP(ip); got.To4() != nil {
			t.Errorf("#%d: NextIP(%s) collapsed to IPv4: %v", i, c.in, got)
		}
	}
}

func TestBigIntRoundTrip(t *testing.T) {
	cases := []struct {
		in string
		v6 bool
	}{
		{in: "2001:db8::dead:beef", v6: true},
		{in: "::1", v6: true},
		{in: "10.0.0.1", v6: false},
	}
	for _, c := range cases {
		ip := net.ParseIP(c.in)
		back := BigIntToIP(IPToBigInt(ip), c.v6)
		if !back.Equal(ip) {
			t.Errorf("round trip %s: got %v", c.in, back)
		}
	}
}

func TestBroadcastIPv6IsNil(t *testing.T) {
	_, n, _ := net.ParseCIDR("2001:db8::/64")
	if got := BroadcastIP(n); got != nil {
		t.Errorf("BroadcastIP for IPv6 should be nil, got %v", got)
	}
}

func TestIPMaskToIPv6(t *testing.T) {
	got := IPMaskToIP(net.CIDRMask(64, 128))
	want := net.ParseIP("ffff:ffff:ffff:ffff::")
	if !got.Equal(want) {
		t.Errorf("IPMaskToIP(/64) = %v, want %v", got, want)
	}
}
