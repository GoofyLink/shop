package v1

import "github.com/gogf/gf/v2/frame/g"

type AddCollectionReq struct {
	g.Meta   `path:"/v1/add/collection" method:"post" tags:"前台收藏" summary:"添加收藏"`
	UserId   uint `json:"userId"    dc:"登录用户ID"    `                           // 用户id
	ObjectId uint `json:"objectId"  dc:"对象ID" v:"required#收藏的对象id必传" `         // 对象id
	Type     int8 `json:"type"      dc:"收藏类型：1商品 2文章 范围约束in"  v:"in:1,2"     ` // 收藏类型：1商品 2文章 范围约束in
}

type AddCollectionRes struct {
	Id uint `json:"id"`
}

type DeleteCollectionReq struct {
	g.Meta   `path:"/v1/delete/collection" method:"delete" tags:"前台收藏" summary:"删除收藏"`
	Id       uint  `json:"id" dc:"用户ID"`
	Type     uint8 `json:"type" dc:"类型 只有1和2 "`
	ObjectId uint  `json:"object_id" dc:"对象ID"`
}

type DeleteCollectionRes struct {
	Id uint `json:"id"`
}

type CollectionListReq struct {
	g.Meta `path:"/v1/list/collection" method:"post" tags:"前台收藏" summary:"收藏列表"`
	Type   uint8 `json:"type" v:"in:0,1,2" dc:"收藏类型"` // 默认0 收藏取消收藏就是1,2
	CommonPaginationReq
}

type CollectionListRes struct {
	Page  int         `json:"page"`  // 分页码
	Size  int         `json:"size"`  // 分页数量
	List  interface{} `json:"list"`  // 列表
	Total int         `json:"total"` // 数据总数
}

type ListCollectionItem struct {
	Id       int         `json:"id"        dc:"id"         ` //
	UserId   int         `json:"userId"    dc:"user_id"    ` // 用户id
	ObjectId int         `json:"objectId"  dc:"object_id"  ` // 对象id
	Type     int         `json:"type"      dc:"type"       ` // 收藏类型：1商品 2文章
	Goods    interface{} `json:"goods"`
	Article  interface{} `json:"article"`
}
