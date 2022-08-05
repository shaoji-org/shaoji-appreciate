package api

import (
	"github.com/gin-gonic/gin"
	"github.com/shaoji-org/shaoji-appreciate/service"
)

// CreateShaoJi 烧鸡投稿
func CreateShaoJi(c *gin.Context) {
	service := service.CreateShaoJiService{}
	if err := c.ShouldBind(&service); err == nil {
		res := service.Create()
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}

// ShowShaoJi 烧鸡详情接口
func ShowShaoJi(c *gin.Context) {
	service := service.ShowShaoJiService{}
	res := service.Show(c.Param("id"))
	c.JSON(200, res)
}

// ListShaoJi 烧鸡列表接口
func ListShaoJi(c *gin.Context) {
	service := service.ListShaoJiService{}
	if err := c.ShouldBind(&service); err == nil {
		res := service.List()
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}

// UpdateShaoJi 更新烧鸡的接口
func UpdateShaoJi(c *gin.Context) {
	service := service.UpdateShaoJiService{}
	if err := c.ShouldBind(&service); err == nil {
		res := service.Update(c.Param("id"))
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}

// DeleteShaoJi 删除烧鸡的接口
func DeleteShaoJi(c *gin.Context) {
	service := service.DeleteShaoJiService{}
	res := service.Delete(c.Param("id"))
	c.JSON(200, res)
}
