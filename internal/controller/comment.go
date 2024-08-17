package controller

import (
	"context"
	"github.com/gogf/gf/v2/util/gconv"
	v1 "shop-v2/api/frontend/v1"
	"shop-v2/internal/model"
	"shop-v2/internal/service"
)

var Comment = cComment{}

type cComment struct{}

func (c *cComment) AddComment(ctx context.Context, req *v1.AddCommentReq) (res *v1.AddCommentRes, err error) {
	data := model.AddCommentInput{}
	err = gconv.Struct(ctx, &data)
	if err != nil {
		return nil, err
	}
	out, err := service.Comment().AddComment(ctx, data)
	if err != nil {
		return nil, err
	}

	return &v1.AddCommentRes{Id: out.Id}, nil
}

func (c *cComment) DeleteComment(ctx context.Context, req *v1.DeleteCommentReq) (res *v1.DeleteCommentRes, err error) {
	data := model.DeleteCommentInput{}
	err = gconv.Struct(ctx, &data)
	if err != nil {
		return nil, err
	}
	collection, err := service.Comment().DeleteComment(ctx, data)
	if err != nil {
		return nil, err
	}
	return &v1.DeleteCommentRes{Id: collection.Id}, nil
}

func (c *cComment) CommentList(ctx context.Context, req *v1.CommentListReq) (res *v1.CommentListRes, err error) {
	getListRes, err := service.Comment().CommentList(ctx, model.CommentListInput{
		Page: req.Page,
		Size: req.Size,
		Type: req.Type,
	})
	if err != nil {
		return nil, err
	}
	return &v1.CommentListRes{
		List:  getListRes.List,
		Page:  getListRes.Page,
		Size:  getListRes.Size,
		Total: getListRes.Total,
	}, nil
}
