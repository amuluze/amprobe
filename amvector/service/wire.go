//go:build wireinject
// +build wireinject

package service

import (
	"github.com/google/wire"
)

func BuildInjector(configFile string, prefix Prefix) (*Injector, func(), error) {
	wire.Build(
		NewConfig,
		NewLogger,
		NewRPCServer,
		InjectorSet,
	)
	return new(Injector), nil, nil
}
