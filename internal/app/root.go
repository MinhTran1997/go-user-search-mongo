package app

import (
	"github.com/core-go/log"
	"github.com/core-go/log/middleware"
	"github.com/core-go/service"
)

type Root struct {
	Server     service.ServerConf    `mapstructure:"server"`
	Mongo      MongoConfig     		 `mapstructure:"mongo"`
	Log        log.Config       	 `mapstructure:"log"`
	MiddleWare middleware.LogConfig  `mapstructure:"middleware"`
}

type MongoConfig struct {
	Uri      string `mapstructure:"uri"`
	Database string `mapstructure:"database"`
}