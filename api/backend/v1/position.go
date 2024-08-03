package v1

import (
	"github.com/gogf/gf/v2/frame/g"
)

type PositionGetListCommonReq struct {
	g.Meta `path:"/backend/v1/position/List" method:"get" tags:"内容" summary:"商品列表"`
	Sort   int `json:"sort"   in:"query" dc:"排序类型"`
	CommonPaginationReq
}
type PositionGetListCommonRes struct {
	//todo
	List  interface{}    `json:"list"`  // 列表
	Stats map[string]int `json:"stats"` // 搜索统计
	Page  int            `json:"page"`  // 分页码
	Size  int            `json:"size"`  // 分页数量
	Total int            `json:"total"` // 数据总数
	//g.Meta `mime:"text/html" type:"string" example:"<html/>"`
}

type PositionCreateReq struct {
	g.Meta    `path:"/backend/v1/position/add" method:"post" tags:"内容" summary:"图片接口"`
	Id        int    `json:"id"    v:"required#内容Id不能为空" dc:"ID"`
	Pic_url   string `json:"pic_url"    v:"required#图片"      dc:"图片地址"`
	Link      string `json:"link"   v:"required#图片跳转链接"      dc:"图片跳转连接"`
	Sort      int    `json:"sort" dc:"排序"`
	GoodsName string `json:"goods_name" v:"required#商品名称" dc:"商品名称"`
	GoodsId   int    `json:"goods_id" v:"required#商品id" dc:"商品id"`
}

type PositionCreateRes struct {
	PositionId int `json:"position_id"`
}

type PositionDeleteReq struct {
	g.Meta `path:"/backend/v1/position/delete" method:"delete" tags:"商品" summary:"删除商品接口"`
	Id     uint `v:"min:1#请选择需要删除的内容" dc:"内容id"`
}

// 自定义返回内容
type PositionDeleteRes struct {
	Id int `json:"id"`
}

type PositionUpdateReq struct {
	g.Meta    `path:"/backend/v1/position/{Id}" method:"post" tags:"商品" summary:"修改商品接口"`
	Id        int    `json:"id"    v:"required#内容Id不能为空" dc:"ID"`
	Pic_url   string `json:"pic_url"    v:"required#图片链接"      dc:"图片地址"`
	Link      string `json:"link"   v:"required#图片跳转链接"      dc:"图片跳转连接"`
	Sort      int    `json:"sort" dc:"排序"`
	GoodsName string `json:"goods_name" v:"required#商品名称" dc:"商品名称"`
	GoodsId   int    `json:"goods_id" v:"required#商品id" dc:"商品id"`
}

// 自定义返回内容
type PositionUpdateRes struct {
	Id int `json:"id"`
}
