package main

import (
	"github.com/gin-gonic/gin"
	"xuetang/media_api"
)

// InitRouter 函数用于初始化路由及其对应的处理函数。
// 参数r是一个gin.Engine实例，用于构建路由。
func InitRouter(r *gin.Engine) {

	// 创建一个名为/media_api/ningmeng的路由组，用于将相关的API分组。
	apiRouter := r.Group("/media")

	apiRouter.POST("/files/", media_api.ListMediaFiles)               // 分页获取媒资文件的API
	apiRouter.POST("/upload/coursefile/", media_api.UploadMediaFiles) // 上传普通文件的的API
	apiRouter.GET("/open/preview/", media_api.GetPlayUrlByMediaId)    //预览媒资文件API

	//大文件分片接口——视频
	apiRouter.POST("/upload/checkfile/", media_api.CheckFile)         //大文件存在判断API
	apiRouter.POST("/upload/bigfile/", media_api.UploadBigFile)       //大文件上传API
	apiRouter.POST("/upload/process/", media_api.UploadProcessResult) //获取大文件上传进度API
}
