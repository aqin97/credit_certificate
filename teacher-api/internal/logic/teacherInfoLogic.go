package logic

import (
	"context"
	"errors"

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
	teacher1 := types.Teacher{
		Id:     1,
		Name:   "赵老师",
		Gender: 1,
		Mobile: "187 8888 8888",
		Courses: []string{
			"高等数学",
			"线性代数",
		},
	}
	teacher2 := types.Teacher{
		Id:     1,
		Name:   "王老师",
		Gender: 0,
		Mobile: "187 8888 9999",
		Courses: []string{
			"近代史",
			"马哲",
		},
	}
	m := map[int64]types.Teacher{
		1: teacher1,
		2: teacher2,
	}
	if teacherInfo, ok := m[req.TeacherId]; ok {
		return &types.TeacherInfoResp{
			TeacherInfo: teacherInfo,
		}, nil
	}

	return nil, errors.New("该教师不存在")
}
