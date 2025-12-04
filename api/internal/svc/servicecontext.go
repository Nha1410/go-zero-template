package svc

import (
	"database/sql"

	"github.com/Nha1410/go-zero-template/api/internal/config"
	"github.com/Nha1410/go-zero-template/common/auth"
	"github.com/Nha1410/go-zero-template/common/cache"
	"github.com/Nha1410/go-zero-template/common/database"
	"github.com/Nha1410/go-zero-template/common/queue"
	"github.com/zeromicro/go-zero/core/logx"
)

type ServiceContext struct {
	Config   config.Config
	DB       *sql.DB
	Redis    *cache.RedisClient
	RabbitMQ *queue.RabbitMQClient
	Zitadel *auth.ZitadelClient
}

func NewServiceContext(c config.Config) *ServiceContext {
	db, err := database.NewPostgresConnection(c.Database.Postgres)
	if err != nil {
		logx.Errorf("Failed to connect to database: %v", err)
		panic(err)
	}

	redisClient, err := cache.NewRedisClient(c.Redis)
	if err != nil {
		logx.Errorf("Failed to connect to Redis: %v", err)
		panic(err)
	}

	rabbitmqClient, err := queue.NewRabbitMQClient(queue.RabbitMQConfig{
		Host:     c.RabbitMQ.Host,
		Port:     c.RabbitMQ.Port,
		User:     c.RabbitMQ.User,
		Password: c.RabbitMQ.Password,
		VHost:    c.RabbitMQ.VHost,
	})
	if err != nil {
		logx.Errorf("Failed to connect to RabbitMQ: %v", err)
		panic(err)
	}

	zitadelClient, err := auth.NewZitadelClient(c.Zitadel)
	if err != nil {
		logx.Errorf("Failed to initialize Zitadel client: %v", err)
		panic(err)
	}

	return &ServiceContext{
		Config:   c,
		DB:       db,
		Redis:    redisClient,
		RabbitMQ: rabbitmqClient,
		Zitadel:  zitadelClient,
	}
}
