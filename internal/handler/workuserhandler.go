package handler

import (
	"net/http"

	"github.com/suhanyujie/workUser/internal/logic"
	"github.com/suhanyujie/workUser/internal/svc"
	"github.com/suhanyujie/workUser/internal/types"

	"github.com/tal-tech/go-zero/rest/httpx"
)

func WorkUserHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.Request
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewWorkUserLogic(r.Context(), ctx)
		resp, err := l.WorkUser(req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
