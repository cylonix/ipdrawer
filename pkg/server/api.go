package server

import (
	"encoding/json"
	"fmt"
	"net"
	"sort"
	"strings"

	"github.com/pkg/errors"
	"golang.org/x/net/context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/peer"
	"google.golang.org/grpc/status"

	pvg "github.com/bufbuild/protovalidate-go"
	"github.com/hatena/ipdrawer/gen/go/model"
	server "github.com/hatena/ipdrawer/gen/go/serverpb"
	pm "github.com/hatena/ipdrawer/pkg/model"
	"github.com/hatena/ipdrawer/pkg/utils/netutil"
	"github.com/sirupsen/logrus"
)

var (
	DrawIPSuccessMsg           = "success"
	DrawIPActivationSuccessMsg = "activation success"
)

// maskBits returns the prefix bit-length for the address family of ip: 32 for
// IPv4 and 128 for IPv6.
func maskBits(ip net.IP) int {
	if ip != nil && ip.To4() != nil {
		return 32
	}
	return 128
}

// ListNetwork is an endpoints returning all networks
func (api *APIServer) ListNetwork(
	ctx context.Context,
	req *server.ListNetworkRequest,
) (*server.ListNetworkResponse, error) {
	log := logger.WithFields(logrus.Fields{
		"namespace": req.Namespace,
		"handle":    "list network",
	})
	log.Infoln("list network request")

	if err := pvg.Validate(req); err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	networks, err := api.manager.GetNetworks(ctx, req.Namespace)
	if err != nil {
		return nil, errors.Wrap(err, "failed to get network list")
	}

	// Sort network by prefix
	sort.Slice(networks, func(i, j int) bool {
		return networks[i].Prefix < networks[j].Prefix
	})

	return &server.ListNetworkResponse{
		Networks: networks,
	}, nil
}

