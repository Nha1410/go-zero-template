package config

import (
	"github.com/Nha1410/go-zero-template/common/database"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	zrpc.RpcServerConf `yaml:",inline"`
	Database           struct {
		Postgres database.PostgresConfig
		Type     string `yaml:",default=postgres"`
	}
	AppRedis struct {
		Host     string
		Port     int
		Password string
		DB       int `yaml:",default=0"`
		PoolSize int `yaml:",default=10"`
	} `yaml:"AppRedis" json:"AppRedis"`
	RabbitMQ struct {
		Host     string
		Port     int
		User     string
		Password string
		VHost    string `yaml:",default=/"`
	}
}
