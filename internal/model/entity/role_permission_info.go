// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// RolePermissionInfo is the golang structure for table role_permission_info.
type RolePermissionInfo struct {
	Id           int         `json:"id"           orm:"id"            ` //
	RoleId       int         `json:"roleId"       orm:"role_id"       ` // 角色id
	PermissionId int         `json:"permissionId" orm:"permission_id" ` // 权限id
	CreatedAt    *gtime.Time `json:"createdAt"    orm:"created_at"    ` //
	UpdatedAt    *gtime.Time `json:"updatedAt"    orm:"updated_at"    ` //
}