// DrawIP returns new IP.
func (api *APIServer) DrawIP(
	ctx context.Context,
	req *server.DrawIPRequest,
) (*server.DrawIPResponse, error) {
	var (
		n         *model.Network
		err       error
		pools     []*model.Pool
		namespace = req.Namespace
	)
	log := logger.WithFields(logrus.Fields{
		"uuid":      req.Uuid,
		"namespace": namespace,
		"handle":    "draw ip",
		"ip":        req.Ip,
	})
	log.Infoln("draw-ip request")
	if err := pvg.Validate(req); err != nil {
		v, _ := json.Marshal(req)
		log.WithField("request", string(v)).
			WithError(err).
			Errorln("invalid request")
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	if req.RangeStart != "" && req.RangeEnd != "" {
		s := net.ParseIP(req.RangeStart)
		e := net.ParseIP(req.RangeEnd)
		if s == nil || e == nil {
			return nil, status.Error(codes.InvalidArgument, "invalid range")
		}
		pool, err := api.manager.GetPool(ctx, namespace, s, e)
		if err != nil {
			return nil, err
		}
		pools = []*model.Pool{pool}
	}
	if len(pools) <= 0 && (req.Name != "" || req.Ip != "") {
		if req.Ip != "" {
			mask := net.CIDRMask(int(req.Mask), maskBits(net.ParseIP(req.Ip)))
			ip := &net.IPNet{
				IP:   net.ParseIP(req.Ip).Mask(mask),
				Mask: mask,
			}
			n, err = api.manager.GetNetworkByIP(ctx, namespace, ip)
		} else {
			n, err = api.manager.GetNetworkByName(ctx, namespace, req.Name)
		}
		if err != nil {
			return nil, err
		}
		pools, err = api.manager.GetPoolsInNetwork(ctx, namespace, n)
		if err != nil {
			return nil, err
		}
	}

	if len(pools) == 0 {
		return nil, status.Error(codes.NotFound, "no address pool")
	}

	target := make([]*model.Pool, 0)
	foundAvails := false
	for _, p := range pools {
		if p.Status == model.Pool_AVAILABLE {
			foundAvails = true
		}
		if p.Status == model.Pool_AVAILABLE &&
			(req.PoolTag == nil || pm.PoolMatchTags(p, []*model.Tag{req.PoolTag})) {
			target = append(target, p)
		}
	}
	if !foundAvails {
		return nil, status.Error(codes.NotFound, "no available address pool")
	}
	if len(target) == 0 {
		return nil, status.Errorf(codes.NotFound, "no matched tags: %v", req.PoolTag.String())
	}

	// If req.Ip is not a prefix, try to draw req.Ip
	wantIP := ""
	if req.Ip != "" {
		ip := net.ParseIP(req.Ip)
		mask := net.CIDRMask(int(req.Mask), maskBits(ip))
		if ip != nil {
			log.WithField("masked-ip", ip.Mask(mask).String()).
				WithField("mask", mask.String()).
				WithField("is-network", ip.Mask(mask).Equal(ip)).
				Infoln("want-ip check")
			if !ip.Mask(mask).Equal(ip) {
				wantIP = req.Ip
			}
		}
	}
	if wantIP == "" && req.MustHaveWantIp {
		return nil, status.Error(codes.InvalidArgument, "must have valid ip if must assign the specific address")
	}

	// Parse exclude prefix if provided
	var excludeNet *net.IPNet
	if req.Exclude != "" {
		_, excludeNet, err = net.ParseCIDR(req.Exclude)
		if err != nil {
			return nil, status.Errorf(codes.InvalidArgument, "invalid exclude prefix: %v", err)
		}
	}

	for _, p := range target {
		ret, err := api.manager.DrawIP(
			ctx, namespace, p, req.Uuid, wantIP, !req.Sequential, /* random */
			true, false, req.MustHaveWantIp, excludeNet,
		)
		if err != nil {
			// log or return error?
			continue
		}
		res := &server.DrawIPResponse{
			Ip:      ret.String(),
			Message: DrawIPSuccessMsg,
		}
		if req.TemporaryReserved {
			return res, nil
		}
		_, err = api.ActivateIP(ctx, &server.ActivateIPRequest{
			Ip:        ret.String(),
			Uuid:      req.Uuid,
			Namespace: namespace,
		})
		if err != nil {
			continue
		}
		res.Message = DrawIPActivationSuccessMsg
		return res, nil
	}

	return nil, status.Error(codes.NotFound, "no address is available")
}

func (api *APIServer) DrawIPEstimatingNetwork(
	ctx context.Context,
	req *server.DrawIPEstimatingNetworkRequest,
) (*server.DrawIPResponse, error) {
	namespace := req.Namespace
	log := logger.WithFields(logrus.Fields{
		"handle":    "DrawIPEstimatingNetwork",
		"namespace": namespace,
	})
	log.Infoln("draw ip estimating network request")
	if err := pvg.Validate(req); err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	var ip net.IP

	// In case that request is passed through grpc-gateway
	if md, ok := metadata.FromIncomingContext(ctx); ok {
		if ips, ok := md["x-forwarded-for"]; ok {
			ip = net.ParseIP(strings.Split(ips[0], ",")[0])
		}
	}
	if ip == nil {
		if pr, ok := peer.FromContext(ctx); ok {
			if tcpAddr, ok := pr.Addr.(*net.TCPAddr); ok {
				ip = tcpAddr.IP
			}
		}
	}
	if ip == nil {
		return nil, status.Error(codes.Internal, "cannot find remote addr")
	}

	n, err := api.manager.GetNetworkIncludingIP(ctx, namespace, ip)
	if err != nil {
		return nil, status.Error(codes.NotFound, err.Error())
	}

	_, pre, err := net.ParseCIDR(n.Prefix)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	ones, _ := pre.Mask.Size()
	return api.DrawIP(ctx, &server.DrawIPRequest{
		Namespace:         namespace,
		Ip:                pre.IP.String(),
		Mask:              int32(ones),
		PoolTag:           req.PoolTag,
		TemporaryReserved: req.TemporaryReserved,
		Sequential:        req.Sequential,
	})
}

func (api *APIServer) GetNetworkIncludingIP(
	ctx context.Context,
	req *server.GetNetworkIncludingIPRequest,
) (*server.GetNetworkResponse, error) {
	if err := pvg.Validate(req); err != nil {
		return nil, status.Errorf(codes.InvalidArgument, err.Error())
	}

	namespace := req.Namespace
	n, err := api.manager.GetNetworkIncludingIP(ctx, namespace, net.ParseIP(req.Ip))
	if err != nil {
		return nil, status.Error(codes.NotFound, err.Error())
	}

	gws := make([]string, len(n.Gateways))
	copy(gws, n.Gateways)

	return &server.GetNetworkResponse{
		Network:         n.Prefix,
		Broadcast:       n.Broadcast,
		Netmask:         n.Netmask,
		DefaultGateways: gws,
		Tags:            n.Tags,
	}, nil
}

func (api *APIServer) GetEstimatedNetwork(
	ctx context.Context,
	req *server.GetEstimatedNetworkRequest,
) (*server.GetNetworkResponse, error) {
	if err := pvg.Validate(req); err != nil {
		return nil, status.Errorf(codes.InvalidArgument, err.Error())
	}

	// In case that request is passed through grpc-gateway
	if md, ok := metadata.FromIncomingContext(ctx); ok {
		if ips, ok := md["x-forwarded-for"]; ok {
			return api.GetNetworkIncludingIP(ctx, &server.GetNetworkIncludingIPRequest{
				Ip:        strings.Split(ips[0], ",")[0],
				Namespace: req.Namespace,
			})
		}
	}

	var ip net.IP
	if pr, ok := peer.FromContext(ctx); ok {
		if tcpAddr, ok := pr.Addr.(*net.TCPAddr); ok {
			ip = tcpAddr.IP
		}
	}
	if ip == nil {
		return nil, status.Error(codes.Internal, "Not support remote addr")
	}

	return api.GetNetworkIncludingIP(ctx, &server.GetNetworkIncludingIPRequest{
		Ip:        ip.String(),
		Namespace: req.Namespace,
	})
}

func (api *APIServer) CreateIP(
	ctx context.Context,
	addr *model.IPAddr,
) (*server.CreateIPResponse, error) {
	if err := pvg.Validate(addr); err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	namespace := addr.Namespace
	ip := net.ParseIP(addr.Ip)

	n, err := api.manager.GetNetworkIncludingIP(ctx, namespace, ip)
	if err != nil {
		return nil, err
	}

	pools, err := api.manager.GetPoolsInNetwork(ctx, namespace, n)
	if err != nil {
		return nil, err
	}

	if err := api.manager.CreateIP(ctx, pools, addr); err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &server.CreateIPResponse{}, nil
}

func (api *APIServer) ActivateIP(
	ctx context.Context,
	req *server.ActivateIPRequest,
) (*server.CreateIPResponse, error) {
	if err := pvg.Validate(req); err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	addr := &model.IPAddr{
		Ip:        req.Ip,
		Status:    model.IPAddr_ACTIVE,
		Tags:      req.Tags,
		Uuid:      req.Uuid,
		Namespace: req.Namespace,
	}

	return api.CreateIP(ctx, addr)
}

func (api *APIServer) DeactivateIP(
	ctx context.Context,
	req *server.DeactivateIPRequest,
) (*server.DeactivateIPResponse, error) {
	if err := pvg.Validate(req); err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	namespace := req.Namespace
	ip := net.ParseIP(req.Ip)
	addr := &model.IPAddr{
		Ip: req.Ip,
	}

	n, err := api.manager.GetNetworkIncludingIP(ctx, namespace, ip)
	if err != nil {
		return nil, err
	}

	pools, err := api.manager.GetPoolsInNetwork(ctx, namespace, n)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	if len(pools) == 0 {
		return nil, status.Errorf(
			codes.NotFound, "pool not found: %s", ip.String(),
		)
	}

	if err := api.manager.Deactivate(ctx, namespace, pools, addr); err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &server.DeactivateIPResponse{}, nil
}

func (api *APIServer) UpdateIP(
	ctx context.Context,
	addr *model.IPAddr,
) (*server.UpdateIPResponse, error) {
	if err := pvg.Validate(addr); err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	if err := api.manager.UpdateIP(ctx, addr); err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &server.UpdateIPResponse{}, nil
}

func (api *APIServer) GetNetwork(
	ctx context.Context,
	req *server.GetNetworkRequest,
) (*server.GetNetworkResponse, error) {
	log := logger.WithFields(logrus.Fields{
		"namespace": req.Namespace,
		"handle":    "get network",
		"ip":        req.Ip,
	})
	log.Infoln("get network request")

	if err := pvg.Validate(req); err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}
	namespace := req.Namespace
	var n *model.Network
	var err error
	if req.Name == "" {
		ip := net.ParseIP(req.Ip)
		n, err = api.manager.GetNetworkByIP(ctx, namespace, &net.IPNet{
			IP:   ip,
			Mask: net.CIDRMask(int(req.Mask), maskBits(ip)),
		})
	} else {
		n, err = api.manager.GetNetworkByName(ctx, namespace, req.Name)
	}
	if err != nil {
		return nil, err
	}

	gws := make([]string, len(n.Gateways))
	copy(gws, n.Gateways)

	return &server.GetNetworkResponse{
		Network:         n.Prefix,
		Broadcast:       n.Broadcast,
		Netmask:         n.Netmask,
		DefaultGateways: gws,
		Tags:            n.Tags,
	}, nil
}

