package main

import (
	"fmt"

	"github.com/Nha1410/go-zero-template/api/internal/config"
	"github.com/Nha1410/go-zero-template/api/internal/handler"
	"github.com/Nha1410/go-zero-template/api/internal/svc"
	envConfig "github.com/Nha1410/go-zero-template/common/config"

	"github.com/zeromicro/go-zero/rest"
)

func main() {
	_ = envConfig.LoadEnv()
	c := config.LoadFromEnv()

	server := rest.MustNewServer(c.RestConf)
	defer server.Stop()

	ctx := svc.NewServiceContext(c)
	handler.RegisterHandlers(server, ctx)

	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	server.Start()
}
