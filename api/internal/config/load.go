package config

import (
	"time"

	envConfig "github.com/Nha1410/go-zero-template/common/config"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/service"
	"github.com/zeromicro/go-zero/rest"
)

func LoadFromEnv() Config {
	c := Config{
		RestConf: rest.RestConf{
			ServiceConf: service.ServiceConf{
				Name: envConfig.GetString("API_NAME", "api-gateway"),
				Log: logx.LogConf{
					ServiceName: envConfig.GetString("API_NAME", "api-gateway"),
					Mode:        envConfig.GetString("LOG_MODE", "file"),
					Path:        envConfig.GetString("LOG_PATH", "logs"),
					Level:       envConfig.GetString("LOG_LEVEL", "info"),
					Compress:    envConfig.GetBool("LOG_COMPRESS", true),
					KeepDays:    envConfig.GetInt("LOG_KEEP_DAYS", 7),
				},
			},
			Host: envConfig.GetString("API_HOST", "0.0.0.0"),
			Port: envConfig.GetInt("API_PORT", 8888),
		},
	}

	c.Database.Type = envConfig.GetString("DATABASE_TYPE", "postgres")
	c.Database.Postgres.Host = envConfig.GetString("DATABASE_HOST", "localhost")
	c.Database.Postgres.Port = envConfig.GetInt("DATABASE_PORT", 5432)
	c.Database.Postgres.User = envConfig.GetString("DATABASE_USER", "postgres")
	c.Database.Postgres.Password = envConfig.GetString("DATABASE_PASSWORD", "postgres")
	c.Database.Postgres.Database = envConfig.GetString("DATABASE_NAME", "gozero_template")
	c.Database.Postgres.SSLMode = envConfig.GetString("DATABASE_SSLMODE", "disable")
	c.Database.Postgres.MaxOpenConns = envConfig.GetInt("DATABASE_MAX_OPEN_CONNS", 100)
	c.Database.Postgres.MaxIdleConns = envConfig.GetInt("DATABASE_MAX_IDLE_CONNS", 10)
	c.Database.Postgres.ConnMaxLifetime = time.Duration(envConfig.GetInt("DATABASE_CONN_MAX_LIFETIME", 3600)) * time.Second
	c.Database.Postgres.ConnMaxIdleTime = time.Duration(envConfig.GetInt("DATABASE_CONN_MAX_IDLE_TIME", 600)) * time.Second

	c.Redis.Host = envConfig.GetString("REDIS_HOST", "localhost")
	c.Redis.Port = envConfig.GetInt("REDIS_PORT", 6379)
	c.Redis.Password = envConfig.GetString("REDIS_PASSWORD", "")
	c.Redis.DB = envConfig.GetInt("REDIS_DB", 0)
	c.Redis.PoolSize = envConfig.GetInt("REDIS_POOL_SIZE", 10)

	c.RabbitMQ.Host = envConfig.GetString("RABBITMQ_HOST", "localhost")
	c.RabbitMQ.Port = envConfig.GetInt("RABBITMQ_PORT", 5672)
	c.RabbitMQ.User = envConfig.GetString("RABBITMQ_USER", "guest")
	c.RabbitMQ.Password = envConfig.GetString("RABBITMQ_PASSWORD", "guest")
	c.RabbitMQ.VHost = envConfig.GetString("RABBITMQ_VHOST", "/")

	c.Zitadel.Issuer = envConfig.GetString("ZITADEL_ISSUER", "")
	c.Zitadel.ClientID = envConfig.GetString("ZITADEL_CLIENT_ID", "")
	c.Zitadel.ClientSecret = envConfig.GetString("ZITADEL_CLIENT_SECRET", "")
	scopes := envConfig.GetStringSlice("ZITADEL_SCOPES", []string{"openid", "profile", "email"})
	if len(scopes) > 0 {
		c.Zitadel.Scopes = scopes
	}

	rpcHost := envConfig.GetString("USER_RPC_HOST", "localhost:9000")
	c.UserRpc.Etcd.Hosts = []string{rpcHost}
	c.UserRpc.Etcd.Key = envConfig.GetString("USER_RPC_KEY", "user.rpc")
	c.UserRpc.Timeout = int64(envConfig.GetInt("USER_RPC_TIMEOUT", 5000))
	c.UserRpc.NonBlock = envConfig.GetBool("USER_RPC_NON_BLOCK", true)

	return c
}
