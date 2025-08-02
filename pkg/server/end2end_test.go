package server

import (
	"net"
	"net/http"
	"testing"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	client "github.com/hatena/ipdrawer/gen/client"
	"github.com/sirupsen/logrus"
)

func (te *test) startServer() {
	la := "localhost:0"
	lis, err := net.Listen("tcp", la)
	if err != nil {
		te.t.Fatalf("Failed to listen: %v", err)
	}
	te.api.lis = lis
	_, port, err := net.SplitHostPort(lis.Addr().String())
	if err != nil {
		te.t.Fatalf("Failed to parse listener address: %v", err)
	}
	addr := "localhost:" + port
	go te.api.Start()
	te.srvAddr = addr
	te.scheme = "http"
	te.host = addr
	te.base = te.scheme + "://" + addr
	logrus.Infof("started server at %v", te.base)
}

func (te *test) ClientConn() *grpc.ClientConn {
	if te.cc != nil {
		return te.cc
	}

	var err error
	te.cc, err = grpc.NewClient(te.srvAddr, []grpc.DialOption{
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	}...)
	if err != nil {
		te.t.Fatalf("Dial(%q) = %v", te.srvAddr, err)
	}
	return te.cc
}

func TestDrawIPEstimatingNetwork_E2E_ViaGW(t *testing.T) {
	te := newTest(t)
	te.startServer()
	defer te.tearDown()

	// Use testdata
	te.manager.CreateNetwork(te.ctx, testNS, testNetwork)
	te.manager.CreatePool(te.ctx, testNS, testNetwork, testPool)

	testCases := []struct {
		// Input
		remote   string
		tagKey   string
		tagValue string

		// Expected
		status int
		ip     string

		desc string
	}{
		{
			remote:   "192.168.0.1",
			tagKey:   "Role",
			tagValue: "test",

			status: http.StatusOK,
			ip:     "192.168.0.2",

			desc: "Ideal case",
		},
		{
			remote:   "192.168.0.1",
			tagKey:   "Role",
			tagValue: "nothing",

			status: http.StatusNotFound,

			desc: "Pool tag Role=nothing does not exists",
		},
		{
			remote:   "192.168.1.1",
			tagKey:   "Role",
			tagValue: "test",

			status: http.StatusNotFound,

			desc: "Network does not exists",
		},
	}
	var md runtime.ServerMetadata
	ctx := runtime.NewServerMetadataContext(te.ctx, md)

	for i, tc := range testCases {
		cfg := client.NewConfiguration()
		cfg.Host = te.host
		cfg.Scheme = te.scheme
		cfg.AddDefaultHeader("X-Forwarded-For", tc.remote)
		cl := client.NewAPIClient(cfg).NetworkServiceV0API
		req := cl.NetworkServiceV0DrawIPEstimatingNetwork(ctx, testNS)
		req = req.
			PoolTagKey(tc.tagKey).
			PoolTagValue(tc.tagValue).
			TemporaryReserved(false).
			Sequential(true)
		resp, apiresp, err := req.Execute()

		if tc.status != http.StatusOK {
			if apiresp != nil && apiresp.StatusCode == tc.status {
				continue
			}
		}

		if err != nil {
			t.Errorf("#%d(desc=%s): cl.DrawIPEstimatingNetwork failed with %v; want success",
				i, tc.desc, err)
			continue
		}

		if apiresp.StatusCode != tc.status {
			t.Errorf("#%d(desc=%s): cl.DrawIPEstimatingNetwork returns %v; want %d %v",
				i, tc.desc, apiresp.Status, tc.status, http.StatusText(tc.status))
		}

		if *resp.IP != tc.ip {
			t.Errorf("#%d(desc=%s): cl.DrawIPEstimatingNetwork returns unexpected IP(%v); want IP(%v)",
				i, tc.desc, resp.IP, tc.ip)
		}
	}
}

func TestGetEstimatedNetwork_E2E_ViaGW(t *testing.T) {
	te := newTest(t)
	te.startServer()
	defer te.tearDown()

	// Use testdata
	err1 := te.manager.CreateNetwork(te.ctx, testNS, testNetwork)
	err2 := te.manager.CreatePool(te.ctx, testNS, testNetwork, testPool)
	if err1 != nil || err2 != nil {
		t.Errorf("set up failed: %v %v", err1, err2)
	}

	testCases := []struct {
		// Input
		remote string

		// Expected
		status  int
		network string

		desc string
	}{
		{
			remote: "192.168.0.1",

			status:  http.StatusOK,
			network: "192.168.0.0/24",

			desc: "Ideal case",
		},
		{
			remote: "192.168.1.1",

			status: http.StatusNotFound,

			desc: "Network including 192.168.1.1 does not exist",
		},
	}

	ctx := runtime.NewServerMetadataContext(te.ctx, runtime.ServerMetadata{})
	for i, tc := range testCases {
		cfg := client.NewConfiguration()
		cfg.Host = te.host
		cfg.Scheme = te.scheme
		cfg.AddDefaultHeader("X-Forwarded-For", tc.remote)
		cl := client.NewAPIClient(cfg).NetworkServiceV0API
		req := cl.NetworkServiceV0GetEstimatedNetwork(ctx, testNS)
		resp, apiresp, err := req.Execute()

		// res error is set if apiresp status is not 200.
		if tc.status != http.StatusOK {
			if apiresp != nil && apiresp.StatusCode == tc.status {
				continue
			}
		}
		if err != nil {
			t.Errorf("#%d(desc=%s): cl.GetEstimatedNetwork() failed with %v; want success",
				i, tc.desc, err)
			continue
		}

		if apiresp.StatusCode != tc.status {
			t.Errorf("#%d(desc=%s): cl.GetEstimatedNetwork() returns %v; want %d %v",
				i, tc.desc, apiresp.Status, tc.status, http.StatusText(tc.status))
		}

		if *resp.Network != tc.network {
			t.Errorf("#%d(desc=%s): cl.GetEstimatedNetwork() returns unexpected network(%v); want network(%v)",
				i, tc.desc, resp.Network, tc.network)
		}
	}
}

func TestActivateIP_E2E_ViaGW(t *testing.T) {
	te := newTest(t)
	te.startServer()
	defer te.tearDown()

	err1 := te.manager.CreateNetwork(te.ctx, testNS, testNetwork)
	err2 := te.manager.CreatePool(te.ctx, testNS, testNetwork, testPool)
	if err1 != nil || err2 != nil {
		t.Errorf("set up failed: %v %v", err1, err2)
	}

	ctx := runtime.NewServerMetadataContext(te.ctx, runtime.ServerMetadata{})
	cfg := client.NewConfiguration()
	cfg.Host = te.host
	cfg.Scheme = te.scheme
	cl := client.NewAPIClient(cfg).IPServiceV0API
	req := cl.IPServiceV0ActivateIP(ctx, testNS, "192.168.0.111")
	key := "Role"
	val := "test"
	req = req.Body(client.IPServiceV0ActivateIPBody{
		Tags: []client.ModelTag{
			{
				Key:   &key,
				Value: &val,
			},
		},
	})
	_, apiresp, err := req.Execute()

	if err != nil {
		t.Fatalf("cl.ActivateIP failed with %v; want success", err)
	}
	if apiresp.StatusCode != http.StatusOK {
		t.Errorf("cl.ActivateIP returns %v; want %v", apiresp.Status, http.StatusOK)
	}
}
