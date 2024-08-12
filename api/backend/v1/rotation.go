package v1

import (
	"github.com/gogf/gf/v2/frame/g"
)

type RotationGetListCommonReq struct {
	g.Meta `path:"/v1/rotation/List" method:"get" tags:"内容" summary:"轮播图列表"`
	Sort   int `json:"sort"   in:"query" dc:"排序类型"`
	CommonPaginationReq
}
type RotationGetListCommonRes struct {
	//todo
	List  interface{}    `json:"list"`  // 列表
	Stats map[string]int `json:"stats"` // 搜索统计
	Page  int            `json:"page"`  // 分页码
	Size  int            `json:"size"`  // 分页数量
	Total int            `json:"total"` // 数据总数
	//g.Meta `mime:"text/html" type:"string" example:"<html/>"`
}

type RotationCreateReq struct {
	g.Meta  `path:"/v1/rotation/add" method:"post" tags:"内容" summary:"图片接口"`
	Id      int    `json:"id"    v:"required#内容Id不能为空" dc:"ID"`
	Pic_url string `json:"pic_url"    v:"required#图片"      dc:"图片地址"`
	Link    string `json:"link"   v:"required#图片跳转链接"      dc:"图片跳转连接"`
	Sort    int    `json:"sort" dc:"排序"`
}

type RotationCreateRes struct {
	RotationId int `json:"rotation_id"`
}

type RotationDeleteReq struct {
	g.Meta `path:"/v1/rotation/delete" method:"delete" tags:"轮播图" summary:"删除轮播图接口"`
	Id     uint `v:"min:1#请选择需要删除的内容" dc:"内容id"`
}

// 自定义返回内容
type RotationDeleteRes struct {
	Id int `json:"id"`
}

type RotationUpdateReq struct {
	g.Meta  `path:"/v1/rotation/{Id}" method:"post" tags:"轮播图" summary:"修改轮播图接口"`
	Id      int    `json:"id"    v:"required#内容Id不能为空" dc:"ID"`
	Pic_url string `json:"pic_url"    v:"required#图片链接"      dc:"图片地址"`
	Link    string `json:"link"   v:"required#图片跳转链接"      dc:"图片跳转连接"`
	Sort    int    `json:"sort" dc:"排序"`
}

// 自定义返回内容
type RotationUpdateRes struct {
	Id int `json:"id"`
}
