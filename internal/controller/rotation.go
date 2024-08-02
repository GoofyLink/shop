package controller

import (
	"context"

	"shop-v2/api/backend/v1"
	"shop-v2/internal/model"
	"shop-v2/internal/service"
)

// Rotation 内容管理
var Rotation = cRotation{}

type cRotation struct{}

func (a *cRotation) Create(ctx context.Context, req *v1.RotationCreateReq) (res *v1.RotationCreateRes, err error) {
	out, err := service.Rotation().Create(ctx, model.RotationCreateInput{
		RotationCreateUpdateBase: model.RotationCreateUpdateBase{
			Id:     req.Id,
			PicUrl: req.Pic_url,
			Link:   req.Link,
			Sort:   req.Sort,
		},
	})
	if err != nil {
		return nil, err
	}
	return &v1.RotationCreateRes{RotationId: out.RotationId}, nil
}

func (a *cRotation) Delete(ctx context.Context, req *v1.RotationDeleteReq) (res *v1.RotationDeleteRes, err error) {
	err = service.Rotation().Delete(ctx, req.Id)
	return
}
