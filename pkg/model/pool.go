package model

import (
	"bytes"
	"fmt"
	"net"

	"github.com/hatena/ipdrawer/gen/go/model"
)

func PoolKey(p *model.Pool) string {
	return fmt.Sprintf("%s,%s", p.Start, p.End)
}

// PoolContains reports whether ip falls within the pool's [Start, End] range.
// It compares the canonical 16-byte forms so it works for both IPv4 and IPv6.
func PoolContains(p *model.Pool, ip net.IP) bool {
	s := net.ParseIP(p.Start)
	e := net.ParseIP(p.End)
	if s == nil || e == nil || ip == nil {
		return false
	}
	si, ei, i := s.To16(), e.To16(), ip.To16()
	return bytes.Compare(i, si) >= 0 && bytes.Compare(i, ei) <= 0
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
