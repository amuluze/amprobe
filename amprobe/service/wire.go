//go:build wireinject
// +build wireinject

package service

import (
	"amprobe/service/account"
	"amprobe/service/alarm"
	"amprobe/service/audit"
	"amprobe/service/auth"
	"amprobe/service/container"
	"amprobe/service/host"
	"amprobe/service/mail"
	"amprobe/service/model"

	"github.com/google/wire"
)

func BuildInjector(configFile string, modelFile ModeConf) (*Injector, func(), error) {
	wire.Build(
		NewConfig,
		NewLogger,
		NewDB,
		NewRPCClient,
		InitAuthStore,
		InitAuth,
		InitAdapter,
		InitCasbin,
		container.Set,
		host.Set,
		model.Set,
		auth.Set,
		audit.Set,
		account.Set,
		alarm.Set,
		mail.Set,
		NewLoggerHandler,
		NewTermHandler,
		NewTimedTask,
		RouterSet,
		NewFiberApp,
		PrepareSet,
		InjectorSet,
	)
	return new(Injector), nil, nil
}
