package config

import (
	"github.com/Nha1410/go-zero-template/common/database"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	zrpc.RpcServerConf
	Database struct {
		Postgres database.PostgresConfig
		Type     string
	}
	AppRedis struct {
		Host     string
		Port     int
		Password string
		DB       int
		PoolSize int
	}
	RabbitMQ struct {
		Host     string
		Port     int
		User     string
		Password string
		VHost    string
	}
}
