package v1

import "github.com/gogf/gf/v2/frame/g"

type AddPraiseReq struct {
	g.Meta   `path:"/v1/add/praise" method:"post" tags:"前台点赞" summary:"添加点赞"`
	ObjectId uint `json:"objectId"  dc:"对象ID" v:"required#点赞的对象id必传" `         // 对象id
	Type     int8 `json:"type"      dc:"点赞类型：1商品 2文章 范围约束in"  v:"in:1,2"     ` // 点赞类型：1商品 2文章 范围约束in
}

type AddPraiseRes struct {
	Id uint `json:"id"`
}

type DeletePraiseReq struct {
	g.Meta   `path:"/v1/delete/praise" method:"delete" tags:"前台点赞" summary:"删除点赞"`
	Id       uint  `json:"id" dc:"用户ID"`
	Type     uint8 `json:"type" dc:"类型 只有1和2 "`
	ObjectId uint  `json:"object_id" dc:"对象ID"`
}

type DeletePraiseRes struct {
	Id uint `json:"id"`
}

type PraiseListReq struct {
	g.Meta `path:"/v1/list/praise" method:"post" tags:"前台点赞" summary:"点赞列表"`
	Type   uint8 `json:"type" v:"in:0,1,2" dc:"点赞类型"` // 默认0 点赞取消点赞就是1,2
	CommonPaginationReq
}

type PraiseListRes struct {
	Page  int         `json:"page"`  // 分页码
	Size  int         `json:"size"`  // 分页数量
	List  interface{} `json:"list"`  // 列表
	Total int         `json:"total"` // 数据总数
}

type ListPraiseItem struct {
	Id       int         `json:"id"        dc:"id"         ` //
	UserId   int         `json:"userId"    dc:"user_id"    ` // 用户id
	ObjectId int         `json:"objectId"  dc:"object_id"  ` // 对象id
	Type     int         `json:"type"      dc:"type"       ` // 点赞类型：1商品 2文章
	Goods    interface{} `json:"goods"`
	Article  interface{} `json:"article"`
}
