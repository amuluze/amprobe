// Package task
// Date: 2022/11/9 10:18
// Author: Amu
// Description:
package task

import (
	"amprobe/service/model"
	rpcSchema "common/rpc"
	"context"
	"log/slog"
	"time"
)

func (a *Task) DockerSummary(ctx context.Context, timestamp time.Time) error {
	args := rpcSchema.DockerSummaryArgs{
		Timestamp: timestamp,
	}
	var reply rpcSchema.DockerSummaryReply
	err := a.rpcClient.Call(ctx, "DockerSummary", args, &reply)
	if err != nil {
		return err
	}

	if err := a.db.Unscoped().Where("1 = 1").Delete(&model.Docker{}).Error; err != nil {
		slog.Warn("failed to clear docker summary", "error", err)
	}
	if err := a.db.Model(&model.Docker{}).Create(&model.Docker{
		Timestamp:     timestamp,
		DockerVersion: reply.Data.DockerVersion,
		APIVersion:    reply.Data.APIVersion,
		MinAPIVersion: reply.Data.MinAPIVersion,
		GitCommit:     reply.Data.GitCommit,
		GoVersion:     reply.Data.GoVersion,
		Os:            reply.Data.Os,
		Arch:          reply.Data.Arch,
	}).Error; err != nil {
		return err
	}
	return nil
}

func (a *Task) ContainerSummary(ctx context.Context, timestamp time.Time) error {
	args := rpcSchema.ContainerSummaryArgs{
		Timestamp: timestamp,
	}
	var reply rpcSchema.ContainerSummaryReply
	err := a.rpcClient.Call(ctx, "ContainerSummary", args, &reply)
	if err != nil {
		return err
	}

	if err := a.db.Unscoped().Where("1 = 1").Delete(&model.Container{}).Error; err != nil {
		slog.Warn("failed to clear container summary", "error", err)
	}

	var containers []model.Container
	for _, item := range reply.Data {
		containers = append(containers, model.Container{
			Timestamp:    timestamp,
			ContainerID:  item.ContainerID,
			Name:         item.Name,
			Image:        item.Image,
			Network:      item.Network,
			IP:           item.IP,
			Ports:        item.Ports,
			State:        item.State,
			Uptime:       item.Uptime,
			Volumes:      item.Volumes,
			Environments: item.Environments,
			Commands:     item.Commands,
			Labels:       item.Labels,
			CPUPercent:   item.CPUPercent,
			MemPercent:   item.MemPercent,
			MemUsage:     item.MemUsage,
			MemLimit:     item.MemLimit,
		})
	}
	if err := a.db.Model(&model.Container{}).Create(&containers).Error; err != nil {
		return err
	}
	return nil
}

func (a *Task) ImageSummary(ctx context.Context, timestamp time.Time) error {
	args := rpcSchema.ImageSummaryArgs{
		Timestamp: timestamp,
	}
	var reply rpcSchema.ImageSummaryReply
	err := a.rpcClient.Call(ctx, "ImageSummary", args, &reply)
	if err != nil {
		return err
	}

	if err := a.db.Unscoped().Where("1 = 1").Delete(&model.Image{}).Error; err != nil {
		slog.Warn("failed to clear image summary", "error", err)
	}
	var images []model.Image
	for _, image := range reply.Data {
		images = append(images, model.Image{
			Timestamp: timestamp,
			ImageID:   image.ImageID,
			Name:      image.Name,
			Tag:       image.Tag,
			Created:   image.Created,
			Size:      image.Size,
		})
	}
	if err := a.db.Model(&model.Image{}).Create(&images).Error; err != nil {
		return err
	}
	return nil
}

func (a *Task) NetworkSummary(ctx context.Context, timestamp time.Time) error {
	args := rpcSchema.NetworkSummaryArgs{
		Timestamp: timestamp,
	}
	var reply rpcSchema.NetworkSummaryReply
	err := a.rpcClient.Call(ctx, "NetworkSummary", args, &reply)
	if err != nil {
		return err
	}

	if err := a.db.Unscoped().Where("1 = 1").Delete(&model.Network{}).Error; err != nil {
		slog.Warn("failed to clear network summary", "error", err)
	}
	var networks []model.Network
	for _, item := range reply.Data {
		networks = append(networks, model.Network{
			Timestamp: timestamp,
			NetworkID: item.NetworkID,
			Name:      item.Name,
			Driver:    item.Driver,
			Created:   item.Created,
			Internal:  item.Internal,
			Subnet:    item.Subnet,
			Gateway:   item.Gateway,
			Labels:    item.Labels,
		})
	}
	if err := a.db.Model(&model.Network{}).Create(&networks).Error; err != nil {
		return err
	}
	return nil
}
