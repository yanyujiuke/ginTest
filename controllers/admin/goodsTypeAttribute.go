package admin

import (
	"net/http"
	"shop/models"
	"strings"

	"github.com/gin-gonic/gin"
)

type GoodsTypeAttributeController struct {
	BaseController
}

func (con GoodsTypeAttributeController) Index(c *gin.Context) {

	typeId, err := models.Int(c.Query("type_id"))
	if err != nil {
		con.Error(c, "传入的参数不正确", "/admin/goodsType")
		return
	}
	// 获取商品类型属性
	goodsTypeAttributeList := []models.GoodsTypeAttribute{}
	models.DB.Where("type_id=?", typeId).Find(&goodsTypeAttributeList)

	// 获取商品类型属性对应的类型
	goodsType := models.GoodsType{}
	models.DB.Where("id=?", typeId).Find(&goodsType)

	c.HTML(http.StatusOK, "admin/goodsTypeAttribute/index.html", gin.H{
		"typeId":                 typeId,
		"goodsTypeAttributeList": goodsTypeAttributeList,
		"goodsType":              goodsType,
	})

}
func (con GoodsTypeAttributeController) Add(c *gin.Context) {
	// 获取当前商品类型属性对应的类型id
	typeId, err := models.Int(c.Query("type_id"))
	if err != nil {
		con.Error(c, "传入的参数不正确", "/admin/goodsType")
		return
	}

	// 获取所有的商品类型
	goodsTypeList := []models.GoodsType{}
	models.DB.Find(&goodsTypeList)
	c.HTML(http.StatusOK, "admin/goodsTypeAttribute/add.html", gin.H{
		"goodsTypeList": goodsTypeList,
		"typeId":        typeId,
	})
}

func (con GoodsTypeAttributeController) DoAdd(c *gin.Context) {

	title := strings.Trim(c.PostForm("title"), " ")
	typeId, err1 := models.Int(c.PostForm("type_id"))
	attrType, err2 := models.Int(c.PostForm("attr_type"))
	attrValue := c.PostForm("attr_value")
	sort, err3 := models.Int(c.PostForm("sort"))

	if err1 != nil || err2 != nil {
		con.Error(c, "非法请求", "/admin/goodsType")
		return
	}
	if title == "" {
		con.Error(c, "商品类型属性名称不能为空", "/admin/goodsTypeAttribute/add?type_id="+models.String(typeId))
		return
	}

	if err3 != nil {
		con.Error(c, "排序值不对", "/admin/goodsTypeAttribute/add?type_id="+models.String(typeId))
		return
	}

	goodsTypeAttr := models.GoodsTypeAttribute{
		Title:     title,
		TypeId:    typeId,
		AttrType:  attrType,
		AttrValue: attrValue,
		Status:    1,
		Sort:      sort,
		AddTime:   int(models.GetUnix()),
	}
	err := models.DB.Create(&goodsTypeAttr).Error
	if err != nil {
		con.Error(c, "增加商品类型属性失败 请重试", "/admin/goodsTypeAttribute/add?type_id="+models.String(typeId))
	} else {
		con.Success(c, "增加商品类型属性成功", "/admin/goodsTypeAttribute?type_id="+models.String(typeId))
	}
}

func (con GoodsTypeAttributeController) Edit(c *gin.Context) {
	id, err1 := models.Int(c.Query("id"))
	if err1 != nil {
		con.Error(c, "参数错误", "/admin/goodsTypeAttribute")
		return
	}

	goodsTypeAttribute := models.GoodsTypeAttribute{Id: id}
	models.DB.Find(&goodsTypeAttribute)

	// 获取所有的商品类型
	goodsTypeList := []models.GoodsType{}
	models.DB.Find(&goodsTypeList)

	c.HTML(http.StatusOK, "admin/goodsTypeAttribute/edit.html", gin.H{
		"goodsTypeAttribute": goodsTypeAttribute,
		"goodsTypeList":      goodsTypeList,
	})
}

func (con GoodsTypeAttributeController) DoEdit(c *gin.Context) {
	id, err1 := models.Int(c.PostForm("id"))
	if err1 != nil {
		con.Error(c, "参数错误", "/admin/goodsTypeAttribute")
		return
	}

	title := strings.Trim(c.PostForm("title"), " ")
	typeId, err1 := models.Int(c.PostForm("type_id"))
	attrType, err2 := models.Int(c.PostForm("attr_type"))
	attrValue := c.PostForm("attr_value")
	sort, err3 := models.Int(c.PostForm("sort"))

	if err1 != nil || err2 != nil {
		con.Error(c, "非法请求", "/admin/goodsTypeAttribute")
		return
	}
	if title == "" {
		con.Error(c, "商品类型属性名称不能为空", "/admin/goodsTypeAttribute/edit?id="+models.String(id))
		return
	}

	if err3 != nil {
		con.Error(c, "排序值不对", "/admin/goodsTypeAttribute/edit?id="+models.String(id))
		return
	}

	goodsTypeAttribute := models.GoodsTypeAttribute{Id: id}
	models.DB.Find(&goodsTypeAttribute)
	goodsTypeAttribute.Title = title
	goodsTypeAttribute.TypeId = typeId
	goodsTypeAttribute.AttrType = attrType
	goodsTypeAttribute.AttrValue = attrValue
	goodsTypeAttribute.Sort = sort
	err4 := models.DB.Save(&goodsTypeAttribute).Error
	if err4 != nil {
		con.Error(c, "编辑商品类型属性失败 请重试", "/admin/goodsTypeAttribute?type_id="+models.String(typeId))
	} else {
		con.Success(c, "编辑商品类型属性成功", "/admin/goodsTypeAttribute?type_id="+models.String(typeId))
	}
}

func (con GoodsTypeAttributeController) Delete(c *gin.Context) {
	id, err := models.Int(c.Query("id"))
	if err != nil {
		con.Error(c, "参数错误", "/admin/goodsTypeAttribute")
		return
	}

	goodsTypeAttribute := models.GoodsTypeAttribute{Id: id}
	models.DB.Find(&goodsTypeAttribute)
	err = models.DB.Delete(&goodsTypeAttribute).Error
	if err != nil {
		con.Error(c, "编辑商品类型属性失败 请重试", "/admin/goodsTypeAttribute?type_id="+models.String(goodsTypeAttribute.TypeId))
	} else {
		con.Success(c, "编辑商品类型属性成功", "/admin/goodsTypeAttribute?type_id="+models.String(goodsTypeAttribute.TypeId))
	}
}
