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
	IPraise interface {
		// 添加
		AddPraise(ctx context.Context, in model.AddPraiseInput) (out model.AddPraiseOutPut, err error)
		// 兼容处理 优先根据收藏ID删除收藏ID为0 再根据对象ID和Type删除
		DeletePraise(ctx context.Context, in model.DeletePraiseInput) (out model.DeletePraiseOutPut, err error)
		// 获取collectionList列表
		PraiseList(ctx context.Context, in model.PraiseListInput) (out *model.PraiseListOutput, err error)
		// 抽取获得收藏数量的方法 for 文章详情/商品详情
		PraiseCount(ctx context.Context, objectId uint, collectionType uint8) (count int, err error)
		// 抽取方法判断当前用户是否收藏 文章详情/商品详情
		PraiseCheck(ctx context.Context, in model.CheckIsPraiseInput) (bool, error)
	}
)

var (
	localPraise IPraise
)

func Praise() IPraise {
	if localPraise == nil {
		panic("implement not found for interface IPraise, forgot register?")
	}
	return localPraise
}

func RegisterPraise(i IPraise) {
	localPraise = i
}
