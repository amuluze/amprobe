// Package schema
// Date: 2024/3/6 13:20
// Author: Amu
// Description:
package schema

type Usage struct {
	Timestamp int64   `json:"timestamp"`
	Value     float64 `json:"value"`
}

type NetIO struct {
	Timestamp int64   `json:"timestamp"`
	BytesSent float64 `json:"bytes_sent"`
	BytesRecv float64 `json:"bytes_recv"`
}

type DiskIO struct {
	Timestamp int64   `json:"timestamp"`
	IORead    float64 `json:"io_read"`
	IOWrite   float64 `json:"io_write"`
}

type HostArgs struct{}

type HostInfoReply struct {
	Timestamp       int64  `json:"timestamp"`
	Uptime          string `json:"uptime"`
	Hostname        string `json:"hostname"`
	OS              string `json:"os"`
	Platform        string `json:"platform"`
	PlatformVersion string `json:"platform_version"`
	KernelVersion   string `json:"kernel_version"`
	KernelArch      string `json:"kernel_arch"`
}

type CPUArgs struct{}

type CPUInfoReply struct {
	Percent float64 `json:"percent"`
}

type CPUUsageArgs struct {
	StartTime int64 `query:"start_time"`
	EndTime   int64 `query:"end_time"`
}

type CPUUsageReply struct {
	Data []Usage `json:"data"`
}

type MemoryArgs struct{}

type MemoryInfoReply struct {
	Percent float64 `json:"percent"`
	Total   float64 `json:"total"`
	Used    float64 `json:"used"`
}

type MemoryUsageArgs struct {
	StartTime int64 `query:"start_time"`
	EndTime   int64 `query:"end_time"`
}

type MemoryUsageReply struct {
	Data []Usage `json:"data"`
}

type DiskArgs struct{}

type DiskInfo struct {
	Device  string  `json:"device"`
	Percent float64 `json:"percent"`
	Total   float64 `json:"total"`
	Used    float64 `json:"used"`
}

type DiskUsage struct {
	Device string   `json:"device"`
	Data   []DiskIO `json:"data"`
}

type NetUsage struct {
	Ethernet string  `json:"ethernet"`
	Data     []NetIO `json:"data"`
}

type DiskInfoReply struct {
	Info []DiskInfo `json:"info"`
}

type DiskUsageArgs struct {
	StartTime int64 `query:"start_time"`
	EndTime   int64 `query:"end_time"`
}

type DiskUsageReply struct {
	Usage []DiskUsage `json:"usage"`
}

type NetworkUsageArgs struct {
	StartTime int64 `query:"start_time"`
	EndTime   int64 `query:"end_time"`
}

type NetworkUsageReply struct {
	Usage []NetUsage `json:"usage"`
}

type FilesSearchArgs struct {
	Path string
}
type FileInfo struct {
	Name    string
	Size    int64
	Mode    string
	ModTime int64
	IsDir   bool
}
type FilesSearchReply struct {
	Files []FileInfo `json:"files"`
}

type FileUploadArgs struct{}

type FIleUploadReply struct{}

type FileDownloadArgs struct{}

type FileDownloadReply struct{}

type FileDeleteArgs struct{}

type FileDeleteReply struct{}

type GetDNSSettingsArgs struct{}

type GetDNSSettingsReply struct {
	DNS []string `json:"dns"`
}

type SetDNSSettingsArgs struct {
	DNS []string `json:"dns"`
}

type SetDNSSettingsReply struct{}

type GetSystemTimeArgs struct {
	SystemTime string `json:"system_time"`
}

type GetSystemTimeReply struct{}

type SetSystemTimeArgs struct {
	SystemTime string `json:"system_time"`
}

type SetSystemTimeReply struct{}

type GetSystemTimezoneArgs struct{}

type GetSystemTimezoneReply struct {
	SystemTimeZone string `json:"system_timezone"`
}

type SetSystemTimezoneArgs struct {
	SystemTimeZone string `json:"system_timezone"`
}

type SetSystemTimezoneReply struct{}

type RebootArgs struct{}

type RebootReply struct{}

type ShutdownArgs struct{}

type ShutdownReply struct{}
