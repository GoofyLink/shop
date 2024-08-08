package model

import "github.com/gogf/gf/v2/os/gtime"

// RoleCreateUpdateBase 创建/修改内容基类
type RoleCreateUpdateBase struct {
	Name string
	Desc string
}

// RoleCreateInput 创建内容
type RoleCreateInput struct {
	RoleCreateUpdateBase
	UserId int
}

// RoleCreateOutput 创建内容返回结果
type RoleCreateOutput struct {
	RoleId int `json:"role_id"`
}

// RoleUpdateInput 修改内容
type RoleUpdateInput struct {
	RoleCreateUpdateBase
	Id int
}

// RoleGetListInput 获取内容列表
type RoleGetListInput struct {
	Page int // 分页号码
	Size int // 分页数量，最大50
	Sort int // 排序类型(0:最新, 默认。1:活跃, 2:热度)
}

// RoleGetListOutput 查询列表结果
type RoleGetListOutput struct {
	List  []RoleGetListOutputItem `json:"list" description:"列表"`
	Page  int                     `json:"page" description:"分页码"`
	Size  int                     `json:"size" description:"分页数量"`
	Total int                     `json:"total" description:"数据总数"`
}

// 返回接口数据
type RoleGetListOutputItem struct {
	//admin *AdminListItem `json:"rotation"`
	//Category *AdminListCategoryItem `json:"category"`
	//User     *AdminListUserItem     `json:"user"`
	Id        uint        `json:"id"` // 自增ID
	Name      string      `json:"name"`
	RoleIds   string      `json:"role_ids"`
	IsAdmin   int         `json:"is_admin"`
	CreatedAt *gtime.Time `json:"created_at"` // 创建时间
	UpdatedAt *gtime.Time `json:"updated_at"` // 修改时间
}
