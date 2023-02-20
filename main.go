package main

import (
	"github.com/kataras/iris/v12"
)

func main() {
	app := iris.New()

	app.Get("/", func(ctx iris.Context) {
		ctx.WriteString("Hello World")
	})

	app.Run(iris.Addr(":18080"))
}

// import (
// 	// _ "easygoadmin/boot"
// 	// "easygoadmin/conf"
// 	// "easygoadmin/router"

// 	"github.com/kataras/iris/v12"
// 	"github.com/kataras/iris/v12/middleware/logger"
// 	"github.com/kataras/iris/v12/middleware/recover"
// )

// func main() {
// 	// 创建app结构体对象
// 	app := iris.New()
// 	// 设置调试模式
// 	app.Logger().SetLevel("debug")
// 	// 可选项添加两个内置的句柄（handlers）
// 	// 捕获相对于http产生的异常行为
// 	app.Use(recover.New())
// 	//记录请求日志
// 	app.Use(logger.New())
// 	// 初始化配置
// 	config := iris.WithConfiguration(iris.YAML("./conf/config.yml"))
// 	// 路由注册
// 	router.RegisterRouter(app)
// 	// 启动应用服务
// 	app.Run(iris.Addr(conf.CONFIG.EGAdmin.Addr+":"+conf.CONFIG.EGAdmin.Port), config)
// }
