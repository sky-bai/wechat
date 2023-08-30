//go:build wireinject
// +build wireinject

// The build tag makes sure the stub is not built in the final build.

package main

import (
	"wx-base/internal/biz"
	"wx-base/internal/conf"
	"wx-base/internal/data"
	"wx-base/internal/data/mini_program_data"
	"wx-base/internal/data/offiaiacl_account_data"
	"wx-base/internal/server"
	"wx-base/internal/service"

	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
)

// wireApp init kratos application.
func wireApp(*conf.Server, *conf.Data, log.Logger) (*kratos.App, func(), error) {
	panic(wire.Build(server.ProviderSet, data.ProviderSet, mini_program_data.ProviderSet, offiaiacl_account_data.ProviderSet, biz.ProviderSet, service.ProviderSet, newApp))
}
