package admin

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"shop/models"
	"strings"
)

type ManagerController struct {
	BaseController
}

func (con ManagerController) Index(c *gin.Context) {
	managerList := []models.Manager{}
	models.DB.Preload("Role").Find(&managerList)
	// fmt.Printf("%#v", managerList)

	c.HTML(http.StatusOK, "admin/manager/index.html", gin.H{
		"managerList": managerList,
	})
}

func (con ManagerController) Search(c *gin.Context) {
	username := strings.Trim(c.PostForm("username"), "")

	managerList := []models.Manager{}
	// if username == "" {
	// 	c.HTML(http.StatusOK, "admin/manager/index.html", gin.H{
	// 		"managerList": managerList,
	// 	})
	// 	return
	// }
	models.DB.Preload("Role").Where("username like ?", "%"+username+"%").Find(&managerList)
	// fmt.Printf("%#v", managerList)
	c.HTML(http.StatusOK, "admin/manager/index.html", gin.H{
		"managerList": managerList,
	})
}

func (con ManagerController) Add(c *gin.Context) {
	// 获取所有的角色
	roleList := []models.Role{}
	models.DB.Find(&roleList)

	c.HTML(http.StatusOK, "admin/manager/add.html", gin.H{
		"roleList": roleList,
	})
}

func (con ManagerController) DoAdd(c *gin.Context) {

	username := strings.Trim(c.PostForm("username"), " ")
	password := strings.Trim(c.PostForm("password"), " ")
	mobile := strings.Trim(c.PostForm("mobile"), " ")
	email := strings.Trim(c.PostForm("email"), " ")
	roleId, err := models.Int(c.PostForm("role_id"))
	if err != nil {
		con.Error(c, "参数错误", "/admin/manager/add")
		return
	}

	if len(username) < 2 || len(password) < 6 {
		con.Error(c, "用户名长度不小于2，密码长度不小于6", "/admin/manager/add")
		return
	}

	// 判断管理员是否存在
	managerList := []models.Manager{}
	models.DB.Where("username=?", username).Find(&managerList)
	if len(managerList) > 0 {
		con.Error(c, "此用户名已存在", "/admin/manager/add")
		return
	}

	manager := models.Manager{
		Username: username,
		Password: models.Md5(password),
		Mobile:   mobile,
		Email:    email,
		Status:   1,
		RoleId:   roleId,
		AddTime:  int(models.GetUnix()),
		IsSuper:  1,
	}

	err = models.DB.Create(&manager).Error
	if err != nil {
		con.Error(c, "添加管理员失败，请重试", "/admin/manager/add")
	} else {
		con.Success(c, "添加管理员成功", "/admin/manager")
	}
}

func (con ManagerController) Edit(c *gin.Context) {
	id, err := models.Int(c.Query("id"))
	if err != nil {
		con.Error(c, "参数错误", "/admin/manager")
		return
	}
	managerinfo := models.Manager{Id: id}
	models.DB.First(&managerinfo)

	// 获取所有的角色
	roleList := []models.Role{}
	models.DB.Find(&roleList)

	c.HTML(http.StatusOK, "admin/manager/edit.html", gin.H{
		"managerinfo": managerinfo,
		"roleList":    roleList,
	})
}

func (con ManagerController) DoEdit(c *gin.Context) {
	id, err1 := models.Int(c.PostForm("id"))
	if err1 != nil {
		con.Error(c, "参数错误", "/admin/role")
		return
	}

	roleId, err2 := models.Int(c.PostForm("role_id"))
	if err2 != nil {
		con.Error(c, "参数错误", "/admin/manager/add")
		return
	}

	username := strings.Trim(c.PostForm("username"), " ")
	password := strings.Trim(c.PostForm("password"), " ")
	mobile := strings.Trim(c.PostForm("mobile"), " ")
	email := strings.Trim(c.PostForm("email"), " ")

	managerinfo := &models.Manager{Id: id}
	models.DB.First(&managerinfo)
	managerinfo.Username = username
	managerinfo.Mobile = mobile
	managerinfo.Email = email
	managerinfo.RoleId = roleId

	// 判断密码是否为空
	if password != "" {
		// 判断密码长度6
		if len(password) < 6 {
			con.Error(c, "密码长度不小于6", "/admin/manager/edit")
			return
		}
		managerinfo.Password = models.Md5(password)
	}

	err3 := models.DB.Save(&managerinfo).Error
	if err3 != nil {
		con.Error(c, "修改数据失败", "/admin/manager")
		return
	}
	con.Success(c, "数据修改成功", "/admin/manager")
}

func (con ManagerController) Delete(c *gin.Context) {
	id, err1 := models.Int(c.Query("id"))
	if err1 != nil {
		con.Error(c, "参数错误", "/admin/role")
		return
	}
	managerinfo := models.Manager{Id: id}
	err := models.DB.Delete(&managerinfo).Error
	if err != nil {
		con.Error(c, "管理员删除失败", "/admin/manager")
		return
	}
	con.Success(c, "管理员删除成功", "/admin/manager")
}
