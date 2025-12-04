package logic

import (
	"context"

	"github.com/Nha1410/go-zero-template/api/internal/svc"
	"github.com/Nha1410/go-zero-template/api/internal/types"
	"github.com/Nha1410/go-zero-template/common/errors"
	"github.com/Nha1410/go-zero-template/common/validator"
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

func (l *CreateUserLogic) CreateUser(req *types.CreateUserRequest) (*types.BaseResponse, error) {
	if err := validator.Validate(req); err != nil {
		return nil, errors.ErrBadRequest.WithDetails(err.Error())
	}

	return &types.BaseResponse{
		Code:    200,
		Message: "User created successfully",
		Data:    nil,
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

func (l *GetUserLogic) GetUser(req *types.GetUserRequest) (*types.BaseResponse, error) {
	return &types.BaseResponse{
		Code:    200,
		Message: "Success",
		Data:    nil,
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

func (l *GetUsersLogic) GetUsers(req *types.GetUsersRequest) (*types.BaseResponse, error) {
	return &types.BaseResponse{
		Code:    200,
		Message: "Success",
		Data:    nil,
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

func (l *UpdateUserLogic) UpdateUser(req *types.UpdateUserRequest) (*types.BaseResponse, error) {
	if err := validator.Validate(req); err != nil {
		return nil, errors.ErrBadRequest.WithDetails(err.Error())
	}

	return &types.BaseResponse{
		Code:    200,
		Message: "User updated successfully",
		Data:    nil,
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

func (l *DeleteUserLogic) DeleteUser(req *types.DeleteUserRequest) (*types.BaseResponse, error) {
	return &types.BaseResponse{
		Code:    200,
		Message: "User deleted successfully",
	}, nil
}
