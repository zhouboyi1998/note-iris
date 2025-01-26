package main

import (
	"github.com/kataras/iris/v12"
	"net/http"
)

func main() {
	// 新建 Iris 实例
	app := iris.New()

	// 设置日志打印级别
	app.Logger().SetLevel("debug")

	// Hello World
	app.Get("/hello/{name}", func(ctx iris.Context) {
		name := ctx.Params().Get("name")
		ctx.JSON(iris.Map{
			"code":    http.StatusOK,
			"message": "Hello, " + name,
		})
	})

	// 启动服务
	app.Run(iris.Addr(":18098"), iris.WithCharset("UTF-8"))
}
