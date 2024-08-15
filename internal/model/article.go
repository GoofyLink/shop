package model

import "shop-v2/internal/model/entity"

// ArticleCreateUpdateBase 创建/修改内容基类
type ArticleCreateUpdateBase struct {
	UserId  int
	Title   string
	Desc    int
	PicUrl  string
	IsAdmin uint
	Detail  int
	Praise  int
	Stock   int
}

// ArticleCreateInput 创建内容
type ArticleCreateInput struct {
	ArticleCreateUpdateBase
	UserId int
}

// ArticleCreateOutput 创建内容返回结果
type ArticleCreateOutput struct {
	ArticleId int `json:"id"`
}

// ArticleUpdateInput 修改内容
type ArticleUpdateInput struct {
	ArticleCreateUpdateBase
	Id int
}

// ArticleGetListInput 获取内容列表
type ArticleGetListInput struct {
	Page        int // 分页号码
	Size        int // 分页数量，最大50
	Sort        int // 排序类型(0:最新, 默认。1:活跃, 2:热度)
	ArticleName string
	ArticleId   int
}

// ArticleGetListOutput 查询列表结果
type ArticleGetListOutput struct {
	List  []ArticleGetListOutputItem `json:"list" description:"列表"`
	Page  int                        `json:"page" description:"分页码"`
	Size  int                        `json:"size" description:"分页数量"`
	Total int                        `json:"total" description:"数据总数"`
}

// 查询列表结果
type ArticleGetListOutputItem struct {
	entity.ArticleInfo //复用entity todo
}
