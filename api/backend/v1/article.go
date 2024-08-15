package v1

import "github.com/gogf/gf/v2/frame/g"

type ArticleCommon struct {
	UserId  int    `json:"user_id"  v:"required#文章ID必传"          dc:"userId"     summary:"文章管理"  `
	Title   string `json:"title"           dc:"文章标题"`
	Desc    int    `json:"desc"            dc:"摘要"  `
	PicUrl  string `json:"pic_url"          dc:"封面图" `
	IsAdmin uint   `json:"is_admin"   d:"1"      dc:"1后台管理员发布 2前台用户发布" `
	Detail  int    `json:"detail"   v:"required#文章详情必填"          dc:"文章详情"    `
	Praise  int    `json:"praise"             dc:"点赞数量"               `
}

type ArticleReq struct {
	g.Meta `path:"/v1/article/add" method:"post" tags:"文章" summary:"添加文章"`
	GoodsCommon
}
type ArticleRes struct {
	Id int `json:"id"`
}

type ArticleDeleteReq struct {
	g.Meta `path:"/v1/article/delete" method:"delete" tags:"文章" summary:"删除文章管理接口"`
	Id     uint `v:"min:1#请选择需要删除的文章" dc:"文章id"`
}

// 自定义返回内容
type ArticleDeleteRes struct {
	Id int `json:"id"`
}

type ArticleUpdateReq struct {
	g.Meta `path:"/v1/article/{Id}" method:"post" tags:"文章" summary:"修改文章管理接口"`
	Id     int `json:"id"    v:"required#内容Id不能为空" dc:"ID"`
	GoodsCommon
}

// 自定义返回内容
type ArticleUpdateRes struct {
	Id int `json:"id"`
}

type ArticleGetListCommonReq struct {
	g.Meta `path:"/v1/article/list" method:"get" tags:"文章" summary:"文章列表"`
	CommonPaginationReq
}
type ArticleGetListCommonRes struct {
	//todo
	List  interface{}    `json:"list"`  // 列表
	Stats map[string]int `json:"stats"` // 搜索统计
	Page  int            `json:"page"`  // 分页码
	Size  int            `json:"size"`  // 分页数量
	Total int            `json:"total"` // 数据总数
}
