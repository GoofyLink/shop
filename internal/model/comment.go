package model

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// 添加收藏
type AddCommentInput struct {
	UserId   uint
	ObjectId uint
	Type     int8
	ParentId int
	Content  string
}

type AddCommentOutPut struct {
	Id uint
}

// 移除收藏
type DeleteCommentInput struct {
	Id uint
}

type DeleteCommentOutPut struct {
	Id uint
}

type CommentListInput struct {
	Page int   // 分页号码
	Size int   // 分页数量，最大50
	Type uint8 //
}

type CommentListOutput struct {
	List  []CommentListOutputItem
	Page  int
	Size  int
	Total int
}

type CommentListOutputItem struct {
	Id        uint // 自增ID
	UserId    uint
	Type      uint8
	ObjectId  uint
	ParentId  int // 父级评论
	Content   string
	Goods     GoodsItem   `json:"goods" orm:"with:id=object_id"`
	Article   ArticleItem `json:"article" orm:"with:id=object_id"`
	CreatedAt *gtime.Time // 创建时间
	UpdatedAt *gtime.Time // 修改时间
}

// 校验当前用户是否被收藏
type CheckIsCommentInput struct {
	UserId   uint
	ObjectId uint
	Type     uint8
}
