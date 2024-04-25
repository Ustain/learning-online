package init_config

import (
	"github.com/qiniu/go-sdk/v7/auth"
	"github.com/qiniu/go-sdk/v7/storage"
)

var UpToken string

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
}
