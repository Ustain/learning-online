package main

import (
	"context"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/server"
	"github.com/kitex-contrib/registry-nacos/registry"
	"github.com/nacos-group/nacos-sdk-go/clients/naming_client"
	"log"
	"net"
	"xuetang/kitex_gen/xuetang"
	"xuetang/kitex_gen/xuetang/bigfile"
	"xuetang/kitex_gen/xuetang/media"
	"xuetang/media_model"
	"xuetang/media_service"
)

// MediaImpl implements the last service interface defined in the IDL.
type MediaImpl struct{}

// QueryMediaFiles implements the MediaImpl interface.
func (s *MediaImpl) QueryMediaFiles(ctx context.Context, req *xuetang.PageParams) (resp *xuetang.PageResult_, err error) {
	// 假设公司 ID 为 1232141425L
	companyId := 1232141425

	// 调用服务层查询媒资文件
	result, err := media_service.QueryMediaFiles(companyId, *req)
	if err != nil {
		return nil, err
	}
	// 返回结果
	return &result, nil
}

// UploadMediaFiles implements the MediaImpl interface.
func (s *MediaImpl) UploadMediaFiles(ctx context.Context, req *xuetang.UploadFileParamsDto, filePath string) (resp *xuetang.UploadFileResultDto, err error) {
	// 假设公司 ID 为 1232141425L
	companyId := 1232141425

	// 调用服务层上传媒资文件
	result, err := media_service.UploadMedia(companyId, *req, filePath)
	if err != nil {
		return nil, err
	}
	// 返回结果
	return &result, nil
}

// GetPlayUrlByMediaId implements the MediaImpl interface.
func (s *MediaImpl) GetPlayUrlByMediaId(ctx context.Context, mediaId string) (resp *xuetang.RestResponse, err error) {
	mediaFiles, err := media_service.GetPlayUrlById(mediaId)
	if err != nil {
		return nil, err
	}

	url := mediaFiles.GetUrl()
	if url == "" {
		return media_model.ValidFail("视频正在处理中"), nil
	}
	return media_model.Success(url), nil
}

// BigFileImpl 大文件分片上传服务
type BigFileImpl struct{}

// Checkfile implements the BigFileImpl interface.
func (s *BigFileImpl) Checkfile(ctx context.Context, fileMd5 string) (resp *xuetang.RestResponse, err error) {
	// TODO: Your code here...
	res, err := media_service.CheckFile(fileMd5)
	if err != nil {
		return nil, err
	}

	return &res, nil
}

// UploadBigFile implements the BigFileImpl interface.
func (s *BigFileImpl) UploadBigFile(ctx context.Context, req *xuetang.UploadFileParamsDto, filePath string) (resp *xuetang.UploadFileResultDto, err error) {
	// 假设公司 ID 为 1232141425L
	companyId := 1232141425

	// 调用服务层上传媒资文件
	result, err := media_service.UploadBigFile(companyId, *req, filePath)
	if err != nil {
		return nil, err
	}
	// 返回结果
	return &result, nil
}

// GetUploadProcess implements the BigFileImpl interface.
func (s *BigFileImpl) GetUploadProcess(ctx context.Context, filepath string, fileSize float64) (resp *xuetang.UploadProcessResult_, err error) {
	result, err := media_service.GetUploadProcess(filepath, fileSize)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func Regis(cli *naming_client.INamingClient) {
	// 初始化原服务
	mediaService := media.NewServer(
		new(MediaImpl),
		server.WithRegistry(registry.NewNacosRegistry(*cli)),
		server.WithServerBasicInfo(&rpcinfo.EndpointBasicInfo{
			ServiceName: "mediaService",
		}),
		// 配置当前RPC服务对外暴露的IP和端口
		server.WithServiceAddr(&net.TCPAddr{
			IP:   net.IPv4(127, 0, 0, 1),
			Port: 8089,
		}),
	)

	// 初始化 BigFile 服务
	bigFileService := bigfile.NewServer(
		new(BigFileImpl),
		server.WithRegistry(registry.NewNacosRegistry(*cli)),
		server.WithServerBasicInfo(&rpcinfo.EndpointBasicInfo{
			ServiceName: "bigFileService",
		}),
		// 配置当前RPC服务对外暴露的IP和端口
		server.WithServiceAddr(&net.TCPAddr{
			IP:   net.IPv4(127, 0, 0, 1),
			Port: 8099, // 可以根据需要修改端口号
		}),
	)

	go func() {
		err := mediaService.Run()
		if err != nil {
			log.Println("Failed to start userService:", err.Error())
		}
	}()

	go func() {
		err := bigFileService.Run()
		if err != nil {
			log.Println("Failed to start bigFileService:", err.Error())
		}
	}()
}
