package service

import (
	"github.com/shaoji-org/shaoji-appreciate/model"
	"github.com/shaoji-org/shaoji-appreciate/serializer"
)

// DeleteShaoJiService 删除投稿的服务
type DeleteShaoJiService struct {
}

// Delete 删除烧鸡
func (service *DeleteShaoJiService) Delete(id string) serializer.Response {
	var shaoji model.ShaoJi
	err := model.DB.First(&shaoji, id).Error
	if err != nil {
		return serializer.Response{
			Status: 404,
			Msg:    "烧鸡不存在",
			Error:  err.Error(),
		}
	}

	err = model.DB.Delete(&shaoji).Error
	if err != nil {
		return serializer.Response{
			Status: 50000,
			Msg:    "烧鸡删除失败",
			Error:  err.Error(),
		}
	}

	return serializer.Response{}
}