func (api *APIServer) CreateNetwork(
	ctx context.Context,
	req *server.CreateNetworkRequest,
) (*server.CreateNetworkResponse, error) {
	if err := pvg.Validate(req); err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	namespace := req.Namespace
	_, ipnet, err := net.ParseCIDR(fmt.Sprintf("%s/%d", req.Ip, req.Mask))
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	netmask := netutil.IPMaskToIP(ipnet.Mask)
	// IPv6 has no broadcast address; BroadcastIP returns nil for IPv6 networks.
	broadcastStr := ""
	if broadcast := netutil.BroadcastIP(ipnet); broadcast != nil {
		broadcastStr = broadcast.String()
	}

	n := &model.Network{
		Prefix:    ipnet.String(),
		Broadcast: broadcastStr,
		Netmask:   netmask.String(),
		Gateways:  req.DefaultGateways,
		Tags:      req.Tags,
		Status:    req.Status,
	}

	if err := api.manager.CreateNetwork(ctx, namespace, n); err != nil {
		return nil, err
	}

	return &server.CreateNetworkResponse{}, nil
}

// DeleteNetwork deletes the network.
func (api *APIServer) DeleteNetwork(
	ctx context.Context,
	req *server.DeleteNetworkRequest,
) (*server.DeleteNetworkResponse, error) {
	if err := pvg.Validate(req); err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	namespace := req.Namespace
	ip := net.ParseIP(req.Ip)
	network, err := api.manager.GetNetworkByIP(ctx, namespace, &net.IPNet{
		IP:   ip,
		Mask: net.CIDRMask(int(req.Mask), maskBits(ip)),
	})
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	if err := api.manager.DeleteNetwork(ctx, namespace, network); err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &server.DeleteNetworkResponse{}, nil
}

