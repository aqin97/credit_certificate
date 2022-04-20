package handler

import (
	"net/http"

	"credit_certificate/course-api/internal/logic"
	"credit_certificate/course-api/internal/svc"
	"credit_certificate/course-api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func courseInfoHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.CourseInfoReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewCourseInfoLogic(r.Context(), svcCtx)
		resp, err := l.CourseInfo(&req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
