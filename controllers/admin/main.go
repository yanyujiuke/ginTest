package admin

import (
	"encoding/json"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"net/http"
	"shop/models"
)

type MainController struct{}

func (con MainController) Index(c *gin.Context) {
	// 获取用户登陆的session信息
	session := sessions.Default(c)
	managerinfo := session.Get("managerinfo")
	// 类型断言 判断 managerinfo 是不是一个 string
	managerStr, ok := managerinfo.(string)
	if ok {
		var managerinfoStruct []models.Manager
		json.Unmarshal([]byte(managerStr), &managerinfoStruct)

		c.HTML(http.StatusOK, "admin/main/index.html", gin.H{
			"username": managerinfoStruct[0].Username,
		})
	} else {
		c.Redirect(302, "admin/login")
	}
}

func (con MainController) Welcome(c *gin.Context) {
	c.HTML(http.StatusOK, "admin/main/welcome.html", gin.H{})
}
