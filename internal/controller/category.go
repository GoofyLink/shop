package controller

import (
	"context"
	"shop-v2/api/backend/v1"
	"shop-v2/internal/model"
	"shop-v2/internal/service"
)

// Category 内容管理
var Category = cCategory{}

type cCategory struct{}

func (a *cCategory) Create(ctx context.Context, req *v1.CategoryReq) (res *v1.CategoryRes, err error) {
	out, err := service.Category().Create(ctx, model.CategoryCreateInput{
		CategoryCreateUpdateBase: model.CategoryCreateUpdateBase{
			ParentId: req.ParentId,
			PicUrl:   req.PicUrl,
			Level:    req.Level,
			Sort:     req.Sort,
			Name:     req.Name,
		},
	})
	if err != nil {
		return nil, err
	}
	return &v1.CategoryRes{CategoryId: out.CategoryId}, nil
}

func (a *cCategory) Delete(ctx context.Context, req *v1.CategoryDeleteReq) (res *v1.CategoryDeleteRes, err error) {
	err = service.Category().Delete(ctx, req.Id)
	return
}

func (a *cCategory) Update(ctx context.Context, req *v1.CategoryUpdateReq) (res *v1.CategoryUpdateRes, err error) {
	err = service.Category().Update(ctx, model.CategoryUpdateInput{
		Id: req.Id,
		CategoryCreateUpdateBase: model.CategoryCreateUpdateBase{
			ParentId: uint(req.ParentId),
			PicUrl:   req.PicUrl,
			Level:    req.Level,
			Sort:     req.Sort,
			Name:     req.Name,
		},
	})
	return &v1.CategoryUpdateRes{Id: req.Id}, nil
}

// Index article list
func (a *cCategory) List(ctx context.Context, req *v1.CategoryGetListCommonReq) (res *v1.CategoryGetListCommonRes, err error) {
	getListRes, err := service.Category().GetList(ctx, model.CategoryGetListInput{
		Page: req.Page,
		Size: req.Size,
		Sort: req.Sort,
	})
	if err != nil {
		return nil, err
	}
	return &v1.CategoryGetListCommonRes{
		List:  getListRes.List,
		Page:  getListRes.Page,
		Size:  getListRes.Size,
		Total: getListRes.Total,
	}, nil
}
