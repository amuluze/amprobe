// Package service
// Date: 2024/3/6 12:53
// Author: Amu
// Description:
package service

import (
	"context"

	"amprobe/service/host/repository"
	"amprobe/service/schema"

	"github.com/google/wire"
)

var HostServiceSet = wire.NewSet(NewHostService, wire.Bind(new(IHostService), new(*HostService)))

var _ IHostService = (*HostService)(nil)

type IHostService interface {
	HostInfo(ctx context.Context) (schema.HostInfoReply, error)
	CPUInfo(ctx context.Context) (schema.CPUInfoReply, error)
	CPUUsage(ctx context.Context, args schema.CPUUsageArgs) (schema.CPUUsageReply, error)
	MemInfo(ctx context.Context) (schema.MemoryInfoReply, error)
	MemUsage(ctx context.Context, args schema.MemoryUsageArgs) (schema.MemoryUsageReply, error)
	DiskInfo(ctx context.Context) (schema.DiskInfoReply, error)
	DiskUsage(ctx context.Context, args schema.DiskUsageArgs) (schema.DiskUsageReply, error)
	NetUsage(ctx context.Context, args schema.NetworkUsageArgs) (schema.NetworkUsageReply, error)

	FilesSearch(ctx context.Context, args schema.FilesSearchArgs) (schema.FilesSearchReply, error)
	FileUpload(ctx context.Context, args schema.FileUploadArgs) error
	FileDownload(ctx context.Context, args schema.FileDownloadArgs) (schema.FileDownloadReply, error)
	FileDelete(ctx context.Context, args schema.FileDeleteArgs) error
	FileCreate(ctx context.Context, args schema.FileCreateArgs) error
	FolderCreate(ctx context.Context, args schema.FolderCreateArgs) error
	GetDNSSettings(ctx context.Context, args schema.GetDNSSettingsArgs) (schema.GetDNSSettingsReply, error)
	SetDNSSettings(ctx context.Context, args schema.SetDNSSettingsArgs) error
	GetSystemTime(ctx context.Context, args schema.GetSystemTimeArgs) (schema.GetSystemTimeReply, error)
	SetSystemTime(ctx context.Context, args schema.SetSystemTimeArgs) error
	GetSystemTimezoneList(ctx context.Context, args schema.GetSystemTimezoneListArgs) (schema.GetSystemTimezoneListReply, error)
	GetSystemTimezone(ctx context.Context, args schema.GetSystemTimezoneArgs) (schema.GetSystemTimezoneReply, error)
	SetSystemTimezone(ctx context.Context, args schema.SetSystemTimezoneArgs) error
	Reboot(ctx context.Context, args schema.RebootArgs) error
	Shutdown(ctx context.Context, args schema.ShutdownArgs) error
}

type HostService struct {
	HostRepo repository.IHostRepo
}

func NewHostService(hostRepo repository.IHostRepo) *HostService {
	return &HostService{HostRepo: hostRepo}
}

func (h *HostService) HostInfo(ctx context.Context) (schema.HostInfoReply, error) {
	info, err := h.HostRepo.HostInfo(ctx)
	if err != nil {
		return schema.HostInfoReply{}, err
	}
	return schema.HostInfoReply{
		Timestamp:       info.Timestamp.Unix(),
		Uptime:          info.Uptime,
		Hostname:        info.Hostname,
		OS:              info.Os,
		Platform:        info.Platform,
		PlatformVersion: info.PlatformVersion,
		KernelVersion:   info.KernelVersion,
		KernelArch:      info.KernelArch,
	}, err
}

func (h *HostService) CPUInfo(ctx context.Context) (schema.CPUInfoReply, error) {
	cpuInfo, err := h.HostRepo.CPUInfo(ctx)
	if err != nil {
		return schema.CPUInfoReply{}, err
	}
	return schema.CPUInfoReply{Percent: cpuInfo.CPUPercent}, nil
}

