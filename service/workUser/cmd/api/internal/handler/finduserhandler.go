package handler

import (
	"net/http"

	"github.com/suhanyujie/workUser/service/workUser/cmd/api/internal/logic"
	"github.com/suhanyujie/workUser/service/workUser/cmd/api/internal/svc"
	"github.com/tal-tech/go-zero/rest/httpx"
)

func FindUserHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		l := logic.NewFindUserLogic(r.Context(), ctx)
		resp, err := l.FindUser()
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
