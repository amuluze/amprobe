// Package psutil
// Date: 2024/06/10 18:58:51
// Author: Amu
// Description:
package psutil

import (
	"context"
	"log/slog"
	"strings"
	"time"

	"github.com/shirou/gopsutil/v3/common"
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
	ctx := context.WithValue(context.Background(), common.EnvKey, common.EnvMap{})
	v, err := mem.VirtualMemoryWithContext(ctx)
	if err != nil {
		return 0.0, 0, 0, err
	}
	return v.UsedPercent, v.Total, v.Used, nil
}

func GetCPUPercent() (float64, error) {
	ctx := context.WithValue(context.Background(), common.EnvKey, common.EnvMap{})
	totalPercent, err := cpu.PercentWithContext(ctx, time.Second, false)
	if err != nil {
		return 0.0, err
	}
	return totalPercent[0], nil
}

func GetDiskInfo(devices map[string]struct{}) (map[string]DiskInfo, error) {
	ctx := context.WithValue(context.Background(), common.EnvKey, common.EnvMap{})
	diskMap := make(map[string]DiskInfo)
	infos, _ := disk.PartitionsWithContext(ctx, false)
	for _, info := range infos {
		usedInfo, err := disk.UsageWithContext(ctx, info.Mountpoint)
		if usedInfo == nil {
			slog.Error("get disk usage err", "err", err)
			continue
		}
		dev := strings.Split(info.Device, "/")[len(strings.Split(info.Device, "/"))-1]
		if _, ok := devices[dev]; !ok {
			continue
		}
		diskMap[dev] = DiskInfo{
			Total:   usedInfo.Total,
			Percent: usedInfo.UsedPercent,
			Used:    usedInfo.Used,
		}
	}
	return diskMap, nil
}

func GetAllDiskInfo() (map[string]DiskInfo, error) {
	ctx := context.WithValue(context.Background(), common.EnvKey, common.EnvMap{})
	diskMap := make(map[string]DiskInfo)
	infos, _ := disk.PartitionsWithContext(ctx, false) // false表示只获取物理分区
	for _, info := range infos {
		// 过滤掉 loop 设备，这些通常是 Linux 系统中 snap 包挂载点，不是真正的物理磁盘
		if strings.Contains(info.Device, "/loop") {
			continue
		}

		usedInfo, err := disk.UsageWithContext(ctx, info.Mountpoint)
		if usedInfo == nil {
			slog.Error("get disk usage err", "err", err)
			continue
		}
		dev := strings.Split(info.Device, "/")[len(strings.Split(info.Device, "/"))-1]
		diskMap[dev] = DiskInfo{
			Total:   usedInfo.Total,
			Percent: usedInfo.UsedPercent,
			Used:    usedInfo.Used,
		}
	}
	return diskMap, nil
}

func GetDiskIO(devices map[string]struct{}) (map[string]DiskIO, error) {
	ctx := context.WithValue(context.Background(), common.EnvKey, common.EnvMap{})
	diskMap := make(map[string]DiskIO)
	// 实现磁盘IO的获取逻辑
	var names []string
	for k := range devices {
		names = append(names, k)
	}
	stat, err := disk.IOCountersWithContext(ctx, names...)
	if err != nil {
		slog.Error("get disk io err", "err", err)
		return diskMap, err
	}
	// 将获取到的磁盘IO信息填充到map中
	for k, v := range stat {
		if _, ok := devices[k]; !ok {
			continue
		}
		diskMap[k] = DiskIO{
			Read:  v.ReadBytes,
			Write: v.WriteBytes,
		}
	}
	return diskMap, nil
}

func GetNetworkIO(eth map[string]struct{}) (map[string]NetIO, error) {
	ctx := context.WithValue(context.Background(), common.EnvKey, common.EnvMap{})

	netMap := make(map[string]NetIO)
	IOCountersStat, err := net.IOCountersWithContext(ctx, true)
	if err != nil {
		slog.Error("get network io err", "err", err)
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
	ctx := context.WithValue(context.Background(), common.EnvKey, common.EnvMap{})

	info, err := host.InfoWithContext(ctx)
	if err != nil {
		return nil, err
	}
	systemInfo := &SystemInfo{
		Uptime:          time.Unix(int64(info.BootTime), 0).Local().Format("2006-01-02 15:04:05"),
		Hostname:        info.Hostname,
		Os:              info.OS,
		Platform:        info.Platform,
		PlatformVersion: info.PlatformVersion,
		KernelVersion:   info.KernelVersion,
		KernelArch:      info.KernelArch,
	}
	return systemInfo, nil
}
