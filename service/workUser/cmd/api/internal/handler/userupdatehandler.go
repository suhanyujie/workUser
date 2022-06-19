package handler

import (
	"net/http"

	"github.com/suhanyujie/workUser/service/workUser/cmd/api/internal/logic"
	"github.com/suhanyujie/workUser/service/workUser/cmd/api/internal/svc"
	"github.com/suhanyujie/workUser/service/workUser/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func UserUpdateHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.Request
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewUserUpdateLogic(r.Context(), ctx)
		resp, err := l.UserUpdate(req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
