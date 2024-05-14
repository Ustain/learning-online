package init_config

import (
	"encoding/json"
	"fmt"
	"github.com/go-redis/redis"
	"gorm.io/gorm"
	"log"
	"xuetang/kitex_gen/xuetang"
)

func Preheat() {
	// 从数据库中获取所有媒资文件数据
	mediaFiles, err := getMediaFiles(Db)
	if err != nil {
		log.Fatal("Error getting media files:", err)
	}

	// 遍历媒资文件数据，并将其存储到Redis中
	for _, mediaFile := range mediaFiles {
		err := setMediaFileToRedis(Rd, mediaFile)
		if err != nil {
			log.Fatal("Error setting media file to Redis:", err)
		}
	}

	fmt.Println("Media files cached to Redis successfully.")
}

// 从数据库中批量获取媒资文件数据
func getMediaFiles(db *gorm.DB) ([]xuetang.MediaFiles, error) {
	var mediaFiles []xuetang.MediaFiles
	if err := db.Limit(10000).Find(&mediaFiles).Error; err != nil {
		return nil, err
	}
	return mediaFiles, nil
}

// 将媒资文件数据存储到Redis中
func setMediaFileToRedis(rdb *redis.Client, mediaFile xuetang.MediaFiles) error {
	mediaFileJSON, err := json.Marshal(mediaFile)
	if err != nil {
		return err
	}

	// 使用SET命令将媒资文件数据存储到Redis中
	err = rdb.Set("media_files:"+mediaFile.Id, mediaFileJSON, 0).Err()
	if err != nil {
		return err
	}

	return nil
}
