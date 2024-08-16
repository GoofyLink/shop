package model

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// 添加收藏
type AddPraiseInput struct {
	UserId   uint
	ObjectId uint
	Type     int8
}

type AddPraiseOutPut struct {
	Id uint
}

// 移除收藏
type DeletePraiseInput struct {
	Id       uint
	UserId   uint
	Type     uint8
	ObjectId uint
}

type DeletePraiseOutPut struct {
	Id uint
}

type PraiseListInput struct {
	Page int   // 分页号码
	Size int   // 分页数量，最大50
	Type uint8 //
}

type PraiseListOutput struct {
	List  []PraiseListOutputItem
	Page  int
	Size  int
	Total int
}

type PraiseListOutputItem struct {
	Id        uint // 自增ID
	UserId    uint
	Type      uint8
	ObjectId  uint
	Goods     GoodsItem   `json:"goods" orm:"with:id=object_id"`
	Article   ArticleItem `json:"article" orm:"with:id=object_id"`
	CreatedAt *gtime.Time // 创建时间
	UpdatedAt *gtime.Time // 修改时间
}

// 校验当前用户是否被收藏
type CheckIsPraiseInput struct {
	UserId   uint
	ObjectId uint
	Type     uint8
}
