package consts

const (
	ProjectName        = "shop"
	ProjectUsage       = "ProjectUsage"
	ProjectBrief       = "ProjectBrief"
	Version            = "v0.2.0"             // 当前服务版本(用于模板展示)
	CaptchaDefaultName = "CaptchaDefaultName" // 验证码默认存储空间名称
	ContextKey         = "ContextKey"         // 上下文变量存储键名，前后端系统共享
	//文件相关
	FileMaxUploadCountMinute = 10       // 同一用户1分钟之内最大上传数量
	GtoKenAdminPrefix        = "admin:" // 唯一标识
	GtoKenUserPrefix         = "user:"
	CtxAdminId               = "CtxAdminId"
	CtxAdminName             = "CtxAdminName"
	CtxAdminIsAdmin          = "CtxAdminIsAdmin"
	CtxAdminRoleIds          = "CtxAdminRoleIds"
	// 用户登录
	CtxUserId            = "CtxUserId"
	CtxUserName          = "CtxUserName"
	CtxUserAvatar        = "CtxUserAvatar"
	CtxUserSex           = "CtxUserSex"
	CtxUserSign          = "CtxUserSign"
	CtxUserStatus        = "CtxUserStatus"
	CtxRoleId            = "CtxRoleId"
	CtxRoleName          = "CtxRoleName"
	CtxRoleDesc          = "CtxRoleDesc"
	CodeMissingParameter = "请检查是否缺少参数"
	TokenType            = "Bearer"
	CacheModeRedis       = 2
	GTokenExpireIn       = 10 * 24 * 60 * 60
	BackendServerName    = "shop"
	MultiLogin           = true
	FrontendMultiLogin   = false
	ErrorLoginFailedMsg  = "账号或密码错误"
)
