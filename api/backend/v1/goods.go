package v1

import (
	"github.com/gogf/gf/v2/frame/g"
)

type GoodsCommon struct {
	PicUrl           string `json:"pic_url"           dc:"pic_url"     summary:"商品管理"        ` // 图片
	Name             string `json:"name"  v:"required#商品名称必传"            dc:"name"`            // 商品名称
	Price            int    `json:"price" v:"required#商品价格必传"          dc:"price"  `           // 价格 单位分
	Level1CategoryId int    `json:"level1CategoryId" dc:"level1_category_id" `                 // 1级分类id
	Level2CategoryId int    `json:"level2CategoryId" dc:"level2_category_id" `                 // 2级分类id
	Level3CategoryId int    `json:"level3CategoryId" dc:"level3_category_id" `                 // 3级分类id
	Brand            string `json:"brand" v:"max-length:30#品牌名称最大30字"   dc:"brand" `           // 品牌
	Stock            int    `json:"stock"            dc:"stock"              `                 // 库存
	Sale             int    `json:"sale"             dc:"sale"               `                 // 销量
	Tags             string `json:"tags"             dc:"tags"               `                 // 标签
	DetailInfo       string `json:"detailInfo"       dc:"detail_info"        `                 // 商品详情
}

type GoodsReq struct {
	g.Meta `path:"/v1/goods/add" method:"post" tags:"商品" summary:"添加商品"`
	GoodsCommon
}
type GoodsRes struct {
	Id int `json:"id"`
}

type GoodsDeleteReq struct {
	g.Meta `path:"/v1/goods/delete" method:"delete" tags:"商品" summary:"删除商品管理接口"`
	Id     uint `v:"min:1#请选择需要删除的商品" dc:"商品id"`
}

// 自定义返回内容
type GoodsDeleteRes struct {
	Id int `json:"id"`
}

type GoodsUpdateReq struct {
	g.Meta `path:"/v1/goods/{Id}" method:"post" tags:"商品" summary:"修改商品管理接口"`
	Id     int `json:"id"    v:"required#内容Id不能为空" dc:"ID"`
	GoodsCommon
}

// 自定义返回内容
type GoodsUpdateRes struct {
	Id int `json:"id"`
}

type GoodsGetListCommonReq struct {
	g.Meta `path:"/v1/goods/list" method:"get" tags:"商品" summary:"商品列表"`
	CommonPaginationReq
}
type GoodsGetListCommonRes struct {
	//todo
	List  interface{}    `json:"list"`  // 列表
	Stats map[string]int `json:"stats"` // 搜索统计
	Page  int            `json:"page"`  // 分页码
	Size  int            `json:"size"`  // 分页数量
	Total int            `json:"total"` // 数据总数
}
