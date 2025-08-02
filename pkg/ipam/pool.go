package ipam

import (
	"net"
	"strings"
	"time"

	"github.com/bufbuild/protovalidate-go"
	"github.com/hatena/ipdrawer/gen/go/model"
	"github.com/hatena/ipdrawer/pkg/storage"
	"github.com/pkg/errors"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func setTSToPool(r *storage.Redis, namespace string, pool *model.Pool) error {
	if err := protovalidate.Validate(pool); err != nil {
		return err
	}

	if existsPool(r, namespace, pool) {
		stored, _ := getPool(r, namespace, net.ParseIP(pool.Start), net.ParseIP(pool.End))
		pool.CreatedAt = stored.CreatedAt
	}

	now := timestamppb.New(time.Now())
	if pool.CreatedAt == nil {
		pool.CreatedAt = now
	} else {
		pool.LastModifiedAt = now
	}

	return nil
}

func setPool(r *storage.Redis, namespace string, pool *model.Pool) error {
	if err := protovalidate.Validate(pool); err != nil {
		return err
	}

	pipe := r.Client.TxPipeline()
	s := net.ParseIP(pool.Start)
	e := net.ParseIP(pool.End)

	k := makePoolDetailsKey(namespace, s, e)
	data, err := proto.Marshal(pool)
	if err != nil {
		return err
	}
	pipe.Set(k, string(data), 0)

	_, err = pipe.Exec()

	return err
}

func getPool(r *storage.Redis, namespace string, start, end net.IP) (*model.Pool, error) {
	k := makePoolDetailsKey(namespace, start, end)

	check, err := r.Client.Exists(k).Result()
	if err != nil || check == 0 {
		return nil, errors.New("not found pool")
	}

	data, err := r.Client.Get(k).Result()
	if err != nil {
		return nil, err
	}
	pool := &model.Pool{}
	if err := proto.Unmarshal([]byte(data), pool); err != nil {
		return nil, err
	}

	return pool, nil
}

func getPools(r *storage.Redis, keys []string) ([]*model.Pool, error) {
	pools := make([]*model.Pool, len(keys))

	if len(keys) == 0 {
		return pools, nil
	}

	data, err := r.Client.MGet(keys...).Result()
	if err != nil {
		return nil, err
	}

	for i, d := range data {
		if s, ok := d.(string); ok {
			pools[i] = &model.Pool{}
			if err := proto.Unmarshal([]byte(s), pools[i]); err != nil {
				return nil, err
			}
		}
	}
	return pools, nil
}

func getPoolsInNetwork(r *storage.Redis, namespace string, prefix *model.Network) ([]*model.Pool, error) {
	_, pre, err := net.ParseCIDR(prefix.Prefix)
	if err != nil {
		return nil, err
	}
	poolKey := makeNetworkPoolKey(namespace, pre)
	keys, err := r.Client.SMembers(poolKey).Result()
	if err != nil {
		return nil, err
	}
	pools := make([]*model.Pool, len(keys))
	for i, key := range keys {
		start := net.ParseIP(key[:strings.Index(key, ",")])
		end := net.ParseIP(key[strings.Index(key, ",")+1:])
		pool, err := getPool(r, namespace, start, end)
		if err != nil {
			return nil, err
		}
		pools[i] = pool
	}
	return pools, nil
}

func existsPool(r *storage.Redis, namespace string, pool *model.Pool) bool {
	s := net.ParseIP(pool.Start)
	e := net.ParseIP(pool.End)

	k := makePoolDetailsKey(namespace, s, e)
	check, _ := r.Client.Exists(k).Result()
	return check != 0
}
