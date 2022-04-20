package logic

import (
	"context"

	"credit_certificate/student-api/internal/svc"
	"credit_certificate/student-api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type StudentInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewStudentInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *StudentInfoLogic {
	return &StudentInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *StudentInfoLogic) StudentInfo(req *types.StudentInfoReq) (resp *types.StudentInfoResp, err error) {
	// todo: add your logic here and delete this line

	return
}
