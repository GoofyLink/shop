package v1

import (
	"github.com/gogf/gf/v2/frame/g"
)

type RotationCreateReq struct {
	g.Meta  `path:"/rotation/backend/v1/rotation" method:"post" tags:"内容" summary:"图片接口"`
	Id      int    `json:"id"    v:"required#内容Id不能为空" dc:"ID"`
	Pic_url string `json:"pic_url"    v:"required#图片"      dc:"图片地址"`
	Link    string `json:"link"   v:"required#请输入标题"      dc:"图片跳转连接"`
	Sort    int    `json:"sort" dc:"排序"`
}

type RotationCreateRes struct {
	RotationId int `json:"RotationId"`
}
