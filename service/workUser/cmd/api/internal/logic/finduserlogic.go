package logic

import (
	"context"
	"github.com/suhanyujie/workUser/service/workUser/cmd/api/internal/utils/constvar"

	"github.com/suhanyujie/workUser/service/workUser/cmd/api/internal/svc"
	"github.com/suhanyujie/workUser/service/workUser/cmd/api/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type FindUserLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewFindUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) FindUserLogic {
	return FindUserLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *FindUserLogic) FindUser() (*types.FindUserResp, error) {
	// todo: add your logic here and delete this line
	userIdStr := l.ctx.Value("userId")

	return &types.FindUserResp{
		Code:    constvar.Ok,
		Message: "success",
		Data:    userIdStr.(string),
	}, nil
}
