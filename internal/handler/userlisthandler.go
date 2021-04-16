package handler

import (
	"net/http"

	"github.com/suhanyujie/workUser/internal/logic"
	"github.com/suhanyujie/workUser/internal/svc"
	"github.com/suhanyujie/workUser/internal/types"

	"github.com/tal-tech/go-zero/rest/httpx"
)

func UserListHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.Request
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewUserListLogic(r.Context(), ctx)
		resp, err := l.UserList(req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
