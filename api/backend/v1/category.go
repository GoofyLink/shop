package v1

import (
	"github.com/gogf/gf/v2/frame/g"
)

type CategoryCommon struct {
	ParentId int    `json:"parent_id"   dc:"排序类型"`
	Name     string `json:"name"  v:"required#商品分类名称不能为空" dc:"商品分类名称"`
	PicUrl   string `json:"pic_url" dc:"商品分类图片链接"`
	Level    uint8  `json:"level" dc:"商品分类等级 默认等级1级"`
	Sort     uint8  `json:"sort" dc:"商品分类排序"`
}

type CategoryReq struct {
	g.Meta   `path:"/v1/category/add" method:"post" tags:"商品分类" summary:"添加商品分类"`
	ParentId uint   `json:"parent_id"   dc:"排序类型"`
	Name     string `json:"name"  v:"required#商品分类名称不能为空" dc:"商品分类名称"`
	PicUrl   string `json:"pic_url" dc:"商品分类图片链接"`
	Level    uint8  `json:"level" dc:"商品分类等级 默认等级1级"`
	Sort     uint8  `json:"sort" dc:"商品分类排序"`
}
type CategoryRes struct {
	CategoryId int `json:"category_id"`
}

type CategoryDeleteReq struct {
	g.Meta `path:"/v1/category/delete" method:"delete" tags:"商品分类" summary:"删除商品分类分类接口"`
	Id     uint `v:"min:1#请选择需要删除的商品分类" dc:"商品分类id"`
}

// 自定义返回内容
type CategoryDeleteRes struct {
	Id int `json:"id"`
}

type CategoryUpdateReq struct {
	g.Meta `path:"/v1/category/{Id}" method:"post" tags:"商品分类" summary:"修改商品分类分类接口"`
	Id     int `json:"id"    v:"required#内容Id不能为空" dc:"ID"`
	CategoryCommon
}

// 自定义返回内容
type CategoryUpdateRes struct {
	Id int `json:"id"`
}

type CategoryGetListCommonReq struct {
	g.Meta `path:"/v1/category/list" method:"get" tags:"商品分类" summary:"商品列表"`
	Sort   int `json:"sort"   in:"query" dc:"排序类型"`
	CommonPaginationReq
}
type CategoryGetListCommonRes struct {
	//todo
	List  interface{}    `json:"list"`  // 列表
	Stats map[string]int `json:"stats"` // 搜索统计
	Page  int            `json:"page"`  // 分页码
	Size  int            `json:"size"`  // 分页数量
	Total int            `json:"total"` // 数据总数
	//g.Meta `mime:"text/html" type:"string" example:"<html/>"`
}
