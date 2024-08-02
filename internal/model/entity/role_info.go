// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// RoleInfo is the golang structure for table role_info.
type RoleInfo struct {
	Id        int         `json:"id"        orm:"id"         ` //
	Name      string      `json:"name"      orm:"name"       ` // 角色名称
	Desc      string      `json:"desc"      orm:"desc"       ` // 描述
	CreatedAt *gtime.Time `json:"createdAt" orm:"created_at" ` //
	UpdatedAt *gtime.Time `json:"updatedAt" orm:"updated_at" ` //
	DeletedAt *gtime.Time `json:"deletedAt" orm:"deleted_at" ` //
}
