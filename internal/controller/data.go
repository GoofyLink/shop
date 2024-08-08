package controller

import (
	"context"
	v1 "shop-v2/api/backend/v1"
	"shop-v2/internal/service"
)

// Data 内容管理
var Data = cData{}

type cData struct{}

func (c *cData) DataHead(ctx context.Context, req *v1.DataHeadCardReq) (res *v1.DataHeadRes, err error) {
	dataHead, err := service.Data().DataHead(ctx)
	if err != nil {
		return nil, err
	}
	return &v1.DataHeadRes{
		ToDayOrderCount: dataHead.ToDayOrderCount,
		DAU:             dataHead.DAU,
		ConversionRate:  dataHead.ConversionRate,
	}, nil
}
