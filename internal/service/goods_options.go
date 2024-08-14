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
	IGoodsOptions interface {
		// Create 创建分类
		Create(ctx context.Context, in model.GoodsOptionsCreateInput) (out model.GoodsOptionsCreateOutput, err error)
		// Delete 删除 这样删除是软删除还会留着原来的信息
		Delete(ctx context.Context, id uint) (err error)
		// Update 修改
		Update(ctx context.Context, in model.GoodsOptionsUpdateInput) (err error)
		// GetList 查询分类列表
		GetList(ctx context.Context, in model.GoodsOptionsGetListInput) (out *model.GoodsOptionsGetListOutput, err error)
	}
)

var (
	localGoodsOptions IGoodsOptions
)

func GoodsOptions() IGoodsOptions {
	if localGoodsOptions == nil {
		panic("implement not found for interface IGoodsOptions, forgot register?")
	}
	return localGoodsOptions
}

func RegisterGoodsOptions(i IGoodsOptions) {
	localGoodsOptions = i
}
