package main

import (
	"github.com/kataras/iris/v12"
	"note-iris/src/application"
	"note-iris/src/router"
)

func main() {
	// 新建 Iris 实例
	app := iris.New()

	// 设置日志打印级别
	app.Logger().SetLevel("debug")

	// 注册路由
	router.RegisterRouter(app)
	// 启动服务
	app.Listen(
		application.App.Server.Host+":"+application.App.Server.Port,
		iris.WithCharset("UTF-8"),
	)
}
