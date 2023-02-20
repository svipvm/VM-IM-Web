package router

// import (
// 	"../app/middleware"
// 	"../app/widget"

// 	"github.com/kataras/iris/v12"
// 	"github.com/kataras/iris/v12/sessions"
// )

// // 注册路由
// func RegisterRouter(app *iris.Application) {
// 	// 注册SESSION中间件
// 	session := sessions.New(sessions.Config{
// 		Cookie: sessions.DefaultCookieName,
// 	})
// 	// SESSION中间件
// 	app.Use(session.Handler())
// 	// 登录验证中间件
// 	app.Use(middleware.CheckLogin)

// 	// 视图文件目录 每次请求时自动重载模板
// 	tmpl := iris.HTML("./view", ".html").Reload(true)

// 	// 注册框架核心组件

// 	// 安全组件
// 	tmpl.AddFunc("safe", widget.Safe)
// 	// 日期选择组件
// 	tmpl.AddFunc("date", widget.Date)
// 	// 自定义按钮组件
// 	tmpl.AddFunc("widget", widget.Widget)
// 	// 查询按钮组件
// 	tmpl.AddFunc("query", widget.Query)
// 	// 添加按钮组件
// 	tmpl.AddFunc("add", widget.Add)
// 	// 编辑按钮组件
// 	tmpl.AddFunc("edit", widget.Edit)
// 	// 删除按钮组件(支持大、小按钮)
// 	tmpl.AddFunc("delete", widget.Delete)
// 	// 批量删除按钮组件
// 	tmpl.AddFunc("dall", widget.Dall)
// 	// 展开按钮组件
// 	tmpl.AddFunc("expand", widget.Expand)
// 	// 收缩按钮组件
// 	tmpl.AddFunc("collapse", widget.Collapse)
// 	// 添加子级按钮组件
// 	tmpl.AddFunc("addz", widget.Addz)
// 	// 开关按钮组件
// 	tmpl.AddFunc("switch", widget.Switch)
// 	// 单选下拉组件
// 	tmpl.AddFunc("select", widget.Select)
// 	// 提交表单按钮组件
// 	tmpl.AddFunc("submit", widget.Submit)
// 	// 图标选择按钮组件
// 	tmpl.AddFunc("icon", widget.Icon)
// 	// 穿梭组件
// 	tmpl.AddFunc("transfer", widget.Transfer)
// 	// 单图上传组件
// 	tmpl.AddFunc("upload_image", widget.UploadImage)
// 	// 多图上传组件
// 	tmpl.AddFunc("album", widget.Album)
// 	// 站点选择组件
// 	tmpl.AddFunc("item", widget.Item)
// 	// 富文本编辑器组件
// 	tmpl.AddFunc("kindeditor", widget.Kindeditor)
// 	// 复选框组件
// 	tmpl.AddFunc("checkbox", widget.Checkbox)
// 	// 单选按钮组件
// 	tmpl.AddFunc("radio", widget.Radio)
// 	// 行政区划选择组件
// 	tmpl.AddFunc("city", widget.City)

// 	// 注册视图
// 	app.RegisterView(tmpl)

// 	// 静态文件
// 	app.HandleDir("/static", "./public/static")

// 	// 错误请求配置
// 	app.OnErrorCode(iris.StatusNotFound, func(ctx iris.Context) {
// 		ctx.View("error/404.html")
// 	})
// 	app.OnErrorCode(iris.StatusInternalServerError, func(ctx iris.Context) {
// 		ctx.View("error/500.html")
// 	})

// 	// 登录、主页
// 	index := app.Party("/")
// 	{
// 		index.Get("/", controllers.Index.Index)
// 		index.Any("/login", controllers.Login.Login)
// 		index.Get("/captcha", controllers.Login.Captcha)
// 		index.Get("/index", controllers.Index.Index)
// 		index.Get("/main", controllers.Index.Main)
// 		index.Any("/userInfo", controllers.Index.UserInfo)
// 		index.Any("/updatePwd", controllers.Index.UpdatePwd)
// 		index.Get("/logout", controllers.Index.Logout)
// 	}

// 	// 文件上传
// 	upload := app.Party("/upload")
// 	{
// 		upload.Post("/uploadImage", controllers.Upload.UploadImage)
// 		upload.Post("/uploadEditImage", controllers.Upload.UploadEditImage)
// 	}

