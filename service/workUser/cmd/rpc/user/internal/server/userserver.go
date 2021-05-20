// Code generated by goctl. DO NOT EDIT!
// Source: user.proto

package server

import (
	"context"

	"github.com/suhanyujie/workUser/service/workUser/cmd/rpc/user/internal/logic"
	"github.com/suhanyujie/workUser/service/workUser/cmd/rpc/user/internal/svc"
	"github.com/suhanyujie/workUser/service/workUser/cmd/rpc/user/user"
)

type UserServer struct {
	svcCtx *svc.ServiceContext
}

func NewUserServer(svcCtx *svc.ServiceContext) *UserServer {
	return &UserServer{
		svcCtx: svcCtx,
	}
}

func (s *UserServer) Create(ctx context.Context, in *user.CreateUserReq) (*user.CreateUserResp, error) {
	l := logic.NewCreateLogic(ctx, s.svcCtx)
	return l.Create(in)
}