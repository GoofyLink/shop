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
		AddCollection(ctx context.Context, in model.AddCollectionInput) (out model.AddCollectionOutPut, err error)
		// 兼容处理 优先根据收藏ID删除收藏ID为0 再根据对象ID和Type删除
		DeleteCollection(ctx context.Context, in model.DeleteCollectionInput) (out model.DeleteCollectionOutPut, err error)
		CollectionList(ctx context.Context, in model.CollectionListInput) (out *model.CollectionListOutput, err error)
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
