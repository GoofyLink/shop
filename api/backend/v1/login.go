package v1

import (
	"github.com/gogf/gf/v2/frame/g"
	"time"
)

type LoginDoReq struct {
	g.Meta   `path:"/backend/v1/login" method:"post" summary:"执行登录请求" tags:"登录"`
	Name     string `json:"name" v:"required#请输入账号"   dc:"账号"`
	Password string `json:"password" v:"required#请输入密码"   dc:"密码(明文)"`
	IsAdmin  string `json:"is_admin"  dc:"是否管理员"`
}
type LoginDoRes struct {
	//Referer string `json:"referer" dc:"引导客户端跳转地址"`
	//Info interface{} `json:"info"`
	// todo token
	Token  string    `json:"token"`
	Expire time.Time `json:"expire"`
}

type RefreshTokenReq struct {
	g.Meta `path:"/backend/v1/refresh_token" method:"post"`
}

type RefreshTokenRes struct {
	Token  string    `json:"token"`
	Expire time.Time `json:"expire"`
}

type LogoutReq struct {
	g.Meta `path:"/backend/v1/logout" method:"post"`
}

type LogoutRes struct {
}
