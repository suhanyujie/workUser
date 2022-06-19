package main

import (
	"flag"
	"fmt"

	"github.com/suhanyujie/workUser/service/workUser/cmd/rpc/user/internal/config"
	"github.com/suhanyujie/workUser/service/workUser/cmd/rpc/user/internal/server"
	"github.com/suhanyujie/workUser/service/workUser/cmd/rpc/user/internal/svc"
	"github.com/suhanyujie/workUser/service/workUser/cmd/rpc/user/user"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

var configFile = flag.String("f", "etc/user.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)
	ctx := svc.NewServiceContext(c)
	srv := server.NewUserServer(ctx)

	s := zrpc.MustNewServer(c.RpcServerConf, func(grpcServer *grpc.Server) {
		user.RegisterUserServer(grpcServer, srv)
	})
	defer s.Stop()

	fmt.Printf("Starting rpc server at %s...\n", c.ListenOn)
	s.Start()
}
