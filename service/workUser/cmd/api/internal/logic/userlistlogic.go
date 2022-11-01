package logic

import (
	"context"
	"github.com/suhanyujie/go-utils/helper/copyer"

	"github.com/suhanyujie/workUser/service/workUser/cmd/api/internal/svc"
	"github.com/suhanyujie/workUser/service/workUser/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserListLogic(ctx context.Context, svcCtx *svc.ServiceContext) UserListLogic {
	return UserListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserListLogic) UserList(req types.UserListReq) (*types.UserListRespVo, error) {
	param := map[string]interface{}{
		"page": 1,
		"size": 20,
	}
	total, list, err := l.svcCtx.WorkUserModel.GetList(param)
	if err != nil {
		return nil, err
	}
	userList := make([]types.OneUser, 0)
	copyer.Copy(list, &userList)
	return &types.UserListRespVo{
		Code:    0,
		Message: "",
		Data: types.UserListRespVoData{
			List:  userList,
			Total: total,
			Page:  param["page"].(int),
			Size:  param["size"].(int),
		},
	}, nil
}
