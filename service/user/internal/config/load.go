package config

import (
	"time"

	envConfig "github.com/Nha1410/go-zero-template/common/config"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/service"
	"github.com/zeromicro/go-zero/zrpc"
)

func LoadFromEnv() Config {
	c := Config{
		RpcServerConf: zrpc.RpcServerConf{
			ServiceConf: service.ServiceConf{
				Name: envConfig.GetString("USER_SERVICE_NAME", "user-service"),
				Log: logx.LogConf{
					ServiceName: envConfig.GetString("USER_SERVICE_NAME", "user-service"),
					Mode:        envConfig.GetString("LOG_MODE", "file"),
					Path:        envConfig.GetString("LOG_PATH", "logs"),
					Level:       envConfig.GetString("LOG_LEVEL", "info"),
					Compress:    envConfig.GetBool("LOG_COMPRESS", true),
					KeepDays:    envConfig.GetInt("LOG_KEEP_DAYS", 7),
				},
			},
			ListenOn: envConfig.GetString("USER_SERVICE_LISTEN_ON", "0.0.0.0:9000"),
		},
	}

	c.Mode = envConfig.GetString("USER_SERVICE_MODE", "dev")
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

	c.AppRedis.Host = envConfig.GetString("REDIS_HOST", "localhost")
	c.AppRedis.Port = envConfig.GetInt("REDIS_PORT", 6379)
	c.AppRedis.Password = envConfig.GetString("REDIS_PASSWORD", "")
	c.AppRedis.DB = envConfig.GetInt("REDIS_DB", 0)
	c.AppRedis.PoolSize = envConfig.GetInt("REDIS_POOL_SIZE", 10)

	c.RabbitMQ.Host = envConfig.GetString("RABBITMQ_HOST", "localhost")
	c.RabbitMQ.Port = envConfig.GetInt("RABBITMQ_PORT", 5672)
	c.RabbitMQ.User = envConfig.GetString("RABBITMQ_USER", "guest")
	c.RabbitMQ.Password = envConfig.GetString("RABBITMQ_PASSWORD", "guest")
	c.RabbitMQ.VHost = envConfig.GetString("RABBITMQ_VHOST", "/")

	c.Prometheus.Host = envConfig.GetString("PROMETHEUS_HOST", "0.0.0.0")
	c.Prometheus.Port = envConfig.GetInt("PROMETHEUS_PORT", 9092)
	c.Prometheus.Path = envConfig.GetString("PROMETHEUS_PATH", "/metrics")

	return c
}
