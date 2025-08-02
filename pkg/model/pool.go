package model

import (
	"fmt"
	"net"

	"github.com/hatena/ipdrawer/gen/go/model"
	nu "github.com/hatena/ipdrawer/pkg/utils/netutil"
)

func PoolKey(p *model.Pool) string {
	return fmt.Sprintf("%s,%s", p.Start, p.End)
}

func PoolContains(p *model.Pool, ip net.IP) bool {
	s := nu.IP2Uint(net.ParseIP(p.Start))
	e := nu.IP2Uint(net.ParseIP(p.End))
	i := nu.IP2Uint(ip)
	return s <= i && i <= e
}

func PoolMatchTags(p *model.Pool, tags []*model.Tag) bool {
	var flag bool
	for _, target := range tags {
		flag = false
		for _, t := range p.Tags {
			if target.Key == t.Key && target.Value == t.Value {
				flag = true
			}
		}
		if !flag {
			return false
		}
	}
	return true
}
