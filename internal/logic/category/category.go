package position

import (
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"shop-v2/internal/model/entity"

	"shop-v2/internal/dao"
	"shop-v2/internal/model"
	"shop-v2/internal/service"
)

type sCategory struct{}

func init() {
	service.RegisterCategory(New())
}

func New() *sCategory {
	return &sCategory{}
}

// Create 创建分类
func (s *sCategory) Create(ctx context.Context, in model.CategoryCreateInput) (out model.CategoryCreateOutput, err error) {
	lastInsertID, err := dao.CategoryInfo.Ctx(ctx).Data(in).InsertAndGetId()
	if err != nil {
		return out, err
	}
	return model.CategoryCreateOutput{CategoryId: int(lastInsertID)}, err
}

// Delete 删除 这样删除是软删除还会留着原来的信息
func (s *sCategory) Delete(ctx context.Context, id uint) (err error) {
	// 删除分类
	_, err = dao.CategoryInfo.Ctx(ctx).Where(g.Map{
		dao.CategoryInfo.Columns().Id: id,
	}).Delete()
	//}).Unscoped().Delete()  unscoped直接删除信息 不会软删除
	if err != nil {
		return err
	}
	return
}

// Update 修改
func (s *sCategory) Update(ctx context.Context, in model.CategoryUpdateInput) (err error) {
	//FieldsEx(dao.CategoryInfo.Columns().Id). 排除id的其它字段进行数据更新
	_, err = dao.CategoryInfo.
		Ctx(ctx).
		Data(in).
		FieldsEx(dao.CategoryInfo.Columns().Id).
		Where(dao.CategoryInfo.Columns().Id, in.Id).
		Update()
	return
}

// GetList 查询分类列表
func (s *sCategory) GetList(ctx context.Context, in model.CategoryGetListInput) (out *model.CategoryGetListOutput, err error) {
	var (
		m = dao.CategoryInfo.Ctx(ctx)
	)
	out = &model.CategoryGetListOutput{
		Page: in.Page,
		Size: in.Size,
	}

	// 分页查询
	listModel := m.Page(in.Page, in.Size)
	listModel = listModel.OrderDesc(dao.CategoryInfo.Columns().Sort)

	// 执行查询 倒序
	var list []*entity.CategoryInfo
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
