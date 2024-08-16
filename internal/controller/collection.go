package controller

import (
	"context"
	"github.com/gogf/gf/v2/util/gconv"
	v1 "shop-v2/api/frontend/v1"
	"shop-v2/internal/model"
	"shop-v2/internal/service"
)

var Collection = cCollection{}

type cCollection struct{}

func (c *cCollection) AddCollection(ctx context.Context, req *v1.AddCollectionReq) (res *v1.AddCollectionRes, err error) {
	data := model.AddCollectionInput{}
	err = gconv.Struct(ctx, &data)
	if err != nil {
		return nil, err
	}
	out, err := service.Collection().AddCollection(ctx, data)
	if err != nil {
		return nil, err
	}

	return &v1.AddCollectionRes{Id: out.Id}, nil
}

func (c *cCollection) DeleteCollection(ctx context.Context, req *v1.DeleteCollectionReq) (res *v1.DeleteCollectionRes, err error) {
	data := model.DeleteCollectionInput{}
	err = gconv.Struct(ctx, &data)
	if err != nil {
		return nil, err
	}
	collection, err := service.Collection().DeleteCollection(ctx, data)
	if err != nil {
		return nil, err
	}
	return &v1.DeleteCollectionRes{Id: collection.Id}, nil
}

func (c *cCollection) CollectionList(ctx context.Context, req *v1.CollectionListReq) (res *v1.CollectionListRes, err error) {
	getListRes, err := service.Collection().CollectionList(ctx, model.CollectionListInput{
		Page: req.Page,
		Size: req.Size,
		Type: req.Type,
	})
	if err != nil {
		return nil, err
	}
	return &v1.CollectionListRes{
		List:  getListRes.List,
		Page:  getListRes.Page,
		Size:  getListRes.Size,
		Total: getListRes.Total,
	}, nil
}
