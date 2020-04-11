package main

import (
	"fmt"
	"net"
)

func main() {
	ip:=GetServerIP()
	fmt.Println(ip)
}
// GetServerIP 获取服务器真实ip
func GetServerIP() (ip string) {
	addrSlice, err := net.InterfaceAddrs()
	if nil != err {
		fmt.Println("get ip err", err)
		return "127.0.0.1"
	}
	for _, addr := range addrSlice {
		// 检查ip地址判断是否回环地址
		if ipnet, ok := addr.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if nil != ipnet.IP.To4() {
				ip := ipnet.IP.String()
				return ip
			}
		}
	}
	return ""
}
