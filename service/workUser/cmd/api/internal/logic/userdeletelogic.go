package logic

import (
	"context"
	"github.com/pkg/errors"
	"strconv"

	"github.com/suhanyujie/workUser/service/workUser/cmd/api/internal/svc"
	"github.com/suhanyujie/workUser/service/workUser/cmd/api/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
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

func (l *UserDeleteLogic) UserDelete() (*types.ResponseVo, error) {
	resp := &types.ResponseVo{
		Code:    0,
		Message: "",
		Data:    false,
	}
	userIdStr, isOk := l.ctx.Value("userId").(string)
	if !isOk {
		return nil, errors.New("参数 userId 不合法。")
	}
	userId, err := strconv.ParseInt(userIdStr, 10, 64)
	if err != nil {
		return nil, errors.Wrap(err, "parse int error")
	}
	// 检查用户是否存在
	_, err = l.svcCtx.WorkUserModel.FindOne(userId)
	if err != nil {
		return nil, err
	}
	err = l.svcCtx.WorkUserModel.Delete(userId)
	if err != nil {
		return nil, err
	}
	resp.Data = true
	return resp, nil
}
