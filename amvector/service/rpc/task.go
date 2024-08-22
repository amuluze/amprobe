// Package rpc
// Date: 2022/11/9 10:18
// Author: Amu
// Description:
package rpc

import (
	"amvector/pkg/psutil"
	"common/rpc"
	"context"
	"encoding/json"
	"strings"
)

const (
	LatestDiskReadKey   = "latest_disk_io_read_"
	LatestDisKWriteKey  = "latest_disk_io_write_"
	LatestNetReceiveKey = "latest_net_io_receive_"
	LatestNetSendKey    = "latest_net_io_send_"
)

/************************* Docker *************************/

func (s *Service) ContainerSummary(ctx context.Context, args rpc.ContainerSummaryArgs, reply *rpc.ContainerSummaryReply) error {
	containerList, err := s.Manager.ListContainer(ctx)
	if err != nil {
		return err
	}
	var containers []rpc.Container
	for _, info := range containerList {
		cpuPercent, err := s.Manager.GetContainerCpu(ctx, info.ID[:6])
		if err != nil {
			return err
		}
		memPercent, used, limit, err := s.Manager.GetContainerMem(ctx, info.ID[:6])
		if err != nil {
			return err
		}
		labels, _ := json.Marshal(info.Labels)
		container := rpc.Container{
			Timestamp:    args.Timestamp,
			Name:         info.Name,
			State:        info.State,
			Image:        info.Image,
			Uptime:       info.Uptime,
			IP:           info.IP,
			Ports:        strings.Join(info.Ports, ","),
			Volumes:      strings.Join(info.Volumes, ","),
			Environments: strings.Join(info.Environments, ","),
			Labels:       string(labels),
			CPUPercent:   cpuPercent,
			MemPercent:   memPercent,
			MemUsage:     used,
			MemLimit:     limit,
		}
		containers = append(containers, container)
	}
	reply.Data = containers
	
	return nil
}

func (s *Service) DockerSummary(ctx context.Context, args rpc.DockerSummaryArgs, reply *rpc.DockerSummaryReply) error {
	dockerVersion, err := s.Manager.Version(ctx)
	if err != nil {
		return err
	}
	reply.Data.Timestamp = args.Timestamp
	reply.Data.DockerVersion = dockerVersion.DockerVersion
	reply.Data.APIVersion = dockerVersion.APIVersion
	reply.Data.MinAPIVersion = dockerVersion.MinAPIVersion
	reply.Data.GitCommit = dockerVersion.GitCommit
	reply.Data.GoVersion = dockerVersion.GoVersion
	reply.Data.Os = dockerVersion.OS
	reply.Data.Arch = dockerVersion.Arch
	return nil
}

func (s *Service) ImageSummary(ctx context.Context, args rpc.ImageSummaryArgs, reply *rpc.ImageSummaryReply) error {
	imageList, err := s.Manager.ListImage(ctx)
	if err != nil {
		return err
	}
	var images []rpc.Image
	for _, info := range imageList {
		images = append(images, rpc.Image{
			Timestamp: args.Timestamp,
			ImageID:   info.ID,
			Name:      info.Name,
			Tag:       info.Tag,
			Created:   info.Created,
			Size:      info.Size,
		})
	}
	reply.Data = images
	return nil
}

func (s *Service) NetworkSummary(ctx context.Context, args rpc.NetworkSummaryArgs, reply *rpc.NetworkSummaryReply) error {
	networkList, err := s.Manager.ListNetwork(ctx)
	if err != nil {
		return err
	}
	var networks []rpc.Network
	for _, info := range networkList {
		subnet := ""
		gateway := ""
		labels, _ := json.Marshal(info.Labels)
		if len(info.SubNet) > 0 {
			subnet = info.SubNet[0].Subnet
			gateway = info.SubNet[0].Gateway
		}
		networks = append(networks, rpc.Network{
			Timestamp: args.Timestamp,
			NetworkID: info.ID,
			Name:      info.Name,
			Driver:    info.Driver,
			Scope:     info.Scope,
			Created:   info.Created,
			Internal:  info.Internal,
			Subnet:    subnet,
			Gateway:   gateway,
			Labels:    string(labels),
		})
	}
	reply.Data = networks
	return nil
}

/************************* Host *************************/

