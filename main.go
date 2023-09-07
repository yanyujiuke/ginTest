package main

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"html/template"
	"shop/models"
	"shop/routers"
)

func main() {
	// 创建；路由引擎
	r := gin.Default()
	// 自定义模板函数 放在加载模板之前
	r.SetFuncMap(template.FuncMap{
		"UnixToTime": models.UnixToTime,
	})

	// 加载模板 放在配置路由前面
	r.LoadHTMLGlob("templates/**/**/*")
	// 配置静态web目录 1路由 2映射的目录
	r.Static("/static", "./static")

	// 创建基于 cookie 的存储引擎，secret11111 参数是用于加密的密钥
	store := cookie.NewStore([]byte("secret111"))
	// 配置session的中间件 store是前面创建的存储引擎，我们可以替换成其他存储引擎
	r.Use(sessions.Sessions("mysession", store))

	// 初始化路由
	routers.AdminRoutersInit(r)

	r.Run()
}
