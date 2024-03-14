// Package api
// Date: 2024/3/6 12:49
// Author: Amu
// Description:
package api

import (
	"github.com/amuluze/amprobe/pkg/fiberx"
	"github.com/amuluze/amprobe/service/container/service"
	"github.com/gofiber/fiber/v2"
)

type ContainerAPI struct {
	ContainerService service.IContainerService
}

func NewContainerAPI(containerService service.IContainerService) *ContainerAPI {
	return &ContainerAPI{
		ContainerService: containerService,
	}
}

func (a *ContainerAPI) ContainerList(ctx *fiber.Ctx) error {
	c := ctx.UserContext()
	
	container, err := a.ContainerService.ContainerList(c)
	if err != nil {
		return fiberx.Failure(ctx, err)
	}
	return fiberx.Success(ctx, container)
}
