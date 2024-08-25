// Package service
// Date: 2024/3/6 13:34
// Author: Amu
// Description:
package service

import (
	"amprobe/pkg/database"
	"amprobe/service/task"
	"fmt"
	"time"
	
	"github.com/amuluze/amutool/timex"
	"github.com/amuluze/docker"
)

type TimedTask struct {
	task   *task.Task
	ticker timex.Ticker
	stopCh chan struct{}
}

func NewTimedTask(conf *Config, db *database.DB) *TimedTask {
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
	
	newTask := task.NewTask(interval, db, manager, dev, eth)
	
	return &TimedTask{
		task:   newTask,
		ticker: tk,
		stopCh: make(chan struct{}),
	}
}

func (a *TimedTask) Execute() {
	timestamp := time.Now()
	// 处理数组指标
	go a.task.Host(timestamp)
	go a.task.Cpu(timestamp)
	go a.task.Memory(timestamp)
	go a.task.Disk(timestamp)
	go a.task.Network(timestamp)
	
	// 处理 Docker 容器指标
	go a.task.Container(timestamp)
	go func() {
		a.task.Docker(timestamp)
		a.task.Image(timestamp)
		a.task.Net(timestamp)
	}()
	
	//go a.task.ClearOldRecord()
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
