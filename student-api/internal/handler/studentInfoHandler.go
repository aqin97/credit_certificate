package handler

import (
	"net/http"

	"credit_certificate/student-api/internal/logic"
	"credit_certificate/student-api/internal/svc"
	"credit_certificate/student-api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func studentInfoHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.StudentInfoReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewStudentInfoLogic(r.Context(), svcCtx)
		resp, err := l.StudentInfo(&req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
