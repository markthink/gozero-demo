package logic

import (
	"context"

	"blog/users/rpc/internal/svc"
	"blog/users/rpc/user"

	"github.com/tal-tech/go-zero/core/logx"
)

type LoginLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *LoginLogic) Login(in *user.ReqUser) (*user.RespLogin, error) {
	// todo: add your logic here and delete this line

	return &user.RespLogin{}, nil
}
