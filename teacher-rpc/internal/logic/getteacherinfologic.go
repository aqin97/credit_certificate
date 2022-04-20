package logic

import (
	"context"

	"credit_certificate/teacher-rpc/internal/svc"
	"credit_certificate/teacher-rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetTeacherInfoLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetTeacherInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetTeacherInfoLogic {
	return &GetTeacherInfoLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetTeacherInfoLogic) GetTeacherInfo(in *pb.GetTeacherInfoReq) (*pb.GetTeacherInfoResp, error) {
	// todo: add your logic here and delete this line

	return &pb.GetTeacherInfoResp{}, nil
}
