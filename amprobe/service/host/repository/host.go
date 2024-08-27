// Package repository
// Date: 2024/3/6 12:53
// Author: Amu
// Description:
package repository

import (
	"context"
	"fmt"
	"strings"
	"time"

	"amprobe/pkg/database"
	"amprobe/pkg/rpc"
	"amprobe/service/model"
	"amprobe/service/schema"
	rpcSchema "common/rpc"

	"github.com/google/wire"
)

var HostRepoSet = wire.NewSet(NewHostRepo, wire.Bind(new(IHostRepo), new(*HostRepo)))

var _ IHostRepo = (*HostRepo)(nil)

type IHostRepo interface {
	HostInfo(ctx context.Context) (model.Host, error)
	CPUInfo(ctx context.Context) (model.CPU, error)
	CPUUsage(ctx context.Context, args schema.CPUUsageArgs) ([]model.CPU, error)
	MemInfo(ctx context.Context) (model.Memory, error)
	MemUsage(ctx context.Context, args schema.MemoryUsageArgs) ([]model.Memory, error)
	DiskInfo(ctx context.Context) ([]model.Disk, error)
	DiskUsage(ctx context.Context, args schema.DiskUsageArgs) ([]model.Disk, error)
	NetUsage(ctx context.Context, args schema.NetworkUsageArgs) ([]model.Net, error)

	FilesSearch(ctx context.Context, args schema.FilesSearchArgs) (rpcSchema.FilesSearchReply, error)
	FileUpload(ctx context.Context, args schema.FileUploadArgs) error
	FileDownload(ctx context.Context, args schema.FileDownloadArgs) (rpcSchema.FileDownloadReply, error)
	FileDelete(ctx context.Context, args schema.FileDeleteArgs) error
	FileCreate(ctx context.Context, args schema.FileCreateArgs) error
	FolderCreate(ctx context.Context, args schema.FolderCreateArgs) error
	GetDNSSettings(ctx context.Context, args schema.GetDNSSettingsArgs) (rpcSchema.GetDNSReply, error)
	SetDNSSettings(ctx context.Context, args schema.SetDNSSettingsArgs) error
	GetSystemTime(ctx context.Context, args schema.GetSystemTimeArgs) (rpcSchema.GetSystemTimeReply, error)
	SetSystemTime(ctx context.Context, args schema.SetSystemTimeArgs) error
	GetSystemTimezoneList(ctx context.Context, args schema.GetSystemTimezoneListArgs) (rpcSchema.GetSystemTimeZoneListReply, error)
	GetSystemTimezone(ctx context.Context, args schema.GetSystemTimezoneArgs) (rpcSchema.GetSystemTimeZoneReply, error)
	SetSystemTimezone(ctx context.Context, args schema.SetSystemTimezoneArgs) error
	Reboot(ctx context.Context, args schema.RebootArgs) error
	Shutdown(ctx context.Context, args schema.ShutdownArgs) error
}

type HostRepo struct {
	DB        *database.DB
	RPCClient *rpc.Client
}

func NewHostRepo(db *database.DB, client *rpc.Client) *HostRepo {
	return &HostRepo{DB: db, RPCClient: client}
}

func (h *HostRepo) HostInfo(ctx context.Context) (model.Host, error) {
	var hostInfo model.Host
	if err := h.DB.Order("timestamp desc").Take(&hostInfo).Error; err != nil {
		return hostInfo, err
	}
	return hostInfo, nil
}

func (h *HostRepo) CPUInfo(ctx context.Context) (model.CPU, error) {
	var cpuInfo model.CPU
	if err := h.DB.Order("timestamp desc").Take(&cpuInfo).Error; err != nil {
		return cpuInfo, err
	}
	return cpuInfo, nil
}

func (h *HostRepo) CPUUsage(ctx context.Context, args schema.CPUUsageArgs) ([]model.CPU, error) {
	var cpuInfos []model.CPU
	if err := h.DB.Model(&model.CPU{}).Where("timestamp > ? and timestamp < ?", time.Unix(args.StartTime, 0), time.Unix(args.EndTime, 0)).Order("timestamp asc").Find(&cpuInfos).Error; err != nil {
		return cpuInfos, err
	}
	return cpuInfos, nil
}

