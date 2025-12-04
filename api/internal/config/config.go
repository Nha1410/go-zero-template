package config

import (
	"github.com/Nha1410/go-zero-template/common/auth"
	redisCache "github.com/Nha1410/go-zero-template/common/cache"
	"github.com/Nha1410/go-zero-template/common/database"
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	rest.RestConf
	Database struct {
		Postgres database.PostgresConfig
		MySQL    database.MySQLConfig
		Type     string `json:"type"` // postgres or mysql
	}
	Redis    redisCache.RedisConfig
	RabbitMQ struct {
		Host     string
		Port     int
		User     string
		Password string
		VHost    string `json:"vhost"`
	}
	Zitadel auth.ZitadelConfig
	UserRpc zrpc.RpcClientConf
}
