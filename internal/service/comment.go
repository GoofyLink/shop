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
	IComment interface {
		// 添加
		AddComment(ctx context.Context, in model.AddCommentInput) (out model.AddCommentOutPut, err error)
		// 兼容处理 优先根据收藏ID删除收藏ID为0 再根据对象ID和Type删除
		DeleteComment(ctx context.Context, in model.DeleteCommentInput) (out model.DeleteCommentOutPut, err error)
		// 获取collectionList列表
		CommentList(ctx context.Context, in model.CommentListInput) (out *model.CommentListOutput, err error)
		// 抽取获得收藏数量的方法 for 文章详情/商品详情
		CommentCount(ctx context.Context, objectId uint, collectionType uint8) (count int, err error)
		// 抽取方法判断当前用户是否收藏 文章详情/商品详情
		CommentCheck(ctx context.Context, in model.CheckIsCommentInput) (bool, error)
	}
)

var (
	localComment IComment
)

func Comment() IComment {
	if localComment == nil {
		panic("implement not found for interface IComment, forgot register?")
	}
	return localComment
}

func RegisterComment(i IComment) {
	localComment = i
}
