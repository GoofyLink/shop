package v1

import "github.com/gogf/gf/v2/frame/g"

// 创建角色
type RoleReq struct {
	g.Meta `path:"/v1/role/add" method:"post" desc:"添加角色" tags:"role"`
	Name   string `json:"name"  v:"required#姓名必填" dc:"角色名称"`
	Desc   string `json:"desc" dc:"角色描述"`
}

type RoleRes struct {
	RoleId int `json:"role_id"`
}

// 修改编辑角色
type RoleUpdateReq struct {
	g.Meta `path:"/v1/role/update" method:"post" desc:"修改角色" tags:"role"`
	Id     int    `json:"id" v:"required#id必填" dc:"角色id"`
	Name   string `json:"name"  v:"required#姓名必填" dc:"角色名称"`
	Desc   string `json:"desc" dc:"角色描述"`
}

type RoleUpdateRes struct {
	RoleId int `json:"role_id"`
}

// 删除角色
type RoleDeleteReq struct {
	g.Meta `path:"/v1/role/delete" method:"delete" tags:"角色" summary:"删除角色接口"`
	Id     uint `v:"min:1#请选择需要删除的内容" dc:"内容id"`
}

// 自定义返回内容
type RoleDeleteRes struct {
	Id int `json:"id"`
}

// 获取角色列表
type RoleGetListCommonReq struct {
	g.Meta `path:"/v1/role/List" method:"get" tags:"角色" summary:"商品列表"`
	CommonPaginationReq
}
type RoleGetListCommonRes struct {
	//todo
	List  interface{}    `json:"list"`  // 列表
	Stats map[string]int `json:"stats"` // 搜索统计
	Page  int            `json:"page"`  // 分页码
	Size  int            `json:"size"`  // 分页数量
	Total int            `json:"total"` // 数据总数
}

// 用户信息
type RoleGetInfoReq struct {
	g.Meta `path:"/v1/role/info" method:"post"`
}

type RoleGetInfoRes struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
	Desc string `json:"desc"`
}
