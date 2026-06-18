package netutil

import (
	"encoding/binary"
	"fmt"
	"math/big"
	"net"
	"os"
	"runtime"
	"sync"
	"time"

	"github.com/pkg/errors"
	"golang.org/x/net/icmp"
	"golang.org/x/net/ipv4"
	"golang.org/x/net/ipv6"
)

var (
	pingLock = &sync.Mutex{}
)

// IPMaskToIP converts a net.IPMask to its net.IP representation. It supports
// both 4-byte (IPv4) and 16-byte (IPv6) masks.
func IPMaskToIP(m net.IPMask) net.IP {
	if len(m) == net.IPv4len {
		return net.IPv4(m[0], m[1], m[2], m[3])
	}
	ip := make(net.IP, len(m))
	copy(ip, m)
	return ip
}

// BroadcastIP returns the broadcast address of an IPv4 network. IPv6 has no
// concept of a broadcast address, so nil is returned for IPv6 networks; callers
// must treat a nil/empty result as "not applicable".
func BroadcastIP(n *net.IPNet) net.IP {
	ipv4 := n.IP.To4()
	if ipv4 == nil {
		return nil
	}
	ret := make([]byte, len(ipv4))
	for i, v := range n.Mask {
		ret[i] = ipv4[i] | (v ^ 255)
	}
	return ret
}

func Ping(addr string) error {
	switch runtime.GOOS {
	case "darwin":
	case "linux":
	default:
		return errors.New(fmt.Sprintf("not supported on: %v", runtime.GOOS))
	}

	pingLock.Lock()
	defer pingLock.Unlock()

	target := net.ParseIP(addr)
	if target == nil {
		return errors.New(fmt.Sprintf("invalid ip: %v", addr))
	}

	var (
		network   string
		listen    string
		echoType  icmp.Type
		replyType icmp.Type
		proto     int
	)
	if target.To4() != nil {
		network, listen, proto = "udp4", "0.0.0.0", 1 // iana.ProtocolICMP
		echoType, replyType = ipv4.ICMPTypeEcho, ipv4.ICMPTypeEchoReply
	} else {
		network, listen, proto = "udp6", "::", 58 // iana.ProtocolIPv6ICMP
		echoType, replyType = ipv6.ICMPTypeEchoRequest, ipv6.ICMPTypeEchoReply
	}

	c, err := icmp.ListenPacket(network, listen)
	if err != nil {
		return err
	}
	defer c.Close()

	wm := icmp.Message{
		Type: echoType,
		Code: 0,
		Body: &icmp.Echo{
			ID: os.Getpid() & 0xffff, Seq: 1,
			Data: []byte("HELLO-R-U-THERE"),
		},
	}
	wb, err := wm.Marshal(nil)
	if err != nil {
		return err
	}
	if _, err := c.WriteTo(wb, &net.UDPAddr{IP: target}); err != nil {
		return err
	}

	rb := make([]byte, 1500)
	c.SetReadDeadline(time.Now().Add(2 * time.Second))
	n, _, err := c.ReadFrom(rb)
	if err != nil {
		return err
	}
	rm, err := icmp.ParseMessage(proto, rb[:n])
	if err != nil {
		return err
	}
	if rm.Type != replyType {
		return errors.New(fmt.Sprintf("got %+v; want echo reply", rm))
	}
	return nil
}

// IP2Uint returns the IPv4 (last 32 bits) representation of an IP as a uint32.
// It is retained for IPv4 fast paths and must not be used where IPv6 addresses
// may flow through; use IPToBigInt for family-agnostic arithmetic.
func IP2Uint(ip net.IP) uint32 {
	if len(ip) == 16 {
		return binary.BigEndian.Uint32(ip[12:16])
	}
	return binary.BigEndian.Uint32(ip)
}

func Int2IP(nn uint32) net.IP {
	ip := make(net.IP, 4)
	binary.BigEndian.PutUint32(ip, nn)
	return ip
}

// IPToBigInt converts an IP address (v4 or v6) to a big.Int built from its
// canonical 16-byte big-endian representation.
func IPToBigInt(ip net.IP) *big.Int {
	return new(big.Int).SetBytes(ip.To16())
}

// BigIntToIP converts a big.Int back to a net.IP. When v6 is true a 16-byte
// IPv6 address is produced, otherwise a 4-byte IPv4 address.
func BigIntToIP(n *big.Int, v6 bool) net.IP {
	size := net.IPv4len
	if v6 {
		size = net.IPv6len
	}
	b := n.Bytes()
	ip := make(net.IP, size)
	// Right-align the big-endian bytes, left-padding with zeros.
	if len(b) > size {
		b = b[len(b)-size:]
	}
	copy(ip[size-len(b):], b)
	return ip
}

// isV6 reports whether an IP should be treated as IPv6 (i.e. it is not a
// representable IPv4 address).
func isV6(ip net.IP) bool {
	return ip.To4() == nil
}

// NextIP returns the IP immediately following ip, preserving its address family.
func NextIP(ip net.IP) net.IP {
	n := new(big.Int).Add(IPToBigInt(ip), big.NewInt(1))
	return BigIntToIP(n, isV6(ip))
}

// PrevIP returns the IP immediately preceding ip, preserving its address family.
func PrevIP(ip net.IP) net.IP {
	n := new(big.Int).Sub(IPToBigInt(ip), big.NewInt(1))
	return BigIntToIP(n, isV6(ip))
}
