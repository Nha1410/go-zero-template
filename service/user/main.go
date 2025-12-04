package main

import (
	"flag"
	"fmt"

	"github.com/Nha1410/go-zero-template/service/user/internal/config"
	// "github.com/Nha1410/go-zero-template/service/user/internal/handler" // Will be used after proto generation
	"github.com/Nha1410/go-zero-template/service/user/internal/svc"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var configFile = flag.String("f", "etc/user.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)

	svcCtx := svc.NewServiceContext(c)
	_ = svcCtx // Will be used after proto generation

	s := zrpc.MustNewServer(c.RpcServerConf, func(grpcServer *grpc.Server) {
		// Register your gRPC services here
		// This will be auto-generated after running: ./scripts/generate-service.sh user
		// handler.RegisterHandlers(grpcServer, svcCtx)

		if c.Mode == "dev" {
			reflection.Register(grpcServer)
		}
	})
	defer s.Stop()

	fmt.Printf("Starting rpc server at %s...\n", c.ListenOn)
	s.Start()
}
