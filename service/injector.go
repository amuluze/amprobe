// Package service
// Date: 2024/3/6 11:08
// Author: Amu
// Description:
package service

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/wire"
)

var InjectorSet = wire.NewSet(NewInjector)

type Injector struct {
	App    *fiber.App
	Router *Router
	Config *Config
	Task   *TimedTask
}

func NewInjector(app *fiber.App, router *Router, config *Config, task *TimedTask) (*Injector, error) {
	return &Injector{
		App:    app,
		Router: router,
		Config: config,
		Task:   task,
	}, nil
}
