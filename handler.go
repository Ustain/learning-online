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
	"xuetang/kitex_gen/xuetang/media"
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

func Regis(cli *naming_client.INamingClient) {
	// 初始化server服务
	svr := media.NewServer(
		new(MediaImpl),
		server.WithRegistry(registry.NewNacosRegistry(*cli)),
		server.WithServerBasicInfo(&rpcinfo.EndpointBasicInfo{
			ServiceName: "userService",
		}),
		// 配置当前RPC服务对外暴露的IP和端口
		server.WithServiceAddr(&net.TCPAddr{
			IP:   net.IPv4(127, 0, 0, 1),
			Port: 8089,
		}),
	)

	err := svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
