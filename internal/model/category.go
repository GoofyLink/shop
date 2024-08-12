package model

import "github.com/gogf/gf/v2/os/gtime"

// CategoryCreateUpdateBase 创建/修改内容基类
type CategoryCreateUpdateBase struct {
	ParentId uint
	PicUrl   string
	Name     string
	Sort     uint8
	Level    uint8
}

// CategoryCreateInput 创建内容
type CategoryCreateInput struct {
	CategoryCreateUpdateBase
	UserId int
}

// CategoryCreateOutput 创建内容返回结果
type CategoryCreateOutput struct {
	CategoryId int `json:"position_id"`
}

// CategoryUpdateInput 修改内容
type CategoryUpdateInput struct {
	CategoryCreateUpdateBase
	Id int
}

// CategoryGetListInput 获取内容列表
type CategoryGetListInput struct {
	Page      int // 分页号码
	Size      int // 分页数量，最大50
	Sort      int // 排序类型(0:最新, 默认。1:活跃, 2:热度)
	GoodsName string
	GoodsId   int
}

// CategoryGetListOutput 查询列表结果
type CategoryGetListOutput struct {
	List  []CategoryGetListOutputItem `json:"list" description:"列表"`
	Page  int                         `json:"page" description:"分页码"`
	Size  int                         `json:"size" description:"分页数量"`
	Total int                         `json:"total" description:"数据总数"`
}

// CategorySearchInput 搜索列表
type CategorySearchInput struct {
	Key        string // 关键字
	Type       string // 内容模型
	CategoryId uint   // 栏目ID
	Page       int    // 分页号码
	Size       int    // 分页数量，最大50
	Sort       int    // 排序类型(0:最新, 默认。1:活跃, 2:热度)
}

// CategorySearchOutput 搜索列表结果
type CategorySearchOutput struct {
	List  []CategorySearchOutputItem `json:"list"`  // 列表
	Stats map[string]int             `json:"stats"` // 搜索统计
	Page  int                        `json:"page"`  // 分页码
	Size  int                        `json:"size"`  // 分页数量
	Total int                        `json:"total"` // 数据总数
}

// ContentListItem 主要用于列表展示
//type CategoryListItem struct {
//	Id        uint        `json:"id"` // 自增ID
//	PicUrl    string      `json:"pic_Url"`
//	Link      string      `json:"link"`
//	Sort      int         `json:"sort"` // 排序，数值
//	GoodsName string      `json:"goods_name"`
//	GoodsId   int         `json:"goods_id"`
//	CreatedAt *gtime.Time `json:"created_at"` // 创建时间
//	UpdatedAt *gtime.Time `json:"updated_at"` // 修改时间
//}

type CategoryGetListOutputItem struct {
	//position *CategoryListItem `json:"position"`
	//Category *CategoryListCategoryItem `json:"category"`
	//User     *CategoryListUserItem     `json:"user"`
	ParentId  uint8       `json:"parent_id"` // 自增ID
	PicUrl    string      `json:"pic_url"`
	Name      string      `json:"name"`
	Sort      uint8       `json:"sort"` // 排序，数值
	Level     uint8       `json:"level"`
	CreatedAt *gtime.Time `json:"created_at"` // 创建时间
	UpdatedAt *gtime.Time `json:"updated_at"` // 修改时间
}

type CategorySearchOutputItem struct {
	CategoryGetListOutputItem
}
