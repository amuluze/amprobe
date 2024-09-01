// Package service
// Date: 2024/3/6 13:34
// Author: Amu
// Description:
package service

import (
	"amprobe/pkg/database"
	"amprobe/pkg/rpc"
	"amprobe/service/task"
	"context"
	"fmt"
	"log/slog"
	"time"

	"github.com/amuluze/amutool/timex"
	"github.com/amuluze/docker"
)

type TimedTask struct {
	task   *task.Task
	ticker timex.Ticker
	stopCh chan struct{}
}

func NewTimedTask(conf *Config, db *database.DB, client *rpc.Client) *TimedTask {
	interval := conf.Task.Interval
	tk := timex.NewTicker(time.Duration(interval) * time.Second)
	manager, err := docker.NewManager()
	if err != nil {
		return nil
	}

	dev := make(map[string]struct{})
	for _, d := range conf.Task.Devices {
		dev[d] = struct{}{}
	}

	eth := make(map[string]struct{})
	for _, d := range conf.Task.Ethernets {
		eth[d] = struct{}{}
	}

	newTask := task.NewTask(interval, db, client, manager, dev, eth)

	return &TimedTask{
		task:   newTask,
		ticker: tk,
		stopCh: make(chan struct{}),
	}
}

func (a *TimedTask) Execute() {
	timestamp := time.Now()
	// ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	// defer cancel()
	ctx := context.TODO()
	// 处理 Host 指标
	go func() {
		if err := a.task.HostSummary(ctx, timestamp); err != nil {
			slog.Error("host summary failed: ", "error", err)
		}
		if err := a.task.CPUSummary(ctx, timestamp); err != nil {
			slog.Error("cpu summary failed: ", "error", err)
		}
		if err := a.task.MemorySummary(ctx, timestamp); err != nil {
			slog.Error("memory summary failed: ", "error", err)
		}
		if err := a.task.DiskSummary(ctx, timestamp); err != nil {
			slog.Error("disk summary failed: ", "error", err)
		}
		if err := a.task.NetSummary(ctx, timestamp); err != nil {
			slog.Error("net summary failed: ", "error", err)
		}
	}()

	// 处理 Docker 容器指标
	go func() {
		if err := a.task.DockerSummary(ctx, timestamp); err != nil {
			slog.Error("docker summary failed", "error", err)
		}
		if err := a.task.ContainerSummary(ctx, timestamp); err != nil {
			slog.Error("containers summary failed", "error", err)
		}
		if err := a.task.ImageSummary(ctx, timestamp); err != nil {
			slog.Error("image summary failed", "error", err)
		}
		if err := a.task.NetworkSummary(ctx, timestamp); err != nil {
			slog.Error("network summary failed", "error", err)
		}
	}()
}

func (a *TimedTask) Run() {
	for {
		select {
		case <-a.ticker.Chan():
			go a.Execute()
		case <-a.stopCh:
			fmt.Println("task exit")
			return
		}
	}
}

func (a *TimedTask) Stop() {
	close(a.stopCh)
}
