package media_api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"net/http"
	"os"
	"strconv"
	"xuetang/kitex_gen/xuetang"
	"xuetang/media_model"
	"xuetang/media_service"
)

// CheckFile 检查文件是否存在
func CheckFile(c *gin.Context) {
	fileMd5 := c.Query("fileMd5")
	res, err := media_service.CheckFile(fileMd5)
	if err != nil {
		c.JSON(http.StatusBadRequest, media_model.ValidFail("找不到视频"))
		return
	}

	c.JSON(http.StatusOK, res)
}

// UploadBigFile 分片上传大文件
func UploadBigFile(c *gin.Context) {
	file, err := c.FormFile("data")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// 创建临时文件
	tempFile, err := os.CreateTemp("", "upload-*.tmp")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create temporary file"})
		return
	}
	defer os.Remove(tempFile.Name())

	// 将上传的文件内容拷贝到临时文件
	src, err := file.Open()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to open uploaded file"})
		return
	}
	defer src.Close()
	_, err = io.Copy(tempFile, src)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to copy uploaded file content"})
		return
	}

	// 获取临时文件路径
	filePath := tempFile.Name()

	fmt.Println(filePath)

	//准备上传信息
	uploadFileParam := xuetang.NewUploadFileParamsDto()
	uploadFileParam.SetFilename(file.Filename)
	uploadFileParam.SetFileSize(file.Size)
	uploadFileParam.SetFileType("001001")

	// 假设公司 ID 为 1232141425L
	companyId := 1232141425

	// 调用服务层上传媒资文件
	result, err := media_service.UploadBigFile(companyId, *uploadFileParam, filePath)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// 返回结果
	c.JSON(http.StatusOK, result)
}

// UploadProcessResult 获取上传进度
func UploadProcessResult(c *gin.Context) {
	filepath := c.Query("filepath")
	Size := c.Query("fileSize")
	fileSize, err := strconv.ParseFloat(Size, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	res, err := media_service.GetUploadProcess(filepath, fileSize)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, res)
}
