// Package task
// Date: 2022/11/9 10:18
// Author: Amu
// Description:
package task

import (
	"amprobe/pkg/database"
	"amprobe/pkg/rpc"
	"context"
	"time"
	
	"github.com/amuluze/docker"
	"github.com/patrickmn/go-cache"
)

var _ ITask = (*Task)(nil)

type ITask interface {
	DockerSummary(context.Context, time.Time) error
	ContainerSummary(context.Context, time.Time) error
	ImageSummary(context.Context, time.Time) error
	NetworkSummary(context.Context, time.Time) error
	
	HostSummary(context.Context, time.Time) error
	CPUSummary(context.Context, time.Time) error
	MemorySummary(context.Context, time.Time) error
	DiskSummary(context.Context, time.Time) error
	NetSummary(context.Context, time.Time) error
}
type Task struct {
	interval  int
	db        *database.DB
	rpcClient *rpc.Client
	manager   *docker.Manager
	devices   map[string]struct{}
	ethernet  map[string]struct{}
	cache     *cache.Cache
}

func NewTask(interval int, db *database.DB, client *rpc.Client, manager *docker.Manager, dev map[string]struct{}, eth map[string]struct{}) *Task {
	return &Task{
		interval:  interval,
		db:        db,
		rpcClient: client,
		manager:   manager,
		devices:   dev,
		ethernet:  eth,
		cache:     cache.New(5*time.Minute, 60*time.Second),
	}
}
