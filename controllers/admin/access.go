package admin

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"shop/models"
	"strings"
)

type AccessController struct {
	BaseController
}

func (con AccessController) Index(c *gin.Context) {
	accessList := []models.Access{}
	models.DB.Where("module_id = ?", 0).Preload("AccessItem").Find(&accessList)
	// fmt.Printf("%#v", accessList)
	c.HTML(http.StatusOK, "admin/access/index.html", gin.H{
		"accessList": accessList,
	})
}

func (con AccessController) Add(c *gin.Context) {
	accessList := []models.Access{}
	models.DB.Where("module_id = ?", 0).Find(&accessList)

	c.HTML(http.StatusOK, "admin/access/add.html", gin.H{
		"accessList": accessList,
	})
}

func (con AccessController) DoAdd(c *gin.Context) {

	moduleName := strings.Trim(c.PostForm("module_name"), " ")
	accessType, err1 := models.Int(c.PostForm("type"))
	actionName := strings.Trim(c.PostForm("action_name"), " ")
	url := strings.Trim(c.PostForm("url"), " ")
	moduleId, err2 := models.Int(c.PostForm("module_id"))
	sort, err3 := models.Int(c.PostForm("sort"))
	desc := strings.Trim(c.PostForm("desc"), " ")
	status, err4 := models.Int(c.PostForm("status"))

	if err1 != nil || err2 != nil || err3 != nil || err4 != nil {
		con.Error(c, "参数错误", "admin/access/add")
		return
	}

	if moduleName == "" {
		con.Error(c, "模块名称不能为空", "admin/access/add")
		return
	}

	access := models.Access{
		ModuleName: moduleName,
		Type:       accessType,
		ActionName: actionName,
		Url:        url,
		ModuleId:   moduleId,
		Sort:       sort,
		Desc:       desc,
		Status:     status,
		AddTime:    int(models.GetUnix()),
	}
	err5 := models.DB.Create(&access).Error
	if err5 != nil {
		con.Error(c, "添加权限菜单失败，请重试", "/admin/access/add")
	} else {
		con.Success(c, "添加权限菜单成功", "/admin/access")
	}
}

func (con AccessController) Edit(c *gin.Context) {
	id, err := models.Int(c.Query("id"))
	if err != nil {
		con.Error(c, "参数错误", "/admin/access")
		return
	}

	accessList := []models.Access{}
	models.DB.Where("module_id = ?", 0).Find(&accessList)

	access := models.Access{Id: id}
	models.DB.Find(&access)

	c.HTML(http.StatusOK, "admin/access/edit.html", gin.H{
		"accessList": accessList,
		"access":     access,
	})
}

func (con AccessController) DoEdit(c *gin.Context) {

	id, err := models.Int(c.PostForm("id"))
	if err != nil {
		con.Error(c, "参数错误", "/admin/access")
		return
	}

	moduleName := strings.Trim(c.PostForm("module_name"), " ")
	accessType, err1 := models.Int(c.PostForm("type"))
	actionName := strings.Trim(c.PostForm("action_name"), " ")
	url := strings.Trim(c.PostForm("url"), " ")
	moduleId, err2 := models.Int(c.PostForm("module_id"))
	sort, err3 := models.Int(c.PostForm("sort"))
	desc := strings.Trim(c.PostForm("desc"), " ")
	status, err4 := models.Int(c.PostForm("status"))

	if err != nil || err1 != nil || err2 != nil || err3 != nil || err4 != nil {
		con.Error(c, "参数错误", "admin/access/edit?id="+models.String(id))
		return
	}

	if moduleName == "" {
		con.Error(c, "模块名称不能为空", "admin/access/edit?id="+models.String(id))
		return
	}

	access := models.Access{Id: id}
	models.DB.Find(&access)
	access.ModuleName = moduleName
	access.Type = accessType
	access.ActionName = actionName
	access.Url = url
	access.ModuleId = moduleId
	access.Sort = sort
	access.Desc = desc
	access.Status = status

	fmt.Printf("%#v", access)
	err5 := models.DB.Save(&access).Error
	if err5 != nil {
		con.Error(c, "编辑权限菜单失败，请重试", "/admin/access")
	} else {
		con.Success(c, "编辑权限菜单成功", "/admin/access")
	}
}

func (con AccessController) Delete(c *gin.Context) {
	id, err := models.Int(c.Query("id"))
	if err != nil {
		con.Error(c, "参数错误", "/admin/access")
		return
	}
	access := models.Access{Id: id}
	models.DB.First(&access)
	// 判断顶级模块是否有子菜单
	if access.ModuleId == 0 {
		chadrenList := []models.Access{}
		models.DB.Where("module_id = ?", access.Id).Find(&chadrenList)
		if len(chadrenList) > 0 {
			con.Error(c, "该模块有子模块，不能删除", "/admin/access")
			return
		}
	}
	models.DB.Delete(&access)
	con.Success(c, "删除角色成功", "/admin/access")
}
