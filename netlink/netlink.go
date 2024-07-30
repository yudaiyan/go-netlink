package netlink

import (
	"crypto/rand"
	"fmt"
	"net"
)

// 获取本地所有IP
func GetAllIps() (cidrs []*net.IPNet, err error) {
	interfaces, err := net.Interfaces()
	if err != nil {
		return
	}

	for _, iface := range interfaces {
		var addrs []net.Addr
		addrs, err = iface.Addrs()
		if err != nil {
			return
		}
		for _, addr := range addrs {
			if ipnet, ok := addr.(*net.IPNet); ok && !ipnet.IP.IsLoopback() && ipnet.IP.To4() != nil {
				cidrs = append(cidrs, ipnet)
			}
		}
	}
	return
}

// 获取网卡所有IP
func GetIps(name string) (cidrs []*net.IPNet, err error) {
	var iface net.Interface
	iface, err = GetInterface(name)
	if err != nil {
		return
	}

	var addrs []net.Addr
	addrs, err = iface.Addrs()
	if err != nil {
		return
	}

	for _, addr := range addrs {
		if ipnet, ok := addr.(*net.IPNet); ok && !ipnet.IP.IsLoopback() && ipnet.IP.To4() != nil {
			cidrs = append(cidrs, ipnet)
		}
	}
	return
}

// 获取本地网卡
func GetInterface(name string) (iface net.Interface, err error) {
	interfaces, err := net.Interfaces()
	if err != nil {
		return
	}

	for _, iface = range interfaces {
		if iface.Name == name {
			return
		}
	}
	err = fmt.Errorf("没有找到网卡[%s]", name)
	return
}

// 随机生成cidr
func RandomCIDR() (cidr *net.IPNet, err error) {
	ip := make(net.IP, 4)
	rand.Read(ip)
	addr := fmt.Sprintf("%s/%d", ip.String(), 24)

	ip, cidr, err = net.ParseCIDR(addr)
	if err != nil {
		return
	}
	cidr.IP = ip.To4()
	return
}

// 给网卡添加一个随机的cidr
func RandAddrAdd(iface string) (out *net.IPNet, err error) {
	// 随机cidr
	out, err = RandomCIDR()
	if err != nil {
		return
	}

	// 获取本地所有cidr，本检测是否冲突
	var cidrs []*net.IPNet
	if cidrs, err = GetAllIps(); err != nil {
		return
	} else {
		for _, cidr := range cidrs {
			if cidr.Contains(out.IP) || out.Contains(cidr.IP) {
				err = fmt.Errorf("随机的[%s]与已有的[%s]冲突", out, cidr)
				return
			}
		}
	}

	// cidr 添加到 回环网卡 lo
	err = AddrAdd(iface, out.String())
	return
}

func noImplemention() error {
	return fmt.Errorf("no implemention")
}
