package ipam

import (
	"net"
	"testing"

	"github.com/hatena/ipdrawer/gen/go/model"
	"github.com/hatena/ipdrawer/pkg/storage"
	"github.com/hatena/ipdrawer/pkg/utils/netutil"
	"google.golang.org/protobuf/proto"
)

var (
	testPrefix = &net.IPNet{
		IP:   net.ParseIP("192.168.0.0"),
		Mask: net.CIDRMask(24, 32),
	}

	testNetwork = &model.Network{
		Prefix:    testPrefix.String(),
		Broadcast: netutil.BroadcastIP(testPrefix).String(),
		Netmask:   netutil.IPMaskToIP(net.CIDRMask(24, 32)).String(),
		Gateways: []string{
			"192.168.0.1",
		},
		Status: model.Network_AVAILABLE,
	}
)

func TestSetNetwork(t *testing.T) {
	r, deferFunc := storage.NewTestRedis()
	defer deferFunc()

	err := setNetwork(r, testNS, testNetwork)
	if err != nil {
		t.Fatalf("Get error: %v", err)
	}
}

func TestGetNetwork(t *testing.T) {
	r, def := storage.NewTestRedis()
	defer def()

	_ = setNetwork(r, testNS, testNetwork)

	resp, err := getNetwork(r, testNS, testPrefix)
	if err != nil {
		t.Fatalf("Got error %v; want success", err)
	}
	if !proto.Equal(resp, testNetwork) {
		t.Errorf("Got wrong Network %v; want %v", resp, testNetwork)
	}
}

func TestGetNetworks(t *testing.T) {
	r, def := storage.NewTestRedis()
	defer def()

	networks := []*model.Network{
		{
			Prefix: "192.168.0.0/24",
		},
		{
			Prefix: "192.168.0.0/22",
		},
	}

	for _, n := range networks {
		if err := setNetwork(r, testNS, n); err != nil {
			t.Fatalf("Got error %v; want success", err)
		}
	}

	resp, err := getNetworks(r, testNS)
	if err != nil {
		t.Fatalf("Got error %v; want success", err)
	}

	if len(resp) != len(networks) {
		t.Errorf("Got wrong number of networks %d; want %d", len(resp), len(networks))
	}
}

func TestSetNetworkWithInvalidModel(t *testing.T) {
	r, def := storage.NewTestRedis()
	defer def()

	testCases := []struct {
		model *model.Network

		desc string
	}{
		{
			model: &model.Network{
				Prefix: "192.168.0.0",
			},
			desc: "Prefix is invalid, should has a prefix",
		},
		{
			model: &model.Network{
				Prefix:    "192.168.0.0/24",
				Broadcast: "invalid-broadcast",
			},
			desc: "Has invalid broadcast",
		},
		{
			model: &model.Network{
				Broadcast: "invalid-broadcast",
			},
			desc: "Must has Prefix",
		},
	}

	for i, tc := range testCases {
		err := setNetwork(r, testNS, tc.model)
		if err == nil {
			t.Errorf("#%d(%s): got no error; want error", i, tc.desc)
		}
	}
}

func TestGetNetworkIncludingPool(t *testing.T) {
	r, def := storage.NewTestRedis()
	defer def()

	err := setNetwork(r, testNS, testNetwork)
	if err != nil {
		t.Fatalf("Get error: %v", err)
	}

	n, err := getNetworkIncludingPool(r, testNS, net.ParseIP("192.168.0.10"), net.ParseIP("192.168.0.12"))
	if err != nil {
		t.Fatalf("Got error: %v; want success", err)
	}
	if !proto.Equal(n, testNetwork) {
		t.Fatalf("Got wrong network: %v", n)
	}
}
