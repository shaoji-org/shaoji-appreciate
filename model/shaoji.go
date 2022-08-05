package model

import (
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"github.com/jinzhu/gorm"
	"github.com/shaoji-org/shaoji-appreciate/cache"
	"os"
	"strconv"
)

type ShaoJi struct {
	gorm.Model
	Title string
	Info  string
	URL   string
}

// ShaoJiURL 烧鸡地址
func (s *ShaoJi) ShaoJiURL() string {
	client, _ := oss.New(os.Getenv("OSS_END_POINT"), os.Getenv("OSS_ACCESS_KEY_ID"), os.Getenv("OSS_ACCESS_KEY_SECRET"))
	bucket, _ := client.Bucket(os.Getenv("OSS_BUCKET"))
	signedGetURL, _ := bucket.SignURL(s.URL, oss.HTTPGet, 600)
	return signedGetURL
}

// View 点击数
func (s *ShaoJi) View() uint64 {
	countStr, _ := cache.RedisClient.Get(cache.ShaoJiViewKey(s.ID)).Result()
	count, _ := strconv.ParseUint(countStr, 10, 64)
	return count
}

// AddView 视频游览
func (s *ShaoJi) AddView() {
	// 增加视频点击数
	cache.RedisClient.Incr(cache.ShaoJiViewKey(s.ID))
	// 增加排行点击数
	cache.RedisClient.ZIncrBy(cache.DailyRankKey, 1, strconv.Itoa(int(s.ID)))
}
