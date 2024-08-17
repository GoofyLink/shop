package model

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// 添加收藏
type AddCollectionInput struct {
	UserId   uint
	ObjectId uint
	Type     int8
}

type AddCollectionOutPut struct {
	Id uint
}

// 移除收藏
type DeleteCollectionInput struct {
	Id       uint
	UserId   uint
	Type     uint8
	ObjectId uint
}

type DeleteCollectionOutPut struct {
	Id uint
}

type CollectionListInput struct {
	Page int   // 分页号码
	Size int   // 分页数量，最大50
	Type uint8 //
}

type CollectionListOutput struct {
	List  []CollectionListOutputItem
	Page  int
	Size  int
	Total int
}

type CollectionListOutputItem struct {
	Id        uint // 自增ID
	UserId    uint
	Type      uint8
	ObjectId  uint
	Goods     GoodsItem   `json:"goods" orm:"with:id=object_id"`
	Article   ArticleItem `json:"article" orm:"with:id=object_id"`
	CreatedAt *gtime.Time // 创建时间
	UpdatedAt *gtime.Time // 修改时间
}

type GoodsItem struct {
	g.Meta `orm:"table:goods_info"`
	Id     uint   `json:"id"`
	Name   string `json:"name"`
	PicUrl string `json:"pic_url"`
	Price  int    `json:"price"`
}

type ArticleItem struct {
	g.Meta `orm:"table:article_info"`
	Id     uint   `json:"id"`
	Title  string `json:"title"`
	Desc   string `json:"desc"`
	PicUrl string `json:"pic_url"`
}

// 校验当前用户是否被收藏
type CheckIsCollectionInput struct {
	UserId   uint
	ObjectId uint
	Type     uint8
}