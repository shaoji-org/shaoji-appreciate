package service

import (
	"github.com/shaoji-org/shaoji-appreciate/model"
	"github.com/shaoji-org/shaoji-appreciate/serializer"
)

// UpdateShaoJiService 更新烧鸡的服务
type UpdateShaoJiService struct {
	Title string `form:"title" json:"title" binding:"required,min=2,max=30"`
	Info  string `form:"info" json:"info" binding:"max=300"`
}

// Update 更新烧鸡
func (service *UpdateShaoJiService) Update(id string) serializer.Response {
	var shaoji model.ShaoJi
	err := model.DB.First(&shaoji, id).Error
	if err != nil {
		return serializer.Response{
			Status: 404,
			Msg:    "烧鸡不存在",
			Error:  err.Error(),
		}
	}

	shaoji.Title = service.Title
	shaoji.Info = service.Info
	err = model.DB.Save(&shaoji).Error
	if err != nil {
		return serializer.Response{
			Status: 50002,
			Msg:    "烧鸡保存失败",
			Error:  err.Error(),
		}
	}

	return serializer.Response{
		Data: serializer.BuildShaoJi(shaoji),
	}
}
