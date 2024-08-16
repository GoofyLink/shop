package controller

import (
	"context"
	"github.com/gogf/gf/v2/util/gconv"
	v1 "shop-v2/api/frontend/v1"
	"shop-v2/internal/consts"
	"shop-v2/internal/model"
	"shop-v2/internal/service"
)

// User 内容管理
var User = cUser{}

type cUser struct{}

func (c *cUser) Register(ctx context.Context, req *v1.RegisterReq) (res *v1.RegisterRes, err error) {
	data := model.RegisterInput{}
	err = gconv.Struct(req, &data)
	if err != nil {
		return nil, err
	}
	out, err := service.User().Register(ctx, data)
	if err != nil {
		return nil, err
	}
	return &v1.RegisterRes{Id: out.Id}, nil
}

func (c *cUser) Info(ctx context.Context, req *v1.UserInfoReq) (res *v1.UserInfoRes, err error) {
	res = &v1.UserInfoRes{}
	res.Id = gconv.Uint(ctx.Value(consts.CtxUserId))
	res.Name = gconv.String(ctx.Value(consts.CtxUserName))
	res.Sex = gconv.Uint8(ctx.Value(consts.CtxUserSex))
	res.Avatar = gconv.String(ctx.Value(consts.CtxUserAvatar))
	res.Sign = gconv.String(ctx.Value(consts.CtxUserSign))
	res.Status = gconv.Uint8(ctx.Value(consts.CtxUserStatus))
	return res, nil
}
