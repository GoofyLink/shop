// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"context"
	"shop-v2/internal/model"
)

type (
	IUser interface {
		// Register 创建内容
		Register(ctx context.Context, in model.RegisterInput) (out model.RegisterOutPut, err error)
		// UpdatePassword 修改密码
		UpdatePassword(ctx context.Context, in model.UpdatePasswordInput) (out model.UpdatePasswordOutPut, err error)
	}
)

var (
	localUser IUser
)

func User() IUser {
	if localUser == nil {
		panic("implement not found for interface IUser, forgot register?")
	}
	return localUser
}

func RegisterUser(i IUser) {
	localUser = i
}
