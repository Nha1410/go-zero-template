package main

import (
	"fmt"

	envConfig "github.com/Nha1410/go-zero-template/common/config"
	"github.com/Nha1410/go-zero-template/service/user/internal/config"
	"github.com/Nha1410/go-zero-template/service/user/internal/svc"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	_ = envConfig.LoadEnv()
	c := config.LoadFromEnv()

	svcCtx := svc.NewServiceContext(c)
	_ = svcCtx

	s := zrpc.MustNewServer(c.RpcServerConf, func(grpcServer *grpc.Server) {

		if c.Mode == "dev" {
			reflection.Register(grpcServer)
		}
	})
	defer s.Stop()

	fmt.Printf("Starting rpc server at %s...\n", c.ListenOn)
	s.Start()
}
