package model

import (
	"net"
	"testing"

	"github.com/hatena/ipdrawer/gen/go/model"
)

func TestPoolContains(t *testing.T) {
	cases := []struct {
		name string
		pool *model.Pool
		ip   string
		want bool
	}{
		{"v4 in range", &model.Pool{Start: "10.0.0.1", End: "10.0.0.254"}, "10.0.0.100", true},
		{"v4 below start", &model.Pool{Start: "10.0.0.10", End: "10.0.0.254"}, "10.0.0.1", false},
		{"v4 above end", &model.Pool{Start: "10.0.0.1", End: "10.0.0.10"}, "10.0.0.11", false},
		{"v4 boundary start", &model.Pool{Start: "10.0.0.1", End: "10.0.0.10"}, "10.0.0.1", true},
		{"v4 boundary end", &model.Pool{Start: "10.0.0.1", End: "10.0.0.10"}, "10.0.0.10", true},
		{"v6 in range", &model.Pool{Start: "2001:db8::1", End: "2001:db8::ff"}, "2001:db8::80", true},
		{"v6 below start", &model.Pool{Start: "2001:db8::10", End: "2001:db8::ff"}, "2001:db8::1", false},
		{"v6 above end", &model.Pool{Start: "2001:db8::1", End: "2001:db8::ff"}, "2001:db8::100", false},
		{"v6 boundary end", &model.Pool{Start: "2001:db8::1", End: "2001:db8::ff"}, "2001:db8::ff", true},
	}
	for _, c := range cases {
		got := PoolContains(c.pool, net.ParseIP(c.ip))
		if got != c.want {
			t.Errorf("%s: PoolContains(%s) = %v, want %v", c.name, c.ip, got, c.want)
		}
	}
}
