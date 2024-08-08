package role

import (
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"shop-v2/internal/dao"
	"shop-v2/internal/model"
	"shop-v2/internal/model/entity"
	"shop-v2/internal/service"
	utility "shop-v2/utility"
)

type sPermission struct{}

func init() {
	service.RegisterPermission(New())
}

func New() *sPermission {
	return &sPermission{}
}

// Create 创建内容
func (s *sPermission) Create(ctx context.Context, in model.PermissionCreateInput) (out model.PermissionCreateOutput, err error) {
	// 插入数据
	lastInsertID, err := dao.PermissionInfo.Ctx(ctx).Data(in).InsertAndGetId()
	if err != nil {
		return out, err
	}
	return model.PermissionCreateOutput{PermissionId: int(lastInsertID)}, err
}

// Delete 删除 这样删除是软删除还会留着原来的信息
func (s *sPermission) Delete(ctx context.Context, id uint) error {
	//return dao.PermissionInfo.Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
	// 删除内容
	_, err := dao.PermissionInfo.Ctx(ctx).Where(g.Map{
		dao.PermissionInfo.Columns().Id: id,
	}).Unscoped().Delete()
	//}).Unscoped().Delete()  unscoped直接删除信息 不会软删除
	return err
	//})
}

// Update 修改
func (s *sPermission) Update(ctx context.Context, in model.PermissionUpdateInput) error {
	//更新密码操作
	_, err := dao.PermissionInfo.
		Ctx(ctx).
		Data(in).
		FieldsEx(dao.PermissionInfo.Columns().Id).
		Where(dao.PermissionInfo.Columns().Id, in.Id).
		Update()
	return err

}

// GetList 查询内容列表
func (s *sPermission) GetList(ctx context.Context, in model.PermissionGetListInput) (out *model.PermissionGetListOutput, err error) {
	var (
		m = dao.PermissionInfo.Ctx(ctx)
	)
	out = &model.PermissionGetListOutput{
		Page: in.Page,
		Size: in.Size,
	}

	// 分配查询
	listModel := m.Page(in.Page, in.Size)

	// 执行查询
	var list []*entity.PermissionInfo
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

// 获取role密码
func (s *sPermission) GetPermissionByUserNamePassword(ctx context.Context, in model.LoginInput) map[string]interface{} {
	// todo 对接DB
	//验证账号密码是否正确
	roleInfo := entity.PermissionInfo{}
	err := dao.PermissionInfo.Ctx(ctx).Where("name", in.Name).Scan(&roleInfo)
	if err != nil {
		return nil
	}
	// 获取的加密字符和DB里面的字符比较
	if utility.EncryptPassword(in.Password, roleInfo.Name) != roleInfo.Name {
		return nil
	} else {
		return g.Map{
			"id":       roleInfo.Id,
			"username": roleInfo.Name,
		}
	}
}
