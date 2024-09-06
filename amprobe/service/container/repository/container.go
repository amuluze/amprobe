// Package repository
// Date: 2024/3/6 12:46
// Author: Amu
// Description:
package repository

import (
	"context"
	"log/slog"

	"amprobe/pkg/database"
	"amprobe/pkg/rpc"
	"amprobe/service/model"
	"amprobe/service/schema"

	"github.com/google/wire"

	rpcSchema "common/rpc"
)

var ContainerRepoSet = wire.NewSet(NewContainerRepo, wire.Bind(new(IContainerRepo), new(*ContainerRepo)))

var _ IContainerRepo = (*ContainerRepo)(nil)

type IContainerRepo interface {
	Version(ctx context.Context) (model.Docker, error)
	ContainerList(ctx context.Context, args schema.ContainerQueryArgs) (model.Containers, error)
	ContainersByImage(ctx context.Context, image string) (int, error)
	ContainerCount(ctx context.Context) (int, error)
	ImageList(ctx context.Context, args schema.ImageQueryArgs) (model.Images, error)
	ImageCount(ctx context.Context) (int, error)
	NetworkList(ctx context.Context, args schema.NetworkListArgs) (model.Networks, error)
	NetworkCount(ctx context.Context) (int, error)

	ContainerCreate(context.Context, rpcSchema.ContainerCreateArgs) (rpcSchema.ContainerCreateReply, error)
	ContainerUpdate(context.Context, rpcSchema.ContainerUpdateArgs) (rpcSchema.ContainerUpdateReply, error)
	ContainerDelete(context.Context, rpcSchema.ContainerDeleteArgs) (rpcSchema.ContainerDeleteReply, error)
	ContainerStart(context.Context, rpcSchema.ContainerStartArgs) (rpcSchema.ContainerStartReply, error)
	ContainerStop(context.Context, rpcSchema.ContainerStopArgs) (rpcSchema.ContainerStopReply, error)
	ContainerRestart(context.Context, rpcSchema.ContainerRestartArgs) (rpcSchema.ContainerRestartReply, error)
	ImagePull(context.Context, rpcSchema.ImagePullArgs) (rpcSchema.ImagePullReply, error)
	ImageTag(context.Context, rpcSchema.ImageTagArgs) (rpcSchema.ImageTagReply, error)
	ImageDelete(context.Context, rpcSchema.ImageDeleteArgs) (rpcSchema.ImageDeleteReply, error)
	ImagesPrune(ctx context.Context) error
	ImageImport(context.Context, rpcSchema.ImageImportArgs) (rpcSchema.ImageImportReply, error)
	ImageExport(context.Context, rpcSchema.ImageExportArgs) (rpcSchema.ImageExportReply, error)
	NetworkCreate(context.Context, rpcSchema.NetworkCreateArgs) (rpcSchema.NetworkCreateReply, error)
	NetworkDelete(context.Context, rpcSchema.NetworkDeleteArgs) (rpcSchema.NetworkDeleteReply, error)

	GetDockerRegistryMirrors(context.Context, rpcSchema.GetDockerRegistryMirrorsArgs) (rpcSchema.GetDockerRegistryMirrorsReply, error)
	SetDockerRegistryMirrors(context.Context, rpcSchema.SetDockerRegistryMirrorsArgs) (rpcSchema.SetDockerRegistryMirrorsReply, error)
}

type ContainerRepo struct {
	DB        *database.DB
	RPCClient *rpc.Client
}

func NewContainerRepo(db *database.DB, client *rpc.Client) *ContainerRepo {
	return &ContainerRepo{DB: db, RPCClient: client}
}

func (a *ContainerRepo) Version(ctx context.Context) (model.Docker, error) {
	var res model.Docker
	if err := a.DB.Model(&model.Docker{}).First(&res).Error; err != nil {
		return res, err
	}
	return res, nil
}

func (a *ContainerRepo) ContainerList(ctx context.Context, args schema.ContainerQueryArgs) (model.Containers, error) {
	var containers model.Containers
	if err := a.DB.Model(&model.Container{}).Order("created_at desc").Offset((args.Page - 1) * args.Size).Limit(args.Size).Find(&containers).Error; err != nil {
		return containers, err
	}
	return containers, nil
}

func (a *ContainerRepo) ContainersByImage(ctx context.Context, image string) (int, error) {
	var count int64
	if err := a.DB.Model(&model.Container{}).Where("image = ?", image).Count(&count).Error; err != nil {
		return 0, err
	}
	return int(count), nil
}

func (a *ContainerRepo) ContainerCount(ctx context.Context) (int, error) {
	var total int64
	if err := a.DB.Model(&model.Container{}).Order("created_at desc").Count(&total).Error; err != nil {
		return int(total), err
	}
	return int(total), nil
}

func (a *ContainerRepo) ImageList(ctx context.Context, args schema.ImageQueryArgs) (model.Images, error) {
	var images model.Images
	if err := a.DB.Model(&model.Image{}).Order("created_at desc").Offset((args.Page - 1) * args.Size).Limit(args.Size).Find(&images).Error; err != nil {
		return images, err
	}
	return images, nil
}

func (a *ContainerRepo) ImageCount(ctx context.Context) (int, error) {
	var total int64
	if err := a.DB.Model(&model.Images{}).Order("created_at desc").Count(&total).Error; err != nil {
		return int(total), err
	}
	return int(total), nil
}

func (a *ContainerRepo) NetworkList(ctx context.Context, args schema.NetworkListArgs) (model.Networks, error) {
	var networks model.Networks
	if err := a.DB.Model(&model.Network{}).Order("created_at desc").Offset((args.Page - 1) * args.Size).Limit(args.Size).Find(&networks).Error; err != nil {
		return networks, err
	}
	return networks, nil
}

