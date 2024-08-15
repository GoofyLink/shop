package controller

import (
	"context"
	"github.com/gogf/gf/v2/util/gconv"
	"shop-v2/api/backend/v1"
	"shop-v2/internal/consts"
	"shop-v2/internal/model"
	"shop-v2/internal/service"
)

// Article 内容管理
var Article = cArticle{}

type cArticle struct{}

func (a *cArticle) Create(ctx context.Context, req *v1.ArticleReq) (res *v1.ArticleRes, err error) {
	data := model.ArticleCreateInput{}
	// 用scan扫描
	err = gconv.Scan(req, &data)
	if err != nil {
		return nil, err
	}
	data.UserId = gconv.Int(ctx.Value(consts.CtxAdminId))
	out, err := service.Article().Create(ctx, data)
	if err != nil {
		return nil, err
	}
	return &v1.ArticleRes{Id: out.ArticleId}, nil
}

func (a *cArticle) Delete(ctx context.Context, req *v1.ArticleDeleteReq) (res *v1.ArticleDeleteRes, err error) {
	err = service.Article().Delete(ctx, req.Id)
	return
}

func (a *cArticle) Update(ctx context.Context, req *v1.ArticleUpdateReq) (res *v1.ArticleUpdateRes, err error) {
	data := model.ArticleUpdateInput{}
	err = gconv.Struct(req, &data)
	if err != nil {
		return nil, err
	}
	// 获取当前登录用户
	data.UserId = gconv.Int(ctx.Value(consts.CtxAdminId))
	err = service.Article().Update(ctx, data)
	return &v1.ArticleUpdateRes{Id: req.Id}, nil
}

// Index article list
func (a *cArticle) List(ctx context.Context, req *v1.ArticleGetListCommonReq) (res *v1.ArticleGetListCommonRes, err error) {
	getListRes, err := service.Article().GetList(ctx, model.ArticleGetListInput{
		Page: req.Page,
		Size: req.Size,
	})
	if err != nil {
		return nil, err
	}
	return &v1.ArticleGetListCommonRes{
		List:  getListRes.List,
		Page:  getListRes.Page,
		Size:  getListRes.Size,
		Total: getListRes.Total,
	}, nil
}
