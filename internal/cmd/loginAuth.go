package cmd

import (
	"context"
	"github.com/goflyfox/gtoken/gtoken"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/text/gstr"
	"github.com/gogf/gf/v2/util/gconv"
	v1 "shop-v2/api/backend/v1"
	v1f "shop-v2/api/frontend/v1"
	"shop-v2/internal/consts"
	"shop-v2/internal/dao"
	"shop-v2/internal/model/entity"
	library "shop-v2/utility"
	"shop-v2/utility/response"
	"strconv"
)

// 管理后台相关
func StartBackendGToken() (gfAdminToken *gtoken.GfToken, err error) {
	// 启动gtoken
	gfAdminToken = &gtoken.GfToken{
		ServerName: consts.BackendServerName,
		CacheMode:  consts.CacheModeRedis, //gredis  防止关闭服务后token失效  将token缓存到redis中
		//Timeout:          10 * 1000,
		LoginPath:        "/v1/login",
		LoginBeforeFunc:  login,
		LoginAfterFunc:   loginAfterFunc,
		LogoutPath:       "/v1/logout",
		AuthPaths:        g.SliceStr{"/backend/v1/admin/info"},
		AuthExcludePaths: g.SliceStr{"/admin/user/info", "/admin/system/user/info"}, // 不拦截路径 /user/info,/system/user/info,/system/user,
		MultiLogin:       consts.MultiLogin,                                         // 是否启用多端登录
		AuthAfterFunc:    authAfterFunc,
	}
	// todo 去掉全局校验只用路由组的校验
	err = gfAdminToken.Start()
	if err != nil {
		return nil, err
	}
	return
}

// todo 迁移合适的位置
// 后台 登录处理
func login(r *ghttp.Request) (string, interface{}) {
	name := r.Get("name").String()
	password := r.Get("password").String()

	if name == "" || password == "" {
		r.Response.WriteJson(gtoken.Fail(consts.ErrorLoginFailedMsg))
		r.ExitAll()
	}
	ctx := context.TODO()
	//验证账号密码是否正确
	adminInfo := entity.AdminInfo{}
	err := dao.AdminInfo.Ctx(ctx).Where(dao.AdminInfo.Columns().Name, name).Scan(&adminInfo)
	if err != nil {
		r.Response.WriteJson(gtoken.Fail(consts.ErrorLoginFailedMsg))
		r.ExitAll()
	}
	if library.EncryptPassword(password, adminInfo.UserSalt) != adminInfo.Password {
		r.Response.WriteJson(gtoken.Fail(consts.ErrorLoginFailedMsg))
		r.ExitAll()
	}

	// 唯一标识
	return consts.GtoKenAdminPrefix + strconv.Itoa(adminInfo.Id), adminInfo
}

// 后台 登录后的处理
func loginAfterFunc(r *ghttp.Request, respData gtoken.Resp) {
	//g.Dump("respData:", respData)
	if !respData.Success() {
		respData.Code = 0
		r.Response.WriteJson(respData)
		return
	} else {
		respData.Code = 1
		//获得登录用户id
		userKey := respData.GetString("userKey")
		adminId := gstr.StrEx(userKey, consts.GtoKenAdminPrefix)
		//根据id获得登录用户其他信息
		adminInfo := entity.AdminInfo{}
		err := dao.AdminInfo.Ctx(context.TODO()).WherePri(adminId).Scan(&adminInfo)
		if err != nil {
			return
		}
		//通过角色查询权限
		//先通过角色查询权限id
		var rolePermissionInfos []entity.RolePermissionInfo
		err = dao.RolePermissionInfo.Ctx(context.TODO()).WhereIn(dao.RolePermissionInfo.Columns().RoleId, g.Slice{adminInfo.RoleIds}).Scan(&rolePermissionInfos)
		if err != nil {
			return
		}
		permissionIds := g.Slice{}
		for _, info := range rolePermissionInfos {
			permissionIds = append(permissionIds, info.PermissionId)
		}

		var permissions []entity.PermissionInfo
		err = dao.PermissionInfo.Ctx(context.TODO()).WhereIn(dao.PermissionInfo.Columns().Id, permissionIds).Scan(&permissions)
		if err != nil {
			return
		}
		data := &v1.LoginRes{
			Type:        consts.TokenType,
			Token:       respData.GetString("token"),
			ExpireIn:    consts.GTokenExpireIn, //单位秒,
			IsAdmin:     adminInfo.IsAdmin,
			RoleIds:     adminInfo.RoleIds,
			Permissions: permissions,
		}
		response.JsonExit(r, 0, "", data)
	}
	return
}

