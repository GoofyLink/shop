// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// PositionInfo is the golang structure for table position_info.
type PositionInfo struct {
	Id        int         `json:"id"        orm:"id"         ` //
	PicUrl    string      `json:"picUrl"    orm:"pic_url"    ` // 图片链接
	GoodsName string      `json:"goodsName" orm:"goods_name" ` // 商品名称
	Link      string      `json:"link"      orm:"link"       ` // 跳转连接
	Sort      int         `json:"sort"      orm:"sort"       ` // 排序
	GoodsId   int         `json:"goodsId"   orm:"goods_id"   ` // 商品id
	CreatedAt *gtime.Time `json:"createdAt" orm:"created_at" ` //
	UpdatedAt *gtime.Time `json:"updatedAt" orm:"updated_at" ` //
	DeletedAt *gtime.Time `json:"deletedAt" orm:"deleted_at" ` //
}
