package v1

import "github.com/gogf/gf/v2/frame/g"

type AddCommentReq struct {
	g.Meta   `path:"/v1/add/comment" method:"post" tags:"前台评论" summary:"添加评论"` // 用户id
	ObjectId uint                                                              `json:"objectId"  dc:"对象ID" v:"required#评论的对象id必传" `         // 对象id
	Type     int8                                                              `json:"type"      dc:"评论类型：1商品 2文章 范围约束in"  v:"in:1,2"     ` // 评论类型：1商品 2文章 范围约束in
	Content  string                                                            `json:"content" dc:"评论内容"`
	ParentId int                                                               `json:"parent_id" dc:"父级内容"`
}

type AddCommentRes struct {
	Id uint `json:"id"`
}

type DeleteCommentReq struct {
	g.Meta `path:"/v1/delete/comment" method:"delete" tags:"前台评论" summary:"删除评论"`
	Id     uint `json:"id" dc:"用户ID"`
}

type DeleteCommentRes struct {
	Id uint `json:"id"`
}

type CommentListReq struct {
	g.Meta `path:"/v1/list/comment" method:"post" tags:"前台评论" summary:"评论列表"`
	Type   uint8 `json:"type" v:"in:0,1,2" dc:"评论类型"` // 默认0 评论取消评论就是1,2
	CommonPaginationReq
}

type CommentListRes struct {
	Page  int         `json:"page"`  // 分页码
	Size  int         `json:"size"`  // 分页数量
	List  interface{} `json:"list"`  // 列表
	Total int         `json:"total"` // 数据总数
}

type ListCommentItem struct {
	Id       int         `json:"id"        dc:"id"         ` //
	UserId   int         `json:"userId"    dc:"user_id"    ` // 用户id
	ObjectId int         `json:"objectId"  dc:"object_id"  ` // 对象id
	Type     int         `json:"type"      dc:"type"       ` // 评论类型：1商品 2文章
	Goods    interface{} `json:"goods"`
	Article  interface{} `json:"article"`
}
