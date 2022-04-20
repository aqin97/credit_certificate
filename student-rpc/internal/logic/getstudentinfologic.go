package logic

import (
	"context"

	"credit_certificate/student-rpc/internal/svc"
	"credit_certificate/student-rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetStudentInfoLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetStudentInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetStudentInfoLogic {
	return &GetStudentInfoLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetStudentInfoLogic) GetStudentInfo(in *pb.GetStudentInfoReq) (*pb.GetStudentInfoResp, error) {
	// todo: add your logic here and delete this line

	return &pb.GetStudentInfoResp{}, nil
}
