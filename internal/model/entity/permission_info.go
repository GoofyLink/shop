// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// PermissionInfo is the golang structure for table permission_info.
type PermissionInfo struct {
	Id        int         `json:"id"        orm:"id"         ` //
	Name      string      `json:"name"      orm:"name"       ` // 权限名称
	Path      string      `json:"path"      orm:"path"       ` // 路径
	CreatedAt *gtime.Time `json:"createdAt" orm:"created_at" ` //
	UpdatedAt *gtime.Time `json:"updatedAt" orm:"updated_at" ` //
	DeletedAt *gtime.Time `json:"deletedAt" orm:"deleted_at" ` //
}
