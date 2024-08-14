package model

import "shop-v2/internal/model/entity"

// GoodsCreateUpdateBase 创建/修改内容基类
type GoodsCreateUpdateBase struct {
	PicUrl           string // 图片
	Name             string // 商品名称
	Price            int    // 价格 单位分
	Level1CategoryId int    // 1级分类id
	Level2CategoryId int    // 2级分类id
	Level3CategoryId int    // 3级分类id
	Brand            string // 品牌
	Stock            int    // 库存
	Sale             int    // 销量
	Tags             string // 标签
	DetailInfo       string // 商品详情
}

// GoodsCreateInput 创建内容
type GoodsCreateInput struct {
	GoodsCreateUpdateBase
	UserId int
}

// GoodsCreateOutput 创建内容返回结果
type GoodsCreateOutput struct {
	GoodsId int `json:"id"`
}

// GoodsUpdateInput 修改内容
type GoodsUpdateInput struct {
	GoodsCreateUpdateBase
	Id int
}

// GoodsGetListInput 获取内容列表
type GoodsGetListInput struct {
	Page      int // 分页号码
	Size      int // 分页数量，最大50
	Sort      int // 排序类型(0:最新, 默认。1:活跃, 2:热度)
	GoodsName string
	GoodsId   int
}

// GoodsGetListOutput 查询列表结果
type GoodsGetListOutput struct {
	List  []GoodsGetListOutputItem `json:"list" description:"列表"`
	Page  int                      `json:"page" description:"分页码"`
	Size  int                      `json:"size" description:"分页数量"`
	Total int                      `json:"total" description:"数据总数"`
}

// 查询列表结果
type GoodsGetListOutputItem struct {
	entity.GoodsInfo //复用entity todo
}