// 	// 职级管理
// 	level := app.Party("/level")
// 	{
// 		level.Get("/index", controllers.Level.Index)
// 		level.Post("/list", controllers.Level.List)
// 		level.Get("/edit/{id:int}", controllers.Level.Edit)
// 		level.Post("/add", controllers.Level.Add)
// 		level.Post("/update", controllers.Level.Update)
// 		level.Post("/delete/{id:int}", controllers.Level.Delete)
// 		level.Post("/setStatus", controllers.Level.Status)
// 	}

// 	// 岗位管理
// 	position := app.Party("/position")
// 	{
// 		position.Get("/index", controllers.Position.Index)
// 		position.Post("/list", controllers.Position.List)
// 		position.Get("/edit/{id:int}", controllers.Position.Edit)
// 		position.Post("/add", controllers.Position.Add)
// 		position.Post("/update", controllers.Position.Update)
// 		position.Post("/delete/{id:int}", controllers.Position.Delete)
// 		position.Post("/setStatus", controllers.Position.Status)
// 	}

// 	/* 角色路由 */
// 	role := app.Party("role")
// 	{
// 		role.Get("/index", controllers.Role.Index)
// 		role.Post("/list", controllers.Role.List)
// 		role.Get("/edit/{id:int}", controllers.Role.Edit)
// 		role.Post("/add", controllers.Role.Add)
// 		role.Post("/update", controllers.Role.Update)
// 		role.Post("/delete/{id:int}", controllers.Role.Delete)
// 		role.Post("/setStatus", controllers.Role.Status)
// 		role.Get("/getRoleList", controllers.Role.GetRoleList)
// 	}

// 	/* 角色菜单权限 */
// 	roleMenu := app.Party("rolemenu")
// 	{
// 		roleMenu.Get("/index/{roleId:int}", controllers.RoleMenu.Index)
// 		roleMenu.Post("/save", controllers.RoleMenu.Save)
// 	}

// 	/* 部门管理 */
// 	dept := app.Party("dept")
// 	{
// 		dept.Get("/index", controllers.Dept.Index)
// 		dept.Post("/list", controllers.Dept.List)
// 		dept.Get("/edit/{id:int}", controllers.Dept.Edit)
// 		dept.Get("/edit/{id:int}/{pid:string}", controllers.Dept.Edit)
// 		dept.Post("/add", controllers.Dept.Add)
// 		dept.Post("/update", controllers.Dept.Update)
// 		dept.Post("/delete/{id:int}", controllers.Dept.Delete)
// 		dept.Get("/getDeptList", controllers.Dept.GetDeptList)
// 	}

// 	/* 用户管理 */
// 	user := app.Party("user")
// 	{
// 		user.Get("/index", controllers.User.Index)
// 		user.Post("/list", controllers.User.List)
// 		user.Get("/edit/{id:int}", controllers.User.Edit)
// 		user.Post("/add", controllers.User.Add)
// 		user.Post("/update", controllers.User.Update)
// 		user.Post("/delete/{id:int}", controllers.User.Delete)
// 		user.Post("/setStatus", controllers.User.Status)
// 		user.Post("/resetPwd", controllers.User.ResetPwd)
// 	}

// 	/* 菜单管理 */
// 	menu := app.Party("menu")
// 	{
// 		menu.Get("/index", controllers.Menu.Index)
// 		menu.Post("/list", controllers.Menu.List)
// 		menu.Get("/edit/{id:int}", controllers.Menu.Edit)
// 		menu.Get("/edit/{id:int}/{pid:string}", controllers.Menu.Edit)
// 		menu.Post("/add", controllers.Menu.Add)
// 		menu.Post("/update", controllers.Menu.Update)
// 		menu.Post("/delete/{id:int}", controllers.Menu.Delete)
// 	}

// 	/* 友链管理 */
// 	link := app.Party("link")
// 	{
// 		link.Get("/index", controllers.Link.Index)
// 		link.Post("/list", controllers.Link.List)
// 		link.Get("/edit/{id:int}", controllers.Link.Edit)
// 		link.Post("/add", controllers.Link.Add)
// 		link.Post("/update", controllers.Link.Update)
// 		link.Post("/delete/{id:int}", controllers.Link.Delete)
// 		link.Post("/setStatus", controllers.Link.Status)
// 	}

