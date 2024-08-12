package v1

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
)

type UploadImgToCloudReq struct {
	g.Meta `path:"/v1/upload" method:"post"  mime:"multipart/form-data" tags:"上传" summary:"文件上传到云"`
	File   *ghttp.UploadFile `json:"file" type:"file" dc:"选择上传文件"`
}

type UploadImgToCloudRes struct {
	Url string `json:"url" dc:"文件地址"`
}
