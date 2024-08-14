package v1

import (
	"github.com/gogf/gf/v2/frame/g"
)

type GoodsOptionsCommon struct {
	GoodsId int    `json:"goods_id" dc:"主商品规格ID" summary:"商品规格"`               // 1级分类id
	PicUrl  string `json:"pic_url"  dc:"pic_url"            `                  // 图片
	Name    string `json:"name"  v:"required#商品规格名称必传"           dc:"name"`    // 商品规格名称
	Price   int    `json:"price" v:"required#商品规格价格必传"           dc:"price"  ` // 价格 单位分
	Brand   string `json:"brand" v:"max-length:30#品牌名称最大30字"   dc:"brand" `    // 品牌
	Stock   int    `json:"stock"    dc:"stock"              `                  // 库存
}

type GoodsOptionsReq struct {
	g.Meta `path:"/v1/goods/options/add" method:"post" tags:"商品规格" summary:"添加商品规格"`
	GoodsOptionsCommon
}
type GoodsOptionsRes struct {
	Id int `json:"id"`
}

type GoodsOptionsDeleteReq struct {
	g.Meta `path:"/v1/goods/options/delete" method:"delete" tags:"商品规格" summary:"删除商品规格分类接口"`
	Id     uint `v:"min:1#请选择需要删除的商品规格" dc:"商品规格id"`
}

// 自定义返回内容
type GoodsOptionsDeleteRes struct {
	Id int `json:"id"`
}

type GoodsOptionsUpdateReq struct {
	g.Meta `path:"/v1/goods/options/{Id}" method:"post" tags:"商品规格" summary:"修改商品规格分类接口"`
	Id     int `json:"id"    v:"required#内容Id不能为空" dc:"ID"`
	GoodsOptionsCommon
}

// 自定义返回内容
type GoodsOptionsUpdateRes struct {
	Id int `json:"id"`
}

type GoodsOptionsGetListCommonReq struct {
	g.Meta `path:"/v1/goods/options/list" method:"get" tags:"商品规格" summary:"商品规格列表"`
	CommonPaginationReq
}
type GoodsOptionsGetListCommonRes struct {
	//todo
	List  interface{}    `json:"list"`  // 列表
	Stats map[string]int `json:"stats"` // 搜索统计
	Page  int            `json:"page"`  // 分页码
	Size  int            `json:"size"`  // 分页数量
	Total int            `json:"total"` // 数据总数
}
