package middlewares

import (
	"encoding/json"
	"fmt"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"gopkg.in/ini.v1"
	"net/http"
	"os"
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
		} else {
			// 登陆成功 判断权限
			// 超级管理员不用判断
			urlPath := strings.Replace(pathname, "/admin/", "", 1)
			if managerStruct[0].IsSuper == 0 && !excludeAuthPath("/"+urlPath) {
				// 1 获取当前角色所拥有的权限，并把权限id放在一个map对象里面
				roleAccessList := []models.RoleAccess{}
				models.DB.Where("role_id=?", managerStruct[0].RoleId).Find(&roleAccessList)
				roleAccessMap := make(map[int]int)
				for _, v := range roleAccessList {
					roleAccessMap[v.AccessId] = v.AccessId
				}

				// 2 获取当前访问的 url 对应的权限 id，判断用户是否有该权限
				// 获取 url 获取权限 id
				access := models.Access{}
				models.DB.Where("url=?", urlPath).First(&access)

				// 3 判断当前访问的 url 是否有权限
				if _, ok := roleAccessMap[access.Id]; !ok {
					c.String(http.StatusOK, "没有权限")
					c.Abort() // 终止访问
				}
			}
		}
	} else {
		// 用户没有登陆
		if pathname != "/admin/login" && pathname != "/admin/doLogin" && pathname != "/admin/captcha" {
			c.Redirect(302, "/admin/login")
		}
	}
}

// 排除权限判断方法
func excludeAuthPath(urlPath string) bool {

	config, iniErr := ini.Load("./conf/app.ini")
	if iniErr != nil {
		fmt.Printf("Fail to read file: %v", iniErr)
		os.Exit(1)
	}

	excludeAuthPath := config.Section("").Key("excludeAuthPath").String()
	excludeAuthPathSlice := strings.Split(excludeAuthPath, ",")

	for _, v := range excludeAuthPathSlice {
		if v == urlPath {
			return true
		}
	}
	return false
}
