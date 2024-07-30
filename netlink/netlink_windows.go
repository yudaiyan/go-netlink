package netlink

import (
	"fmt"
	"net"
)

func noImplemention() error {
	return fmt.Errorf("no implemention")
}

// 根据网卡名，获取网卡的ip、mask、mac。其中ip、mask默认只获取第一个。
func GetLocalInterface(ifname string) (ip net.IP, mask net.IPMask, mac net.HardwareAddr, err error) {
	return nil, nil, nil, noImplemention()
}

// 根据网卡名，获取网卡的mac。
func GetMac(ifname string) (mac net.HardwareAddr, err error) {
	return nil, noImplemention()
}
