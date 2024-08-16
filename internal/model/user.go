package model

type RegisterInput struct {
	Name         string `json:"name"         dc:"用户名"     v:"required#用户名必填"     ` // 用户名
	Avatar       string `json:"avatar"       dc:"头像"        `                      // 头像
	Password     string `json:"password"     dc:"password"  v:"password"    `      //
	UserSalt     string `json:"userSalt"     dc:"加密盐 生成密码用"     `                  // 加密盐 生成密码用
	Sex          int    `json:"sex"          dc:"1男 2女"           `                // 1男 2女
	Status       int    `json:"status"       dc:"1正常 2拉黑冻结"        `               // 1正常 2拉黑冻结
	Sign         string `json:"sign"         dc:"个性签名"          `                  // 个性签名
	SecretAnswer string `json:"secretAnswer" dc:"密保问题的答案" `                        // 密保问题的答案
}

type RegisterOutPut struct {
	Id uint
}

type LoginUserInput struct {
	Name     string
	Password string
}
