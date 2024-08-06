package controller

import (
	"context"
	"github.com/gogf/gf/v2/util/gconv"
	"shop-v2/api/backend/v1"
	"shop-v2/internal/consts"
	"shop-v2/internal/model"
	"shop-v2/internal/service"
)

// Admin 内容管理
var Admin = cAdmin{}

type cAdmin struct{}

func (a *cAdmin) Create(ctx context.Context, req *v1.AdminCreateReq) (res *v1.AdminCreateRes, err error) {
	out, err := service.Admin().Create(ctx, model.AdminCreateInput{
		AdminCreateUpdateBase: model.AdminCreateUpdateBase{
			Name:     req.Name,
			Password: req.Password,
			RoleIds:  req.RoleIds,
			IsAdmin:  req.IsAdmin,
		},
	})
	if err != nil {
		return nil, err
	}
	return &v1.AdminCreateRes{AdminId: out.AdminId}, nil
}

func (a *cAdmin) Delete(ctx context.Context, req *v1.AdminDeleteReq) (res *v1.AdminDeleteRes, err error) {
	err = service.Admin().Delete(ctx, req.Id)
	return
}

func (a *cAdmin) Update(ctx context.Context, req *v1.AdminUpdateReq) (res *v1.AdminUpdateRes, err error) {
	err = service.Admin().Update(ctx, model.AdminUpdateInput{
		Id: req.Id,
		AdminCreateUpdateBase: model.AdminCreateUpdateBase{
			Name:     req.Name,
			Password: req.Password,
			RoleIds:  req.RoleIds,
			IsAdmin:  req.IsAdmin,
		},
	})
	return &v1.AdminUpdateRes{Id: req.Id}, nil
}

// Index article list
func (a *cAdmin) List(ctx context.Context, req *v1.AdminGetListCommonReq) (res *v1.AdminGetListCommonRes, err error) {
	getListRes, err := service.Admin().GetList(ctx, model.AdminGetListInput{
		Page: req.Page,
		Size: req.Size,
	})
	if err != nil {
		return nil, err
	}
	return &v1.AdminGetListCommonRes{
		List:  getListRes.List,
		Page:  getListRes.Page,
		Size:  getListRes.Size,
		Total: getListRes.Total,
	}, nil
}

//for jwt 获取用户信息
//func (c *cAdmin) Info(ctx context.Context, req *v1.AdminGetInfoReq) (res *v1.AdminGetInfoRes, err error) {
//	return &v1.AdminGetInfoRes{
//		Id:          gconv.Int(service.Auth().GetIdentity(ctx)),
//		IdentityKey: service.Auth().IdentityKey,
//		Payload:     service.Auth().GetPayload(ctx),
//	}, nil
//}

// for gtoken
func (c *cAdmin) Info(ctx context.Context, req *v1.AdminGetInfoReq) (res *v1.AdminGetInfoRes, err error) {
	return &v1.AdminGetInfoRes{
		Id:      gconv.Int(ctx.Value(consts.CtxAdminId)),
		Name:    gconv.String(ctx.Value(consts.CtxAdminName)),
		RoleIds: gconv.String(ctx.Value(consts.CtxAdminRoleIds)),
		IsAdmin: gconv.Int(ctx.Value(consts.CtxAdminId)),
	}, nil
}
