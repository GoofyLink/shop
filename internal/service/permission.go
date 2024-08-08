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
	IPermission interface {
		// Create 创建内容
		Create(ctx context.Context, in model.PermissionCreateInput) (out model.PermissionCreateOutput, err error)
		// Delete 删除 这样删除是软删除还会留着原来的信息
		Delete(ctx context.Context, id uint) error
		// Update 修改
		Update(ctx context.Context, in model.PermissionUpdateInput) error
		// GetList 查询内容列表
		GetList(ctx context.Context, in model.PermissionGetListInput) (out *model.PermissionGetListOutput, err error)
		// 获取role密码
		GetPermissionByUserNamePassword(ctx context.Context, in model.LoginInput) map[string]interface{}
	}
)

var (
	localPermission IPermission
)

func Permission() IPermission {
	if localPermission == nil {
		panic("implement not found for interface IPermission, forgot register?")
	}
	return localPermission
}

func RegisterPermission(i IPermission) {
	localPermission = i
}
