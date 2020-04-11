package cpuinfo

import (
	"fmt"
	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/disk"
	"github.com/shirou/gopsutil/host"
	"github.com/shirou/gopsutil/net"
)

func GetCpuInfo()  {
	logicalcounts, err := cpu.Counts(true)
	phycicalcounts, err := cpu.Counts(false)
	counters, err := net.IOCounters(false)
	time, err := host.BootTime()
	stat, err := host.Info()
	//interfaces, err := net.Interfaces()
	//for _,v:=range interfaces{
	//	fmt.Println(v)
	//}
	usage, err := disk.Usage("/")
	partitions, err := disk.Partitions(true)
	if err != nil {
		fmt.Println(err)
	}
	info, err := cpu.Info()
	fmt.Println(info)
	fmt.Println("cpu逻辑核数量",logicalcounts)
	fmt.Println("cpu物理核数量",phycicalcounts)
	fmt.Println("开机时间物理核数量",time)
	fmt.Println("机器信息",stat)
	fmt.Println("网络信息",counters)
	//fmt.Println("网卡",interfaces)
	fmt.Println("分区/的信息",usage)
	fmt.Println("分区/信息",partitions)
}
