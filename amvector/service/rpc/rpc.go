// Package rpc
// Date: 2022/11/9 10:18
// Author: Amu
// Description:
package rpc

import (
	"common/rpc"
	"context"
	
	"github.com/amuluze/docker"
	"github.com/patrickmn/go-cache"
)

var _ IService = (*Service)(nil)

type IService interface {
	ContainerCreate(context.Context, rpc.ContainerCreateArgs, *rpc.ContainerCreateReply) error
	ContainerUpdate(context.Context, rpc.ContainerUpdateArgs, *rpc.ContainerUpdateReply) error
	ContainerDelete(context.Context, rpc.ContainerDeleteArgs, *rpc.ContainerDeleteReply) error
	ContainerStart(context.Context, rpc.ContainerStartArgs, *rpc.ContainerStartReply) error
	ContainerStop(context.Context, rpc.ContainerStopArgs, *rpc.ContainerStopReply) error
	ContainerRestart(context.Context, rpc.ContainerRestartArgs, *rpc.ContainerRestartReply) error
	ImagePull(context.Context, rpc.ImagePullArgs, *rpc.ImagePullReply) error
	ImageTag(context.Context, rpc.ImageTagArgs, *rpc.ImageTagReply) error
	ImageDelete(context.Context, rpc.ImageDeleteArgs, *rpc.ImageDeleteReply) error
	ImagesPrune(ctx context.Context) error
	ImageImport(context.Context, rpc.ImageImportArgs, *rpc.ImageImportReply) error
	ImageExport(context.Context, rpc.ImageExportArgs, *rpc.ImageExportReply) error
	NetworkCreate(context.Context, rpc.NetworkCreateArgs, *rpc.NetworkCreateReply) error
	NetworkDelete(context.Context, rpc.NetworkDeleteArgs, *rpc.NetworkDeleteReply) error
	
	DockerSummary(context.Context, rpc.DockerSummaryArgs, *rpc.DockerSummaryReply) error
	ContainerSummary(context.Context, rpc.ContainerSummaryArgs, *rpc.ContainerSummaryReply) error
	ImageSummary(context.Context, rpc.ImageSummaryArgs, *rpc.ImageSummaryReply) error
	NetworkSummary(context.Context, rpc.NetworkSummaryArgs, *rpc.NetworkSummaryReply) error
	HostSummary(context.Context, rpc.HostSummaryArgs, *rpc.HostSummaryReply) error
	CPUSummary(context.Context, rpc.CPUSummaryArgs, *rpc.CPUSummaryReply) error
	MemorySummary(context.Context, rpc.MemorySummaryArgs, *rpc.MemorySummaryReply) error
	DiskSummary(context.Context, rpc.DiskSummaryArgs, *rpc.DiskSummaryReply) error
	NetSummary(context.Context, rpc.NetSummaryArgs, *rpc.NetSummaryReply) error
	
	FilesSearch(context.Context, rpc.FilesSearchArgs, *rpc.FilesSearchReply) error
	DirSize(context.Context, rpc.DirSizeArgs, *rpc.DirSizeReply) error
	FileCreate(context.Context, rpc.FileCreateArgs, *rpc.FileCreateReply) error
	FileDelete(context.Context, rpc.FileDeleteArgs, *rpc.FileDeleteReply) error
	FileUpload(context.Context, rpc.FileUploadArgs, *rpc.FileUploadReply) error
	FileDownload(context.Context, rpc.FileDownloadArgs, *rpc.FileDownloadReply) error
	FolderCreate(context.Context, rpc.FolderCreateArgs, *rpc.FolderCreateReply) error
	Reboot(context.Context, rpc.RebootArgs, *rpc.RebootReply) error
	Shutdown(context.Context, rpc.ShutdownArgs, *rpc.ShutdownReply) error
	GetDNS(context.Context, rpc.GetDNSArgs, *rpc.GetDNSReply) error
	SetDNS(context.Context, rpc.SetDNSArgs, *rpc.SetDNSReply) error
	GetSystemTime(context.Context, rpc.GetSystemTimeArgs, *rpc.GetSystemTimeReply) error
	SetSystemTime(context.Context, rpc.SetSystemTimeArgs, *rpc.SetSystemTimeReply) error
	GetSystemTimeZone(context.Context, rpc.GetSystemTimeZoneArgs, *rpc.GetSystemTimeZoneReply) error
	SetSystemTimeZone(context.Context, rpc.SetSystemTimeZoneArgs, *rpc.SetSystemTimeZoneReply) error
	GetDockerRegistryMirrors(context.Context, rpc.GetDockerRegistryMirrorsArgs, *rpc.GetDockerRegistryMirrorsReply) error
	SetDockerRegistryMirrors(context.Context, rpc.SetDockerRegistryMirrorsArgs, *rpc.SetDockerRegistryMirrorsReply) error
}

type Service struct {
	Manager *docker.Manager
	cache   *cache.Cache
}

func NewService(manager *docker.Manager) *Service {
	return &Service{
		Manager: manager,
		cache:   cache.New(0, 0),
	}
}