func (h *HostRepo) MemInfo(ctx context.Context) (model.Memory, error) {
	var memInfo model.Memory
	if err := h.DB.Order("timestamp desc").Take(&memInfo).Error; err != nil {
		return memInfo, err
	}
	return memInfo, nil
}

func (h *HostRepo) MemUsage(ctx context.Context, args schema.MemoryUsageArgs) ([]model.Memory, error) {
	var memInfos []model.Memory
	if err := h.DB.Model(&model.Memory{}).Where("timestamp > ?", time.Unix(args.StartTime, 0)).Order("timestamp asc").Find(&memInfos).Error; err != nil {
		return memInfos, err
	}
	return memInfos, nil
}

func (h *HostRepo) DiskInfo(ctx context.Context) ([]model.Disk, error) {
	var diskInfos []model.Disk
	if err := h.DB.Model(&model.Disk{}).Group("device").Order("timestamp desc").Find(&diskInfos).Error; err != nil {
		return diskInfos, err
	}
	return diskInfos, nil
}

func (h *HostRepo) DiskUsage(ctx context.Context, args schema.DiskUsageArgs) ([]model.Disk, error) {
	var diskInfos []model.Disk
	if err := h.DB.Model(&model.Disk{}).Where("timestamp > ?", time.Unix(args.StartTime, 0)).Order("timestamp asc").Find(&diskInfos).Error; err != nil {
		return diskInfos, err
	}
	return diskInfos, nil
}

func (h *HostRepo) NetUsage(ctx context.Context, args schema.NetworkUsageArgs) ([]model.Net, error) {
	var netInfos []model.Net
	if err := h.DB.Model(&model.Net{}).Where("timestamp > ?", time.Unix(args.StartTime, 0)).Order("timestamp asc").Find(&netInfos).Error; err != nil {
		return netInfos, err
	}
	return netInfos, nil
}

func (h *HostRepo) FilesSearch(ctx context.Context, args schema.FilesSearchArgs) (rpcSchema.FilesSearchReply, error) {
	rpcArgs := rpcSchema.FilesSearchArgs{
		Path: args.Path,
	}
	var reply rpcSchema.FilesSearchReply
	err := h.RPCClient.Call(ctx, "FilesSearch", rpcArgs, &reply)
	return reply, err
}

func (h *HostRepo) FileUpload(ctx context.Context, args schema.FileUploadArgs) error {
	rpcArgs := rpcSchema.FileUploadArgs{
		SourceFilePath: args.SourceFilePath,
		TargetFilePath: args.TargetFilePath,
	}
	var reply rpcSchema.FileUploadReply
	return h.RPCClient.Call(ctx, "FileUpload", rpcArgs, &reply)
}

func (h *HostRepo) FileDownload(ctx context.Context, args schema.FileDownloadArgs) (rpcSchema.FileDownloadReply, error) {
	filename := strings.Split(args.Filepath, "/")[len(strings.Split(args.Filepath, "/"))-1]
	rpcArgs := rpcSchema.FileDownloadArgs{
		SourceFilePath: args.Filepath,
		TargetFilePath: fmt.Sprintf("/tmp/%s", filename),
	}
	var reply rpcSchema.FileDownloadReply
	err := h.RPCClient.Call(ctx, "FileDownload", rpcArgs, &reply)
	return reply, err
}

func (h *HostRepo) FileDelete(ctx context.Context, args schema.FileDeleteArgs) error {
	rpcArgs := rpcSchema.FileDeleteArgs{
		Filepath: args.Filepath,
	}
	var reply rpcSchema.FileDeleteReply
	return h.RPCClient.Call(ctx, "FileDelete", rpcArgs, &reply)
}

func (h *HostRepo) FileCreate(ctx context.Context, args schema.FileCreateArgs) error {
	rpcArgs := rpcSchema.FileCreateArgs{
		Path:     args.Path,
		FileName: args.FileName,
	}
	var reply rpcSchema.FileCreateReply
	return h.RPCClient.Call(ctx, "FileCreate", rpcArgs, &reply)
}

