// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// ConsigneeInfo is the golang structure for table consignee_info.
type ConsigneeInfo struct {
	Id        int         `json:"id"        orm:"id"         ` // 收货地址表
	UserId    int         `json:"userId"    orm:"user_id"    ` //
	IsDefault int         `json:"isDefault" orm:"is_default" ` // 默认地址1  非默认0
	Name      string      `json:"name"      orm:"name"       ` //
	Phone     string      `json:"phone"     orm:"phone"      ` //
	Province  string      `json:"province"  orm:"province"   ` //
	City      string      `json:"city"      orm:"city"       ` //
	Town      string      `json:"town"      orm:"town"       ` // 县区
	Street    string      `json:"street"    orm:"street"     ` // 街道乡镇
	Detail    string      `json:"detail"    orm:"detail"     ` // 地址详情
	CreatedAt *gtime.Time `json:"createdAt" orm:"created_at" ` //
	UpdatedAt *gtime.Time `json:"updatedAt" orm:"updated_at" ` //
	DeletedAt *gtime.Time `json:"deletedAt" orm:"deleted_at" ` //
}
