package logic

import (
	"context"

	"blog/users/api/internal/svc"
	"blog/users/api/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type GetAllUsersLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetAllUsersLogic(ctx context.Context, svcCtx *svc.ServiceContext) GetAllUsersLogic {
	return GetAllUsersLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetAllUsersLogic) GetAllUsers() ([]types.User, error) {
	// todo: add your logic here and delete this line

	return []types.User{}, nil
}
