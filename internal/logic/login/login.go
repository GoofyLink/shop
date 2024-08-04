package login

import (
	"context"
	"github.com/gogf/gf/v2/errors/gerror"
	"shop-v2/internal/dao"
	"shop-v2/internal/model"
	"shop-v2/internal/model/entity"
	"shop-v2/internal/service"
	library "shop-v2/utility"
)

type sLogin struct{}

func init() {
	service.RegisterLogin(New())
}

func New() *sLogin {
	return &sLogin{}
}

// 执行登录
func (s *sLogin) Login(ctx context.Context, in model.LoginInput) error {
	//验证账号密码是否正确
	adminInfo := entity.AdminInfo{}
	err := dao.AdminInfo.Ctx(ctx).Where("name", in.Name).Scan(&adminInfo)
	if err != nil {
		return err
	}
	if library.EncryptPassword(in.Password, adminInfo.UserSalt) != adminInfo.Password {
		return gerror.New("账号或者密码不正确")
	}
	if err := service.Session().SetUser(ctx, &adminInfo); err != nil {
		return err
	}
	// 自动更新上线
	service.BizCtx().SetUser(ctx, &model.ContextUser{
		Id:      uint(adminInfo.Id),
		Name:    adminInfo.Name,
		IsAdmin: uint8(adminInfo.IsAdmin),
	})
	return nil
}
