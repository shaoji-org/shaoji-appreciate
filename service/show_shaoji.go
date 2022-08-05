package service

import (
	"github.com/shaoji-org/shaoji-appreciate/model"
	"github.com/shaoji-org/shaoji-appreciate/serializer"
)

// ShowShaoJiService 投稿详情的服务
type ShowShaoJiService struct {
}

// Show 烧鸡
func (service *ShowShaoJiService) Show(id string) serializer.Response {
	var shaoji model.ShaoJi
	err := model.DB.First(&shaoji, id).Error
	if err != nil {
		return serializer.Response{
			Status: 404,
			Msg:    "烧鸡不存在",
			Error:  err.Error(),
		}
	}

	//处理烧鸡被观看的一系问题
	shaoji.AddView()

	return serializer.Response{
		Data: serializer.BuildShaoJi(shaoji),
	}
}
