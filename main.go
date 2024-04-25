package main

import (
	"github.com/gin-gonic/gin"
	"xuetang/init_config"
)

func main() {
	//初始化数据库
	go init_config.InitDb()
	//初始化Redis
	go init_config.InitRd()
	//初始化Nacos
	init_config.Initnacos()
	//注册服务到Nacos
	go Regis(init_config.NacosCli)

	// 创建一个默认的gin.Engine实例，用于构建HTTP路由。
	r := gin.Default()

	// 初始化路由，调用initRouter函数，用于设置API路由及其处理函数。
	InitRouter(r)

	//启动HTTP服务
	r.Run()
}
