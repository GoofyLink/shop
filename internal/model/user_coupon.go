package model

import "github.com/gogf/gf/v2/os/gtime"

// UserCouponCreateUpdateBase 创建/修改内容基类
type UserCouponCreateUpdateBase struct {
	UserId   int
	CouponId int
	Status   uint8
}

// UserCouponCreateInput 创建内容
type UserCouponCreateInput struct {
	UserCouponCreateUpdateBase
	UserId int
}

// UserCouponCreateOutput 创建内容返回结果
type UserCouponCreateOutput struct {
	UserCouponId int `json:"id"`
}

// UserCouponUpdateInput 修改内容
type UserCouponUpdateInput struct {
	UserCouponCreateUpdateBase
	Id int
}

// UserCouponGetListInput 获取内容列表
type UserCouponGetListInput struct {
	Page      int // 分页号码
	Size      int // 分页数量，最大50
	Sort      int // 排序类型(0:最新, 默认。1:活跃, 2:热度)
	GoodsName string
	GoodsId   int
}

// UserCouponGetListOutput 查询列表结果
type UserCouponGetListOutput struct {
	List  []UserCouponGetListOutputItem `json:"list" description:"列表"`
	Page  int                           `json:"page" description:"分页码"`
	Size  int                           `json:"size" description:"分页数量"`
	Total int                           `json:"total" description:"数据总数"`
}

// 查询列表结果
type UserCouponGetListOutputItem struct {
	//position *UserCouponListItem `json:"position"`
	//UserCoupon *UserCouponListUserCouponItem `json:"category"`
	//User     *UserCouponListUserItem     `json:"user"`
	Id        int         `json:"id"`
	UserId    int         `json:"user_id"`
	CouponId  int         `json:"coupon_id"` // 排序，数值
	Status    int         `json:"status"`
	CreatedAt *gtime.Time `json:"created_at"` // 创建时间
	UpdatedAt *gtime.Time `json:"updated_at"` // 修改时间
}
