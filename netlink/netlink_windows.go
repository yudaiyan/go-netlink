package netlink

import (
	"net"

	"github.com/yudaiyan/gonetsh/netsh"
	"k8s.io/utils/exec"
)

// 根据网卡名，获取网卡的ip、mask、mac。其中ip、mask默认只获取第一个。
func GetLocalInterface(ifname string) (ip net.IP, mask net.IPMask, mac net.HardwareAddr, err error) {
	return nil, nil, nil, noImplemention()
}

// 根据网卡名，获取网卡的mac。
func GetMac(ifname string) (mac net.HardwareAddr, err error) {
	return nil, noImplemention()
}

// 添加网卡ip
func AddrAdd(name string, addr string) error {
	h := netsh.New(exec.New())
	return h.SetIPAddress(name, addr)
}

// 删除网卡ip
func AddrDel(name string, addr string) error {
	return noImplemention()
}

// 清空lo上的其他IP
func LoAddrClear() error {
	return noImplemention()
}

// 修改网卡mac
func LinkSetHardwareAddr(name string, mac net.HardwareAddr) error {
	return noImplemention()
}

// 启用网卡
func LinkSetUp(name string) error {
	return noImplemention()
}

// 删除网卡
func LinkDel(name string) error {
	return noImplemention()
}