func (h *HostService) CPUUsage(ctx context.Context, args schema.CPUUsageArgs) (schema.CPUUsageReply, error) {
	mHosts, err := h.HostRepo.CPUUsage(ctx, args)
	if err != nil {
		return schema.CPUUsageReply{}, err
	}
	var list []schema.Usage
	for _, item := range mHosts {
		list = append(list, schema.Usage{
			Timestamp: item.Timestamp.Unix(),
			Value:     item.CPUPercent,
		})
	}
	return schema.CPUUsageReply{Data: list}, nil
}

func (h *HostService) MemInfo(ctx context.Context) (schema.MemoryInfoReply, error) {
	memInfo, err := h.HostRepo.MemInfo(ctx)
	if err != nil {
		return schema.MemoryInfoReply{}, err
	}
	return schema.MemoryInfoReply{Percent: memInfo.MemPercent, Total: memInfo.MemTotal, Used: memInfo.MemUsed}, nil
}

func (h *HostService) MemUsage(ctx context.Context, args schema.MemoryUsageArgs) (schema.MemoryUsageReply, error) {
	memInfos, err := h.HostRepo.MemUsage(ctx, args)
	if err != nil {
		return schema.MemoryUsageReply{}, err
	}
	var list []schema.Usage
	for _, item := range memInfos {
		list = append(list, schema.Usage{
			Timestamp: item.Timestamp.Unix(),
			Value:     item.MemPercent,
		})
	}
	return schema.MemoryUsageReply{Data: list}, nil
}

func (h *HostService) DiskInfo(ctx context.Context) (schema.DiskInfoReply, error) {
	diskInfos, err := h.HostRepo.DiskInfo(ctx)
	if err != nil {
		return schema.DiskInfoReply{}, err
	}
	var list []schema.DiskInfo
	for _, di := range diskInfos {
		list = append(list, schema.DiskInfo{
			Device:  di.Device,
			Percent: di.DiskPercent,
			Total:   di.DiskTotal,
			Used:    di.DiskUsed,
		})
	}

	return schema.DiskInfoReply{Info: list}, err
}

func (h *HostService) DiskUsage(ctx context.Context, args schema.DiskUsageArgs) (schema.DiskUsageReply, error) {
	diskInfos, err := h.HostRepo.DiskUsage(ctx, args)
	if err != nil {
		return schema.DiskUsageReply{}, err
	}
	var reply schema.DiskUsageReply
	usageMap := make(map[string][]schema.DiskIO)

	for _, item := range diskInfos {
		if _, ok := usageMap[item.Device]; !ok {
			usageMap[item.Device] = []schema.DiskIO{{
				Timestamp: item.Timestamp.Unix(),
				IORead:    item.DiskRead,
				IOWrite:   item.DiskWrite,
			}}
		} else {
			usageMap[item.Device] = append(usageMap[item.Device], schema.DiskIO{
				Timestamp: item.Timestamp.Unix(),
				IORead:    item.DiskRead,
				IOWrite:   item.DiskWrite,
			})
		}
	}
	for key, item := range usageMap {
		reply.Usage = append(reply.Usage, schema.DiskUsage{
			Device: key,
			Data:   item,
		})
	}
	return reply, nil
}

func (h *HostService) NetUsage(ctx context.Context, args schema.NetworkUsageArgs) (schema.NetworkUsageReply, error) {
	netInfos, err := h.HostRepo.NetUsage(ctx, args)
	if err != nil {
		return schema.NetworkUsageReply{}, err
	}

	var reply schema.NetworkUsageReply
	usageMap := make(map[string][]schema.NetIO)

	for _, item := range netInfos {
		if _, ok := usageMap[item.Ethernet]; !ok {
			usageMap[item.Ethernet] = []schema.NetIO{{
				Timestamp: item.Timestamp.Unix(),
				BytesSent: item.NetSend,
				BytesRecv: item.NetRecv,
			}}
		} else {
			usageMap[item.Ethernet] = append(usageMap[item.Ethernet], schema.NetIO{
				Timestamp: item.Timestamp.Unix(),
				BytesSent: item.NetSend,
				BytesRecv: item.NetRecv,
			})
		}
	}
	for key, item := range usageMap {
		reply.Usage = append(reply.Usage, schema.NetUsage{
			Ethernet: key,
			Data:     item,
		})
	}
	return reply, nil
}

