package controller

import (
	"context"
	"shop-v2/api/backend/v1"
	"shop-v2/internal/model"
	"shop-v2/internal/service"
)

// Coupon 内容管理
var Coupon = cCoupon{}

type cCoupon struct{}

func (a *cCoupon) Create(ctx context.Context, req *v1.CouponReq) (res *v1.CouponRes, err error) {
	out, err := service.Coupon().Create(ctx, model.CouponCreateInput{
		CouponCreateUpdateBase: model.CouponCreateUpdateBase{
			Name:       req.Name,
			Price:      req.Price,
			GoodsId:    req.GoodsIds,
			CategoryId: req.CategoryId,
		},
	})
	if err != nil {
		return nil, err
	}
	return &v1.CouponRes{CouponId: out.CouponId}, nil
}

func (a *cCoupon) Delete(ctx context.Context, req *v1.CouponDeleteReq) (res *v1.CouponDeleteRes, err error) {
	err = service.Coupon().Delete(ctx, req.Id)
	return
}

func (a *cCoupon) Update(ctx context.Context, req *v1.CouponUpdateReq) (res *v1.CouponUpdateRes, err error) {
	err = service.Coupon().Update(ctx, model.CouponUpdateInput{
		Id: req.Id,
		CouponCreateUpdateBase: model.CouponCreateUpdateBase{
			Name:       req.Name,
			Price:      req.Price,
			GoodsId:    req.GoodsIds,
			CategoryId: req.CategoryId,
		},
	})
	return &v1.CouponUpdateRes{Id: req.Id}, nil
}

// Index article list
func (a *cCoupon) List(ctx context.Context, req *v1.CouponGetListCommonReq) (res *v1.CouponGetListCommonRes, err error) {
	getListRes, err := service.Coupon().GetList(ctx, model.CouponGetListInput{
		Page: req.Page,
		Size: req.Size,
		Sort: req.Sort,
	})
	if err != nil {
		return nil, err
	}
	return &v1.CouponGetListCommonRes{
		List:  getListRes.List,
		Page:  getListRes.Page,
		Size:  getListRes.Size,
		Total: getListRes.Total,
	}, nil
}
