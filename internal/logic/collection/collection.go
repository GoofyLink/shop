package collection

import (
	"context"
	"github.com/gogf/gf/v2/util/gconv"
	"shop-v2/internal/consts"
	"shop-v2/internal/dao"
	"shop-v2/internal/model"
	"shop-v2/internal/service"
)

type sCollection struct{}

func init() {
	service.RegisterCollection(New())
}

func New() *sCollection {
	return &sCollection{}
}

func (s *sCollection) AddCollection(ctx context.Context, in model.AddCollectionInput) (out model.AddCollectionOutPut, err error) {
	// 获取到userId
	in.UserId = gconv.Uint(ctx.Value(consts.CtxUserId))
	id, err := dao.CollectionInfo.Ctx(ctx).InsertAndGetId(in)
	if err != nil {
		return model.AddCollectionOutPut{}, err
	}
	return model.AddCollectionOutPut{Id: gconv.Uint(id)}, nil
}

// 兼容处理 优先根据收藏ID删除收藏ID为0 再根据对象ID和Type删除
func (s *sCollection) DeleteCollection(ctx context.Context, in model.DeleteCollectionInput) (out model.DeleteCollectionOutPut, err error) {
	//优先根据收藏ID删除
	if in.Id != 0 {
		_, err = dao.CategoryInfo.Ctx(ctx).WherePri(in.Id).Delete()
		if err != nil {
			return model.DeleteCollectionOutPut{}, err
		}
		return model.DeleteCollectionOutPut{Id: in.Id}, nil
		//为0 再根据对象ID和Type删除
	} else {
		// 获取到userId
		in.UserId = gconv.Uint(ctx.Value(consts.CtxUserId))
		id, err := dao.CollectionInfo.Ctx(ctx).OmitEmpty().Where(in).Delete() // OmitEmpty需要过滤空值
		if err != nil {
			return out, err
		}
		return model.DeleteCollectionOutPut{Id: gconv.Uint(id)}, nil
	}
}

func (s *sCollection) CollectionList(ctx context.Context, in model.CollectionListInput) (out *model.CollectionListOutput, err error) {
	var (
		m = dao.CollectionInfo.Ctx(ctx)
	)
	out = &model.CollectionListOutput{
		Page: 0,
		Size: 0,
		List: []model.CollectionListOutputItem{}, // 数据为空时 返回空数组 而不是null
	}
	// 1.分配查询
	listModel := m.Page(in.Page, in.Size)
	// 2.条件查询 只查询type相关的
	if in.Type != 0 {
		listModel = listModel.Where(dao.CollectionInfo.Columns().Type, in.Type)
	}
	// 优化
	if out.Total, err = listModel.Count(); err != nil {
		return out, err
	}

	if out.Total == 0 {
		return out, nil
	}
	// 优化
	if in.Type == consts.CollectionTypeGoods {
		if err = listModel.With(model.GoodsItem{}).Scan(&out.List); err != nil {
			return out, err
		}
	} else if in.Type == consts.CollectionTypeArticle {
		if err = listModel.With(model.ArticleItem{}).Scan(&out.List); err != nil {
			return out, err
		}
	} else {
		if err := listModel.WithAll().Scan(&out.List); err != nil {
			return out, err
		}
	}
	return
}
