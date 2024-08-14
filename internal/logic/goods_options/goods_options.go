package goods

import (
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"shop-v2/internal/model/entity"

	"shop-v2/internal/dao"
	"shop-v2/internal/model"
	"shop-v2/internal/service"
)

type sGoodsOptions struct{}

func init() {
	service.RegisterGoodsOptions(New())
}

func New() *sGoodsOptions {
	return &sGoodsOptions{}
}

// Create 创建分类
func (s *sGoodsOptions) Create(ctx context.Context, in model.GoodsOptionsCreateInput) (out model.GoodsOptionsCreateOutput, err error) {
	lastInsertID, err := dao.GoodsOptionsInfo.Ctx(ctx).Data(in).InsertAndGetId()
	if err != nil {
		return out, err
	}
	return model.GoodsOptionsCreateOutput{GoodsOptionsId: int(lastInsertID)}, err
}

// Delete 删除 这样删除是软删除还会留着原来的信息
func (s *sGoodsOptions) Delete(ctx context.Context, id uint) (err error) {
	// 删除分类
	_, err = dao.GoodsOptionsInfo.Ctx(ctx).Where(g.Map{
		dao.GoodsOptionsInfo.Columns().Id: id,
	}).Delete()
	//}).Unscoped().Delete()  unscoped直接删除信息 不会软删除
	if err != nil {
		return err
	}
	return
}

// Update 修改
func (s *sGoodsOptions) Update(ctx context.Context, in model.GoodsOptionsUpdateInput) (err error) {
	//FieldsEx(dao.GoodsOptionsInfo.Columns().Id). 排除id的其它字段进行数据更新
	_, err = dao.GoodsOptionsInfo.
		Ctx(ctx).
		Data(in).
		FieldsEx(dao.GoodsOptionsInfo.Columns().Id).
		Where(dao.GoodsOptionsInfo.Columns().Id, in.Id).
		Update()
	return
}

// GetList 查询分类列表
func (s *sGoodsOptions) GetList(ctx context.Context, in model.GoodsOptionsGetListInput) (out *model.GoodsOptionsGetListOutput, err error) {
	var (
		m = dao.GoodsOptionsInfo.Ctx(ctx)
	)
	out = &model.GoodsOptionsGetListOutput{
		Page: in.Page,
		Size: in.Size,
	}

	// 分页查询
	listModel := m.Page(in.Page, in.Size)

	// 执行查询 倒序
	var list []*entity.GoodsOptionsInfo
	if err := listModel.Scan(&list); err != nil {
		return out, err
	}
	// 没有数据
	if len(list) == 0 {
		return out, nil
	}
	out.Total, err = m.Count()
	if err != nil {
		return out, err
	}

	if err := listModel.Scan(&out.List); err != nil {
		return out, err
	}
	return
}
