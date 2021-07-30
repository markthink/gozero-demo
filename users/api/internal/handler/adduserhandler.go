package handler

import (
	"net/http"

	"blog/users/api/internal/logic"
	"blog/users/api/internal/svc"
	"blog/users/api/internal/types"

	"github.com/tal-tech/go-zero/rest/httpx"
)

func AddUserHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.ReqUser
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewAddUserLogic(r.Context(), ctx)
		resp, err := l.AddUser(req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
