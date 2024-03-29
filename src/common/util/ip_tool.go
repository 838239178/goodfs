package util

import (
	"common/logs"
	"net"
	"os"
	"strings"
	"sync"
)

// LookupIP returns ipv4 if success else return empty string.
func LookupIP(addr string) string {
	if ip := ParseIP(addr); ip != nil {
		return ip.String()
	}
	return ""
}

// GetHost get host name from environment variables or os.Hostname or DetectServerIP
func GetHost() string {
	if host, ok := os.LookupEnv("HOST"); ok {
		return host
	}
	var err error
	if host, err := os.Hostname(); err == nil {
		return host
	}
	logs.Std().Error(err)
	return DetectServerIP()
}

func GetHostPort(port string) string {
	return net.JoinHostPort(GetHost(), port)
}

// ServerAddress join given port with outbound ip detecting by udp
func ServerAddress(port string) string {
	return net.JoinHostPort(DetectServerIP(), port)
}

func GetHostFromAddr(addr string) string {
	if _, aft, ok := strings.Cut(addr, "://"); ok {
		addr = aft
	}
	if idx := strings.IndexByte(addr, '/'); idx > 0 {
		addr = addr[:idx]
	}
	if host, _, err := net.SplitHostPort(addr); err == nil {
		return host
	}
	return addr
}

// ParseIP ipv4 > private.ipv4 > loopback.ipv4 > ipv6 > private.ipv6 > loopback.ipv6 > nil
func ParseIP(addr string) net.IP {
	host := GetHostFromAddr(addr)
	if netIP := net.ParseIP(host); netIP != nil {
		return netIP
	}
	ips, err := net.LookupIP(host)
	if err != nil {
		LogErr(err)
		return nil
	}
	var (
		loopback   net.IP
		loopbackV6 net.IP
		private    net.IP
		privateV6  net.IP
		ipv6       net.IP
		ipv4       net.IP
	)
	for _, ip := range ips {
		if ip.IsLoopback() {
			if ip.To4() != nil {
				loopback = ip
			} else if ip.To16() != nil {
				loopbackV6 = ip
			}
		} else if ip.IsPrivate() {
			if ip.To4() != nil {
				private = ip
			} else if ip.To16() != nil {
				privateV6 = ip
			}
		} else if ip.To4() != nil {
			ipv4 = ip
		} else if ip.To16() != nil {
			ipv6 = ip
		}
	}

	if ipv4 != nil {
		return ipv4
	}

	if private != nil {
		return private
	}

	if loopback != nil {
		return loopback
	}

	if ipv6 != nil {
		return ipv6
	}

	if privateV6 != nil {
		return privateV6
	}

	if loopbackV6 != nil {
		return loopbackV6
	}

	return nil
}

var localIP string
var getLocalIP = sync.Once{}

const EnvServerIP = "SERVER_IP"

func DetectServerIP() string {
	getLocalIP.Do(func() {
		defer func() {
			_ = os.Setenv(EnvServerIP, localIP)
			logs.Std().Infof("detected server ip: %s", localIP)
		}()
		if env, ok := os.LookupEnv(EnvServerIP); ok {
			localIP = env
			return
		}
		conn, err := net.Dial("udp", "223.5.5.5:53")
		if err != nil {
			LogErrWithPre("get server ip", err)
			localIP = "127.0.0.1"
			return
		}
		defer CloseAndLog(conn)
		localIP, _, _ = net.SplitHostPort(conn.LocalAddr().String())
		ip := net.ParseIP(localIP)
		if !ip.IsPrivate() && !IsPublicIP(ip) {
			localIP = "127.0.0.1"
			return
		}
	})
	return localIP
}

func IsPublicIP(ip net.IP) bool {
	if ip.IsLoopback() || ip.IsLinkLocalMulticast() || ip.IsLinkLocalUnicast() || ip.IsPrivate() {
		return false
	}
	if ip4 := ip.To4(); ip4 != nil {
		return true
	}
	return false
}
