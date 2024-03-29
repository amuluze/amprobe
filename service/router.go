// Package service
// Date: 2024/3/6 11:09
// Author: Amu
// Description:
package service

import (
	"fmt"
	"github.com/amuluze/amprobe/pkg/auth"
	"github.com/gofiber/contrib/websocket"
	"github.com/gofiber/fiber/v2"

	"github.com/google/wire"

	authAPI "github.com/amuluze/amprobe/service/auth/api"
	containerAPI "github.com/amuluze/amprobe/service/container/api"
	hostAPI "github.com/amuluze/amprobe/service/host/api"
)

var RouterSet = wire.NewSet(wire.Struct(new(Router), "*"), wire.Bind(new(IRouter), new(*Router)))

var _ IRouter = (*Router)(nil)

type IRouter interface {
	Register(app *fiber.App) error
	Prefixes() []string
}

type Router struct {
	config *Config
	auth   auth.Auther

	containerAPI *containerAPI.ContainerAPI
	hostAPI      *hostAPI.HostAPI
	authAPI      *authAPI.AuthAPI

	loggerHandler *LoggerHandler
}

func (a *Router) RegisterAPI(app *fiber.App) {
	// 以下是 websocket service

	// 以下是 http 服务
	g := app.Group("/api")

	//if a.config.Auth.Enable {
	//	g.Use(middleware.UserAuthMiddleware(
	//		a.auth,
	//		middleware.AllowPathPrefixSkipper("/api/v1/auth/login"),
	//	))
	//}

	v1 := g.Group("v1")
	{
		gIndex := v1.Group("index")
		{
			gIndex.Get("/index", func(c *fiber.Ctx) error {
				fmt.Println("i am here")
				return c.SendString("hello world")
			})
		}

		gAuth := v1.Group("auth")
		{
			gAuth.Post("/login", a.authAPI.Login).Name("登录")
			gAuth.Post("/logout", a.authAPI.Logout).Name("登出")
			gAuth.Post("/pass_update", a.authAPI.PassUpdate).Name("更新密码")
			gAuth.Post("/token_update", a.authAPI.TokenUpdate).Name("更新 token")
		}

		gContainer := v1.Group("container")
		{
			gContainer.Get("/containers", a.containerAPI.ContainerList)
			gContainer.Get("/images", a.containerAPI.ImageList)
			gContainer.Get("/version", a.containerAPI.Version)
		}

		gHost := v1.Group("host")
		{
			gHost.Get("/host_info", a.hostAPI.HostInfo)
			gHost.Get("/cpu_info", a.hostAPI.CPUInfo)
			gHost.Get("/mem_info", a.hostAPI.MemInfo)
			gHost.Get("/disk_info", a.hostAPI.DiskInfo)
			gHost.Get("cpu_trending", a.hostAPI.CPUUsage)
			gHost.Get("mem_trending", a.hostAPI.MemUsage)
			gHost.Get("disk_trending", a.hostAPI.DiskUsage)
			gHost.Get("net_trending", a.hostAPI.NetUsage)
		}
	}

	app.Use("/ws/:id", websocket.New(a.loggerHandler.Handler))
}

func (a *Router) Register(app *fiber.App) error {
	a.RegisterAPI(app)
	return nil
}
func (a *Router) Prefixes() []string {
	return []string{"/api/"}
}
