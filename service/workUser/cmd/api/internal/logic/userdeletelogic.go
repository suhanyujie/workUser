package logic

import (
	"context"

	"github.com/suhanyujie/workUser/service/workUser/cmd/api/internal/svc"
	"github.com/suhanyujie/workUser/service/workUser/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserDeleteLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) UserDeleteLogic {
	return UserDeleteLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserDeleteLogic) UserDelete(req types.Request) (*types.Response, error) {
	// todo: add your logic here and delete this line

	return &types.Response{}, nil
}
