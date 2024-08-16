package cmd

import (
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gcmd"
	"shop-v2/internal/consts"
	"shop-v2/internal/controller"
	"shop-v2/internal/service"
)

var (
	Main = gcmd.Command{
		Name:  consts.ProjectName,
		Usage: consts.ProjectUsage,
		Brief: consts.ProjectBrief,
		Func: func(ctx context.Context, parser *gcmd.Parser) (err error) {
			s := g.Server()
			// 认证接口
			//loginFunc := Login
			// 启动gtoken
			gfAdminToken, err := StartBackendGToken()
			if err != nil {
				return err
			}

			// 管理后台路由组
			s.Group("/backend", func(group *ghttp.RouterGroup) {
				// json数据返回设置
				//group.Middleware(ghttp.MiddlewareHandlerResponse)
				group.Middleware(
					service.Middleware().CORS,
					service.Middleware().Ctx,
					service.Middleware().ResponseHandler,
				)

				// 不需要登录验证的路由
				group.Bind(
					controller.Admin.Create, //角色管理 注册
					controller.Admin.List,
				)

				// 需要登录绑定的路由
				group.Group("/", func(group *ghttp.RouterGroup) {
					//group.Middleware(service.Middleware().Auth) jwt
					// 使用gtoken
					err := gfAdminToken.Middleware(ctx, group)
					if err != nil {
						panic(err)
					}
					//group.ALLMap(g.Map{ 这样也可以写
					//	"/v1/admin/info": controller.Admin.Info,
					//})
					group.Bind(
						controller.Admin.Update,
						controller.Admin.Delete,
						controller.Login,        // 登录
						controller.Data,         // 数据大屏
						controller.Role,         // 角色
						controller.Permission,   // 角色权限
						controller.Admin.Info,   // 查询当前管理员信息
						controller.Rotation,     //轮播图
						controller.Position,     //手工位
						controller.File,         //从0到一实现文件入库
						controller.Upload,       //实现可跨项目使用的文件上云工具类
						controller.Category,     //商品分类管理
						controller.Coupon,       //优惠券管理
						controller.UserCoupon,   // 用户优惠券
						controller.Goods,        // 商品
						controller.GoodsOptions, // 商品规格
						controller.Article,      // 文章管理
					)
				})
			})
			// 前台项目路由组
			frontendToken, err := StartFrontendGToken()
			s.Group("/frontend", func(group *ghttp.RouterGroup) {

				group.Middleware(
					service.Middleware().CORS,
					service.Middleware().Ctx,
					service.Middleware().ResponseHandler,
				)

				// 不需要登录验证的路由 todo
				group.Bind(
					controller.User.Register, //用户注册
					// 当前登录用户信息
					//controller.User.Login,    // 用户登录
					//controller.Admin.Create, //角色管理 注册
					//controller.Admin.List,
				)

				// 需要登录鉴权的路由组 todo
				group.Group("/", func(group *ghttp.RouterGroup) {
					err := frontendToken.Middleware(ctx, group)
					if err != nil {
						return
					}
					//需要登录鉴权的接口
					group.Bind(
						controller.User.Info,           // 获取当前用户信息
						controller.User.UpdatePassword, // 当前用户修改密码
						controller.Collection,          // 收藏
					)
				})
			})
			s.Run()
			return nil
		},
	}
)
