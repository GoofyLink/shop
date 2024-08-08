package controller

import (
	"context"
	"github.com/gogf/gf/v2/util/gconv"
	v1 "shop-v2/api/backend/v1"
	"shop-v2/internal/consts"
	"shop-v2/internal/model"
	"shop-v2/internal/service"
)

// Role 内容管理
var Role = cRole{}

type cRole struct{}

func (a *cRole) Create(ctx context.Context, req *v1.RoleReq) (res *v1.RoleRes, err error) {
	// 1.调service服务
	out, err := service.Role().Create(ctx, model.RoleCreateInput{
		RoleCreateUpdateBase: model.RoleCreateUpdateBase{
			Name: req.Name,
			Desc: req.Desc,
		},
	})
	if err != nil {
		return nil, err
	}
	// 2.返回数据
	return &v1.RoleRes{RoleId: out.RoleId}, nil
}

func (a *cRole) Delete(ctx context.Context, req *v1.RoleDeleteReq) (res *v1.RoleDeleteRes, err error) {
	err = service.Role().Delete(ctx, req.Id)
	return
}

func (a *cRole) Update(ctx context.Context, req *v1.RoleUpdateReq) (res *v1.RoleUpdateRes, err error) {
	err = service.Role().Update(ctx, model.RoleUpdateInput{
		Id: req.Id,
		RoleCreateUpdateBase: model.RoleCreateUpdateBase{
			Name: req.Name,
			Desc: req.Desc,
		},
	})
	return &v1.RoleUpdateRes{RoleId: req.Id}, nil
}

// Index article list
func (a *cRole) List(ctx context.Context, req *v1.RoleGetListCommonReq) (res *v1.RoleGetListCommonRes, err error) {
	getListRes, err := service.Role().GetList(ctx, model.RoleGetListInput{
		Page: req.Page,
		Size: req.Size,
	})
	if err != nil {
		return nil, err
	}
	return &v1.RoleGetListCommonRes{
		List:  getListRes.List,
		Page:  getListRes.Page,
		Size:  getListRes.Size,
		Total: getListRes.Total,
	}, nil
}

//for jwt 获取用户信息
//func (c *cRole) Info(ctx context.Context, req *v1.RoleGetInfoReq) (res *v1.RoleGetInfoRes, err error) {
//	return &v1.RoleGetInfoRes{
//		Id:          gconv.Int(service.Auth().GetIdentity(ctx)),
//		IdentityKey: service.Auth().IdentityKey,
//		Payload:     service.Auth().GetPayload(ctx),
//	}, nil
//}

// for gtoken
func (c *cRole) Info(ctx context.Context, req *v1.RoleGetInfoReq) (res *v1.RoleGetInfoRes, err error) {
	return &v1.RoleGetInfoRes{
		Id:   gconv.Int(ctx.Value(consts.CtxRoleId)),
		Name: gconv.String(ctx.Value(consts.CtxRoleName)),
		Desc: gconv.String(ctx.Value(consts.CtxRoleDesc)),
	}, nil
}
