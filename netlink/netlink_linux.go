package netlink

import (
	"fmt"
	"log"
	"net"

	"github.com/vishvananda/netlink"
)

// 根据网卡名，获取网卡的ip、mask、mac。其中ip、mask默认只获取第一个。
func GetLocalInterface(ifname string) (ip net.IP, mask net.IPMask, mac net.HardwareAddr, err error) {
	defer func() {
		ones, _ := mask.Size()
		log.Printf("获取本地网卡%s: [cidr=%s/%d, mac=%s]", ifname, ip, ones, mac)
	}()

	link, err := netlink.LinkByName(ifname)
	if err != nil {
		return
	}
	addrs, err := netlink.AddrList(link, netlink.FAMILY_V4)
	if err != nil {
		return
	}
	if len(addrs) == 0 {
		err = fmt.Errorf("网卡 %s 没有配置ip", ifname)
		return
	}
	return addrs[0].IPNet.IP, addrs[0].IPNet.Mask, link.Attrs().HardwareAddr, nil
}

// 根据网卡名，获取网卡的mac。
func GetMac(ifname string) (mac net.HardwareAddr, err error) {
	defer func() {
		log.Printf("获取本地网卡%s: [mac=%s]", ifname, mac)
	}()

	link, err := netlink.LinkByName(ifname)
	if err != nil {
		return
	}
	return link.Attrs().HardwareAddr, nil
}

// 添加网卡ip
func AddrAdd(name string, addr string) error {
	addr2, err := netlink.ParseAddr(addr)
	if err != nil {
		return err
	}
	link, err := netlink.LinkByName(name)
	if err != nil {
		return err
	}
	return netlink.AddrAdd(link, addr2)
}

// 删除网卡ip
func AddrDel(name string, addr string) error {
	addr2, err := netlink.ParseAddr(addr)
	if err != nil {
		return err
	}
	link, err := netlink.LinkByName(name)
	if err != nil {
		return err
	}
	return netlink.AddrDel(link, addr2)
}

// 清空lo上的其他IP
func LoAddrClear() error {
	cidrs, err := GetIps("lo")
	if err != nil {
		return err
	}

	for _, cidr := range cidrs {
		if cidr.IP.IsLoopback() {
			continue
		}
		err = AddrDel("lo", cidr.String())
		if err != nil {
			return err
		}
	}
	return nil
}

// 修改网卡mac
func LinkSetHardwareAddr(name string, mac net.HardwareAddr) error {
	link, err := netlink.LinkByName(name)
	if err != nil {
		return err
	}
	return netlink.LinkSetHardwareAddr(link, mac)
}

// 启用网卡
func LinkSetUp(name string) error {
	link, err := netlink.LinkByName(name)
	if err != nil {
		return err
	}
	return netlink.LinkSetUp(link)
}

// 删除网卡
func LinkDel(name string) error {
	link, err := netlink.LinkByName(name)
	if err != nil {
		return err
	}
	return netlink.LinkDel(link)
}
