// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// AddressInfo is the golang structure for table address_info.
type AddressInfo struct {
	Id        int         `json:"id"        orm:"id"         ` //
	Name      string      `json:"name"      orm:"name"       ` //
	Pid       int         `json:"pid"       orm:"pid"        ` //
	Status    int         `json:"status"    orm:"status"     ` //
	UpdatedAt *gtime.Time `json:"updatedAt" orm:"updated_at" ` //
}
