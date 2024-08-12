package v1

import "github.com/gogf/gf/v2/frame/g"

type PermissionCreateUpdateBase struct {
	Name string `json:"name"  v:"required#姓名必填" dc:"权限名称"`
	Path string `json:"desc" dc:"权限描述路径"`
}

// 创建权限
type PermissionReq struct {
	g.Meta `path:"/v1/permission/add" method:"post" desc:"添加权限" tags:"permission"`
	PermissionCreateUpdateBase
}

type PermissionRes struct {
	PermissionId int `json:"permission_id"`
}

// 修改编辑权限
type PermissionUpdateReq struct {
	g.Meta `path:"/v1/permission/update" method:"post" desc:"修改权限" tags:"permission"`
	Id     int `json:"id" v:"required#id必填" dc:"权限id"`
	PermissionCreateUpdateBase
}

type PermissionUpdateRes struct {
	PermissionId int `json:"permission_id"`
}

// 删除权限
type PermissionDeleteReq struct {
	g.Meta `path:"/v1/permission/delete" method:"delete" tags:"权限" summary:"删除权限接口"`
	Id     uint `v:"min:1#请选择需要删除的内容" dc:"内容id"`
}

// 自定义返回内容
type PermissionDeleteRes struct {
	Id int `json:"id"`
}

// 获取权限列表
type PermissionGetListCommonReq struct {
	g.Meta `path:"/v1/permission/List" method:"get" tags:"权限" summary:"商品列表"`
	CommonPaginationReq
}
type PermissionGetListCommonRes struct {
	//todo
	List  interface{}    `json:"list"`  // 列表
	Stats map[string]int `json:"stats"` // 搜索统计
	Page  int            `json:"page"`  // 分页码
	Size  int            `json:"size"`  // 分页数量
	Total int            `json:"total"` // 数据总数
}
