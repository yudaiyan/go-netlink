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