// UpdateNetwork updates the network.
func (api *APIServer) UpdateNetwork(
	ctx context.Context,
	network *model.Network,
) (*server.UpdateNetworkResponse, error) {
	if err := pvg.Validate(network); err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	if err := api.manager.UpdateNetwork(ctx, network); err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	return &server.UpdateNetworkResponse{}, nil
}

// GetPoolsInNetwork returns all pools in a given network.
func (api *APIServer) GetPoolsInNetwork(
	ctx context.Context,
	req *server.GetPoolsInNetworkRequest,
) (*server.GetPoolsInNetworkResponse, error) {
	if err := pvg.Validate(req); err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	namespace := req.Namespace
	ip := net.ParseIP(req.Ip)
	n, err := api.manager.GetNetworkByIP(ctx, namespace, &net.IPNet{
		IP:   ip,
		Mask: net.CIDRMask(int(req.Mask), maskBits(ip)),
	})
	if err != nil {
		return nil, err
	}
	pools, err := api.manager.GetPoolsInNetwork(ctx, namespace, n)
	if err != nil {
		return nil, err
	}

	return &server.GetPoolsInNetworkResponse{
		Pools: pools,
	}, nil
}

func (api *APIServer) CreatePool(
	ctx context.Context,
	req *server.CreatePoolRequest,
) (*server.CreatePoolResponse, error) {
	if err := pvg.Validate(req); err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	namespace := req.Namespace
	reqIP := net.ParseIP(req.Ip)
	ip := &net.IPNet{
		IP:   reqIP,
		Mask: net.CIDRMask(int(req.Mask), maskBits(reqIP)),
	}

	pool := &model.Pool{
		Start:  req.Pool.Start,
		End:    req.Pool.End,
		Status: req.Pool.Status,
		Tags:   req.Pool.Tags,
	}

	n, err := api.manager.GetNetworkByIP(ctx, namespace, ip)
	if err != nil {
		return nil, err
	}

	if err := api.manager.CreatePool(ctx, namespace, n, pool); err != nil {
		return nil, err
	}

	return &server.CreatePoolResponse{}, nil
}

