package admin

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"shop/models"
	"strings"
)

type FocusController struct {
	BaseController
}

func (con FocusController) Index(c *gin.Context) {

	focusList := []models.Focus{}
	models.DB.Find(&focusList)

	c.HTML(http.StatusOK, "admin/focus/index.html", gin.H{
		"focusList": focusList,
	})
}

func (con FocusController) Search(c *gin.Context) {
	title := strings.Trim(c.PostForm("title"), "")

	focusList := []models.Focus{}
	models.DB.Where("title LIKE ?", "%"+title+"%").Find(&focusList)

	c.HTML(http.StatusOK, "admin/focus/index.html", gin.H{
		"focusList": focusList,
	})
}

func (con FocusController) Add(c *gin.Context) {
	c.HTML(http.StatusOK, "admin/focus/add.html", gin.H{})
}

func (con FocusController) DoAdd(c *gin.Context) {

	title := strings.Trim(c.PostForm("title"), " ")
	focusType, err1 := models.Int(c.PostForm("focus_type"))
	link := strings.Trim(c.PostForm("link"), " ")
	status, err3 := models.Int(c.PostForm("status"))
	if err1 != nil || err3 != nil {
		con.Error(c, "参数错误", "/admin/focus")
		return
	}

	// 上传图片
	imgSrc, err4 := models.UploadImg(c, "focus_img")
	if err4 != nil {
		fmt.Println(err4)
	}

	focus := &models.Focus{
		Title:     title,
		FocusType: focusType,
		FocusImg:  imgSrc,
		Link:      link,
		Sort:      100,
		Status:    status,
		AddTime:   int(models.GetUnix()),
	}

	err := models.DB.Create(&focus).Error
	if err != nil {
		con.Error(c, "轮播图添加失败", "/admin/focus/add")
	} else {
		con.Success(c, "轮播图添加成功", "/admin/focus")
	}
}

func (con FocusController) Edit(c *gin.Context) {
	id, err := models.Int(c.Query("id"))
	if err != nil {
		con.Error(c, "参数错误", "/admin/focus/edit")
		return
	}

	// 获取列表
	focus := models.Focus{Id: id}
	models.DB.First(&focus)

	c.HTML(http.StatusOK, "admin/focus/edit.html", gin.H{
		"focus": focus,
	})
}

func (con FocusController) DoEdit(c *gin.Context) {
	id, err1 := models.Int(c.PostForm("id"))
	if err1 != nil {
		con.Error(c, "参数 id 错误", "/admin/focus")
		return
	}

	title := strings.Trim(c.PostForm("title"), " ")
	focusType, err2 := models.Int(c.PostForm("focus_type"))
	if err2 != nil {
		con.Error(c, "参数 focusType 错误", "/admin/focus")
		return
	}
	link := strings.Trim(c.PostForm("link"), " ")
	status, err3 := models.Int(c.PostForm("status"))
	if err3 != nil {
		con.Error(c, "参数 status 错误", "/admin/focus")
		return
	}

	focus := &models.Focus{Id: id}
	models.DB.First(&focus)
	focus.Title = title
	focus.FocusType = focusType
	focus.Link = link
	focus.Status = status

	// 上传图片
	imgSrc, err4 := models.UploadImg(c, "focus_img")
	if err4 == nil {
		focus.FocusImg = imgSrc
	}

	err5 := models.DB.Save(&focus).Error
	if err5 != nil {
		con.Error(c, "修改数据失败", "/admin/focus")
		return
	}
	con.Success(c, "数据修改成功", "/admin/focus")
}

func (con FocusController) Delete(c *gin.Context) {

	id, err1 := models.Int(c.Query("id"))
	if err1 != nil {
		con.Error(c, "参数错误", "/admin/focus")
		return
	}

	focus := models.Focus{Id: id}
	// models.DB.First(&focus)
	// // 是否删除图片
	// os.Remove(focus.FocusImg)
	err := models.DB.Delete(&focus).Error
	if err != nil {
		con.Error(c, "删除失败", "/admin/focus")
		return
	}
	con.Success(c, "删除成功", "/admin/focus")
}
