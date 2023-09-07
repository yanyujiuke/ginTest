package middlewares

import (
	"encoding/json"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"shop/models"
	"strings"
)

func InitAdminAuthMiddleware(c *gin.Context) {

	// 1 获取访问的url地址
	pathname := strings.Split(c.Request.URL.String(), "?")[0]

	// 2 获取登陆的session数据
	session := sessions.Default(c)
	managerinfo := session.Get("managerinfo")

	// 类型断言来判断managerinfo是不是一个string
	managerinfoStr, ok := managerinfo.(string)
	if ok {
		// 判断用户信息是否存在
		var managerStruct []models.Manager
		err := json.Unmarshal([]byte(managerinfoStr), &managerStruct)
		if err != nil || !(len(managerStruct) > 0 && managerStruct[0].Username != "") {
			if pathname != "/admin/login" && pathname != "/admin/doLogin" && pathname != "/admin/captcha" {
				c.Redirect(301, "/admin/login")
			}
		}
	} else {
		// 用户没有登陆
		if pathname != "/admin/login" && pathname != "/admin/doLogin" && pathname != "/admin/captcha" {
			c.Redirect(302, "/admin/login")
		}
	}
}