func (s *Service) HostSummary(ctx context.Context, args rpc.HostSummaryArgs, reply *rpc.HostSummaryReply) error {
	info, err := psutil.GetSystemInfo()
	if err != nil {
		return err
	}
	reply.Data = rpc.Host{
		Timestamp:       args.Timestamp,
		Uptime:          info.Uptime,
		Hostname:        info.Hostname,
		Os:              info.Os,
		Platform:        info.Platform,
		PlatformVersion: info.PlatformVersion,
		KernelArch:      info.KernelArch,
		KernelVersion:   info.KernelVersion,
	}
	return nil
}

func (s *Service) CPUSummary(ctx context.Context, args rpc.CPUSummaryArgs, reply *rpc.CPUSummaryReply) error {
	cpuPercent, err := psutil.GetCPUPercent()
	if err != nil {
		return err
	}
	reply.Data = rpc.CPU{
		Timestamp:  args.Timestamp,
		CPUPercent: cpuPercent,
	}
	return nil
}

func (s *Service) MemorySummary(ctx context.Context, args rpc.MemorySummaryArgs, reply *rpc.MemorySummaryReply) error {
	percent, total, used, err := psutil.GetMemInfo()
	if err != nil {
		return err
	}
	reply.Data = rpc.Memory{
		Timestamp:  args.Timestamp,
		MemPercent: percent,
		MemTotal:   float64(total),
		MemUsed:    float64(used),
	}
	return nil
}

func (s *Service) DiskSummary(ctx context.Context, args rpc.DiskSummaryArgs, reply *rpc.DiskSummaryReply) error {
	info, err := psutil.GetDiskInfo(args.Devices)
	if err != nil {
		return err
	}
	ioMap, err := psutil.GetDiskIO(args.Devices)
	if err != nil {
		return err
	}
	var disks []rpc.Disk
	for device, info := range info {
		disk := rpc.Disk{
			Timestamp:   args.Timestamp,
			Device:      device,
			DiskPercent: info.Percent,
			DiskTotal:   float64(info.Total),
			DiskUsed:    float64(info.Used),
		}
		for dev, state := range ioMap {
			if dev == device {
				if latestReadBytes, ok := s.cache.Get(LatestDiskReadKey + device); ok {
					disk.DiskRead = float64((state.Read - latestReadBytes.(uint64)) / uint64(args.Interval))
					s.cache.Set(LatestDiskReadKey+device, state.Read, 0)
				} else {
					s.cache.Set(LatestDiskReadKey+device, state.Read, 0)
					disk.DiskRead = 0
				}
				if latestWriteBytes, ok := s.cache.Get(LatestDisKWriteKey + device); ok {
					disk.DiskWrite = float64((state.Write - latestWriteBytes.(uint64)) / uint64(args.Interval))
					s.cache.Set(LatestDisKWriteKey+device, state.Write, 0)
				} else {
					s.cache.Set(LatestDisKWriteKey+device, state.Write, 0)
					disk.DiskWrite = 0
				}
			}
		}
		disks = append(disks, disk)
	}
	reply.Data = disks
	return nil
}

func (s *Service) NetSummary(ctx context.Context, args rpc.NetSummaryArgs, reply *rpc.NetSummaryReply) error {
	netMap, err := psutil.GetNetworkIO(args.Ethernets)
	if err != nil {
		return err
	}
	var nets []rpc.Net
	for eth, info := range netMap {
		net := rpc.Net{
			Timestamp: args.Timestamp,
			Ethernet:  eth,
		}
		if LatestNetReceiveBytes, ok := s.cache.Get(LatestNetReceiveKey + eth); ok {
			net.NetRecv = float64((info.Recv - LatestNetReceiveBytes.(uint64)) / uint64(args.Interval))
			s.cache.Set(LatestNetReceiveKey+eth, info.Recv, 0)
		} else {
			s.cache.Set(LatestNetReceiveKey+eth, info.Recv, 0)
			net.NetRecv = 0
		}
		if LatestNetSendBytes, ok := s.cache.Get(LatestNetSendKey + eth); ok {
			net.NetSend = float64((info.Send - LatestNetSendBytes.(uint64)) / uint64(args.Interval))
			s.cache.Set(LatestNetSendKey+eth, info.Send, 0)
		} else {
			s.cache.Set(LatestNetSendKey+eth, info.Send, 0)
			net.NetSend = 0
		}
		nets = append(nets, net)
	}
	reply.Data = nets
	return nil
}
