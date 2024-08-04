package rotation

import (
	"context"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
	"shop-v2/internal/model/entity"

	"github.com/gogf/gf/v2/encoding/ghtml"
	"shop-v2/internal/dao"
	"shop-v2/internal/model"
	"shop-v2/internal/service"
)

type sRotation struct{}

func init() {
	service.RegisterRotation(New())
}

func New() *sRotation {
	return &sRotation{}
}

// Create 创建内容
func (s *sRotation) Create(ctx context.Context, in model.RotationCreateInput) (out model.RotationCreateOutput, err error) {

	// 不允许HTML代码
	if err = ghtml.SpecialCharsMapOrStruct(in); err != nil {
		return out, err
	}
	lastInsertID, err := dao.RotationInfo.Ctx(ctx).Data(in).InsertAndGetId()
	if err != nil {
		return out, err
	}
	return model.RotationCreateOutput{RotationId: int(lastInsertID)}, err
}

// Delete 删除 这样删除是软删除还会留着原来的信息
func (s *sRotation) Delete(ctx context.Context, id uint) error {
	return dao.RotationInfo.Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
		// 删除内容
		_, err := dao.RotationInfo.Ctx(ctx).Where(g.Map{
			dao.RotationInfo.Columns().Id: id,
		}).Delete()
		//}).Unscoped().Delete()  unscoped直接删除信息 不会软删除
		return err
	})
}

// Update 修改
func (s *sRotation) Update(ctx context.Context, in model.RotationUpdateInput) error {
	return dao.RotationInfo.Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
		// 不允许HTML代码
		if err := ghtml.SpecialCharsMapOrStruct(in); err != nil {
			return err
		}
		//FieldsEx(dao.RotationInfo.Columns().Id). 排除id的其它字段进行数据更新
		_, err := dao.RotationInfo.
			Ctx(ctx).
			Data(in).
			FieldsEx(dao.RotationInfo.Columns().Id).
			Where(dao.RotationInfo.Columns().Id, in.Id).
			Update()
		return err
	})
}

// GetList 查询内容列表
func (s *sRotation) GetList(ctx context.Context, in model.RotationGetListInput) (out *model.RotationGetListOutput, err error) {
	var (
		m = dao.RotationInfo.Ctx(ctx)
	)
	out = &model.RotationGetListOutput{
		Page: in.Page,
		Size: in.Size,
	}

	// 分配查询
	listModel := m.Page(in.Page, in.Size)
	listModel = listModel.OrderDesc(dao.RotationInfo.Columns().Sort)

	// 执行查询
	var list []*entity.RotationInfo
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
	// rotation
	//if err := listModel.ScanList(&out.List, "rotation"); err != nil {
	//	return out, err
	//}

	if err := listModel.Scan(&out.List); err != nil {
		return out, err
	}
	return
}