// 	/* 城市管理 */
// 	city := app.Party("city")
// 	{
// 		city.Get("/index", controllers.City.Index)
// 		city.Post("/list", controllers.City.List)
// 		city.Get("/edit/{id:int}", controllers.City.Edit)
// 		city.Get("/edit/{id:int}/{pid:string}", controllers.City.Edit)
// 		city.Post("/add", controllers.City.Add)
// 		city.Post("/update", controllers.City.Update)
// 		city.Post("/delete/{id:int}", controllers.City.Delete)
// 		city.Post("/getChilds", controllers.City.GetChilds)
// 	}

// 	/* 通知管理 */
// 	notice := app.Party("notice")
// 	{
// 		notice.Get("/index", controllers.Notice.Index)
// 		notice.Post("/list", controllers.Notice.List)
// 		notice.Get("/edit/{id:int}", controllers.Notice.Edit)
// 		notice.Post("/add", controllers.Notice.Add)
// 		notice.Post("/update", controllers.Notice.Update)
// 		notice.Post("/delete/{id:int}", controllers.Notice.Delete)
// 		notice.Post("/setStatus", controllers.Notice.Status)
// 	}

// 	/* 会员等级 */
// 	memberlevel := app.Party("memberlevel")
// 	{
// 		memberlevel.Get("/index", controllers.MemberLevel.Index)
// 		memberlevel.Post("/list", controllers.MemberLevel.List)
// 		memberlevel.Get("/edit/{id:int}", controllers.MemberLevel.Edit)
// 		memberlevel.Post("/add", controllers.MemberLevel.Add)
// 		memberlevel.Post("/update", controllers.MemberLevel.Update)
// 		memberlevel.Post("/delete/{id:int}", controllers.MemberLevel.Delete)
// 		memberlevel.Get("/getMemberLevelList", controllers.MemberLevel.GetMemberLevelList)
// 	}

// 	/* 会员管理 */
// 	member := app.Party("member")
// 	{
// 		member.Get("/index", controllers.Member.Index)
// 		member.Post("/list", controllers.Member.List)
// 		member.Get("/edit/{id:int}", controllers.Member.Edit)
// 		member.Post("/add", controllers.Member.Add)
// 		member.Post("/update", controllers.Member.Update)
// 		member.Post("/delete/{id:int}", controllers.Member.Delete)
// 		member.Post("/setStatus", controllers.Member.Status)
// 	}

// 	/* 字典管理 */
// 	dict := app.Party("dict")
// 	{
// 		dict.Get("/index", controllers.Dict.Index)
// 		dict.Post("/list", controllers.Dict.List)
// 		dict.Post("/add", controllers.Dict.Add)
// 		dict.Post("/update", controllers.Dict.Update)
// 		dict.Post("/delete/{id:int}", controllers.Dict.Delete)
// 	}

// 	/* 字典项管理 */
// 	dictdata := app.Party("dictdata")
// 	{
// 		dictdata.Post("/list", controllers.DictData.List)
// 		dictdata.Post("/add", controllers.DictData.Add)
// 		dictdata.Post("/update", controllers.DictData.Update)
// 		dictdata.Post("/delete/{id:int}", controllers.DictData.Delete)
// 	}

// 	/* 配置管理 */
// 	config := app.Party("config")
// 	{
// 		config.Get("/index", controllers.Config.Index)
// 		config.Post("/list", controllers.Config.List)
// 		config.Post("/add", controllers.Config.Add)
// 		config.Post("/update", controllers.Config.Update)
// 		config.Post("/delete/{id:int}", controllers.Config.Delete)
// 	}

// 	/* 配置项管理 */
// 	configdata := app.Party("configdata")
// 	{
// 		configdata.Post("/list", controllers.ConfigData.List)
// 		configdata.Post("/add", controllers.ConfigData.Add)
// 		configdata.Post("/update", controllers.ConfigData.Update)
// 		configdata.Post("/delete/{id:int}", controllers.ConfigData.Delete)
// 		configdata.Post("/setStatus", controllers.ConfigData.Status)
// 	}

// 	/* 网站设置 */
// 	configweb := app.Party("configweb")
// 	{
// 		configweb.Any("/index", controllers.ConfigWeb.Index)
// 	}

