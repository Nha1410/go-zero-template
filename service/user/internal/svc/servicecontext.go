package svc

import (
	"database/sql"

	"github.com/Nha1410/go-zero-template/common/cache"
	"github.com/Nha1410/go-zero-template/common/database"
	"github.com/Nha1410/go-zero-template/common/queue"
	"github.com/Nha1410/go-zero-template/service/user/internal/config"
	domainRepo "github.com/Nha1410/go-zero-template/service/user/internal/domain/repository"
	"github.com/Nha1410/go-zero-template/service/user/internal/repository"
	"github.com/Nha1410/go-zero-template/service/user/internal/usecase"

	"github.com/zeromicro/go-zero/core/logx"
)

type ServiceContext struct {
	Config      config.Config
	DB          *sql.DB
	Redis       *cache.RedisClient
	RabbitMQ    *queue.RabbitMQClient
	UserRepo    domainRepo.UserRepository
	UserUsecase *usecase.UserUsecase
}

func NewServiceContext(c config.Config) *ServiceContext {
	// Initialize database
	var db *sql.DB
	var err error
	if c.Database.Type == "mysql" {
		db, err = database.NewMySQLConnection(c.Database.MySQL)
	} else {
		db, err = database.NewPostgresConnection(c.Database.Postgres)
	}
	if err != nil {
		logx.Errorf("Failed to connect to database: %v", err)
		panic(err)
	}

	// Initialize Redis
	redisClient, err := cache.NewRedisClient(cache.RedisConfig{
		Host:     c.Redis.Host,
		Port:     c.Redis.Port,
		Password: c.Redis.Password,
		DB:       c.Redis.DB,
		PoolSize: c.Redis.PoolSize,
	})
	if err != nil {
		logx.Errorf("Failed to connect to Redis: %v", err)
		panic(err)
	}

	// Initialize RabbitMQ
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

	// Initialize repositories
	userRepo := repository.NewUserRepo(db)

	// Initialize use cases
	userUsecase := usecase.NewUserUsecase(userRepo)

	return &ServiceContext{
		Config:      c,
		DB:          db,
		Redis:       redisClient,
		RabbitMQ:    rabbitmqClient,
		UserRepo:    userRepo,
		UserUsecase: userUsecase,
	}
}
