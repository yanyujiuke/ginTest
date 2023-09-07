package admin

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"shop/models"
	"strings"
)

type RoleController struct {
	BaseController
}

func (con RoleController) Index(c *gin.Context) {
	roleList := []models.Role{}
	models.DB.Find(&roleList)
	c.HTML(http.StatusOK, "admin/role/index.html", gin.H{
		"roleList": roleList,
	})
}

func (con RoleController) Add(c *gin.Context) {
	c.HTML(http.StatusOK, "admin/role/add.html", gin.H{})
}

func (con RoleController) DoAdd(c *gin.Context) {
	title := strings.Trim(c.PostForm("title"), " ")
	desc := strings.Trim(c.PostForm("desc"), " ")

	if title == "" {
		con.Error(c, "角色标题不能为空", "admin/role/add")
		return
	}

	role := models.Role{
		Title:   title,
		Desc:    desc,
		Status:  1,
		AddTime: int(models.GetUnix()),
	}
	err := models.DB.Create(&role).Error
	if err != nil {
		con.Error(c, "添加角色失败，请重试", "/admin/role/add")
	} else {
		con.Success(c, "添加角色成功", "/admin/role")
	}
}

func (con RoleController) Edit(c *gin.Context) {
	id, err := models.Int(c.Query("id"))
	if err != nil {
		con.Error(c, "参数错误", "/admin/role")
		return
	}
	role := models.Role{Id: id}
	models.DB.First(&role)
	c.HTML(http.StatusOK, "admin/role/edit.html", gin.H{
		"role": role,
	})
}

func (con RoleController) DoEdit(c *gin.Context) {
	id, err := models.Int(c.PostForm("id"))
	if err != nil {
		con.Error(c, "参数错误", "/admin/role")
		return
	}
	title := strings.Trim(c.PostForm("title"), " ")
	desc := strings.Trim(c.PostForm("desc"), " ")

	role := models.Role{Id: id}
	models.DB.Find(&role)
	role.Title = title
	role.Desc = desc
	err = models.DB.Save(&role).Error
	if err != nil {
		con.Error(c, "修改角色失败，请重试", "/admin/role/edit?id="+models.String(id))
	} else {
		con.Success(c, "修改角色成功", "/admin/role")
	}
}

func (con RoleController) Delete(c *gin.Context) {
	id, err := models.Int(c.Query("id"))
	if err != nil {
		con.Error(c, "参数错误", "/admin/role")
		return
	}
	role := models.Role{Id: id}
	models.DB.Delete(&role)
	con.Success(c, "删除角色成功", "/admin/role")
}

func (con RoleController) Auth(c *gin.Context) {
	// 1 获取角色id
	id, err := models.Int(c.Query("id"))
	if err != nil {
		con.Error(c, "参数错误", "/admin/role")
		return
	}

	// 2 获取所有的权限列表
	accessList := []models.Access{}
	models.DB.Where("module_id=?", 0).Preload("AccessItem").Find(&accessList)

	// 3 获取当前角色所拥有的权限，并把权限id放在一个map对象里面
	roleAccessList := []models.RoleAccess{}
	models.DB.Where("role_id=?", id).Find(&roleAccessList)
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

	c.HTML(http.StatusOK, "admin/role/auth.html", gin.H{
		"roleId":     id,
		"accessList": accessList,
	})
}

func (con RoleController) DoAuth(c *gin.Context) {
	roleId, err := models.Int(c.PostForm("role_id"))
	if err != nil {
		con.Error(c, "参数错误", "/admin/role")
		return
	}

	accessIds := c.PostFormArray("access_node[]")
	fmt.Printf("%#v\n", accessIds)
	roleAccess := models.RoleAccess{}
	// 删除当前角色对应的权限
	models.DB.Where("role_id = ?", roleId).Delete(&roleAccess)

	// 添加当前对应的权限
	for _, v := range accessIds {
		accessId, _ := models.Int(v)

		roleAccess.RoleId = roleId
		roleAccess.AccessId = accessId
		models.DB.Create(&roleAccess)
	}

	con.Success(c, "角色授权成功", "/admin/role/auth?id="+models.String(roleId))
}
