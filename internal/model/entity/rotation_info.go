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
	PicUrl    string      `json:"picUrl"    orm:"pic_url"    ` //
	Link      string      `json:"link"      orm:"link"       ` //
	Sort      int         `json:"sort"      orm:"sort"       ` //
	CreateAt  *gtime.Time `json:"createAt"  orm:"create_at"  ` //
	UpdateAt  *gtime.Time `json:"updateAt"  orm:"update_at"  ` //
	DeletedAt *gtime.Time `json:"deletedAt" orm:"deleted_at" ` //
}
