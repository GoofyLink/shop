package controller

import (
	"context"
	v1 "shop-v2/api/backend/v1"
	"shop-v2/internal/model"
	"shop-v2/internal/service"
)

// 登录管理
var Login = cLogin{}

type cLogin struct{}

func (a *cLogin) Login(ctx context.Context, req *v1.LoginDoReq) (res *v1.LoginDoRes, err error) {
	res = &v1.LoginDoRes{}
	err = service.Login().Login(ctx, model.LoginInput{
		Name:     req.Name,
		Password: req.Password,
	})
	if err != nil {
		return
	}
	// 获取当前session信息返回
	res.Info = service.Session().GetUser(ctx)
	return
}