func (h *HostRepo) FolderCreate(ctx context.Context, args schema.FolderCreateArgs) error {
	rpcArgs := rpcSchema.FolderCreateArgs{
		Path:       args.Path,
		FolderName: args.FolderName,
	}
	var reply rpcSchema.FolderCreateReply
	return h.RPCClient.Call(ctx, "FolderCreate", rpcArgs, &reply)
}

func (h *HostRepo) GetDNSSettings(ctx context.Context, args schema.GetDNSSettingsArgs) (rpcSchema.GetDNSReply, error) {
	rpcArgs := rpcSchema.GetDNSArgs{}
	var reply rpcSchema.GetDNSReply
	err := h.RPCClient.Call(ctx, "GetDNS", rpcArgs, &reply)
	return reply, err
}

func (h *HostRepo) SetDNSSettings(ctx context.Context, args schema.SetDNSSettingsArgs) error {
	rpcArgs := rpcSchema.SetDNSArgs{
		DNS: args.DNS,
	}
	var reply rpcSchema.SetDNSReply
	return h.RPCClient.Call(ctx, "SetDNS", rpcArgs, &reply)
}

func (h *HostRepo) GetSystemTime(ctx context.Context, args schema.GetSystemTimeArgs) (rpcSchema.GetSystemTimeReply, error) {
	rpcArgs := rpcSchema.GetSystemTimeArgs{}
	var reply rpcSchema.GetSystemTimeReply
	err := h.RPCClient.Call(ctx, "GetSystemTime", rpcArgs, &reply)
	return reply, err
}

func (h *HostRepo) SetSystemTime(ctx context.Context, args schema.SetSystemTimeArgs) error {
	rpcArgs := rpcSchema.SetSystemTimeArgs{
		SystemTime: args.SystemTime,
	}
	var reply rpcSchema.SetSystemTimeReply
	return h.RPCClient.Call(ctx, "SetSystemTime", rpcArgs, &reply)
}

func (h *HostRepo) GetSystemTimezoneList(ctx context.Context, args schema.GetSystemTimezoneListArgs) (rpcSchema.GetSystemTimeZoneListReply, error) {
	rpcArgs := rpcSchema.GetSystemTimeZoneListArgs{}
	var reply rpcSchema.GetSystemTimeZoneListReply
	err := h.RPCClient.Call(ctx, "GetSystemTimeZoneList", rpcArgs, &reply)
	return reply, err
}

func (h *HostRepo) GetSystemTimezone(ctx context.Context, args schema.GetSystemTimezoneArgs) (rpcSchema.GetSystemTimeZoneReply, error) {
	rpcArgs := rpcSchema.GetSystemTimeZoneArgs{}
	var reply rpcSchema.GetSystemTimeZoneReply
	err := h.RPCClient.Call(ctx, "GetSystemTimeZone", rpcArgs, &reply)
	return reply, err
}

func (h *HostRepo) SetSystemTimezone(ctx context.Context, args schema.SetSystemTimezoneArgs) error {
	rpcArgs := rpcSchema.SetSystemTimeZoneArgs{
		SystemTimeZone: args.SystemTimeZone,
	}
	var reply rpcSchema.SetSystemTimeZoneReply
	return h.RPCClient.Call(ctx, "SetSystemTimeZone", rpcArgs, &reply)
}

func (h *HostRepo) Reboot(ctx context.Context, args schema.RebootArgs) error {
	rpcArgs := rpcSchema.RebootArgs{}
	var reply rpcSchema.RebootReply
	return h.RPCClient.Call(ctx, "Reboot", rpcArgs, &reply)
}

func (h *HostRepo) Shutdown(ctx context.Context, args schema.ShutdownArgs) error {
	rpcArgs := rpcSchema.ShutdownArgs{}
	var reply rpcSchema.ShutdownReply
	return h.RPCClient.Call(ctx, "Shutdown", rpcArgs, &reply)
}
