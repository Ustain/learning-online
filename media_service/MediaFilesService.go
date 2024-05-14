package media_service

import (
	"context"
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"github.com/go-redis/redis"
	"github.com/qiniu/go-sdk/v7/storage"
	"io"
	"os"
	"time"
	"xuetang/init_config"
	"xuetang/kitex_gen/xuetang"
)

// QueryMediaFiles 分页查询媒资文件
func QueryMediaFiles(companyId int, pageParams xuetang.PageParams) (xuetang.PageResult_, error) {
	var mediaFiles []*xuetang.MediaFiles
	var total int64

	// 构建Redis键名
	keyp := "media_files:*"

	// 从Redis中获取缓存数据
	keys, err := init_config.Rd.Keys(keyp).Result()
	if err != nil {
		fmt.Println("redis错误", err)
		return xuetang.PageResult_{}, err
	}

	// 获取总记录数
	total = int64(len(keys))

	// 解码JSON数据
	for i, key := range keys {
		if i < int(pageParams.PageSize) {
			// 从Redis中获取对应键的值
			value, err := init_config.Rd.Get(key).Result()
			if err != nil {
				fmt.Printf("获取键值失败：%s\n", err)
				return xuetang.PageResult_{}, err
			}

			// 解码JSON
			var mf xuetang.MediaFiles
			err = json.Unmarshal([]byte(value), &mf)
			if err != nil {
				fmt.Printf("解码JSON失败：%s\n", err)
				return xuetang.PageResult_{}, err
			}

			// 将解码后的数据添加到mediaFiles数组中
			mediaFiles = append(mediaFiles, &mf)
		}
	}

	//数据不够再查数据库
	if len(mediaFiles) < int(pageParams.PageSize) {
		fmt.Println("查数据库了！！")
		// 构建查询条件
		result := init_config.Db.Where("company_id = ?", companyId).Find(&mediaFiles)
		if result.Error != nil {
			return xuetang.PageResult_{}, nil
		}
		// 获取总记录数
		if err := result.Model(&xuetang.MediaFiles{}).Count(&total).Error; err != nil {
			return xuetang.PageResult_{}, err
		}

		// 分页查询
		if err := result.Offset(int((pageParams.PageNo - 1) * pageParams.PageSize)).Limit(int(pageParams.PageSize)).Find(&mediaFiles).Error; err != nil {
			return xuetang.PageResult_{}, err
		}
	}

	// 分页处理
	startIndex := (pageParams.PageNo - 1) * pageParams.PageSize
	endIndex := startIndex + pageParams.PageSize
	if endIndex > total {
		endIndex = total
	}
	mediaFiles = mediaFiles[startIndex:endIndex]

	var pageResult *xuetang.PageResult_
	// 构建分页结果
	pageResult = xuetang.NewPageResult_()
	pageResult.SetItems(mediaFiles)
	pageResult.SetCounts(total)
	pageResult.SetPage(pageParams.PageNo)
	pageResult.SetPageSize(pageParams.PageSize)

	return *pageResult, nil
}

// 存储媒资数据到数据库
func saveMediaToDb(companyId int, fileMd5 string, uploadFileParam xuetang.UploadFileParamsDto, bucket string, objectName string) (xuetang.MediaFiles, error) {
	var media xuetang.MediaFiles
	res := init_config.Db.Where("id=?", fileMd5).Find(&media)
	fmt.Println(media.Id)

	if media.Id == "" {
		fmt.Println("xie")
		mediaFiles := xuetang.NewMediaFiles()
		//文件id
		mediaFiles.SetId(fileMd5)
		//机构id
		mediaFiles.SetCompanyId(int64(companyId))
		//桶
		mediaFiles.SetBucket(bucket)
		//file_path
		mediaFiles.SetFilePath(objectName)
		//file_id
		mediaFiles.SetFileId(fileMd5)
		//url
		mediaFiles.SetUrl("http://sce0dlgy0.hn-bkt.clouddn.com/" + objectName)
		//上传时间
		mediaFiles.SetCreateDate(time.Now().Format("2006-01-02 15:04:05"))
		//更改时间
		mediaFiles.SetChangeDate(time.Now().Format("2006-01-02 15:04:05"))
		//状态
		mediaFiles.SetStatus("1")
		//审核状态
		mediaFiles.SetAuditStatus("002003")

		//同步请求体
		mediaFiles.SetFilename(uploadFileParam.Filename)
		mediaFiles.SetFileSize(uploadFileParam.FileSize)
		mediaFiles.SetFileType(uploadFileParam.FileType)

		//插入数据库
		res = init_config.Db.Create(&mediaFiles)
		if res.Error != nil {
			return xuetang.MediaFiles{}, res.Error
		} else {
			return *mediaFiles, nil
		}
	}
	return xuetang.MediaFiles{}, res.Error
}

