// Package rpc
// Date: 2022/11/9 10:18
// Author: Amu
// Description:
package rpc

import (
	"amvector/service/model"
	"context"

	"common/database"

	"amvector/service/schema"

	"github.com/amuluze/docker"
	"github.com/patrickmn/go-cache"
)

var _ IService = (*Service)(nil)

type IService interface {
	DockerVersion(context.Context, schema.VersionArgs, *model.Docker) error
	ContainerList(context.Context, schema.ContainerQueryArgs, *model.Containers) error
	ContainerCount(context.Context, schema.QueryCountArgs, *schema.QueryCountReply) error
	ContainerCreate(context.Context, schema.ContainerCreateArgs, *schema.ContainerCreateReply) error
	ContainerUpdate(context.Context, schema.ContainerUpdateArgs, *schema.ContainerUpdateReply) error
	ContainerDelete(context.Context, schema.ContainerDeleteArgs, *schema.ContainerDeleteReply) error
	ContainerStart(context.Context, schema.ContainerStartArgs, *schema.ContainerStartReply) error
	ContainerStop(context.Context, schema.ContainerStopArgs, *schema.ContainerStopReply) error
	ContainerRestart(context.Context, schema.ContainerRestartArgs, *schema.ContainerRestartReply) error
	ImageList(context.Context, schema.ImageQueryArgs, *model.Images) error
	ImagePull(context.Context, schema.ImagePullArgs, *schema.ImagePullReply) error
	ImageTag(context.Context, schema.ImageTagArgs, *schema.ImageTagReply) error
	ImageCount(context.Context, schema.ImageCountArgs, *schema.ImageCountReply) error
	ImageDelete(context.Context, schema.ImageDeleteArgs, *schema.ImageDeleteReply) error
	ImagesPrune(ctx context.Context) error
	ImageImport(context.Context, schema.ImageImportArgs, *schema.ImageImportReply) error
	ImageExport(context.Context, schema.ImageExportArgs, *schema.ImageExportReply) error
	NetworkList(context.Context, schema.NetworkQueryArgs, *model.Networks) error
	NetworkCreate(context.Context, schema.NetworkCreateArgs, *schema.NetworkCreateReply) error
	NetworkCount(context.Context, schema.NetworkCountArgs, *schema.NetworkCountReply) error
	NetworkDelete(context.Context, schema.NetworkDeleteArgs, *schema.NetworkDeleteReply) error
	FilesSearch(context.Context, schema.FilesSearchArgs, *schema.FilesSearchReply) error
	DirSize(context.Context, schema.DirSizeArgs, *schema.DirSizeReply) error
	FileCreate(context.Context, schema.FileCreateArgs, *schema.FileCreateReply) error
	FileDelete(context.Context, schema.FileDeleteArgs, *schema.FileDeleteReply) error
	FileUpload(context.Context, schema.FileUploadArgs, *schema.FileUploadReply) error
	FileDownload(context.Context, schema.FileDownloadArgs, *schema.FileDownloadReply) error
	FolderCreate(context.Context, schema.FolderCreateArgs, *schema.FolderCreateReply) error
	HostInfo(context.Context, schema.HostArgs, *model.Host) error
	CPUInfo(context.Context, schema.CPUArgs, *model.CPU) error
	CPUUsage(context.Context, schema.CPUUsageArgs, *[]model.CPU) error
	MemInfo(context.Context, schema.MemoryArgs, *model.Memory) error
	MemUsage(context.Context, schema.MemoryUsageArgs, *[]model.Memory) error
	DiskInfo(context.Context, schema.DiskArgs, *[]model.Disk) error
	DiskUsage(context.Context, schema.DiskUsageArgs, *[]model.Disk) error
	NetUsage(context.Context, schema.NetworkUsageArgs, *[]model.Net) error
	Reboot(context.Context, schema.RebootArgs, *schema.RebootReply) error
	Shutdown(context.Context, schema.ShutdownArgs, *schema.ShutdownReply) error
	GetDNS(context.Context, schema.GetDNSArgs, *schema.GetDNSReply) error
	SetDNS(context.Context, schema.SetDNSArgs, *schema.SetDNSReply) error
	GetSystemTime(context.Context, schema.GetSystemTimeArgs, *schema.GetSystemTimeReply) error
	SetSystemTime(context.Context, schema.SetSystemTimeArgs, *schema.SetSystemTimeReply) error
	GetSystemTimeZone(context.Context, schema.GetSystemTimeZoneArgs, *schema.GetSystemTimeZoneReply) error
	SetSystemTimeZone(context.Context, schema.SetSystemTimeZoneArgs, *schema.SetSystemTimeZoneReply) error
	GetDockerRegistryMirrors(context.Context, schema.GetDockerRegistryMirrorsArgs, *schema.GetDockerRegistryMirrorsReply) error
	SetDockerRegistryMirrors(context.Context, schema.SetDockerRegistryMirrorsArgs, *schema.SetDockerRegistryMirrorsReply) error
}

type Service struct {
	DB      *database.DB
	Manager *docker.Manager
	cache   *cache.Cache
}

func NewService(db *database.DB, manager *docker.Manager) *Service {
	return &Service{DB: db, Manager: manager, cache: cache.New(0, 0)}
}
