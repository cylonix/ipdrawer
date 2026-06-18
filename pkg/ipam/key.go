package ipam

import (
	"fmt"
	"net"
	"strings"

	"github.com/pkg/errors"
)

const (
	lockPrefix = "lock:"
	globalLock = "global"

	ipDetails      = "%s:ip:%s:details"
	ipTempReserved = "%s:ip:%s:temporary_reserved"
	networkList    = "%s:network:list"
	networkDetails = "%s:network:%s:details"
	poolDetails    = "%s:pool:%s,%s:details"
	poolUsedIPZSet = "%s:pool:%s,%s:used_ip_zset"
	ipPoolUuid     = "%s:ip:%s,%s,%s,%s:uuid"
)

func makeGlobalLock() string {
	return lockPrefix + globalLock
}

func makeIPDetailsKey(namespace string, ip net.IP) string {
	return fmt.Sprintf(ipDetails, namespace, ip.String())
}

func makeIPListPattern(namespace string) string {
	return fmt.Sprintf(ipDetails, namespace, "*")
}

// parseIPDetailsKey parses a "<namespace>:ip:<ip>:<suffix>" key (suffix is
// "details" or "temporary_reserved"). The IP may be an IPv6 literal containing
// colons, so we anchor on the fixed leading tokens ("<namespace>", "ip"), drop
// the trailing suffix token, and rejoin the middle as the address.
func parseIPDetailsKey(key string) (net.IP, error) {
	d := strings.Split(key, ":")
	if len(d) < 4 || d[1] != "ip" {
		return nil, errors.New("Not matched format")
	}
	ip := strings.Join(d[2:len(d)-1], ":")
	addr := net.ParseIP(ip)
	if addr == nil {
		return nil, errors.New("Failed parse IP")
	}
	return addr, nil
}

func makeIPTempReserved(namespace string, ip net.IP) string {
	return fmt.Sprintf(ipTempReserved, namespace, ip.String())
}

func makeTempReservedIPListPattern(namespace string) string {
	return fmt.Sprintf(ipTempReserved, namespace, "*")
}

func parseTempReservedIPKey(key string) (net.IP, error) {
	return parseIPDetailsKey(key)
}

func makeIPPoolUuid(namespace string, s, e, ip net.IP, uuid string) string {
	return fmt.Sprintf(ipPoolUuid, namespace, s.String(), e.String(), ip.String(), uuid)
}
func makeIPUuidKey(namespace, s, e, ip string) string {
	return fmt.Sprintf(ipPoolUuid, namespace, s, e, ip, "*")
}
func MakeUuidIPKey(namespace, s, e, uuid string) string {
	return fmt.Sprintf(ipPoolUuid, namespace, s, e, "*", uuid)
}
func makePoolUuidIPListPattern(namespace string, s, e net.IP, uuid string) string {
	return fmt.Sprintf(ipPoolUuid, namespace, s.String(), e.String(), "*", uuid)
}

// parsePoolUuidIPKey parses a "<namespace>:ip:<start>,<end>,<ip>,<uuid>:uuid"
// key. The start/end/ip components may be IPv6 literals containing colons, so we
// anchor on the leading ("<namespace>", "ip") and trailing ("uuid") tokens and
// rejoin the middle before splitting it on commas (IPv6 uses colons, not commas,
// so the comma split is unambiguous).
func parsePoolUuidIPKey(key string) (net.IP, net.IP, net.IP, string, error) {
	d := strings.Split(key, ":")
	if len(d) < 4 || d[1] != "ip" || d[len(d)-1] != "uuid" {
		return nil, nil, nil, "", errors.New("Not matched format")
	}
	se := strings.Split(strings.Join(d[2:len(d)-1], ":"), ",")
	if len(se) != 4 {
		return nil, nil, nil, "", errors.New("Not matched format")
	}
	s := net.ParseIP(se[0])
	if s == nil {
		return nil, nil, nil, "", errors.New("Failed parse IP")
	}
	e := net.ParseIP(se[1])
	if e == nil {
		return nil, nil, nil, "", errors.New("Failed parse IP")
	}
	ip := net.ParseIP(se[2])
	if ip == nil {
		return nil, nil, nil, "", errors.New("Failed parse IP")
	}
	return s, e, ip, se[3], nil
}

func makeNetworkListKey(namespace string) string {
	return fmt.Sprintf(networkList, namespace)
}

func makeNetworkDetailsKey(namespace string, ip *net.IPNet) string {
	return fmt.Sprintf(networkDetails, namespace, ip.String())
}

func makeNetworkPoolKey(namespace string, ip *net.IPNet) string {
	return makeNetworkDetailsKey(namespace, ip) + ":pools"
}

func makePoolDetailsKey(namespace string, s, e net.IP) string {
	return fmt.Sprintf(poolDetails, namespace, s.String(), e.String())
}

func makePoolListPattern(namespace string) string {
	return fmt.Sprintf(strings.Replace(poolDetails, "%s,%s", "*", 1), namespace)
}

// ParsePoolDetailsKey parses a "<namespace>:pool:<start>,<end>:details" key. The
// start/end components may be IPv6 literals containing colons, so we anchor on
// the leading ("<namespace>", "pool") and trailing ("details") tokens and rejoin
// the middle before splitting it on the comma.
func ParsePoolDetailsKey(key string) (net.IP, net.IP, error) {
	d := strings.Split(key, ":")
	if len(d) < 4 || d[1] != "pool" || d[len(d)-1] != "details" {
		return nil, nil, errors.New("Not matched format")
	}
	se := strings.Split(strings.Join(d[2:len(d)-1], ":"), ",")
	if len(se) != 2 {
		return nil, nil, errors.New("Not matched format")
	}
	s := net.ParseIP(se[0])
	if s == nil {
		return nil, nil, errors.New("Failed parse IP")
	}
	e := net.ParseIP(se[1])
	if e == nil {
		return nil, nil, errors.New("Failed parse IP")
	}
	return s, e, nil
}

func makePoolUsedIPZSet(namespace string, s, e net.IP) string {
	return fmt.Sprintf(poolUsedIPZSet, namespace, s.String(), e.String())
}
