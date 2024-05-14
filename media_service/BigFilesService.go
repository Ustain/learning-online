package media_service

import (
	"context"
	"fmt"
	"github.com/qiniu/go-sdk/v7/storage"
	"os"
	"os/exec"
	"strings"
	"xuetang/init_config"
	"xuetang/kitex_gen/xuetang"
	"xuetang/media_model"
)

func CheckFile(fileMd5 string) (xuetang.RestResponse, error) {
	//先查询数据库
	mediaFiles := xuetang.NewMediaFiles()
	if err := init_config.Db.Where("id = ?", fileMd5).First(&mediaFiles).Error; err != nil || mediaFiles == nil {
		return *media_model.Success("false"), err
	}

	bucket := "xuetangmedia"
	key := mediaFiles.GetFilePath()

	fileInfo, err := init_config.BucketManager.Stat(bucket, key)
	if err != nil {
		fmt.Println(err)
		return *media_model.Success("false"), err
	}
	fmt.Println(fileInfo.String())
	//可以解析文件的PutTime
	fmt.Println(storage.ParsePutTime(fileInfo.PutTime))

	return *media_model.Success("true"), err
}

// 判断文件路径是否为AVI文件
func isAVI(filePath string) bool {
	// 检查文件路径的后缀名是否为.avi
	return strings.HasSuffix(strings.ToLower(filePath), ".avi")
}

// 视频格式转换
func aviToMp4(inputFile string, outputFile string) error {
	cmd := exec.Command("ffmpeg", "-i", inputFile, "-c:v", "libx264", "-preset", "slow", "-crf", "22", "-c:a", "aac", "-b:a", "192k", "-movflags", "faststart", outputFile)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}

var Recordermy *storage.FileRecorder

func UploadBigFile(companyId int, uploadFileParam xuetang.UploadFileParamsDto, filePath string) (xuetang.UploadFileResultDto, error) {
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

	//上传到七牛云
	init_config.Initqiniu()

	cfg := storage.Config{}
	// 空间对应的机房
	cfg.Region = &storage.ZoneHuanan
	// 是否使用https域名
	cfg.UseHTTPS = false
	// 上传是否使用CDN上传加速
	cfg.UseCdnDomains = false

	resumeUploader := storage.NewResumeUploaderV2(&cfg)
	ret := storage.PutRet{}

	var err error
	Recordermy, err = storage.NewFileRecorder(os.TempDir())
	if err != nil {
		fmt.Println(err)
		return xuetang.UploadFileResultDto{}, err
	}

	putExtra := storage.RputV2Extra{
		Recorder: Recordermy,
	}

	//采用分片上传
	err = resumeUploader.PutFile(context.Background(), &ret, init_config.UpToken, objectName, filePath, &putExtra)
	if err != nil {
		fmt.Println(err)
		return xuetang.UploadFileResultDto{}, err
	}
	fmt.Println(ret.Key, ret.Hash)

	//保存到数据库
	mediaFile, err := saveMediaToDb(companyId, fileMd5, uploadFileParam, "xuetangmedia", objectName)
	if err != nil {
		return xuetang.UploadFileResultDto{}, err
	}
	uploadFileResult := xuetang.NewUploadFileResultDto()
	uploadFileResult.MediaFiles = &mediaFile
	return *uploadFileResult, nil
}

func GetUploadProcess(filepath string, fileSize float64) (xuetang.UploadProcessResult_, error) {
	//先查询数据库
	mediaFiles := xuetang.NewMediaFiles()
	if err := init_config.Db.Where("file_path = ?", filepath).First(&mediaFiles).Error; err != nil {
		return xuetang.UploadProcessResult_{}, err
	}

	if mediaFiles != nil {
		result := xuetang.UploadProcessResult_{
			Filepath: filepath,
			Process:  100,
		}
		return result, nil
	}

	bucket := "xuetangmedia"
	key := mediaFiles.GetFilePath()

	fileInfo, err := init_config.BucketManager.Stat(bucket, key)
	if err != nil {
		fmt.Println(err)
		return xuetang.UploadProcessResult_{}, err
	}
	if fileInfo.Md5 != "" {
		result := xuetang.UploadProcessResult_{
			Filepath: filepath,
			Process:  100,
		}
		return result, nil
	}

	fmt.Println(fileInfo.String())
	//可以解析文件的PutTime
	fmt.Println(storage.ParsePutTime(fileInfo.PutTime))

	offset, err := Recordermy.Get(filepath)
	if err != nil {
		return xuetang.UploadProcessResult_{}, err
	}

	var uploadedSize int64
	if offset != nil {
		uploadedSize = int64(len(offset))
	}

	//计算进度
	progress := float64(uploadedSize) / fileSize * 100

	result := xuetang.UploadProcessResult_{
		Filepath: filepath,
		Process:  progress,
	}

	return result, nil
}
