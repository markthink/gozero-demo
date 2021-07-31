package logic

import (
	"context"

	"blog/users/model"
	"blog/users/rpc/internal/svc"
	"blog/users/rpc/user"

	"github.com/tal-tech/go-zero/core/logx"
)

type UpdateLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateLogic {
	return &UpdateLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateLogic) Update(in *user.User) (*user.CommResp, error) {
	// todo: add your logic here and delete this line
	err := l.svcCtx.Model.Update(model.User{
		Password: in.Password,
		Id:       in.Id,
		Username: in.Username,
	})
	if err != nil {
		return nil, err
	}
	return &user.CommResp{Ok: true}, nil
}