// 	/* 站点管理 */
// 	item := app.Party("item")
// 	{
// 		item.Get("/index", controllers.Item.Index)
// 		item.Post("/list", controllers.Item.List)
// 		item.Get("/edit/{id:int}", controllers.Item.Edit)
// 		item.Post("/add", controllers.Item.Add)
// 		item.Post("/update", controllers.Item.Update)
// 		item.Post("/delete/{id:int}", controllers.Item.Delete)
// 		item.Post("/setStatus", controllers.Item.Status)
// 		item.Get("/getItemList", controllers.Item.GetItemList)
// 	}

// 	/* 栏目管理 */
// 	itemcate := app.Party("itemcate")
// 	{
// 		itemcate.Get("/index", controllers.ItemCate.Index)
// 		itemcate.Post("/list", controllers.ItemCate.List)
// 		itemcate.Get("/edit/{id:int}", controllers.ItemCate.Edit)
// 		itemcate.Get("/edit/{id:int}/{pid:string}", controllers.ItemCate.Edit)
// 		itemcate.Post("/add", controllers.ItemCate.Add)
// 		itemcate.Post("/update", controllers.ItemCate.Update)
// 		itemcate.Post("/delete/{id:int}", controllers.ItemCate.Delete)
// 		itemcate.Get("/getCateList", controllers.ItemCate.GetCateList)
// 		itemcate.Get("/getCateTreeList", controllers.ItemCate.GetCateTreeList)
// 	}

// 	/* 广告位管理 */
// 	adsort := app.Party("adsort")
// 	{
// 		adsort.Get("/index", controllers.AdSort.Index)
// 		adsort.Post("/list", controllers.AdSort.List)
// 		adsort.Get("/edit/{id:int}", controllers.AdSort.Edit)
// 		adsort.Post("/add", controllers.AdSort.Add)
// 		adsort.Post("/update", controllers.AdSort.Update)
// 		adsort.Post("/delete/{id:int}", controllers.AdSort.Delete)
// 		adsort.Get("/getAdSortList", controllers.AdSort.GetAdSortList)
// 	}

// 	/* 广告管理 */
// 	ad := app.Party("ad")
// 	{
// 		ad.Get("/index", controllers.Ad.Index)
// 		ad.Post("/list", controllers.Ad.List)
// 		ad.Get("/edit/{id:int}", controllers.Ad.Edit)
// 		ad.Post("/add", controllers.Ad.Add)
// 		ad.Post("/update", controllers.Ad.Update)
// 		ad.Post("/delete/{id:int}", controllers.Ad.Delete)
// 		ad.Post("/setStatus", controllers.Ad.Status)
// 	}

// 	/* 统计分析 */
// 	analysis := app.Party("analysis")
// 	{
// 		analysis.Get("/index", controllers.Analysis.Index)
// 	}

// 	/* 代码生成器 */
// 	generate := app.Party("generate")
// 	{
// 		generate.Get("/index", controllers.Generate.Index)
// 		generate.Post("/list", controllers.Generate.List)
// 		generate.Post("/generate", controllers.Generate.Generate)
// 		generate.Post("/batchGenerate", controllers.Generate.BatchGenerate)
// 	}

// 	/* 演示一 */
// 	example := app.Party("example")
// 	{
// 		example.Get("/index", controllers.Example.Index)
// 		example.Post("/list", controllers.Example.List)
// 		example.Get("/edit/{id:int}", controllers.Example.Edit)
// 		example.Post("/add", controllers.Example.Add)
// 		example.Post("/update", controllers.Example.Update)
// 		example.Post("/delete/{id:int}", controllers.Example.Delete)
// 		example.Post("/setStatus", controllers.Example.Status)
// 		example.Post("/setIsVip", controllers.Example.IsVip)
// 	}

// 	/* 演示二 */
// 	example2 := app.Party("example2")
// 	{
// 		example2.Get("/index", controllers.Example2.Index)
// 		example2.Post("/list", controllers.Example2.List)
// 		example2.Get("/edit/{id:int}", controllers.Example2.Edit)
// 		example2.Post("/add", controllers.Example2.Add)
// 		example2.Post("/update", controllers.Example2.Update)
// 		example2.Post("/delete/{id:int}", controllers.Example2.Delete)
// 		example2.Post("/setStatus", controllers.Example2.Status)
// 	}
// }
