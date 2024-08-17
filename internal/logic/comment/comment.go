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

type sComment struct{}

func init() {
	service.RegisterComment(New())
}

func New() *sComment {
	return &sComment{}
}

// 添加
func (s *sComment) AddComment(ctx context.Context, in model.AddCommentInput) (out model.AddCommentOutPut, err error) {
	// 获取到userId
	in.UserId = gconv.Uint(ctx.Value(consts.CtxUserId))
	id, err := dao.CommentInfo.Ctx(ctx).InsertAndGetId(in)
	if err != nil {
		return model.AddCommentOutPut{}, err
	}
	return model.AddCommentOutPut{Id: gconv.Uint(id)}, nil
}

// 兼容处理 优先根据收藏ID删除收藏ID为0 再根据对象ID和Type删除
func (s *sComment) DeleteComment(ctx context.Context, in model.DeleteCommentInput) (out model.DeleteCommentOutPut, err error) {
	//优先根据收藏ID删除
	//只允许删除自己的内容
	condition := g.Map{
		dao.CommentInfo.Columns().Id:     in.Id,
		dao.CommentInfo.Columns().UserId: ctx.Value(consts.CtxUserId), //必须是当前用户id
	}
	_, err = dao.CategoryInfo.Ctx(ctx).Where(condition).Delete()
	if err != nil {
		return model.DeleteCommentOutPut{}, err
	}
	return model.DeleteCommentOutPut{Id: gconv.Uint(in.Id)}, nil
}

// 获取collectionList列表
func (s *sComment) CommentList(ctx context.Context, in model.CommentListInput) (out *model.CommentListOutput, err error) {
	var (
		m = dao.CommentInfo.Ctx(ctx)
	)
	out = &model.CommentListOutput{
		Page: 0,
		Size: 0,
		List: []model.CommentListOutputItem{}, // 数据为空时 返回空数组 而不是null
	}
	// 1.分配查询
	listModel := m.Page(in.Page, in.Size)
	// 2.条件查询 只查询type相关的
	if in.Type != 0 {
		listModel = listModel.Where(dao.CommentInfo.Columns().Type, in.Type)
	}
	// 优化
	if out.Total, err = listModel.Count(); err != nil {
		return out, err
	}

	if out.Total == 0 {
		return out, nil
	}
	// 优化
	if in.Type == consts.CommentTypeGoods {
		if err = listModel.With(model.GoodsItem{}).Scan(&out.List); err != nil {
			return out, err
		}
	} else if in.Type == consts.CommentTypeArticle {
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
func (s *sComment) CommentCount(ctx context.Context, objectId uint, collectionType uint8) (count int, err error) {
	condition := g.Map{
		dao.CommentInfo.Columns().ObjectId: objectId,
		dao.CommentInfo.Columns().Type:     collectionType,
	}

	count, err = dao.CommentInfo.Ctx(ctx).Where(condition).Count()
	if err != nil {
		return 0, err
	}
	return
}

// 抽取方法判断当前用户是否收藏 文章详情/商品详情
func (s *sComment) CommentCheck(ctx context.Context, in model.CheckIsCommentInput) (bool, error) {
	condition := g.Map{
		dao.CommentInfo.Columns().UserId:   ctx.Value(consts.CtxUserId),
		dao.CommentInfo.Columns().ObjectId: in.ObjectId,
		dao.CommentInfo.Columns().Type:     in.Type,
	}
	count, err := dao.CommentInfo.Ctx(ctx).Where(condition).Count()
	if err != nil {
		return false, err
	}

	if count > 0 {
		return true, nil
	} else {
		return false, nil
	}
}
