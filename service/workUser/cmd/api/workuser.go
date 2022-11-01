package main

import (
	"flag"
	"fmt"
	"github.com/suhanyujie/workUser/service/workUser/model"

	"github.com/suhanyujie/workUser/service/workUser/cmd/api/internal/config"
	"github.com/suhanyujie/workUser/service/workUser/cmd/api/internal/handler"
	"github.com/suhanyujie/workUser/service/workUser/cmd/api/internal/svc"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/rest"
)

var configFile = flag.String("f", "etc/workUser-api.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)

	ctx := svc.NewServiceContext(c)
	server := rest.MustNewServer(c.RestConf)
	defer server.Stop()

	handler.RegisterHandlers(server, ctx)

	// 数据库初始化的操作
	if err := model.Initial(c.Mysql.DataSource); err != nil {
		fmt.Printf("%+v", err)
		return
	}

	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	server.Start()
}
