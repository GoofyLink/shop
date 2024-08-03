package Admin

import (
	"context"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/grand"
	"shop-v2/internal/model/entity"
	utility "shop-v2/utility"

	"github.com/gogf/gf/v2/encoding/ghtml"
	"shop-v2/internal/dao"
	"shop-v2/internal/model"
	"shop-v2/internal/service"
)

type sAdmin struct{}

func init() {
	service.RegisterAdmin(New())
}

func New() *sAdmin {
	return &sAdmin{}
}

// Create 创建内容
func (s *sAdmin) Create(ctx context.Context, in model.AdminCreateInput) (out model.AdminCreateOutput, err error) {

	// 不允许HTML代码
	if err = ghtml.SpecialCharsMapOrStruct(in); err != nil {
		return out, err
	}
	// 特殊处理加密 盐 和密码
	UserSalt := grand.S(10)
	in.Password = utility.EncryptPassword(in.Password, UserSalt)
	in.UserSalt = UserSalt
	// 插入数据
	lastInsertID, err := dao.AdminInfo.Ctx(ctx).Data(in).InsertAndGetId()
	if err != nil {
		return out, err
	}
	return model.AdminCreateOutput{AdminId: int(lastInsertID)}, err
}

// Delete 删除 这样删除是软删除还会留着原来的信息
func (s *sAdmin) Delete(ctx context.Context, id uint) error {
	return dao.AdminInfo.Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
		// 删除内容
		_, err := dao.AdminInfo.Ctx(ctx).Where(g.Map{
			dao.AdminInfo.Columns().Id: id,
		}).Delete()
		//}).Unscoped().Delete()  unscoped直接删除信息 不会软删除
		return err
	})
}

// Update 修改
func (s *sAdmin) Update(ctx context.Context, in model.AdminUpdateInput) error {
	return dao.AdminInfo.Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
		// 不允许HTML代码
		if err := ghtml.SpecialCharsMapOrStruct(in); err != nil {
			return err
		}
		//FieldsEx(dao.AdminInfo.Columns().Id). 排除id的其它字段进行数据更新
		//判断是否修改了密码
		if in.Password != "" {
			UserSalt := grand.S(10)
			in.Password = utility.EncryptPassword(in.Password, UserSalt)
			in.UserSalt = UserSalt
		}

		//更新密码操作
		_, err := dao.AdminInfo.
			Ctx(ctx).
			Data(in).
			FieldsEx(dao.AdminInfo.Columns().Id).
			Where(dao.AdminInfo.Columns().Id, in.Id).
			Update()
		return err
	})
}

// GetList 查询内容列表
func (s *sAdmin) GetList(ctx context.Context, in model.AdminGetListInput) (out *model.AdminGetListOutput, err error) {
	var (
		m = dao.AdminInfo.Ctx(ctx)
	)
	out = &model.AdminGetListOutput{
		Page: in.Page,
		Size: in.Size,
	}

	// 分配查询
	listModel := m.Page(in.Page, in.Size)

	// 执行查询
	var list []*entity.AdminInfo
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
	// Admin
	//if err := listModel.ScanList(&out.List, "Admin"); err != nil {
	//	return out, err
	//}

	if err := listModel.Scan(&out.List); err != nil {
		return out, err
	}
	return
}
