package model

import "github.com/gogf/gf/v2/os/gtime"

// CouponCreateUpdateBase 创建/修改内容基类
type CouponCreateUpdateBase struct {
	Name       string
	Price      int
	GoodsId    string
	CategoryId int
}

// CouponCreateInput 创建内容
type CouponCreateInput struct {
	CouponCreateUpdateBase
	UserId int
}

// CouponCreateOutput 创建内容返回结果
type CouponCreateOutput struct {
	CouponId int `json:"position_id"`
}

// CouponUpdateInput 修改内容
type CouponUpdateInput struct {
	CouponCreateUpdateBase
	Id int
}

// CouponGetListInput 获取内容列表
type CouponGetListInput struct {
	Page      int // 分页号码
	Size      int // 分页数量，最大50
	Sort      int // 排序类型(0:最新, 默认。1:活跃, 2:热度)
	GoodsName string
	GoodsId   int
}

// CouponGetListOutput 查询列表结果
type CouponGetListOutput struct {
	List  []CouponGetListOutputItem `json:"list" description:"列表"`
	Page  int                       `json:"page" description:"分页码"`
	Size  int                       `json:"size" description:"分页数量"`
	Total int                       `json:"total" description:"数据总数"`
}

// CouponSearchInput 搜索列表
type CouponSearchInput struct {
	Key      string // 关键字
	Type     string // 内容模型
	CouponId uint   // 栏目ID
	Page     int    // 分页号码
	Size     int    // 分页数量，最大50
	Sort     int    // 排序类型(0:最新, 默认。1:活跃, 2:热度)
}

// CouponSearchOutput 搜索列表结果
type CouponSearchOutput struct {
	List  []CouponSearchOutputItem `json:"list"`  // 列表
	Stats map[string]int           `json:"stats"` // 搜索统计
	Page  int                      `json:"page"`  // 分页码
	Size  int                      `json:"size"`  // 分页数量
	Total int                      `json:"total"` // 数据总数
}

// ContentListItem 主要用于列表展示
//type CouponListItem struct {
//	Id        uint        `json:"id"` // 自增ID
//	PicUrl    string      `json:"pic_Url"`
//	Link      string      `json:"link"`
//	Sort      int         `json:"sort"` // 排序，数值
//	GoodsName string      `json:"goods_name"`
//	GoodsId   int         `json:"goods_id"`
//	CreatedAt *gtime.Time `json:"created_at"` // 创建时间
//	UpdatedAt *gtime.Time `json:"updated_at"` // 修改时间
//}

type CouponGetListOutputItem struct {
	//position *CouponListItem `json:"position"`
	//Coupon *CouponListCouponItem `json:"category"`
	//User     *CouponListUserItem     `json:"user"`
	Id         int         `json:"id"`
	Name       string      `json:"name"`
	Price      int         `json:"sort"` // 排序，数值
	GoodsId    int         `json:"goods_id"`
	CategoryId int         `json:"category_id"`
	CreatedAt  *gtime.Time `json:"created_at"` // 创建时间
	UpdatedAt  *gtime.Time `json:"updated_at"` // 修改时间
}

type CouponSearchOutputItem struct {
	CouponGetListOutputItem
}