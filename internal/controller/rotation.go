package controller

import (
	"context"
	"shop-v2/api/backend/v1"
	"shop-v2/internal/model"
	"shop-v2/internal/service"
)

// Rotation 内容管理
var Rotation = cRotation{}

type cRotation struct{}

func (a *cRotation) Create(ctx context.Context, req *v1.RotationCreateReq) (res *v1.RotationCreateRes, err error) {
	out, err := service.Rotation().Create(ctx, model.RotationCreateInput{
		RotationCreateUpdateBase: model.RotationCreateUpdateBase{
			Id:     req.Id,
			PicUrl: req.Pic_url,
			Link:   req.Link,
			Sort:   req.Sort,
		},
	})
	if err != nil {
		return nil, err
	}
	return &v1.RotationCreateRes{RotationId: out.RotationId}, nil
}

func (a *cRotation) Delete(ctx context.Context, req *v1.RotationDeleteReq) (res *v1.RotationDeleteRes, err error) {
	err = service.Rotation().Delete(ctx, req.Id)
	return
}

func (a *cRotation) Update(ctx context.Context, req *v1.RotationUpdateReq) (res *v1.RotationUpdateRes, err error) {
	err = service.Rotation().Update(ctx, model.RotationUpdateInput{
		Id: req.Id,
		RotationCreateUpdateBase: model.RotationCreateUpdateBase{
			PicUrl: req.Pic_url,
			Link:   req.Link,
			Sort:   req.Sort,
		},
	})
	return &v1.RotationUpdateRes{Id: req.Id}, nil
}

// Index article list
func (a *cRotation) List(ctx context.Context, req *v1.RotationGetListCommonReq) (res *v1.RotationGetListCommonRes, err error) {
	getListRes, err := service.Rotation().GetList(ctx, model.RotationGetListInput{
		Page: req.Page,
		Size: req.Size,
		Sort: req.Sort,
	})
	if err != nil {
		return nil, err
	}
	return &v1.RotationGetListCommonRes{
		List:  getListRes.List,
		Page:  getListRes.Page,
		Size:  getListRes.Size,
		Total: getListRes.Total,
	}, nil
}
