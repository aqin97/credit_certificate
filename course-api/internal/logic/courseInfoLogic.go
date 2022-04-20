package logic

import (
	"context"
	"errors"

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
	course1 := types.Course{
		Id:         1,
		CourseName: "高等数学",
		Teacher:    "赵老师",
		BeginAt:    "2022-02-20",
		EndAt:      "2022-06-20",
	}
	course2 := types.Course{
		Id:         2,
		CourseName: "线性代数",
		Teacher:    "赵老师",
		BeginAt:    "2022-02-20",
		EndAt:      "2022-06-20",
	}
	m := map[int64]types.Course{
		1: course1,
		2: course2,
	}
	if req.CourseId != 1 && req.CourseId != 2 {
		return nil, errors.New("该课程不存在")
	}
	courseInfo := m[req.CourseId]
	return &types.CourseInfoResp{
		CourseInfo: courseInfo,
	}, nil
}
