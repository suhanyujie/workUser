package logic

import (
	"context"
	"github.com/pkg/errors"
	"github.com/suhanyujie/workUser/service/workUser/cmd/api/internal/utils/constvar"
	"strconv"

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
	userIdStr, isOk := l.ctx.Value("userId").(string)
	if !isOk {
		return nil, errors.New("参数 userId 不合法。")
	}
	userId, err := strconv.ParseInt(userIdStr, 10, 64)
	if err != nil {
		return nil, errors.Wrap(err, "parse int error")
	}
	res, err := l.svcCtx.WorkUserModel.FindOne(userId)
	if err != nil {
		return nil, err
	}

	return &types.FindUserResp{
		Code:    constvar.Ok,
		Message: "success",
		Data: types.OneUser{
			Id:       int64(res.ID),
			UserName: res.UserName,
		},
	}, nil
}
