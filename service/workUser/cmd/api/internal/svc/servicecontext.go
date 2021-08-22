package svc

import (
	"github.com/suhanyujie/workUser/service/workUser/cmd/api/internal/config"
	"github.com/suhanyujie/workUser/service/workUser/model"
	"github.com/tal-tech/go-zero/core/stores/sqlx"
)

type ServiceContext struct {
	Config        config.Config
	WorkUserModel model.WorkUserModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	conn := sqlx.NewSqlConn("mysql", c.Mysql.DataSource)
	return &ServiceContext{
		Config:        c,
		WorkUserModel: model.NewWorkUserModel(conn, nil),
	}
}
