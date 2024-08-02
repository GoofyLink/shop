package v1

import (
	"github.com/gogf/gf/v2/frame/g"
)

type RotationCreateReq struct {
	g.Meta  `path:"/backend/v1/rotation/add" method:"post" tags:"内容" summary:"图片接口"`
	Id      int    `json:"id"    v:"required#内容Id不能为空" dc:"ID"`
	Pic_url string `json:"pic_url"    v:"required#图片"      dc:"图片地址"`
	Link    string `json:"link"   v:"required#请输入标题"      dc:"图片跳转连接"`
	Sort    int    `json:"sort" dc:"排序"`
}

type RotationCreateRes struct {
	RotationId int `json:"RotationId"`
}

type RotationDeleteReq struct {
	g.Meta `path:"/backend/v1/rotation/delete" method:"delete" tags:"轮播图" summary:"删除轮播图接口"`
	Id     uint `v:"min:1#请选择需要删除的内容" dc:"内容id"`
}
type RotationDeleteRes struct{}
