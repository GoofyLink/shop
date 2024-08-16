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
	ICollection interface {
		// 添加
		AddCollection(ctx context.Context, in model.AddCollectionInput) (out model.AddCollectionOutPut, err error)
		// 兼容处理 优先根据收藏ID删除收藏ID为0 再根据对象ID和Type删除
		DeleteCollection(ctx context.Context, in model.DeleteCollectionInput) (out model.DeleteCollectionOutPut, err error)
		// 获取collectionList列表
		CollectionList(ctx context.Context, in model.CollectionListInput) (out *model.CollectionListOutput, err error)
		// 抽取获得收藏数量的方法 for 文章详情/商品详情
		CollectionCount(ctx context.Context, objectId uint, collectionType uint8) (count int, err error)
		// 抽取方法判断当前用户是否收藏 文章详情/商品详情
		CollectionCheck(ctx context.Context, in model.CheckIsCollectionInput) (bool, error)
	}
)

var (
	localCollection ICollection
)

func Collection() ICollection {
	if localCollection == nil {
		panic("implement not found for interface ICollection, forgot register?")
	}
	return localCollection
}

func RegisterCollection(i ICollection) {
	localCollection = i
}
