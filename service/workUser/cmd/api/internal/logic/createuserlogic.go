package logic

import (
	"context"

	"github.com/pkg/errors"
	"github.com/suhanyujie/workUser/common/stringx"
	"github.com/suhanyujie/workUser/service/workUser/cmd/api/internal/svc"
	"github.com/suhanyujie/workUser/service/workUser/cmd/api/internal/types"
	"github.com/suhanyujie/workUser/service/workUser/cmd/api/internal/utils/constvar"
	"github.com/suhanyujie/workUser/service/workUser/model"
	"github.com/tal-tech/go-zero/core/logx"
)

type CreateUserLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCreateUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) CreateUserLogic {
	return CreateUserLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// CreateUser 新增用户，主要用于注册用户使用
func (l *CreateUserLogic) CreateUser(req types.CreateUserReq) (*types.CreateUserResp, error) {
	insertRes, err := l.svcCtx.WorkUserModel.Insert(model.WorkUser{
		Avatar:   "",
		UserName: req.UserName,
		Pwd:      stringx.Md5Str(req.Pwd),
		Nickname: "",
		Sex:      0,
	})
	if err != nil {
		return nil, err
	}
	dataId, err := insertRes.LastInsertId()
	if err != nil {
		return nil, errors.Wrap(err, "get lastInsertId error")
	}
	user, err := l.svcCtx.WorkUserModel.FindOne(dataId)
	if err != nil {
		return nil, errors.Wrap(err, "FindOne error")
	}

	return &types.CreateUserResp{
		Code:    constvar.Ok,
		Message: "success",
		Data: types.OneUser{
			Id:       int64(user.ID),
			UserName: user.UserName,
		},
	}, nil
}
