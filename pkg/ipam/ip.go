package ipam

import (
	"fmt"
	"net"
	"time"

	pvg "github.com/bufbuild/protovalidate-go"
	"github.com/pkg/errors"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/timestamppb"

	"github.com/hatena/ipdrawer/gen/go/model"
	"github.com/hatena/ipdrawer/pkg/storage"
)

var (
	errNoIPAvailable    = errors.New("no ip available")
	errAddrNotAvailable = errors.New("address is not available")
	ErrNetworkNotFound  = errors.New("network not found")
	errUnmarshalAddr    = errors.New("failed to unmarshal addr")
)

func setIPAddrTimestamp(r *storage.Redis, namespace string, addr *model.IPAddr) error {
	if err := pvg.Validate(addr); err != nil {
		return err
	}

	ip := net.ParseIP(addr.Ip)
	if ip == nil {
		return errors.New("failed to parse ip")
	}

	if existsIP(r, namespace, addr) {
		stored, _ := getIPAddr(r, namespace, ip)
		addr.CreatedAt = stored.CreatedAt
	}

	now := timestamppb.New(time.Now())
	if addr.CreatedAt == nil {
		addr.CreatedAt = now
	} else {
		addr.LastModifiedAt = now
	}

	return nil
}

func setIPAddr(r *storage.Redis, namespace string, addr *model.IPAddr) error {
	if err := pvg.Validate(addr); err != nil {
		return err
	}

	data, err := proto.Marshal(addr)
	if err != nil {
		return errors.Wrap(err, "failed to marshal ip address")
	}

	k := makeIPDetailsKey(namespace, net.ParseIP(addr.Ip))
	if _, err := r.Client.Set(k, string(data), 0).Result(); err != nil {
		return errors.Wrap(err, "failed to save to db")
	}

	return nil
}

func getIPAddr(r *storage.Redis, namespace string, ip net.IP) (*model.IPAddr, error) {
	k := makeIPDetailsKey(namespace, ip)
	data, err := r.Client.Get(k).Result()
	if err != nil {
		return nil, err
	}

	ipaddr := &model.IPAddr{}
	if err := proto.Unmarshal([]byte(data), ipaddr); err != nil {
		return nil, fmt.Errorf("%w: %v", errUnmarshalAddr, err)
	}

	return ipaddr, nil
}

func getIPAddrs(r *storage.Redis, namespace string, ips []net.IP) ([]*model.IPAddr, error) {
	addrs := make([]*model.IPAddr, len(ips))

	if len(ips) == 0 {
		return addrs, nil
	}

	keys := make([]string, len(ips))
	for i, ip := range ips {
		keys[i] = makeIPDetailsKey(namespace, ip)
	}
	data, err := r.Client.MGet(keys...).Result()
	if err != nil {
		return nil, err
	}

	for i, d := range data {
		if s, ok := d.(string); ok {
			addrs[i] = &model.IPAddr{}
			if err := proto.Unmarshal([]byte(s), addrs[i]); err != nil {
				return nil, err
			}
		}
	}
	return addrs, nil
}

func existsIP(r *storage.Redis, namespace string, addr *model.IPAddr) bool {
	ip := net.ParseIP(addr.Ip)
	check, _ := r.Client.Exists(makeIPDetailsKey(namespace, ip)).Result()
	return check != 0
}
