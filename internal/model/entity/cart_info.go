// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// CartInfo is the golang structure for table cart_info.
type CartInfo struct {
	Id             int         `json:"id"             orm:"id"               ` // 购物车表
	UserId         int         `json:"userId"         orm:"user_id"          ` //
	GoodsOptionsId int         `json:"goodsOptionsId" orm:"goods_options_id" ` // 商品规格id
	Count          int         `json:"count"          orm:"count"            ` // 商品数量
	CreatedAt      *gtime.Time `json:"createdAt"      orm:"created_at"       ` //
	UpdatedAt      *gtime.Time `json:"updatedAt"      orm:"updated_at"       ` //
	DeletedAt      *gtime.Time `json:"deletedAt"      orm:"deleted_at"       ` //
}