func (api *APIServer) ListIP(
	ctx context.Context,
	req *server.ListIPRequest,
) (*server.ListIPResponse, error) {
	if err := pvg.Validate(req); err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	namespace := req.Namespace
	addrs, err := api.manager.ListIP(ctx, namespace)
	if err != nil {
		return nil, errors.Wrap(err, "cannot get ip list")
	}

	// Sort by IP
	sort.Slice(addrs, func(i, j int) bool {
		return addrs[i].Ip < addrs[j].Ip
	})

	return &server.ListIPResponse{
		Ips: addrs,
	}, nil
}

func (api *APIServer) ListTemporaryReservedIP(
	ctx context.Context,
	req *server.ListTemporaryReservedIPRequest,
) (*server.ListTemporaryReservedIPResponse, error) {
	if err := pvg.Validate(req); err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	namespace := req.Namespace
	addrs, err := api.manager.GetTemporaryReservedIPs(ctx, namespace)
	if err != nil {
		return nil, errors.Wrap(err, "cannot get ip list")
	}

	// Sort IP
	sort.Slice(addrs, func(i, j int) bool {
		return addrs[i].Ip < addrs[j].Ip
	})

	return &server.ListTemporaryReservedIPResponse{
		Ips: addrs,
	}, nil
}

func (api *APIServer) ListPool(
	ctx context.Context,
	req *server.ListPoolRequest,
) (*server.ListPoolResponse, error) {
	if err := pvg.Validate(req); err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	namespace := req.Namespace
	pools, err := api.manager.GetPools(ctx, namespace)
	if err != nil {
		return nil, errors.Wrap(err, "cannot get pool list")
	}

	// Sort pool by an IP range
	sort.Slice(pools, func(i, j int) bool {
		if pools[i].Start == pools[j].Start {
			return pools[i].End < pools[j].End
		}
		return pools[i].Start < pools[j].End
	})

	return &server.ListPoolResponse{
		Pools: pools,
	}, nil
}

// GetIPInPool is an endpoint to get IPs in a given pool.
func (api *APIServer) GetIPInPool(
	ctx context.Context,
	req *server.GetIPInPoolRequest,
) (*server.GetIPInPoolResponse, error) {
	if err := pvg.Validate(req); err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	namespace := req.Namespace
	pool, err := api.manager.GetPool(ctx, namespace, net.ParseIP(req.RangeStart), net.ParseIP(req.RangeEnd))
	if err != nil {
		return nil, status.Error(codes.NotFound, err.Error())
	}

	addrs, err := api.manager.ListIP(ctx, namespace)
	if err != nil {
		return nil, errors.Wrap(err, "cannot get ip list")
	}

	ret := make([]*model.IPAddr, 0)
	for _, ip := range addrs {
		if pm.PoolContains(pool, net.ParseIP(ip.Ip)) {
			ret = append(ret, ip)
		}
	}

	// Sort by IP
	sort.Slice(ret, func(i, j int) bool {
		return ret[i].Ip < ret[j].Ip
	})

	return &server.GetIPInPoolResponse{
		Pool: pool,
		Ips:  ret,
	}, nil
}

// UpdatePool updates a given pool.
func (api *APIServer) UpdatePool(
	ctx context.Context,
	pool *model.Pool,
) (*server.UpdatePoolResponse, error) {
	if err := pvg.Validate(pool); err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	if err := api.manager.UpdatePool(ctx, pool); err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &server.UpdatePoolResponse{}, nil
}

func (api *APIServer) DeletePool(
	ctx context.Context,
	req *server.DeletePoolRequest,
) (*server.DeletePoolResponse, error) {
	if err := pvg.Validate(req); err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	namespace := req.Namespace
	if err := api.manager.DeletePool(ctx, namespace, net.ParseIP(req.RangeStart), net.ParseIP(req.RangeEnd)); err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &server.DeletePoolResponse{}, nil
}
