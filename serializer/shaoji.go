package serializer

import "github.com/shaoji-org/shaoji-appreciate/model"

// ShaoJi 烧鸡序列化器
type ShaoJi struct {
	ID        uint   `json:"id"`
	Title     string `json:"title"`
	Info      string `json:"info"`
	URL       string `json:"url"`
	View      uint64 `json:"view"`
	CreatedAt int64  `json:"created_at"`
}

// BuildShaoJi 序列化烧鸡
func BuildShaoJi(item model.ShaoJi) ShaoJi {
	return ShaoJi{
		ID:        item.ID,
		Title:     item.Title,
		Info:      item.Info,
		URL:       item.ShaoJiURL(),
		View:      item.View(),
		CreatedAt: item.CreatedAt.Unix(),
	}
}

// BuildShaoJis 序列化烧鸡列表
func BuildShaoJis(items []model.ShaoJi) (shaojis []ShaoJi) {
	for _, item := range items {
		shaoji := BuildShaoJi(item)
		shaojis = append(shaojis, shaoji)
	}
	return shaojis
}
