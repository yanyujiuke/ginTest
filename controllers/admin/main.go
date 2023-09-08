package admin

import (
	"encoding/json"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	"shop/models"
)

type MainController struct {
	BaseController
}

func (con MainController) Index(c *gin.Context) {
	// 获取用户登陆的session信息
	session := sessions.Default(c)
	managerinfo := session.Get("managerinfo")
	// 类型断言 判断 managerinfo 是不是一个 string
	managerStr, ok := managerinfo.(string)
	if ok {
		// 1 获取用户信息
		var managerinfoStruct []models.Manager
		json.Unmarshal([]byte(managerStr), &managerinfoStruct)

		// 2 获取所有的权限
		accessList := []models.Access{}
		models.DB.Where("module_id=?", 0).Preload("AccessItem", func(db *gorm.DB) *gorm.DB {
			return db.Order("access.sort desc")
		}).Order("sort desc").Find(&accessList)

		// 3 获取当前角色所拥有的权限，并把权限id放在一个map对象里面
		roleAccessList := []models.RoleAccess{}
		models.DB.Where("role_id=?", managerinfoStruct[0].RoleId).Find(&roleAccessList)
		roleAccessMap := make(map[int]int)
		for _, v := range roleAccessList {
			roleAccessMap[v.AccessId] = v.AccessId
		}

		// 4 遍历所有的权限数据，判断当前权限的id是否在角色权限的map对象中，如果在，则给当前数据加入checked属性
		for i := 0; i < len(accessList); i++ {
			if _, ok := roleAccessMap[accessList[i].Id]; ok {
				accessList[i].Checked = true
			}
			for j := 0; j < len(accessList[i].AccessItem); j++ {
				if _, ok := roleAccessMap[accessList[i].AccessItem[j].Id]; ok {
					accessList[i].AccessItem[j].Checked = true
				}
			}
		}

		c.HTML(http.StatusOK, "admin/main/index.html", gin.H{
			"username":   managerinfoStruct[0].Username,
			"accessList": accessList,
			"isSuper":    managerinfoStruct[0].IsSuper,
		})
	} else {
		c.Redirect(302, "admin/login")
	}
}

func (con MainController) Welcome(c *gin.Context) {
	c.HTML(http.StatusOK, "admin/main/welcome.html", gin.H{})
}

// 公共修改状态的方法
func (con MainController) ChangeStatus(c *gin.Context) {
	id, err := models.Int(c.Query("id"))
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"success": true,
			"message": "传入的参数错误",
		})
		return
	}

	table := c.Query("table")
	field := c.Query("field")

	err = models.DB.Exec("update "+table+" set "+field+" = ABS("+field+" - 1) where id = ?", id).Error
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"success": true,
			"message": "修改失败",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "修改成功",
	})
}
