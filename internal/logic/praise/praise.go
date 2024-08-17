package collection

import (
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/gconv"
	"shop-v2/internal/consts"
	"shop-v2/internal/dao"
	"shop-v2/internal/model"
	"shop-v2/internal/service"
)

type sPraise struct{}

func init() {
	service.RegisterPraise(New())
}

func New() *sPraise {
	return &sPraise{}
}

// 添加
func (s *sPraise) AddPraise(ctx context.Context, in model.AddPraiseInput) (out model.AddPraiseOutPut, err error) {
	// 获取到userId
	in.UserId = gconv.Uint(ctx.Value(consts.CtxUserId))
	id, err := dao.PraiseInfo.Ctx(ctx).InsertAndGetId(in)
	if err != nil {
		return model.AddPraiseOutPut{}, err
	}
	return model.AddPraiseOutPut{Id: gconv.Uint(id)}, nil
}

// 兼容处理 优先根据收藏ID删除收藏ID为0 再根据对象ID和Type删除
func (s *sPraise) DeletePraise(ctx context.Context, in model.DeletePraiseInput) (out model.DeletePraiseOutPut, err error) {
	//优先根据收藏ID删除
	if in.Id != 0 {
		_, err = dao.CategoryInfo.Ctx(ctx).WherePri(in.Id).Delete()
		if err != nil {
			return model.DeletePraiseOutPut{}, err
		}
		return model.DeletePraiseOutPut{Id: in.Id}, nil
		//为0 再根据对象ID和Type删除
	} else {
		// 获取到userId
		in.UserId = gconv.Uint(ctx.Value(consts.CtxUserId))
		id, err := dao.PraiseInfo.Ctx(ctx).OmitEmpty().Where(in).Delete() // OmitEmpty需要过滤空值
		if err != nil {
			return out, err
		}
		return model.DeletePraiseOutPut{Id: gconv.Uint(id)}, nil
	}
}

// 获取collectionList列表
func (s *sPraise) PraiseList(ctx context.Context, in model.PraiseListInput) (out *model.PraiseListOutput, err error) {
	var (
		m = dao.PraiseInfo.Ctx(ctx)
	)
	out = &model.PraiseListOutput{
		Page: 0,
		Size: 0,
		List: []model.PraiseListOutputItem{}, // 数据为空时 返回空数组 而不是null
	}
	// 1.分配查询
	listModel := m.Page(in.Page, in.Size)
	// 2.条件查询 只查询type相关的
	if in.Type != 0 {
		listModel = listModel.Where(dao.PraiseInfo.Columns().Type, in.Type)
	}
	// 优化
	if out.Total, err = listModel.Count(); err != nil {
		return out, err
	}

	if out.Total == 0 {
		return out, nil
	}
	// 优化
	if in.Type == consts.PraiseTypeGoods {
		if err = listModel.With(model.GoodsItem{}).Scan(&out.List); err != nil {
			return out, err
		}
	} else if in.Type == consts.PraiseTypeArticle {
		if err = listModel.With(model.ArticleItem{}).Scan(&out.List); err != nil {
			return out, err
		}
	} else {
		if err = listModel.WithAll().Scan(&out.List); err != nil {
			return out, err
		}
	}
	return
}

// 抽取获得收藏数量的方法 for 文章详情/商品详情
func (s *sPraise) PraiseCount(ctx context.Context, objectId uint, collectionType uint8) (count int, err error) {
	condition := g.Map{
		dao.PraiseInfo.Columns().ObjectId: objectId,
		dao.PraiseInfo.Columns().Type:     collectionType,
	}

	count, err = dao.PraiseInfo.Ctx(ctx).Where(condition).Count()
	if err != nil {
		return 0, err
	}
	return
}

// 抽取方法判断当前用户是否收藏 文章详情/商品详情
func (s *sPraise) PraiseCheck(ctx context.Context, in model.CheckIsPraiseInput) (bool, error) {
	condition := g.Map{
		dao.PraiseInfo.Columns().UserId:   ctx.Value(consts.CtxUserId),
		dao.PraiseInfo.Columns().ObjectId: in.ObjectId,
		dao.PraiseInfo.Columns().Type:     in.Type,
	}
	count, err := dao.PraiseInfo.Ctx(ctx).Where(condition).Count()
	if err != nil {
		return false, err
	}

	if count > 0 {
		return true, nil
	} else {
		return false, nil
	}
}
