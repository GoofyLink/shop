// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// CategoryInfo is the golang structure for table category_info.
type CategoryInfo struct {
	Id        int         `json:"id"        orm:"id"         ` //
	ParentId  int         `json:"parentId"  orm:"parent_id"  ` // 父级id
	Name      string      `json:"name"      orm:"name"       ` //
	PicUrl    string      `json:"picUrl"    orm:"pic_url"    ` // icon
	Level     int         `json:"level"     orm:"level"      ` // 等级 默认1级分类
	Sort      int         `json:"sort"      orm:"sort"       ` //
	CreatedAt *gtime.Time `json:"createdAt" orm:"created_at" ` //
	UpdatedAt *gtime.Time `json:"updatedAt" orm:"updated_at" ` //
	DeletedAt *gtime.Time `json:"deletedAt" orm:"deleted_at" ` //
}
