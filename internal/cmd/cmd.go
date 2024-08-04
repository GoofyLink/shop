package cmd

import (
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gcmd"
	"shop-v2/internal/controller"
	"shop-v2/internal/service"
)

var (
	Main = gcmd.Command{
		Name:  "main",
		Usage: "main",
		Brief: "start http server",
		Func: func(ctx context.Context, parser *gcmd.Parser) (err error) {
			s := g.Server()
			s.Group("/", func(group *ghttp.RouterGroup) {
				// json数据返回设置
				//group.Middleware(ghttp.MiddlewareHandlerResponse)
				group.Middleware(
					service.Middleware().Ctx,
					service.Middleware().ResponseHandler,
				)
				group.Bind(
					controller.Rotation, //轮播图
					controller.Position, //手工位
					controller.Admin,    //角色管理
					controller.Login,    // 登录
				)
			})
			s.Run()
			return nil
		},
	}
)