// 后台 登录之后用户校验
func authAfterFunc(r *ghttp.Request, respData gtoken.Resp) {
	g.Dump(respData)
	var adminInfo entity.AdminInfo
	err := gconv.Struct(respData.GetString("data"), &adminInfo)
	if err != nil {
		response.Auth(r)
		return
	}

	r.SetCtxVar(consts.CtxAdminId, adminInfo.Id)
	r.SetCtxVar(consts.CtxAdminName, adminInfo.Name)
	r.SetCtxVar(consts.CtxAdminIsAdmin, adminInfo.IsAdmin)
	r.SetCtxVar(consts.CtxAdminRoleIds, adminInfo.RoleIds)
	r.Middleware.Next()
}

// 前台登录相关
func StartFrontendGToken() (gfAdminToken *gtoken.GfToken, err error) {
	// 启动gtoken
	gfAdminToken = &gtoken.GfToken{
		ServerName: consts.BackendServerName,
		CacheMode:  consts.CacheModeRedis, //gredis  防止关闭服务后token失效  将token缓存到redis中
		//Timeout:          10 * 1000,
		LoginPath:       "/v1/user/login",
		LoginBeforeFunc: loginFuncFrontend,
		LoginAfterFunc:  loginAfterFuncFrontend,
		LogoutPath:      "/v1/user/logout",
		//AuthPaths:        g.SliceStr{"/backend/v1/admin/info"},
		//AuthExcludePaths: g.SliceStr{"/admin/user/info", "/admin/system/user/info"}, // 不拦截路径 /user/info,/system/user/info,/system/user,
		AuthAfterFunc: authAfterFuncFrontend,     // 校验通过之后通过gtoken获取到值
		MultiLogin:    consts.FrontendMultiLogin, // 是否启用多端登录
	}
	// todo 去掉全局校验只用路由组的校验
	//err = gfAdminToken.Start()
	if err != nil {
		return nil, err
	}
	return
}

// 前台登录
func loginFuncFrontend(r *ghttp.Request) (string, interface{}) {
	name := r.Get("name").String()
	password := r.Get("password").String()
	if name == "" || password == "" {
		r.Response.WriteJson(gtoken.Fail(consts.ErrorLoginFailedMsg))
		r.ExitAll()
	}
	ctx := context.TODO()
	//验证账号密码是否正确
	userInfo := entity.UserInfo{}
	err := dao.UserInfo.Ctx(ctx).Where(dao.UserInfo.Columns().Name, name).Scan(&userInfo)
	if err != nil {
		r.Response.WriteJson(gtoken.Fail(consts.ErrorLoginFailedMsg))
		r.ExitAll()
	}
	if library.EncryptPassword(password, userInfo.UserSalt) != userInfo.Password {
		r.Response.WriteJson(gtoken.Fail(consts.ErrorLoginFailedMsg))
		r.ExitAll()
	}

	// 唯一标识
	return consts.GtoKenUserPrefix + strconv.Itoa(userInfo.Id), userInfo
}

// 前台 登录后的处理
func loginAfterFuncFrontend(r *ghttp.Request, respData gtoken.Resp) {
	//g.Dump("respData:", respData)
	if !respData.Success() {
		respData.Code = 0
		r.Response.WriteJson(respData)
		return
	} else {
		respData.Code = 1
		//获得登录用户id
		userKey := respData.GetString("userKey") // gtoken定义的
		userId := gstr.StrEx(userKey, consts.GtoKenUserPrefix)
		//根据id获得登录用户其他信息
		userInfo := entity.UserInfo{}
		err := dao.UserInfo.Ctx(context.TODO()).WherePri(userId).Scan(&userInfo)
		if err != nil {
			return
		}
		data := &v1f.LoginRes{
			Type:     consts.TokenType,
			Token:    respData.GetString("token"),
			ExpireIn: consts.GTokenExpireIn, //单位秒,

		}
		data.Name = userInfo.Name
		data.Avatar = userInfo.Avatar
		data.Sign = userInfo.Sign
		data.Status = uint8(userInfo.Status)
		response.JsonExit(r, 0, "", data)
	}
	return
}

// 前台用户信息保存
func authAfterFuncFrontend(r *ghttp.Request, respData gtoken.Resp) {
	//g.Dump(respData)
	var userInfo entity.UserInfo
	err := gconv.Struct(respData.GetString("data"), &userInfo)
	if err != nil {
		response.Auth(r)
		return
	}

	r.SetCtxVar(consts.CtxUserId, userInfo.Id)
	r.SetCtxVar(consts.CtxUserName, userInfo.Name)
	r.SetCtxVar(consts.CtxUserSex, userInfo.Sex)
	r.SetCtxVar(consts.CtxUserAvatar, userInfo.Avatar)
	r.SetCtxVar(consts.CtxUserSign, userInfo.Sign)
	r.SetCtxVar(consts.CtxUserStatus, userInfo.Status)
	r.Middleware.Next()
}
