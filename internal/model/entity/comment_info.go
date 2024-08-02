// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// CommentInfo is the golang structure for table comment_info.
type CommentInfo struct {
	Id        int         `json:"id"        orm:"id"         ` //
	ParentId  int         `json:"parentId"  orm:"parent_id"  ` // 父级评论id
	UserId    int         `json:"userId"    orm:"user_id"    ` //
	ObjectId  int         `json:"objectId"  orm:"object_id"  ` //
	Type      int         `json:"type"      orm:"type"       ` // 评论类型：1商品 2文章
	Content   string      `json:"content"   orm:"content"    ` // 评论内容
	CreatedAt *gtime.Time `json:"createdAt" orm:"created_at" ` //
	UpdatedAt *gtime.Time `json:"updatedAt" orm:"updated_at" ` //
	DeletedAt *gtime.Time `json:"deletedAt" orm:"deleted_at" ` //
}
