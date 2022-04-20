package logic

import (
	"context"

	"credit_certificate/teacher-api/internal/svc"
	"credit_certificate/teacher-api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type TeacherInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewTeacherInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *TeacherInfoLogic {
	return &TeacherInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *TeacherInfoLogic) TeacherInfo(req *types.TeacherInfoReq) (resp *types.TeacherInfoResp, err error) {
	// todo: add your logic here and delete this line

	return
}
