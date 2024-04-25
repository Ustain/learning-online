package media_api

import (
	"github.com/gin-gonic/gin"
	"io"
	"net/http"
	"os"
	"strconv"
	"xuetang/kitex_gen/xuetang"
	"xuetang/media_service"
)

// ListMediaFiles 查询媒资列表的接口处理函数
func ListMediaFiles(c *gin.Context) {
	No := c.Query("pageNo")
	Size := c.Query("pageSize")
	pageNo, _ := strconv.ParseInt(No, 10, 64)
	pageSize, _ := strconv.ParseInt(Size, 10, 64)

	// 从请求中获取页码和查询参数
	pageParams := xuetang.NewPageParams()
	pageParams.SetPageNo(pageNo)
	pageParams.SetPageSize(pageSize)
	if err := c.Bind(&pageParams); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// 假设公司 ID 为 1232141425L
	companyId := 1232141425

	// 调用服务层查询媒资文件
	result, err := media_service.QueryMediaFiles(companyId, *pageParams)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// 返回结果
	c.JSON(http.StatusOK, result)
}

func UploadMediaFiles(c *gin.Context) {
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

	//准备上传信息
	uploadFileParam := xuetang.NewUploadFileParamsDto()
	uploadFileParam.SetFilename(file.Filename)
	uploadFileParam.SetFileSize(file.Size)
	uploadFileParam.SetFileType("001001")

	// 假设公司 ID 为 1232141425L
	companyId := 1232141425

	// 调用服务层上传媒资文件
	result, err := media_service.UploadMedia(companyId, *uploadFileParam, filePath)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// 返回结果
	c.JSON(http.StatusOK, result)
}
