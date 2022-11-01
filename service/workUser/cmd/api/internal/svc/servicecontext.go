package svc

import (
	"github.com/suhanyujie/workUser/service/workUser/cmd/api/internal/config"
	"github.com/suhanyujie/workUser/service/workUser/model"
)

type ServiceContext struct {
	Config        config.Config
	WorkUserModel model.WorkUserModelIf
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:        c,
		WorkUserModel: model.NewWorkUserModel(),
	}
}
