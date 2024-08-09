package controller

import (
	"context"
	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
	v1 "shop-v2/api/backend/v1"
	"shop-v2/internal/consts"
	"shop-v2/utility/upload"
)

var Upload = cUpload{}

type cUpload struct {
}

func (c *cUpload) UploadImgToCloud(ctx context.Context, req *v1.UploadImgToCloudReq) (res *v1.UploadImgToCloudRes, err error) {
	if req.File == nil {
		return nil, gerror.NewCode(gcode.CodeMissingParameter, consts.CodeMissingParameter)
	}

	cloud, err := upload.UploadImgToCloud(ctx, req.File)
	if err != nil {
		return nil, err
	}
	return &v1.UploadImgToCloudRes{Url: cloud}, nil
}
