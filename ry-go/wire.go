//go:build wireinject
// +build wireinject

package main

import (
	"gorm.io/gorm"
	"ry-go/business/controller"
	"ry-go/business/dao/daoImpl"
	"ry-go/business/service/serviceImpl"
	"ry-go/router"
	"ry-go/utils"
	"sync"

	"github.com/go-redis/redis/v8"

	"github.com/google/wire"
	"github.com/labstack/echo/v4"
)

func WireApp(*gorm.DB, *redis.Client, *utils.RedisStore, sync.Map) (*echo.Echo, func(), error) {
	panic(wire.Build(
		daoImpl.ProviderSet,
		serviceImpl.ProviderSet,
		controller.ProviderSet,
		router.ProviderSet,
	))
}
