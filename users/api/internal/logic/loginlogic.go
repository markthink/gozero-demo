package logic

import (
	"context"

	"blog/users/api/internal/svc"
	"blog/users/api/internal/types"
	"blog/users/rpc/users"

	"github.com/tal-tech/go-zero/core/logx"
)

type LoginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) LoginLogic {
	return LoginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *LoginLogic) Login(req types.ReqUser) (*types.RespLogin, error) {
	// todo: add your logic here and delete this line
	// 调用 user rpc 的 login 方法
	resp, err := l.svcCtx.User.Login(l.ctx, &users.ReqUser{Username: req.Username, Password: req.Password})
	if err != nil {
		return nil, err
	}
	return &types.RespLogin{Token: resp.Token}, nil
}
