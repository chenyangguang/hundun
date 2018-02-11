package main

import (
	"github.com/kataras/iris"

	"github.com/kataras/iris/middleware/logger"
	"github.com/kataras/iris/middleware/recover"
)

func main() {
	app := iris.New()
	app.Logger().SetLevel("debug")
	app.Use(recover.New())
	app.Use(logger.New())

	app.Handle("GET", "/", func(ctx iris.Context) {
		ctx.HTML("<h1>Hi</h1>")
	})

	app.Get("/stock", func(ctx iris.Context) {
		ctx.JSON(iris.Map{"msg": "test other go framework"})
	})

	app.Run(iris.Addr(":9000"), iris.WithoutServerError(iris.ErrServerClosed))
}
