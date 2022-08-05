package service

import (
	"github.com/shaoji-org/shaoji-appreciate/model"
	"github.com/shaoji-org/shaoji-appreciate/serializer"
)

// ListShaoJiService 烧鸡列表服务
type ListShaoJiService struct {
	Limit int `form:"limit"`
	Start int `form:"start"`
}

// List 烧鸡列表
func (service *ListShaoJiService) List() serializer.Response {
	var shaojis []model.ShaoJi
	total := 0

	if service.Limit == 0 {
		service.Limit = 6
	}

	if err := model.DB.Model(model.ShaoJi{}).Count(&total).Error; err != nil {
		return serializer.Response{
			Status: 50000,
			Msg:    "数据库连接错误",
			Error:  err.Error(),
		}
	}

	if err := model.DB.Limit(service.Limit).Offset(service.Start).Find(&shaojis).Error; err != nil {
		return serializer.Response{
			Status: 50000,
			Msg:    "数据库连接错误",
			Error:  err.Error(),
		}
	}

	return serializer.BuildListResponse(serializer.BuildShaoJis(shaojis), uint(total))
}