func (h *HostService) FilesSearch(ctx context.Context, args schema.FilesSearchArgs) (schema.FilesSearchReply, error) {
	result := schema.FilesSearchReply{}
	search, err := h.HostRepo.FilesSearch(ctx, args)
	if err != nil {
		return result, err
	}
	var list []schema.FileInfo
	for _, info := range search.Files {
		list = append(list, schema.FileInfo{
			Name:    info.Name,
			Size:    info.Size,
			Mode:    info.Mode,
			ModTime: info.ModTime,
			IsDir:   info.IsDir,
		})
	}
	result.Files = list
	return result, nil
}

func (h *HostService) FileUpload(ctx context.Context, args schema.FileUploadArgs) error {
	return h.HostRepo.FileUpload(ctx, args)
}

func (h *HostService) FileDownload(ctx context.Context, args schema.FileDownloadArgs) (schema.FileDownloadReply, error) {
	result := schema.FileDownloadReply{}
	download, err := h.HostRepo.FileDownload(ctx, args)
	if err != nil {
		return result, err
	}
	result.Filepath = download.Filepath
	return result, err
}

func (h *HostService) FileDelete(ctx context.Context, args schema.FileDeleteArgs) error {
	return h.HostRepo.FileDelete(ctx, args)
}

func (h *HostService) FileCreate(ctx context.Context, args schema.FileCreateArgs) error {
	return h.HostRepo.FileCreate(ctx, args)
}

func (h *HostService) FolderCreate(ctx context.Context, args schema.FolderCreateArgs) error {
	return h.HostRepo.FolderCreate(ctx, args)
}

func (h *HostService) GetDNSSettings(ctx context.Context, args schema.GetDNSSettingsArgs) (schema.GetDNSSettingsReply, error) {
	result := schema.GetDNSSettingsReply{}
	settings, err := h.HostRepo.GetDNSSettings(ctx, args)
	if err != nil {
		return result, err
	}
	result.DNS = settings.DNS
	return result, nil
}

func (h *HostService) SetDNSSettings(ctx context.Context, args schema.SetDNSSettingsArgs) error {
	return h.HostRepo.SetDNSSettings(ctx, args)
}

func (h *HostService) GetSystemTime(ctx context.Context, args schema.GetSystemTimeArgs) (schema.GetSystemTimeReply, error) {
	result := schema.GetSystemTimeReply{}
	reply, err := h.HostRepo.GetSystemTime(ctx, args)
	if err != nil {
		return result, err
	}
	result.SystemTime = reply.SystemTime
	return result, nil
}

func (h *HostService) SetSystemTime(ctx context.Context, args schema.SetSystemTimeArgs) error {
	return h.HostRepo.SetSystemTime(ctx, args)
}

func (h *HostService) GetSystemTimezoneList(ctx context.Context, args schema.GetSystemTimezoneListArgs) (schema.GetSystemTimezoneListReply, error) {
	result := schema.GetSystemTimezoneListReply{}
	reply, err := h.HostRepo.GetSystemTimezoneList(ctx, args)
	if err != nil {
		return result, err
	}
	result.SystemTimeZoneList = reply.SystemTimeZoneList
	return result, nil
}

func (h *HostService) GetSystemTimezone(ctx context.Context, args schema.GetSystemTimezoneArgs) (schema.GetSystemTimezoneReply, error) {
	result := schema.GetSystemTimezoneReply{}
	reply, err := h.HostRepo.GetSystemTimezone(ctx, args)
	if err != nil {
		return result, err
	}
	result.SystemTimeZone = reply.SystemTimeZone
	return result, nil
}

func (h *HostService) SetSystemTimezone(ctx context.Context, args schema.SetSystemTimezoneArgs) error {
	return h.HostRepo.SetSystemTimezone(ctx, args)
}

func (h *HostService) Reboot(ctx context.Context, args schema.RebootArgs) error {
	return h.HostRepo.Reboot(ctx, args)
}

func (h *HostService) Shutdown(ctx context.Context, args schema.ShutdownArgs) error {
	return h.HostRepo.Shutdown(ctx, args)
}
