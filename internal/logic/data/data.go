package data

import (
	"context"
	"github.com/gogf/gf/v2/os/gtime"
	"shop-v2/internal/dao"
	"shop-v2/internal/model"
	"shop-v2/internal/service"
	utility "shop-v2/utility"
)

type sData struct {
}

// 注册服务  给server层生成对应服务
func init() {
	service.RegisterData(New())
}

func New() *sData {
	return &sData{}
}

func (s *sData) DataHead(ctx context.Context) (out *model.DataHeadOutPut, err error) {
	return &model.DataHeadOutPut{
		ToDayOrderCount: todayOrderCount(ctx), // todo
		DAU:             utility.RandInt(1000),
		ConversionRate:  utility.RandInt(50),
	}, nil
}

// 查询订单总数 只在当前包使用
func todayOrderCount(ctx context.Context) (count int) {
	count, err := dao.OrderGoodsInfo.
		Ctx(ctx).
		WhereBetween(dao.OrderInfo.
			Columns().CreatedAt,
			gtime.Now().
				StartOfDay(),
			gtime.Now().EndOfDay()).
		Count(dao.OrderInfo.Columns().Id)
	if err != nil {
		return -1
	}
	return count
}
