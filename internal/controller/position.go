package controller

import (
	"context"
	"shop-v2/api/backend/v1"
	"shop-v2/internal/model"
	"shop-v2/internal/service"
)

// Position 内容管理
var Position = cPosition{}

type cPosition struct{}

func (a *cPosition) Create(ctx context.Context, req *v1.PositionCreateReq) (res *v1.PositionCreateRes, err error) {
	out, err := service.Position().Create(ctx, model.PositionCreateInput{
		PositionCreateUpdateBase: model.PositionCreateUpdateBase{
			Id:        req.Id,
			PicUrl:    req.Pic_url,
			Link:      req.Link,
			Sort:      req.Sort,
			GoodsName: req.GoodsName,
			GoodsId:   req.GoodsId,
		},
	})
	if err != nil {
		return nil, err
	}
	return &v1.PositionCreateRes{PositionId: out.PositionId}, nil
}

func (a *cPosition) Delete(ctx context.Context, req *v1.PositionDeleteReq) (res *v1.PositionDeleteRes, err error) {
	err = service.Position().Delete(ctx, req.Id)
	return
}

func (a *cPosition) Update(ctx context.Context, req *v1.PositionUpdateReq) (res *v1.PositionUpdateRes, err error) {
	err = service.Position().Update(ctx, model.PositionUpdateInput{
		Id: req.Id,
		PositionCreateUpdateBase: model.PositionCreateUpdateBase{
			PicUrl:    req.Pic_url,
			Link:      req.Link,
			Sort:      req.Sort,
			GoodsName: req.GoodsName,
			GoodsId:   req.GoodsId,
		},
	})
	return &v1.PositionUpdateRes{Id: req.Id}, nil
}

// Index article list
func (a *cPosition) List(ctx context.Context, req *v1.PositionGetListCommonReq) (res *v1.PositionGetListCommonRes, err error) {
	getListRes, err := service.Position().GetList(ctx, model.PositionGetListInput{
		Page: req.Page,
		Size: req.Size,
		Sort: req.Sort,
	})
	if err != nil {
		return nil, err
	}
	return &v1.PositionGetListCommonRes{
		List:  getListRes.List,
		Page:  getListRes.Page,
		Size:  getListRes.Size,
		Total: getListRes.Total,
	}, nil
}
