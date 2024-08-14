package controller

import (
	"context"
	"shop-v2/api/backend/v1"
	"shop-v2/internal/model"
	"shop-v2/internal/service"
)

// UserCoupon 内容管理
var UserCoupon = cUserCoupon{}

type cUserCoupon struct{}

func (a *cUserCoupon) Create(ctx context.Context, req *v1.UserCouponReq) (res *v1.UserCouponRes, err error) {
	out, err := service.UserCoupon().Create(ctx, model.UserCouponCreateInput{
		UserCouponCreateUpdateBase: model.UserCouponCreateUpdateBase{
			UserId:   req.UserId,
			CouponId: req.CouponId,
			Status:   req.Status,
		},
	})
	if err != nil {
		return nil, err
	}
	return &v1.UserCouponRes{Id: out.UserCouponId}, nil
}

func (a *cUserCoupon) Delete(ctx context.Context, req *v1.UserCouponDeleteReq) (res *v1.UserCouponDeleteRes, err error) {
	err = service.UserCoupon().Delete(ctx, req.Id)
	return
}

func (a *cUserCoupon) Update(ctx context.Context, req *v1.UserCouponUpdateReq) (res *v1.UserCouponUpdateRes, err error) {
	err = service.UserCoupon().Update(ctx, model.UserCouponUpdateInput{
		Id: req.Id,
		UserCouponCreateUpdateBase: model.UserCouponCreateUpdateBase{
			UserId:   req.UserId,
			CouponId: req.CouponId,
			Status:   req.Status,
		},
	})
	return &v1.UserCouponUpdateRes{Id: req.Id}, nil
}

// Index article list
func (a *cUserCoupon) List(ctx context.Context, req *v1.UserCouponGetListCommonReq) (res *v1.UserCouponGetListCommonRes, err error) {
	getListRes, err := service.UserCoupon().GetList(ctx, model.UserCouponGetListInput{
		Page: req.Page,
		Size: req.Size,
	})
	if err != nil {
		return nil, err
	}
	return &v1.UserCouponGetListCommonRes{
		List:  getListRes.List,
		Page:  getListRes.Page,
		Size:  getListRes.Size,
		Total: getListRes.Total,
	}, nil
}
