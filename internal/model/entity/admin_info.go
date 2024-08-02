// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// AdminInfo is the golang structure for table admin_info.
type AdminInfo struct {
	Id        int         `json:"id"        orm:"id"         ` //
	Name      string      `json:"name"      orm:"name"       ` // 用户名
	Password  string      `json:"password"  orm:"password"   ` // 密码
	RoleIds   string      `json:"roleIds"   orm:"role_ids"   ` // 角色ids
	CreatedAt *gtime.Time `json:"createdAt" orm:"created_at" ` //
	UpdatedAt *gtime.Time `json:"updatedAt" orm:"updated_at" ` //
	UserSalt  string      `json:"userSalt"  orm:"user_salt"  ` // 加密盐
	IsAdmin   int         `json:"isAdmin"   orm:"is_admin"   ` // 是否超级管理员
}
