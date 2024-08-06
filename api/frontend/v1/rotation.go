package v1

import (
	"github.com/gogf/gf/v2/frame/g"
)

type RotationGetListCommonReq struct {
	g.Meta `path:"/frontend/v1/rotation/List" method:"get" tags:"内容" summary:"轮播图列表"`
	Sort   int `json:"sort"   in:"query" dc:"排序类型"`
	CommonPaginationReq
}
type RotationGetListCommonRes struct {
	//todo
	List  interface{}    `json:"list"`  // 列表
	Stats map[string]int `json:"stats"` // 搜索统计
	Page  int            `json:"page"`  // 分页码
	Size  int            `json:"size"`  // 分页数量
	//Total int            `json:"total"` // 数据总数
	//g.Meta `mime:"text/html" type:"string" example:"<html/>"`
}
