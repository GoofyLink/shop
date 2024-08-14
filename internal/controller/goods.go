package controller

import (
	"context"
	"github.com/gogf/gf/v2/util/gconv"
	"shop-v2/api/backend/v1"
	"shop-v2/internal/model"
	"shop-v2/internal/service"
)

// Goods 内容管理
var Goods = cGoods{}

type cGoods struct{}

func (a *cGoods) Create(ctx context.Context, req *v1.GoodsReq) (res *v1.GoodsRes, err error) {
	data := model.GoodsCreateInput{}
	// 用scan扫描
	err = gconv.Scan(req, &data)
	if err != nil {
		return nil, err
	}
	out, err := service.Goods().Create(ctx, data)
	if err != nil {
		return nil, err
	}
	return &v1.GoodsRes{Id: out.GoodsId}, nil
}

func (a *cGoods) Delete(ctx context.Context, req *v1.GoodsDeleteReq) (res *v1.GoodsDeleteRes, err error) {
	err = service.Goods().Delete(ctx, req.Id)
	return
}

func (a *cGoods) Update(ctx context.Context, req *v1.GoodsUpdateReq) (res *v1.GoodsUpdateRes, err error) {
	err = service.Goods().Update(ctx, model.GoodsUpdateInput{
		Id: req.Id,
		GoodsCreateUpdateBase: model.GoodsCreateUpdateBase{
			PicUrl:           req.PicUrl,
			Name:             req.Name,
			Price:            req.Price,
			Level1CategoryId: req.Level1CategoryId,
			Level2CategoryId: req.Level2CategoryId,
			Level3CategoryId: req.Level3CategoryId,
			Brand:            req.Brand,
			Stock:            req.Stock,
			Sale:             req.Sale,
			Tags:             req.Tags,
			DetailInfo:       req.DetailInfo,
		},
	})
	return &v1.GoodsUpdateRes{Id: req.Id}, nil
}

// Index article list
func (a *cGoods) List(ctx context.Context, req *v1.GoodsGetListCommonReq) (res *v1.GoodsGetListCommonRes, err error) {
	getListRes, err := service.Goods().GetList(ctx, model.GoodsGetListInput{
		Page: req.Page,
		Size: req.Size,
	})
	if err != nil {
		return nil, err
	}
	return &v1.GoodsGetListCommonRes{
		List:  getListRes.List,
		Page:  getListRes.Page,
		Size:  getListRes.Size,
		Total: getListRes.Total,
	}, nil
}
