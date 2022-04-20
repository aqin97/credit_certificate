package logic

import (
	"context"

	"credit_certificate/course-api/internal/svc"
	"credit_certificate/course-api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CourseInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCourseInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CourseInfoLogic {
	return &CourseInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CourseInfoLogic) CourseInfo(req *types.CourseInfoReq) (resp *types.CourseInfoResp, err error) {
	// todo: add your logic here and delete this line

	return
}
