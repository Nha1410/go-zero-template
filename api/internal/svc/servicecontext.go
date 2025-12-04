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
	Zitadel  *auth.ZitadelClient
	// UserRpc will be available after generating proto
	// Uncomment after running: ./scripts/generate-service.sh user
	// UserRpc    userclient.User
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
	redisClient, err := cache.NewRedisClient(c.Redis)
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

	// Initialize Zitadel
	zitadelClient, err := auth.NewZitadelClient(c.Zitadel)
	if err != nil {
		logx.Errorf("Failed to initialize Zitadel client: %v", err)
		panic(err)
	}

	// Initialize gRPC clients
	// Note: userclient will be generated from proto file
	// Uncomment after generating proto:
	// userRpc := userclient.NewUser(zrpc.MustNewClient(c.UserRpc).Conn())
	// Uncomment UserRpc in struct definition above as well

	return &ServiceContext{
		Config:   c,
		DB:       db,
		Redis:    redisClient,
		RabbitMQ: rabbitmqClient,
		Zitadel:  zitadelClient,
		// UserRpc will be initialized after generating proto
		// UserRpc:  userRpc,
	}
}
