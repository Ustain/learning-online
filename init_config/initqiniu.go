package init_config

import (
	"github.com/qiniu/go-sdk/v7/auth"
	"github.com/qiniu/go-sdk/v7/storage"
)

var UpToken string

var BucketManager *storage.BucketManager

func Initqiniu() {
	accessKey := "w7JGmTXjwTgYXgeEZQ10svDrfRh9s3HSRDGMO6Pd"
	secretKey := "xcij2mdAgrux_NvFRkF0oZH-RUdP8tdd4vvv87t3"

	mac := auth.New(accessKey, secretKey)
	bucket := "xuetangmedia"
	putPolicy := storage.PutPolicy{
		Scope: bucket,
	}
	putPolicy.Expires = 14400 //示例4小时有效期
	UpToken = putPolicy.UploadToken(mac)

	cfg := storage.Config{
		// 是否使用https域名进行资源管理
		UseHTTPS: false,
	}
	// 指定空间所在的区域，如果不指定将自动探测
	// 如果没有特殊需求，默认不需要指定
	//cfg.Region=&storage.ZoneHuabei
	BucketManager = storage.NewBucketManager(mac, &cfg)
}
