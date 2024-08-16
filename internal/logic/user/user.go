package user

import (
	"context"
	"github.com/gogf/gf/v2/util/grand"
	"shop-v2/internal/dao"
	"shop-v2/internal/model"
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

// Create 创建内容
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
