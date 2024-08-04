package controller

import (
	"context"
	v1 "shop-v2/api/backend/v1"
	"shop-v2/internal/service"
)

// 登录管理
var Login = cLogin{}

type cLogin struct{}

//func (a *cLogin) Login(ctx context.Context, req *v1.LoginDoReq) (res *v1.LoginDoRes, err error) {
//	res = &v1.LoginDoRes{}
//	err = service.Login().Login(ctx, model.LoginInput{
//		Name:     req.Name,
//		Password: req.Password,
//	})
//	if err != nil {
//		return
//	}
//	// 获取当前session信息返回
//	//res.Info = service.Session().GetUser(ctx)
//	return
//}

func (c *cLogin) Login(ctx context.Context, req *v1.LoginDoReq) (res *v1.LoginDoRes, err error) {
	res = &v1.LoginDoRes{}
	res.Token, res.Expire = service.Auth().LoginHandler(ctx)
	return
}

func (c *cLogin) RefreshToken(ctx context.Context, req *v1.RefreshTokenReq) (res *v1.RefreshTokenRes, err error) {
	res = &v1.RefreshTokenRes{}
	res.Token, res.Expire = service.Auth().RefreshHandler(ctx)
	return
}

func (c *cLogin) Logout(ctx context.Context, req *v1.LogoutReq) (res *v1.LogoutRes, err error) {
	service.Auth().LogoutHandler(ctx)
	return
}
