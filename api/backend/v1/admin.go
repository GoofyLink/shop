package v1

import (
	"github.com/gogf/gf/v2/frame/g"
)

type AdminCreateReq struct {
	g.Meta   `path:"/v1/admin/add" method:"post" tags:"用户" summary:"管理员"`
	Name     string `json:"name"        v:"required#姓名"      dc:"姓名"`
	Password string `json:"password"    v:"required#密码"      dc:"密码"`
	RoleIds  string `json:"role_ids"    dc:"角色权限"`
	IsAdmin  int    `json:"goods_name"  dc:"是否超级管理员"`
}

type AdminCreateRes struct {
	AdminId int `json:"admin_id"`
}

type AdminDeleteReq struct {
	g.Meta `path:"/v1/admin/delete" method:"delete" tags:"管理员" summary:"删除管理员接口"`
	Id     uint `v:"min:1#请选择需要删除的内容" dc:"内容id"`
}

// 自定义返回内容
type AdminDeleteRes struct {
	Id int `json:"id"`
}

type AdminUpdateReq struct {
	g.Meta   `path:"/v1/admin/{Id}" method:"post" tags:"管理员" summary:"修改管理员接口"`
	Id       int    `json:"id"    v:"required#内容Id不能为空" dc:"ID"`
	Name     string `json:"name"        v:"required#姓名"      dc:"姓名"`
	Password string `json:"password"    v:"required#密码"      dc:"密码"`
	RoleIds  string `json:"role_ids"    dc:"角色权限"`
	IsAdmin  int    `json:"goods_name"  dc:"是否超级管理员"`
}

// 自定义返回内容
type AdminUpdateRes struct {
	Id int `json:"id"`
}

type AdminGetListCommonReq struct {
	g.Meta `path:"/v1/admin/List" method:"get" tags:"内容" summary:"商品列表"`
	CommonPaginationReq
}
type AdminGetListCommonRes struct {
	//todo
	List  interface{}    `json:"list"`  // 列表
	Stats map[string]int `json:"stats"` // 搜索统计
	Page  int            `json:"page"`  // 分页码
	Size  int            `json:"size"`  // 分页数量
	Total int            `json:"total"` // 数据总数
}

// 用户信息
type AdminGetInfoReq struct {
	g.Meta `path:"/v1/admin/info" method:"post"`
}

// jwt
//type AdminGetInfoRes struct {
//	Id          int    `json:"id"`
//	IdentityKey string `json:"identity_key"`
//	Payload     string `json:"payload"`
//}

// for token
type AdminGetInfoRes struct {
	Id      int    `json:"id"`
	Name    string `json:"name"        v:"required#姓名"      dc:"姓名"`
	RoleIds string `json:"role_ids"    dc:"角色权限"`
	IsAdmin int    `json:"goods_name"  dc:"是否超级管理员"`
}
