// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// GoodsInfo is the golang structure for table goods_info.
type GoodsInfo struct {
	Id               int         `json:"id"               orm:"id"                 ` //
	PicUrl           string      `json:"picUrl"           orm:"pic_url"            ` // 图片
	Name             string      `json:"name"             orm:"name"               ` // 商品名称
	Price            int         `json:"price"            orm:"price"              ` // 价格 单位分
	Level1CategoryId int         `json:"level1CategoryId" orm:"level1_category_id" ` // 1级分类id
	Level2CategoryId int         `json:"level2CategoryId" orm:"level2_category_id" ` // 2级分类id
	Level3CategoryId int         `json:"level3CategoryId" orm:"level3_category_id" ` // 3级分类id
	Brand            string      `json:"brand"            orm:"brand"              ` // 品牌
	Stock            int         `json:"stock"            orm:"stock"              ` // 库存
	Sale             int         `json:"sale"             orm:"sale"               ` // 销量
	Tags             string      `json:"tags"             orm:"tags"               ` // 标签
	DetailInfo       string      `json:"detailInfo"       orm:"detail_info"        ` // 商品详情
	CreatedAt        *gtime.Time `json:"createdAt"        orm:"created_at"         ` //
	UpdatedAt        *gtime.Time `json:"updatedAt"        orm:"updated_at"         ` //
	DeletedAt        *gtime.Time `json:"deletedAt"        orm:"deleted_at"         ` //
}
