// Package task
// Date: 2022/11/9 10:18
// Author: Amu
// Description:
package task

import (
	"amprobe/service/model"
	rpcSchema "common/rpc"
	"context"
	"time"
)

func (a *Task) HostSummary(ctx context.Context, timestamp time.Time) error {
	args := rpcSchema.HostSummaryArgs{
		Timestamp: timestamp,
	}
	var reply rpcSchema.HostSummaryReply
	if err := a.rpcClient.Call(ctx, "HostSummary", args, &reply); err != nil {
		return err
	}

	if err := a.db.Unscoped().Where("1 = 1").Delete(&model.Host{}).Error; err != nil {
		return err
	}

	if err := a.db.Model(&model.Host{}).Create(&model.Host{
		Timestamp:       timestamp,
		Uptime:          reply.Data.Uptime,
		Hostname:        reply.Data.Hostname,
		Os:              reply.Data.Os,
		Platform:        reply.Data.Platform,
		PlatformVersion: reply.Data.PlatformVersion,
		KernelArch:      reply.Data.KernelArch,
		KernelVersion:   reply.Data.KernelVersion,
	}).Error; err != nil {
		return err
	}
	return nil
}

func (a *Task) CPUSummary(ctx context.Context, timestamp time.Time) error {
	args := rpcSchema.CPUSummaryArgs{
		Timestamp: timestamp,
	}
	var reply rpcSchema.CPUSummaryReply
	if err := a.rpcClient.Call(ctx, "CPUSummary", args, &reply); err != nil {
		return err
	}

	if err := a.db.Model(&model.CPU{}).Create(&model.CPU{
		Timestamp:  timestamp,
		CPUPercent: reply.Data.CPUPercent,
	}).Error; err != nil {
		return err
	}
	return nil
}

func (a *Task) MemorySummary(ctx context.Context, timestamp time.Time) error {
	args := rpcSchema.MemorySummaryArgs{
		Timestamp: timestamp,
	}
	var reply rpcSchema.MemorySummaryReply
	if err := a.rpcClient.Call(ctx, "MemorySummary", args, &reply); err != nil {
		return err
	}

	if err := a.db.Model(&model.Memory{}).Create(&model.Memory{
		Timestamp:  timestamp,
		MemPercent: reply.Data.MemPercent,
	}).Error; err != nil {
		return err
	}
	return nil
}

/*DiskSummary
 * disk 函数用于获取磁盘指标，并将其存储到数据库中。
 * 计算方法：两次采样间隔之间磁盘读写的平均速率
 */
func (a *Task) DiskSummary(ctx context.Context, timestamp time.Time) error {
	args := rpcSchema.DiskSummaryArgs{
		Timestamp: timestamp,
		Devices:   a.devices,
	}
	var reply rpcSchema.DiskSummaryReply
	if err := a.rpcClient.Call(ctx, "DiskSummary", args, &reply); err != nil {
		return err
	}

	var infos []model.Disk
	for _, item := range reply.Data {
		infos = append(infos, model.Disk{
			Timestamp:   item.Timestamp,
			Device:      item.Device,
			DiskPercent: item.DiskPercent,
			DiskTotal:   item.DiskTotal,
			DiskUsed:    item.DiskUsed,
			DiskRead:    item.DiskRead,
			DiskWrite:   item.DiskWrite,
		})
	}
	if err := a.db.Model(&model.Disk{}).Create(&infos).Error; err != nil {
		return err
	}
	return nil
}

/*NetSummary
 * network 函数用于获取网络指标，并将其存储到数据库中。
 * 计算方法：两次采样间隔之间发送、接收的平均速率
 */
func (a *Task) NetSummary(ctx context.Context, timestamp time.Time) error {
	args := rpcSchema.NetSummaryArgs{
		Timestamp: timestamp,
		Ethernets: a.ethernet,
	}
	var reply rpcSchema.NetSummaryReply
	if err := a.rpcClient.Call(ctx, "NetSummary", args, &reply); err != nil {
		return err
	}

	var infos []model.Net
	for _, item := range reply.Data {
		infos = append(infos, model.Net{
			Timestamp: item.Timestamp,
			Ethernet:  item.Ethernet,
			NetRecv:   item.NetRecv,
			NetSend:   item.NetSend,
		})
	}
	if err := a.db.Model(&model.Net{}).Create(&infos).Error; err != nil {
		return err
	}
	return nil
}
