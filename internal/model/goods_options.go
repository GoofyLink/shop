package model

import "shop-v2/internal/model/entity"

// GoodsOptionsCreateUpdateBase 创建/修改内容基类
type GoodsOptionsCreateUpdateBase struct {
	GoodsId int    // 商品ID 不可能为负数 uint
	Name    string // 商品名称
	Price   int    // 价格 单位分
	Brand   string // 品牌
	Stock   int    // 库存
	Sale    int    // 销量
}

// GoodsOptionsCreateInput 创建内容
type GoodsOptionsCreateInput struct {
	GoodsOptionsCreateUpdateBase
	Id int
}

// GoodsOptionsCreateOutput 创建内容返回结果
type GoodsOptionsCreateOutput struct {
	GoodsOptionsId int `json:"id"`
}

// GoodsOptionsUpdateInput 修改内容
type GoodsOptionsUpdateInput struct {
	GoodsOptionsCreateUpdateBase
	Id int
}

// GoodsOptionsGetListInput 获取内容列表
type GoodsOptionsGetListInput struct {
	Page             int // 分页号码
	Size             int // 分页数量，最大50
	Sort             int // 排序类型(0:最新, 默认。1:活跃, 2:热度)
	GoodsOptionsName string
	GoodsOptionsId   int
}

// GoodsOptionsGetListOutput 查询列表结果
type GoodsOptionsGetListOutput struct {
	List  []GoodsOptionsGetListOutputItem `json:"list" description:"列表"`
	Page  int                             `json:"page" description:"分页码"`
	Size  int                             `json:"size" description:"分页数量"`
	Total int                             `json:"total" description:"数据总数"`
}

// 查询列表结果
type GoodsOptionsGetListOutputItem struct {
	entity.GoodsOptionsInfo //复用entity todo
}
