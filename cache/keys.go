package cache

import (
	"fmt"
	"strconv"
)

const (
	// DailyRankKey 每日排行
	DailyRankKey = "rank:daily"
)

// ShaoJiViewKey 烧鸡点击数的key
// view:shaoji:1 -> 100
// view:shaoji:2 -> 150
func ShaoJiViewKey(id uint) string {
	return fmt.Sprintf("view:shaoji:%s", strconv.Itoa(int(id)))
}