// 获取文件Md5
func getFileMd5(file os.File) (md5Str string, err error) {
	h := md5.New()
	if _, err := file.Seek(0, 0); err != nil {
		return "", err
	}
	if _, err := io.Copy(h, &file); err != nil {
		return "", err
	}
	md5Str = hex.EncodeToString(h.Sum(nil))

	return md5Str, nil
}

// UploadMedia 上传媒资文件到对象存储
func UploadMedia(companyId int, uploadFileParam xuetang.UploadFileParamsDto, filePath string) (xuetang.UploadFileResultDto, error) {
	//文件名
	filename := uploadFileParam.Filename
	//拓展名
	//index := strings.LastIndex(filename, ".")
	//extension := filename[index+1:]
	//根据拓展名获取mimeType
	//mimeType := mime.TypeByExtension(extension)
	//打开文件
	fileData, _ := os.Open(filePath)
	//文件Md5值
	fileMd5, _ := getFileMd5(*fileData)
	objectName := fileMd5 + filename

	fmt.Println("qimiu")
	//上传到七牛云
	init_config.Initqiniu()

	cfg := storage.Config{}
	// 空间对应的机房
	cfg.Region = &storage.ZoneHuanan
	// 是否使用https域名
	cfg.UseHTTPS = false
	// 上传是否使用CDN上传加速
	cfg.UseCdnDomains = false
	// 构建表单上传的对象
	formUploader := storage.NewFormUploader(&cfg)
	ret := storage.PutRet{}
	putExtra := storage.PutExtra{}

	err := formUploader.PutFile(context.Background(), &ret, init_config.UpToken, objectName, filePath, &putExtra)
	if err != nil {
		return xuetang.UploadFileResultDto{}, err
	}

	mediaFile, err := saveMediaToDb(companyId, fileMd5, uploadFileParam, "xuetangmedia", objectName)
	if err != nil {
		return xuetang.UploadFileResultDto{}, err
	}
	uploadFileResult := xuetang.NewUploadFileResultDto()
	uploadFileResult.MediaFiles = &mediaFile
	return *uploadFileResult, nil
}

func GetPlayUrlById(mediaId string) (xuetang.MediaFiles, error) {
	mediaFiles := xuetang.NewMediaFiles()
	key := "media_files:" + mediaId
	mediaFilesJson, err := init_config.Rd.Get(key).Result()
	if err != nil {
		if err == redis.Nil {
			// Redis中不存在数据，查数据库
			fmt.Println("查数据库了！！")
			if err := init_config.Db.Where("id = ?", mediaId).First(&mediaFiles).Error; err != nil {
				return xuetang.MediaFiles{}, err
			}

			// 将数据写入Redis
			mediaFileJSON, err := json.Marshal(mediaFiles)
			if err != nil {
				return xuetang.MediaFiles{}, err
			}

			// 使用 SET 命令将媒资文件数据存储到 Redis 中
			err = init_config.Rd.Set("media_files:"+mediaFiles.Id, mediaFileJSON, 0).Err()
			if err != nil {
				return xuetang.MediaFiles{}, err
			}
			return *mediaFiles, nil
		} else {
			// Redis出错
			return xuetang.MediaFiles{}, err
		}
	}

	err = json.Unmarshal([]byte(mediaFilesJson), &mediaFiles)
	if err != nil {
		fmt.Printf("解码JSON失败：%s\n", err)
		return xuetang.MediaFiles{}, err
	}
	return *mediaFiles, nil
}
