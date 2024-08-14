package controller

import (
	"context"
	"github.com/gogf/gf/v2/util/gconv"
	"shop-v2/api/backend/v1"
	"shop-v2/internal/model"
	"shop-v2/internal/service"
)

// GoodsOptions 内容管理
var GoodsOptions = cGoodsOptions{}

type cGoodsOptions struct{}

func (a *cGoodsOptions) Create(ctx context.Context, req *v1.GoodsOptionsReq) (res *v1.GoodsOptionsRes, err error) {
	data := model.GoodsOptionsCreateInput{}
	// 用scan扫描
	err = gconv.Scan(req, &data)
	if err != nil {
		return nil, err
	}
	out, err := service.GoodsOptions().Create(ctx, data)
	if err != nil {
		return nil, err
	}
	return &v1.GoodsOptionsRes{Id: out.GoodsOptionsId}, nil
}

func (a *cGoodsOptions) Delete(ctx context.Context, req *v1.GoodsOptionsDeleteReq) (res *v1.GoodsOptionsDeleteRes, err error) {
	err = service.GoodsOptions().Delete(ctx, req.Id)
	return
}

func (a *cGoodsOptions) Update(ctx context.Context, req *v1.GoodsOptionsUpdateReq) (res *v1.GoodsOptionsUpdateRes, err error) {
	data := model.GoodsOptionsUpdateInput{}
	err = gconv.Struct(req, &data)
	if err != nil {
		return nil, err
	}
	err = service.GoodsOptions().Update(ctx, data)
	return &v1.GoodsOptionsUpdateRes{Id: req.Id}, nil
}

// Index article list
func (a *cGoodsOptions) List(ctx context.Context, req *v1.GoodsOptionsGetListCommonReq) (res *v1.GoodsOptionsGetListCommonRes, err error) {
	getListRes, err := service.GoodsOptions().GetList(ctx, model.GoodsOptionsGetListInput{
		Page: req.Page,
		Size: req.Size,
	})
	if err != nil {
		return nil, err
	}
	return &v1.GoodsOptionsGetListCommonRes{
		List:  getListRes.List,
		Page:  getListRes.Page,
		Size:  getListRes.Size,
		Total: getListRes.Total,
	}, nil
}
