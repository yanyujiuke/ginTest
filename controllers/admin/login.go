package admin

import (
	"encoding/json"
	"fmt"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"net/http"
	"shop/models"
)

type LoginController struct {
	BaseController
}

func (con LoginController) Login(c *gin.Context) {
	c.HTML(http.StatusOK, "admin/login/login.html", gin.H{})
}

func (con LoginController) DoLogin(c *gin.Context) {
	captchaId := c.PostForm("captchaId")
	verifyValue := c.PostForm("verifyValue")
	username := c.PostForm("username")
	password := c.PostForm("password")

	// 1 判断验证码是否正确
	if flag := models.VerifyCaptcha(captchaId, verifyValue); flag {
		// 2 判断用户名和密码是否正确
		managerinfo := []models.Manager{}
		password = models.Md5(password)
		models.DB.Where("username=? and password=?", username, password).Find(&managerinfo)
		if len(managerinfo) > 0 {
			// 3 执行登陆 保存用户信息 执行跳转
			session := sessions.Default(c)
			// 把结构体转换成json字符串
			managerinfoSlice, _ := json.Marshal(managerinfo)
			session.Set("managerinfo", string(managerinfoSlice))
			session.Save()
			con.Success(c, "登陆成功", "/admin/index")
		} else {
			con.Error(c, "用户名或密码错误", "/admin/login")
		}
		return
	}

	con.Error(c, "验证码错误", "/admin/login")
}

func (con LoginController) Captcha(c *gin.Context) {
	lid, lb64s, lerr := models.MakeCaptcha()
	if lerr != nil {
		fmt.Println(lerr)
	}
	c.JSON(http.StatusOK, gin.H{
		"captchaId":  lid,
		"captchaImg": lb64s,
	})
}

func (con LoginController) LoginOut(c *gin.Context) {
	sessions := sessions.Default(c)
	sessions.Delete("managerinfo")
	sessions.Save()
	con.Success(c, "退出登陆成功", "/admin/login")
}
