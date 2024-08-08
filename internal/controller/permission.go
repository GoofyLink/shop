package controller

import (
	"context"
	v1 "shop-v2/api/backend/v1"
	"shop-v2/internal/model"
	"shop-v2/internal/service"
)

// Permission 内容管理
var Permission = cPermission{}

type cPermission struct{}

func (a *cPermission) Create(ctx context.Context, req *v1.PermissionReq) (res *v1.PermissionRes, err error) {
	// 1.调service服务
	out, err := service.Permission().Create(ctx, model.PermissionCreateInput{
		PermissionCreateUpdateBase: model.PermissionCreateUpdateBase{
			Name: req.Name,
			Path: req.Path,
		},
	})
	if err != nil {
		return nil, err
	}
	// 2.返回数据
	return &v1.PermissionRes{PermissionId: out.PermissionId}, nil
}

func (a *cPermission) Delete(ctx context.Context, req *v1.PermissionDeleteReq) (res *v1.PermissionDeleteRes, err error) {
	err = service.Permission().Delete(ctx, req.Id)
	return
}

func (a *cPermission) Update(ctx context.Context, req *v1.PermissionUpdateReq) (res *v1.PermissionUpdateRes, err error) {
	err = service.Permission().Update(ctx, model.PermissionUpdateInput{
		Id: req.Id,
		PermissionCreateUpdateBase: model.PermissionCreateUpdateBase{
			Name: req.Name,
			Path: req.Path,
		},
	})
	return &v1.PermissionUpdateRes{PermissionId: req.Id}, nil
}

// Index article list
func (a *cPermission) List(ctx context.Context, req *v1.PermissionGetListCommonReq) (res *v1.PermissionGetListCommonRes, err error) {
	getListRes, err := service.Permission().GetList(ctx, model.PermissionGetListInput{
		Page: req.Page,
		Size: req.Size,
	})
	if err != nil {
		return nil, err
	}
	return &v1.PermissionGetListCommonRes{
		List:  getListRes.List,
		Page:  getListRes.Page,
		Size:  getListRes.Size,
		Total: getListRes.Total,
	}, nil
}
