package model

import (
	"github.com/gogf/gf/v2/net/ghttp"
)

type FileUpdateInput struct {
	File       *ghttp.UploadFile
	Name       string
	RandomName bool
}

type FileUploadOutPut struct {
	Id   uint
	Name string
	Src  string
	Url  string
}