func (a *ContainerRepo) NetworkCount(ctx context.Context) (int, error) {
	var total int64
	if err := a.DB.Model(&model.Network{}).Count(&total).Error; err != nil {
		return int(total), err
	}
	return int(total), nil
}

func (a *ContainerRepo) ContainerCreate(ctx context.Context, args rpcSchema.ContainerCreateArgs) (rpcSchema.ContainerCreateReply, error) {
	var reply rpcSchema.ContainerCreateReply
	err := a.RPCClient.Call(ctx, "ContainerCreate", args, &reply)
	return reply, err
}

func (a *ContainerRepo) ContainerUpdate(ctx context.Context, args rpcSchema.ContainerUpdateArgs) (rpcSchema.ContainerUpdateReply, error) {
	var reply rpcSchema.ContainerUpdateReply
	err := a.RPCClient.Call(ctx, "ContainerUpdate", args, &reply)
	return reply, err
}

func (a *ContainerRepo) ContainerDelete(ctx context.Context, args rpcSchema.ContainerDeleteArgs) (rpcSchema.ContainerDeleteReply, error) {
	var reply rpcSchema.ContainerDeleteReply
	err := a.RPCClient.Call(ctx, "ContainerDelete", args, &reply)
	return reply, err
}

func (a *ContainerRepo) ContainerStart(ctx context.Context, args rpcSchema.ContainerStartArgs) (rpcSchema.ContainerStartReply, error) {
	var reply rpcSchema.ContainerStartReply
	err := a.RPCClient.Call(ctx, "ContainerStart", args, &reply)
	return reply, err
}

func (a *ContainerRepo) ContainerStop(ctx context.Context, args rpcSchema.ContainerStopArgs) (rpcSchema.ContainerStopReply, error) {
	var reply rpcSchema.ContainerStopReply
	err := a.RPCClient.Call(ctx, "ContainerStop", args, &reply)
	return reply, err
}

func (a *ContainerRepo) ContainerRestart(ctx context.Context, args rpcSchema.ContainerRestartArgs) (rpcSchema.ContainerRestartReply, error) {
	var reply rpcSchema.ContainerRestartReply
	err := a.RPCClient.Call(ctx, "ContainerRestart", args, &reply)
	return reply, err
}

func (a *ContainerRepo) ImagePull(ctx context.Context, args rpcSchema.ImagePullArgs) (rpcSchema.ImagePullReply, error) {
	var reply rpcSchema.ImagePullReply
	err := a.RPCClient.Call(ctx, "ImagePull", args, &reply)
	return reply, err
}

func (a *ContainerRepo) ImageTag(ctx context.Context, args rpcSchema.ImageTagArgs) (rpcSchema.ImageTagReply, error) {
	var reply rpcSchema.ImageTagReply
	err := a.RPCClient.Call(ctx, "ImageTag", args, &reply)
	return reply, err
}

func (a *ContainerRepo) ImageDelete(ctx context.Context, args rpcSchema.ImageDeleteArgs) (rpcSchema.ImageDeleteReply, error) {
	var reply rpcSchema.ImageDeleteReply
	err := a.RPCClient.Call(ctx, "ImageDelete", args, &reply)
	return reply, err
}

func (a *ContainerRepo) ImagesPrune(ctx context.Context) error {
	return a.RPCClient.Call(ctx, "ImagesPrune", nil, nil)
}

func (a *ContainerRepo) ImageImport(ctx context.Context, args rpcSchema.ImageImportArgs) (rpcSchema.ImageImportReply, error) {
	var reply rpcSchema.ImageImportReply
	err := a.RPCClient.Call(ctx, "ImageImport", args, &reply)
	return reply, err
}

func (a *ContainerRepo) ImageExport(ctx context.Context, args rpcSchema.ImageExportArgs) (rpcSchema.ImageExportReply, error) {
	var reply rpcSchema.ImageExportReply
	err := a.RPCClient.Call(ctx, "ImageExport", args, &reply)
	return reply, err
}

func (a *ContainerRepo) NetworkCreate(ctx context.Context, args rpcSchema.NetworkCreateArgs) (rpcSchema.NetworkCreateReply, error) {
	var reply rpcSchema.NetworkCreateReply
	err := a.RPCClient.Call(ctx, "NetworkCreate", args, &reply)
	return reply, err
}

func (a *ContainerRepo) NetworkDelete(ctx context.Context, args rpcSchema.NetworkDeleteArgs) (rpcSchema.NetworkDeleteReply, error) {
	var reply rpcSchema.NetworkDeleteReply
	err := a.RPCClient.Call(ctx, "NetworkDelete", args, &reply)
	return reply, err
}

func (a *ContainerRepo) GetDockerRegistryMirrors(ctx context.Context, args rpcSchema.GetDockerRegistryMirrorsArgs) (rpcSchema.GetDockerRegistryMirrorsReply, error) {
	reply := rpcSchema.GetDockerRegistryMirrorsReply{Mirrors: []string{}}
	slog.Info("===>", "args", args, "reply", reply, "client", a.RPCClient)
	err := a.RPCClient.Call(ctx, "GetDockerRegistryMirrors", args, &reply)
	slog.Info("===>", "reply", reply, "error", err)
	return reply, err
}

func (a *ContainerRepo) SetDockerRegistryMirrors(ctx context.Context, args rpcSchema.SetDockerRegistryMirrorsArgs) (rpcSchema.SetDockerRegistryMirrorsReply, error) {
	var reply rpcSchema.SetDockerRegistryMirrorsReply
	err := a.RPCClient.Call(ctx, "SetDockerRegistryMirrors", args, &reply)
	return reply, err
}
