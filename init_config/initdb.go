package init_config

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
)

var Db *gorm.DB

func InitDb() {
	dsn := "root:123456@tcp(localhost:3307)/xcplus_media?charset=utf8mb4&parseTime=True&loc=Local" // 替换为你的数据库信息
	var err error
	Db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		// 启用日志输出，输出包含SQL查询语句
		Logger:                 logger.Default.LogMode(logger.Info),
		SkipDefaultTransaction: true, //关闭默认事务
		PrepareStmt:            true, //缓存预编译语句
	})
	if err != nil {
		log.Fatal(err)
	}

	// 迁移数据库，确保User表存在
	Db.AutoMigrate()
}
