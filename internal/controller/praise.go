package controller

import (
	"context"
	"github.com/gogf/gf/v2/util/gconv"
	v1 "shop-v2/api/frontend/v1"
	"shop-v2/internal/model"
	"shop-v2/internal/service"
)

var Praise = cPraise{}

type cPraise struct{}

func (c *cPraise) AddPraise(ctx context.Context, req *v1.AddPraiseReq) (res *v1.AddPraiseRes, err error) {
	data := model.AddPraiseInput{}
	err = gconv.Struct(ctx, &data)
	if err != nil {
		return nil, err
	}
	out, err := service.Praise().AddPraise(ctx, data)
	if err != nil {
		return nil, err
	}

	return &v1.AddPraiseRes{Id: out.Id}, nil
}

func (c *cPraise) DeletePraise(ctx context.Context, req *v1.DeletePraiseReq) (res *v1.DeletePraiseRes, err error) {
	data := model.DeletePraiseInput{}
	err = gconv.Struct(ctx, &data)
	if err != nil {
		return nil, err
	}
	collection, err := service.Praise().DeletePraise(ctx, data)
	if err != nil {
		return nil, err
	}
	return &v1.DeletePraiseRes{Id: collection.Id}, nil
}

func (c *cPraise) PraiseList(ctx context.Context, req *v1.PraiseListReq) (res *v1.PraiseListRes, err error) {
	getListRes, err := service.Praise().PraiseList(ctx, model.PraiseListInput{
		Page: req.Page,
		Size: req.Size,
		Type: req.Type,
	})
	if err != nil {
		return nil, err
	}
	return &v1.PraiseListRes{
		List:  getListRes.List,
		Page:  getListRes.Page,
		Size:  getListRes.Size,
		Total: getListRes.Total,
	}, nil
}
