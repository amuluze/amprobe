// Package task
// Date: 2022/11/9 10:18
// Author: Amu
// Description:
package task

import (
	"time"

	"github.com/amuluze/docker"
	"github.com/patrickmn/go-cache"
)

const (
	LatestDiskReadKey   = "latest_disk_io_read_"
	LatestDisKWriteKey  = "latest_disk_io_write_"
	LatestNetReceiveKey = "latest_net_io_receive_"
	LatestNetSendKey    = "latest_net_io_send_"
)

type Task struct {
	interval int
	manager  *docker.Manager
	cache    *cache.Cache
}

func NewTask(interval int, manager *docker.Manager) *Task {
	return &Task{
		interval: interval,
		manager:  manager,
		cache:    cache.New(5*time.Minute, 60*time.Second),
	}
}
