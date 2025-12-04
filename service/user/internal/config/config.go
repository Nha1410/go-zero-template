package config

import (
	"github.com/Nha1410/go-zero-template/common/database"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	zrpc.RpcServerConf
	Database struct {
		Postgres database.PostgresConfig
		MySQL    database.MySQLConfig
		Type     string `json:",default=postgres"` // postgres or mysql
	}
	Redis struct {
		Host     string
		Port     int
		Password string
		DB       int `json:",default=0"`
		PoolSize int `json:",default=10"`
	}
	RabbitMQ struct {
		Host     string
		Port     int
		User     string
		Password string
		VHost    string `json:",default=/"`
	}
}
