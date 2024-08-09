package controller

import (
	"context"
	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
	v1 "shop-v2/api/backend/v1"
	"shop-v2/internal/model"
	"shop-v2/internal/service"
)

var File = cFile{}

type cFile struct {
}

func (c *cFile) Upload(ctx context.Context, req *v1.FileUploadReq) (res *v1.FileUploadRes, err error) {
	if req.File == nil {
		return nil, gerror.NewCode(gcode.CodeMissingParameter, "请上传文件") // 自定义错误码
	}
	pload, err := service.File().UploadFile(ctx, model.FileUpdateInput{
		File:       req.File,
		RandomName: true,
	})
	if err != nil {
		return nil, err
	}
	return &v1.FileUploadRes{
		Name: pload.Name,
		Url:  pload.Url,
	}, nil
}
