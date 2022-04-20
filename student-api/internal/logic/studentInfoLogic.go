package logic

import (
	"context"
	"errors"

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
	student1 := types.Student{
		Id:     1,
		Name:   "小钱",
		Gender: 1,
		Mobile: "159 8888 8888",
	}
	student2 := types.Student{
		Id:     2,
		Name:   "小李",
		Gender: 0,
		Mobile: "159 8888 8899",
	}
	m := map[int64]types.Student{
		1: student1,
		2: student2,
	}
	if studentInfo, ok := m[req.StudentId]; ok {
		return &types.StudentInfoResp{
			StudentInfo: studentInfo,
		}, nil
	}
	return nil, errors.New("该学生不存在")
}
