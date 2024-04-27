// Package psutil
// Date: 2024/3/6 10:55
// Author: Amu
// Description:
package psutil

import (
	"fmt"
	"log/slog"
	"strings"
	"time"

	"github.com/shirou/gopsutil/v3/cpu"
	"github.com/shirou/gopsutil/v3/disk"
	"github.com/shirou/gopsutil/v3/host"
	"github.com/shirou/gopsutil/v3/mem"
	"github.com/shirou/gopsutil/v3/net"
)

type DiskInfo struct {
	Total   uint64
	Percent float64
	Used    uint64
}

type NetIO struct {
	Recv uint64 `json:"recv"`
	Send uint64 `json:"send"`
}

type DiskIO struct {
	Read  uint64 `json:"read"`
	Write uint64 `json:"write"`
}

type SystemInfo struct {
	Uptime          string
	Hostname        string
	Os              string
	Platform        string
	PlatformVersion string
	KernelVersion   string
	KernelArch      string
}

func GetMemInfo() (float64, uint64, uint64, error) {
	v, err := mem.VirtualMemory()
	if err != nil {
		return 0.0, 0, 0, err
	}
	return v.UsedPercent, v.Total, v.Used, nil
}

func GetCPUPercent() (float64, error) {
	totalPercent, err := cpu.Percent(3*time.Second, false)
	if err != nil {
		return 0.0, err
	}
	return totalPercent[0], nil
}

func GetDiskInfo(devices map[string]struct{}) (map[string]DiskInfo, error) {
	diskMap := make(map[string]DiskInfo)
	infos, _ := disk.Partitions(false)
	for _, info := range infos {
		usedInfo, _ := disk.Usage(info.Mountpoint)
		if usedInfo == nil {
			continue
		}
		slog.Info("---------------------", "info", info)
		slog.Info("---------------------", "devices", devices)
		dev := strings.Split(info.Device, "/")[len(strings.Split(info.Device, "/"))-1]
		if _, ok := devices[dev]; !ok {
			continue
		}
		if _, ok := diskMap[dev]; !ok {
			diskMap[dev] = DiskInfo{
				Total:   usedInfo.Total,
				Percent: usedInfo.UsedPercent,
				Used:    usedInfo.Used,
			}
		}
	}
	return diskMap, nil
}

func GetDiskIO(devices map[string]struct{}) (map[string]DiskIO, error) {
	diskMap := make(map[string]DiskIO)
	// 实现磁盘IO的获取逻辑
	stat, err := disk.IOCounters()
	if err != nil {
		return diskMap, err
	}
	slog.Error("========> ", "devices", devices)
	// 将获取到的磁盘IO信息填充到map中
	for k, v := range stat {
		slog.Error("=========", k, v)
		if _, ok := devices[k]; !ok {
			continue
		}
		diskMap[k] = DiskIO{
			Read:  v.ReadCount,
			Write: v.WriteCount,
		}
	}
	return diskMap, nil
}

func GetNetworkIO(eth map[string]struct{}) (map[string]NetIO, error) {
	netMap := make(map[string]NetIO)
	IOCountersStat, err := net.IOCounters(true)
	if err != nil {
		return netMap, err
	}
	for _, stat := range IOCountersStat {
		if _, ok := eth[stat.Name]; !ok {
			continue
		}
		netMap[stat.Name] = NetIO{
			Recv: stat.BytesRecv,
			Send: stat.BytesSent,
		}
	}
	return netMap, nil
}

// GetSystemInfo 获取系统信息
func GetSystemInfo() (*SystemInfo, error) {
	info, err := host.Info()
	if err != nil {
		return nil, err
	}
	systemInfo := &SystemInfo{
		Uptime:          time.Unix(int64(info.BootTime), 0).Local().Format("2006-01-02 15:04:05"),
		Hostname:        info.HostID,
		Os:              info.OS,
		Platform:        info.Platform,
		PlatformVersion: info.PlatformVersion,
		KernelVersion:   info.KernelVersion,
		KernelArch:      info.KernelArch,
	}
	return systemInfo, nil
}

// 字节单位转换
func fmtByte(size int64) string {
	if size < 1024 {
		return fmt.Sprintf("%.2fB", float64(size)/float64(1))
	} else if size < 1024*1024 {
		return fmt.Sprintf("%.2fKB", float64(size)/float64(1024))
	} else if size < 1024*1024*1024 {
		return fmt.Sprintf("%.2fMB", float64(size)/float64(1024*1024))
	} else if size < 1024*1024*1024*1024 {
		return fmt.Sprintf("%.2fGB", float64(size)/float64(1024*1024*1024))
	} else if size < 1024*1024*1024*1024*1024 {
		return fmt.Sprintf("%.2fTB", float64(size)/float64(1024*1024*1024*1024))
	} else {
		return fmt.Sprintf("%.2fEB", float64(size)/float64(1024*1024*1024*1024*2014))
	}
}
