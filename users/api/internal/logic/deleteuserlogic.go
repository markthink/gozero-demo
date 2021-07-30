package logic

import (
	"context"

	"blog/users/api/internal/svc"
	"blog/users/api/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type DeleteUserLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeleteUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) DeleteUserLogic {
	return DeleteUserLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeleteUserLogic) DeleteUser(req types.ReqUserId) (*types.CommResp, error) {
	// todo: add your logic here and delete this line

	return &types.CommResp{}, nil
}
