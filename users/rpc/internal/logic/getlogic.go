package logic

import (
	"context"

	"blog/users/rpc/internal/svc"
	"blog/users/rpc/user"

	"github.com/tal-tech/go-zero/core/logx"
)

type GetLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetLogic {
	return &GetLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetLogic) Get(in *user.ReqUserId) (*user.User, error) {
	// todo: add your logic here and delete this line

	one, err := l.svcCtx.Model.FindOne(in.Id)
	if err != nil {
		return nil, err
	}
	return &user.User{
		Id:       one.Id,
		Username: one.Username,
		Password: one.Password,
	}, nil
}
