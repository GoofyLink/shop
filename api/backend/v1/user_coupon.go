package v1

import (
	"github.com/gogf/gf/v2/frame/g"
)

type UserCouponCommon struct {
	UserId   int   `json:"user_id"  v:"required#用户ID必传" dc:"用户优惠券名称"`
	CouponId int   `json:"coupon_id " v:"required#优惠券ID必传" dc:"用户优惠券金额"`
	Status   uint8 `json:"status" dc:"状态：1可用 2已用 3过期"`
}

type UserCouponReq struct {
	g.Meta `path:"/v1/user/coupon/add" method:"post" tags:"用户优惠券" summary:"添加用户优惠券"`
	UserCouponCommon
}
type UserCouponRes struct {
	Id int `json:"id"`
}

type UserCouponDeleteReq struct {
	g.Meta `path:"/v1/user/coupon/delete" method:"delete" tags:"用户优惠券" summary:"删除用户优惠券分类接口"`
	Id     uint `v:"min:1#请选择需要删除的用户优惠券" dc:"用户优惠券id"`
}

// 自定义返回内容
type UserCouponDeleteRes struct {
	Id int `json:"id"`
}

type UserCouponUpdateReq struct {
	g.Meta `path:"/v1/user/coupon/{Id}" method:"post" tags:"用户优惠券" summary:"修改用户优惠券分类接口"`
	Id     int `json:"id"    v:"required#内容Id不能为空" dc:"ID"`
	UserCouponCommon
}

// 自定义返回内容
type UserCouponUpdateRes struct {
	Id int `json:"id"`
}

type UserCouponGetListCommonReq struct {
	g.Meta `path:"/v1/user/coupon/list" method:"get" tags:"用户优惠券" summary:"用户优惠券列表"`
	CommonPaginationReq
}
type UserCouponGetListCommonRes struct {
	//todo
	List  interface{}    `json:"list"`  // 列表
	Stats map[string]int `json:"stats"` // 搜索统计
	Page  int            `json:"page"`  // 分页码
	Size  int            `json:"size"`  // 分页数量
	Total int            `json:"total"` // 数据总数
}
