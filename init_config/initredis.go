package init_config

import (
	"github.com/go-redis/redis"
)

var Rd *redis.Client

func InitRd() {
	Rd = redis.NewClient(&redis.Options{
		Addr:     "127.0.0.1:6379", // url
		Password: "",
		DB:       0, // 0号数据库
	})
}
