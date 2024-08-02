// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// RotationInfo is the golang structure for table rotation_info.
type RotationInfo struct {
	Id        int         `json:"id"        orm:"id"         ` //
	PicUrl    string      `json:"picUrl"    orm:"pic_url"    ` // 轮播图片
	Link      string      `json:"link"      orm:"link"       ` // 跳转链接
	Sort      int         `json:"sort"      orm:"sort"       ` // 排序字段
	CreatedAt *gtime.Time `json:"createdAt" orm:"created_at" ` //
	UpdatedAt *gtime.Time `json:"updatedAt" orm:"updated_at" ` //
	DeletedAt *gtime.Time `json:"deletedAt" orm:"deleted_at" ` //
}
