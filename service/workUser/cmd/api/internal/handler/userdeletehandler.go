package handler

import (
	"net/http"

	"github.com/suhanyujie/workUser/service/workUser/cmd/api/internal/logic"
	"github.com/suhanyujie/workUser/service/workUser/cmd/api/internal/svc"
	"github.com/tal-tech/go-zero/rest/httpx"
)

func UserDeleteHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		l := logic.NewUserDeleteLogic(r.Context(), ctx)
		resp, err := l.UserDelete()
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
