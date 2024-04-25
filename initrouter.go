package main

import (
	"github.com/gin-gonic/gin"
	"xuetang/media_api"
)

// InitRouter 函数用于初始化路由及其对应的处理函数。
// 参数r是一个gin.Engine实例，用于构建路由。
func InitRouter(r *gin.Engine) {

	// 创建一个名为/media的路由组，用于将相关的API分组。
	apiRouter := r.Group("/media")

	apiRouter.POST("/files/", media_api.ListMediaFiles)               // 获取媒资数据的API
	apiRouter.POST("/upload/coursefile/", media_api.UploadMediaFiles) // 上传媒资数据的API
}
