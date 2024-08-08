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
	IData interface {
		DataHead(ctx context.Context) (out *model.DataHeadOutPut, err error)
	}
)

var (
	localData IData
)

func Data() IData {
	if localData == nil {
		panic("implement not found for interface IData, forgot register?")
	}
	return localData
}

func RegisterData(i IData) {
	localData = i
}
