package v1

import (
	"github.com/gogf/gf/v2/frame/g"
)

type CouponCommon struct {
	Name       string `json:"name"  v:"required#优惠券分类名称不能为空" dc:"优惠券分类名称"`
	Price      int    `json:"price " v:"required#优惠券金额" dc:"优惠券金额"`
	GoodsIds   string `json:"goods_ids" dc:"可用商品Id 多个做好分隔"`
	CategoryId int    `json:"category_id" dc:"可用的商品分类"`
}

type CouponReq struct {
	g.Meta `path:"/v1/coupon/add" method:"post" tags:"优惠券分类" summary:"添加优惠券分类"`
	CouponCommon
}
type CouponRes struct {
	CouponId int `json:"coupon_id"`
}

type CouponDeleteReq struct {
	g.Meta `path:"/v1/coupon/delete" method:"delete" tags:"优惠券分类" summary:"删除优惠券分类分类接口"`
	Id     uint `v:"min:1#请选择需要删除的优惠券分类" dc:"优惠券分类id"`
}

// 自定义返回内容
type CouponDeleteRes struct {
	Id int `json:"id"`
}

type CouponUpdateReq struct {
	g.Meta `path:"/v1/coupon/{Id}" method:"post" tags:"优惠券分类" summary:"修改优惠券分类分类接口"`
	Id     int `json:"id"    v:"required#内容Id不能为空" dc:"ID"`
	CouponCommon
}

// 自定义返回内容
type CouponUpdateRes struct {
	Id int `json:"id"`
}

type CouponGetListCommonReq struct {
	g.Meta `path:"/v1/coupon/list" method:"get" tags:"优惠券分类" summary:"优惠券列表"`
	Sort   int `json:"sort"   in:"query" dc:"排序类型"`
	CommonPaginationReq
}
type CouponGetListCommonRes struct {
	//todo
	List  interface{}    `json:"list"`  // 列表
	Stats map[string]int `json:"stats"` // 搜索统计
	Page  int            `json:"page"`  // 分页码
	Size  int            `json:"size"`  // 分页数量
	Total int            `json:"total"` // 数据总数
	//g.Meta `mime:"text/html" type:"string" example:"<html/>"`
}
