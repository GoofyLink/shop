package user

import (
	"context"
	"errors"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/gogf/gf/v2/util/grand"
	"shop-v2/internal/consts"
	"shop-v2/internal/dao"
	"shop-v2/internal/model"
	"shop-v2/internal/model/do"
	"shop-v2/internal/service"
	utility "shop-v2/utility"
)

type sUser struct{}

func init() {
	service.RegisterUser(New())
}

func New() *sUser {
	return &sUser{}
}

// Register 创建内容
func (s *sUser) Register(ctx context.Context, in model.RegisterInput) (out model.RegisterOutPut, err error) {
	// 特殊处理加密 盐 和密码
	UserSalt := grand.S(10)
	in.Password = utility.EncryptPassword(in.Password, UserSalt)
	in.UserSalt = UserSalt
	// 插入数据
	lastInsertID, err := dao.UserInfo.Ctx(ctx).Data(in).InsertAndGetId()
	if err != nil {
		return out, err
	}
	return model.RegisterOutPut{Id: uint(lastInsertID)}, err
}

// UpdatePassword 修改密码
func (s *sUser) UpdatePassword(ctx context.Context, in model.UpdatePasswordInput) (out model.UpdatePasswordOutPut, err error) {
	// 1.验证密保问题
	userInfo := do.UserInfo{}
	userId := gconv.Uint(ctx.Value(consts.CtxUserId))
	err = dao.UserInfo.Ctx(ctx).WherePri(userId).Scan(&userInfo)
	if err != nil {
		return model.UpdatePasswordOutPut{}, err
	}
	// 2.判断是否一致

	// 这里类型必须也一样
	if gconv.String(userInfo.SecretAnswer) != in.SecretAnswer {
		return model.UpdatePasswordOutPut{}, errors.New(consts.ErrorSecretAnswer)
	}
	userSalt := grand.S(10)
	in.UserSalt = userSalt
	in.Password = utility.EncryptPassword(in.Password, userSalt)
	// 3.更新密码
	_, err = dao.UserInfo.Ctx(ctx).WherePri(userId).Update(in)
	if err != nil {
		return model.UpdatePasswordOutPut{}, err
	}
	return model.UpdatePasswordOutPut{Id: userId}, nil
}
