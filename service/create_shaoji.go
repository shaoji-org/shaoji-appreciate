package service

import (
	"github.com/shaoji-org/shaoji-appreciate/model"
	"github.com/shaoji-org/shaoji-appreciate/serializer"
)

// CreateShaoJiService 烧鸡投稿的服务
type CreateShaoJiService struct {
	Title string `form:"title" json:"title" binding:"required,min=2,max=100"`
	Info  string `form:"info" json:"info" binding:"max=3000"`
	URL   string `form:"url" json:"url"`
}

// Create 创建烧鸡
func (service *CreateShaoJiService) Create() serializer.Response {
	shaoji := model.ShaoJi{
		Title: service.Title,
		Info:  service.Info,
		URL:   service.URL,
	}

	err := model.DB.Create(&shaoji).Error
	if err != nil {
		return serializer.Response{
			Status: 50001,
			Msg:    "烧鸡保存失败",
			Error:  err.Error(),
		}
	}

	return serializer.Response{
		Data: serializer.BuildShaoJi(shaoji),
	}
}
