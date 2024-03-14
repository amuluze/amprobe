// Package docker
// Date: 2024/3/6 10:55
// Author: Amu
// Description:
package docker

import (
	"context"
	"fmt"
	"io"
	"strings"

	"github.com/docker/docker/api/types/container"

	"github.com/buger/jsonparser"
	"github.com/docker/docker/client"
)

type Manager struct {
	*client.Client
}

func NewManager() (*Manager, error) {
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		return nil, err
	}
	return &Manager{Client: cli}, nil
}

type ContainerInfo struct {
	ID         string
	Name       string
	State      string
	Image      string
	IP         string
	Uptime     string
	CPUPercent float64
	MemPercent float64
	MemUsage   float64
	MemLimit   float64
}

func (m *Manager) GetContainerInfos() ([]ContainerInfo, error) {
	var infos []ContainerInfo
	ctx := context.Background()
	containers, err := m.Client.ContainerList(ctx, container.ListOptions{All: true})
	if err != nil {
		return infos, err
	}
	for _, cn := range containers {
		var info ContainerInfo
		info.ID = cn.ID
		info.Name = strings.Trim(cn.Names[0], "/")
		info.Image = cn.Image
		// TODO： 获取 IP 的方式有点问题，应该采用遍历 cn.NetworkSettings.Networks map 结构的方式去获取
		//info.IP = cn.NetworkSettings.Networks[cn.HostConfig.NetworkMode].IPAddress
		for _, nt := range cn.NetworkSettings.Networks {
			if nt.IPAddress != "" {
				info.IP = nt.IPAddress
				break
			}
		}

		// 容器状态处理
		state := cn.State
		inspect, err := m.Client.ContainerInspect(ctx, cn.ID)
		if err == nil {
			if inspect.ContainerJSONBase.State.Health != nil && inspect.ContainerJSONBase.State.Health.Status == "healthy" {
				info.State = "running"
			}
		}
		info.State = state
		info.Uptime = cn.Status

		// 容器资源指标处理
		stats, err := m.Client.ContainerStats(ctx, cn.ID, false)
		if err != nil {
			continue
		}
		content, err := io.ReadAll(stats.Body)
		if err != nil {
			continue
		}
		cpuTotalUsage, _ := jsonparser.GetFloat(content, "cpu_stats", "cpu_usage", "total_usage")
		cpuSystemUsage, _ := jsonparser.GetFloat(content, "cpu_stats", "system_cpu_usage")
		// TODO: 这里先这样简单计算，后续有机会优化
		cpuPercent := (cpuTotalUsage / cpuSystemUsage) * 100.0

		memUsd, _ := jsonparser.GetFloat(content, "memory_stats", "usage")
		memLimit, _ := jsonparser.GetFloat(content, "memory_stats", "limit")
		memPercent := (memUsd / memLimit) * 100.0
		info.CPUPercent = cpuPercent
		info.MemPercent = memPercent
		info.MemUsage = memUsd
		info.MemLimit = memLimit
		infos = append(infos, info)
	}
	return infos, err
}

// List 获取容器列表
func (m *Manager) List() {
	ctx := context.Background()
	containers, err := m.Client.ContainerList(ctx, container.ListOptions{All: true})
	if err != nil {
		panic(err)
	}
	for _, cn := range containers {
		fmt.Printf("container id: %s, name: %s, state: %#v\n", cn.ID, strings.Trim(cn.Names[0], "/"), cn.State)
	}
}

func (m *Manager) Inspect(containerID string) {
	ctx := context.Background()
	info, err := m.Client.ContainerInspect(ctx, containerID)
	if err != nil {
		panic(err)
	}
	fmt.Printf("inspect: %#v\n", info.ContainerJSONBase.State.Health.Status)
}

func (m *Manager) Stats(containerID string) {
	ctx := context.Background()
	stats, err := m.Client.ContainerStats(ctx, containerID, false)
	if err != nil {
		panic(err)
	}
	fmt.Printf("stats: %#v\n", stats)
	defer stats.Body.Close()
	content, err := io.ReadAll(stats.Body)
	fmt.Printf("content: %s\n", string(content))
}
