package logic

import (
	"context"
	"time"

	"github.com/Nha1410/go-zero-template/service/user/internal/svc"
	"github.com/Nha1410/go-zero-template/service/user/userclient"
	"github.com/zeromicro/go-zero/core/logx"
)

type CreateUserLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCreateUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateUserLogic {
	return &CreateUserLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CreateUserLogic) CreateUser(req *userclient.CreateUserReq) (*userclient.CreateUserResp, error) {
	user, err := l.svcCtx.UserUsecase.CreateUser(l.ctx, req.Email, req.Name)
	if err != nil {
		return nil, err
	}

	return &userclient.CreateUserResp{
		Id:        user.ID,
		Email:     user.Email,
		Name:      user.Name,
		CreatedAt: user.CreatedAt.Format(time.RFC3339),
	}, nil
}

type GetUserLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserLogic {
	return &GetUserLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetUserLogic) GetUser(req *userclient.GetUserReq) (*userclient.GetUserResp, error) {
	user, err := l.svcCtx.UserUsecase.GetUser(l.ctx, req.Id)
	if err != nil {
		return nil, err
	}

	return &userclient.GetUserResp{
		Id:        user.ID,
		Email:     user.Email,
		Name:      user.Name,
		CreatedAt: user.CreatedAt.Format(time.RFC3339),
		UpdatedAt: user.UpdatedAt.Format(time.RFC3339),
	}, nil
}

type GetUsersLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetUsersLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUsersLogic {
	return &GetUsersLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetUsersLogic) GetUsers(req *userclient.GetUsersReq) (*userclient.GetUsersResp, error) {
	users, total, err := l.svcCtx.UserUsecase.GetUsers(l.ctx, req.Page, req.PageSize)
	if err != nil {
		return nil, err
	}

	var respUsers []*userclient.GetUserResp
	for _, user := range users {
		respUsers = append(respUsers, &userclient.GetUserResp{
			Id:        user.ID,
			Email:     user.Email,
			Name:      user.Name,
			CreatedAt: user.CreatedAt.Format(time.RFC3339),
			UpdatedAt: user.UpdatedAt.Format(time.RFC3339),
		})
	}

	return &userclient.GetUsersResp{
		Users:    respUsers,
		Total:    total,
		Page:     req.Page,
		PageSize: req.PageSize,
	}, nil
}

type UpdateUserLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateUserLogic {
	return &UpdateUserLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateUserLogic) UpdateUser(req *userclient.UpdateUserReq) (*userclient.UpdateUserResp, error) {
	user, err := l.svcCtx.UserUsecase.UpdateUser(l.ctx, req.Id, req.Email, req.Name)
	if err != nil {
		return nil, err
	}

	return &userclient.UpdateUserResp{
		Id:        user.ID,
		Email:     user.Email,
		Name:      user.Name,
		UpdatedAt: user.UpdatedAt.Format(time.RFC3339),
	}, nil
}

type DeleteUserLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeleteUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteUserLogic {
	return &DeleteUserLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeleteUserLogic) DeleteUser(req *userclient.DeleteUserReq) (*userclient.DeleteUserResp, error) {
	err := l.svcCtx.UserUsecase.DeleteUser(l.ctx, req.Id)
	if err != nil {
		return nil, err
	}

	return &userclient.DeleteUserResp{
		Success: true,
	}, nil
}
