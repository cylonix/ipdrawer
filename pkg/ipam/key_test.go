package ipam

import (
	"net"
	"testing"
)

func TestKey(t *testing.T) {
	ip := net.ParseIP("192.168.0.2")
	_, ipnet, _ := net.ParseCIDR("192.168.0.0/24")

	testCases := []struct {
		e string
		a string
	}{
		{
			a: makeIPDetailsKey(testNS, ip),
			e: "test-namespace:ip:192.168.0.2:details",
		},
		{
			a: makeIPTempReserved(testNS, ip),
			e: "test-namespace:ip:192.168.0.2:temporary_reserved",
		},
		{
			a: makeNetworkListKey(testNS),
			e: "test-namespace:network:list",
		},
		{
			a: makeNetworkDetailsKey(testNS, ipnet),
			e: "test-namespace:network:192.168.0.0/24:details",
		},
		{
			a: makeNetworkPoolKey(testNS, ipnet),
			e: "test-namespace:network:192.168.0.0/24:details:pools",
		},
		{
			a: makePoolDetailsKey(testNS, ip, ip),
			e: "test-namespace:pool:192.168.0.2,192.168.0.2:details",
		},
	}

	for i, tc := range testCases {
		if tc.e != tc.a {
			t.Errorf("%d: expected %s, but found %s", i, tc.e, tc.a)
		}
	}
}
