package v1

import (
	"github.com/gogf/gf/v2/frame/g"
)

type UserInfoBase struct {
	Id     uint   `json:"id"`
	Name   string `json:"name"`
	Avatar string `json:"avatar"`
	Sex    uint8  `json:"sex"`
	Sign   string `json:"sign"`
	Status uint8  `json:"status"`
}

// user注册
type RegisterReq struct {
	g.Meta       `path:"/v1/register" method:"post" tags:"前台用户"  summary:"用户注册" `
	Name         string `json:"name"         dc:"用户名"     v:"required#用户名必填"   ` // 用户名
	Password     string `json:"password"     dc:"password"  v:"password"    `
	Avatar       string `json:"avatar"       dc:"头像"        `        // 头像
	UserSalt     string `json:"userSalt"     dc:"加密盐 生成密码用"     `    // 加密盐 生成密码用
	Sex          int    `json:"sex"          dc:"1男 2女"           `  // 1男 2女
	Status       int    `json:"status"       dc:"1正常 2拉黑冻结"        ` // 1正常 2拉黑冻结
	Sign         string `json:"sign"         dc:"个性签名"          `    // 个性签名
	SecretAnswer string `json:"secretAnswer" dc:"密保问题的答案" `          // 密保问题的答案
}

type RegisterRes struct {
	Id uint `json:"id"`
}

type LoginReq struct {
	g.Meta   `path:"/v1/user/login" method:"post" tags:"前台用户" summary:"用户登录"`
	Name     string `json:"name"         dc:"用户名"     v:"required#用户名必填"   ` // 用户名
	Password string `json:"password"     dc:"password"  v:"password"    `
}
type LoginRes struct {
	Type     string `json:"type"` // token的类型
	Token    string `json:"token"`
	ExpireIn int    `json:"expire_in"`
	UserInfoBase
}

type UserInfoReq struct {
	g.Meta `path:"/v1/user/info" method:"get" tags:"前台用户" summary:"当前登录用户信息"`
}

type UserInfoRes struct {
	UserInfoBase
}

// 修改密码
type UpdatePasswordReq struct {
	g.Meta       `path:"/v1/user/update/password" method:"post" tags:"前台用户" summary:"修改密码"`
	Password     string `json:"password" v:"password" dc:"mm"`
	UserSalt     string `json:"user_salt,omitempty" dc:"加盐，加密"`
	SecretAnswer string `json:"secret_answer" dc:"密保问题答案"`
}

type UpdatePasswordRes struct {
	Id uint `json:"id"`
}
