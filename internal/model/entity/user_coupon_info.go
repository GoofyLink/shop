// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// UserCouponInfo is the golang structure for table user_coupon_info.
type UserCouponInfo struct {
	Id        int         `json:"id"        orm:"id"         ` // 用户优惠券表
	UserId    int         `json:"userId"    orm:"user_id"    ` //
	CouponId  int         `json:"couponId"  orm:"coupon_id"  ` //
	Status    int         `json:"status"    orm:"status"     ` // 状态：1可用 2已用 3过期
	CreatedAt *gtime.Time `json:"createdAt" orm:"created_at" ` //
	UpdatedAt *gtime.Time `json:"updatedAt" orm:"updated_at" ` //
	DeletedAt *gtime.Time `json:"deletedAt" orm:"deleted_at" ` //
}
