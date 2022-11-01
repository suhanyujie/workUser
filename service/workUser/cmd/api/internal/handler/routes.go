// Code generated by goctl. DO NOT EDIT.
package handler

import (
	"net/http"

	"github.com/suhanyujie/workUser/service/workUser/cmd/api/internal/svc"

	"github.com/zeromicro/go-zero/rest"
)

func RegisterHandlers(engine *rest.Server, serverCtx *svc.ServiceContext) {
	engine.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodPost,
				Path:    "/user/create",
				Handler: CreateUserHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/user/find/:userId",
				Handler: FindUserHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/user/list",
				Handler: UserListHandler(serverCtx),
			},
			{
				Method:  http.MethodDelete,
				Path:    "/user/delete/:userId",
				Handler: UserDeleteHandler(serverCtx),
			},
		},
	)
}
