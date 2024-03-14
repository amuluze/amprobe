// Package service
// Date: 2024/3/6 12:49
// Author: Amu
// Description:
package service

import (
	"context"
	"github.com/amuluze/amprobe/pkg/errors"
	"github.com/amuluze/amprobe/service/container/repository"
	"github.com/amuluze/amprobe/service/schema"
	"github.com/google/wire"
	"time"
)

var ContainerServiceSet = wire.NewSet(NewContainerService, wire.Bind(new(IContainerService), new(*ContainerService)))

type IContainerService interface {
	ContainerList(ctx context.Context) (*schema.ContainerQueryRely, error)
}

type ContainerService struct {
	ContainerRepo repository.IContainerRepo
}

func NewContainerService(containerRepo repository.IContainerRepo) *ContainerService {
	return &ContainerService{ContainerRepo: containerRepo}
}

func (a *ContainerService) ContainerList(ctx context.Context) (*schema.ContainerQueryRely, error) {
	mContainers, err := a.ContainerRepo.ContainerList(ctx)
	if err != nil {
		return nil, errors.New400Errors(err.Error())
	}
	var list []schema.Container
	var ts time.Time
	for _, item := range mContainers {
		if len(list) == 0 {
			ts = item.Timestamp
			list = append(list, schema.Container{
				ID:            item.ContainerID[:6],
				Name:          item.Name,
				Image:         item.Image,
				State:         item.State,
				Uptime:        item.Uptime,
				IP:            item.IP,
				CPUPercent:    item.CPUPercent,
				MemoryPercent: item.MemPercent,
				MemoryLimit:   item.MemLimit,
				MemoryUsage:   item.MemUsage,
			})
		} else {
			if ts != item.Timestamp {
				continue
			}
			list = append(list, schema.Container{
				ID:            item.ContainerID[:6],
				Name:          item.Name,
				Image:         item.Image,
				State:         item.State,
				Uptime:        item.Uptime,
				IP:            item.IP,
				CPUPercent:    item.CPUPercent,
				MemoryPercent: item.MemPercent,
				MemoryLimit:   item.MemLimit,
				MemoryUsage:   item.MemUsage,
			})
		}
	}
	return &schema.ContainerQueryRely{Containers: list}, nil
}
