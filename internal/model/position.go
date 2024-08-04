package model

import "github.com/gogf/gf/v2/os/gtime"

// PositionCreateUpdateBase 创建/修改内容基类
type PositionCreateUpdateBase struct {
	Id        int
	PicUrl    string
	Link      string
	Sort      int
	GoodsName string
	GoodsId   int
}

// PositionCreateInput 创建内容
type PositionCreateInput struct {
	PositionCreateUpdateBase
	UserId int
}

// PositionCreateOutput 创建内容返回结果
type PositionCreateOutput struct {
	PositionId int `json:"position_id"`
}

// PositionUpdateInput 修改内容
type PositionUpdateInput struct {
	PositionCreateUpdateBase
	Id int
}

// PositionGetListInput 获取内容列表
type PositionGetListInput struct {
	Page      int // 分页号码
	Size      int // 分页数量，最大50
	Sort      int // 排序类型(0:最新, 默认。1:活跃, 2:热度)
	GoodsName string
	GoodsId   int
}

// PositionGetListOutput 查询列表结果
type PositionGetListOutput struct {
	List  []PositionGetListOutputItem `json:"list" description:"列表"`
	Page  int                         `json:"page" description:"分页码"`
	Size  int                         `json:"size" description:"分页数量"`
	Total int                         `json:"total" description:"数据总数"`
}

// PositionSearchInput 搜索列表
type PositionSearchInput struct {
	Key        string // 关键字
	Type       string // 内容模型
	CategoryId uint   // 栏目ID
	Page       int    // 分页号码
	Size       int    // 分页数量，最大50
	Sort       int    // 排序类型(0:最新, 默认。1:活跃, 2:热度)
}

// PositionSearchOutput 搜索列表结果
type PositionSearchOutput struct {
	List  []PositionSearchOutputItem `json:"list"`  // 列表
	Stats map[string]int             `json:"stats"` // 搜索统计
	Page  int                        `json:"page"`  // 分页码
	Size  int                        `json:"size"`  // 分页数量
	Total int                        `json:"total"` // 数据总数
}

// ContentListItem 主要用于列表展示
//type PositionListItem struct {
//	Id        uint        `json:"id"` // 自增ID
//	PicUrl    string      `json:"pic_Url"`
//	Link      string      `json:"link"`
//	Sort      int         `json:"sort"` // 排序，数值
//	GoodsName string      `json:"goods_name"`
//	GoodsId   int         `json:"goods_id"`
//	CreatedAt *gtime.Time `json:"created_at"` // 创建时间
//	UpdatedAt *gtime.Time `json:"updated_at"` // 修改时间
//}

type PositionGetListOutputItem struct {
	//position *PositionListItem `json:"position"`
	//Category *PositionListCategoryItem `json:"category"`
	//User     *PositionListUserItem     `json:"user"`
	Id        uint        `json:"id"` // 自增ID
	PicUrl    string      `json:"pic_Url"`
	Link      string      `json:"link"`
	Sort      int         `json:"sort"` // 排序，数值
	GoodsName string      `json:"goods_name"`
	GoodsId   int         `json:"goods_id"`
	CreatedAt *gtime.Time `json:"created_at"` // 创建时间
	UpdatedAt *gtime.Time `json:"updated_at"` // 修改时间
}

type PositionSearchOutputItem struct {
	PositionGetListOutputItem
}
