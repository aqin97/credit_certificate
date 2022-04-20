package logic

import (
	"context"

	"credit_certificate/course-rpc/internal/svc"
	"credit_certificate/course-rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetCourseInfoLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetCourseInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetCourseInfoLogic {
	return &GetCourseInfoLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetCourseInfoLogic) GetCourseInfo(in *pb.GetCourseInfoReq) (*pb.GetCourseInfoResp, error) {
	// todo: add your logic here and delete this line

	return &pb.GetCourseInfoResp{}, nil
}
