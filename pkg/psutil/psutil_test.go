// Package psutil
// Date: 2024/3/6 10:55
// Author: Amu
// Description:
package psutil

import "testing"

func TestGetCPUPercent(t *testing.T) {
	cpuPercent, err := GetCPUPercent()
	t.Log(cpuPercent, err)
}

func TestGetMemInfo(t *testing.T) {
	memPercent, total, used, err := GetMemInfo()
	t.Log(memPercent, total, used, err)
}

func TestGetDiskPercent(t *testing.T) {
	devices := map[string]struct{}{"/dev/vda3": {}}
	diskMap, err := GetDiskInfo(devices)
	t.Log(diskMap, err)
}

func TestGetNetworkPercent(t *testing.T) {
	netMap, err := GetNetworkIO()
	t.Log(netMap, err)
}

func TestGetDiskIO(t *testing.T) {
	diskMap, err := GetDiskIO()
	t.Log(diskMap, err)
}
