package handler

import (
	"net/http"

	"credit_certificate/teacher-api/internal/logic"
	"credit_certificate/teacher-api/internal/svc"
	"credit_certificate/teacher-api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func teacherInfoHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.TeacherInfoReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewTeacherInfoLogic(r.Context(), svcCtx)
		resp, err := l.TeacherInfo(&req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
