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
	IGoods interface {
		// Create 创建分类
		Create(ctx context.Context, in model.GoodsCreateInput) (out model.GoodsCreateOutput, err error)
		// Delete 删除 这样删除是软删除还会留着原来的信息
		Delete(ctx context.Context, id uint) (err error)
		// Update 修改
		Update(ctx context.Context, in model.GoodsUpdateInput) (err error)
		// GetList 查询分类列表
		GetList(ctx context.Context, in model.GoodsGetListInput) (out *model.GoodsGetListOutput, err error)
	}
)

var (
	localGoods IGoods
)

func Goods() IGoods {
	if localGoods == nil {
		panic("implement not found for interface IGoods, forgot register?")
	}
	return localGoods
}

func RegisterGoods(i IGoods) {
	localGoods = i
}
