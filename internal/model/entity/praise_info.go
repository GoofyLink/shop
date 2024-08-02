// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// PraiseInfo is the golang structure for table praise_info.
type PraiseInfo struct {
	Id        int         `json:"id"        orm:"id"         ` // 点赞表
	UserId    int         `json:"userId"    orm:"user_id"    ` //
	Type      int         `json:"type"      orm:"type"       ` // 点赞类型 1商品 2文章
	ObjectId  int         `json:"objectId"  orm:"object_id"  ` // 点赞对象id 方便后期扩展
	CreatedAt *gtime.Time `json:"createdAt" orm:"created_at" ` //
	UpdatedAt *gtime.Time `json:"updatedAt" orm:"updated_at" ` //
}
