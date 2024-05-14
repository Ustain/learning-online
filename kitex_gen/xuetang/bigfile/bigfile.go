// Code generated by Kitex v0.9.1. DO NOT EDIT.

package bigfile

import (
	"context"
	"errors"
	client "github.com/cloudwego/kitex/client"
	kitex "github.com/cloudwego/kitex/pkg/serviceinfo"
	xuetang "xuetang/kitex_gen/xuetang"
)

var errInvalidMessageType = errors.New("invalid message type for service method handler")

var serviceMethods = map[string]kitex.MethodInfo{
	"checkfile": kitex.NewMethodInfo(
		checkfileHandler,
		newBigFileCheckfileArgs,
		newBigFileCheckfileResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"UploadBigFile": kitex.NewMethodInfo(
		uploadBigFileHandler,
		newBigFileUploadBigFileArgs,
		newBigFileUploadBigFileResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"GetUploadProcess": kitex.NewMethodInfo(
		getUploadProcessHandler,
		newBigFileGetUploadProcessArgs,
		newBigFileGetUploadProcessResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
}

var (
	bigFileServiceInfo                = NewServiceInfo()
	bigFileServiceInfoForClient       = NewServiceInfoForClient()
	bigFileServiceInfoForStreamClient = NewServiceInfoForStreamClient()
)

// for server
func serviceInfo() *kitex.ServiceInfo {
	return bigFileServiceInfo
}

// for client
func serviceInfoForStreamClient() *kitex.ServiceInfo {
	return bigFileServiceInfoForStreamClient
}

// for stream client
func serviceInfoForClient() *kitex.ServiceInfo {
	return bigFileServiceInfoForClient
}

// NewServiceInfo creates a new ServiceInfo containing all methods
func NewServiceInfo() *kitex.ServiceInfo {
	return newServiceInfo(false, true, true)
}

// NewServiceInfo creates a new ServiceInfo containing non-streaming methods
func NewServiceInfoForClient() *kitex.ServiceInfo {
	return newServiceInfo(false, false, true)
}
func NewServiceInfoForStreamClient() *kitex.ServiceInfo {
	return newServiceInfo(true, true, false)
}

func newServiceInfo(hasStreaming bool, keepStreamingMethods bool, keepNonStreamingMethods bool) *kitex.ServiceInfo {
	serviceName := "BigFile"
	handlerType := (*xuetang.BigFile)(nil)
	methods := map[string]kitex.MethodInfo{}
	for name, m := range serviceMethods {
		if m.IsStreaming() && !keepStreamingMethods {
			continue
		}
		if !m.IsStreaming() && !keepNonStreamingMethods {
			continue
		}
		methods[name] = m
	}
	extra := map[string]interface{}{
		"PackageName": "xuetang",
	}
	if hasStreaming {
		extra["streaming"] = hasStreaming
	}
	svcInfo := &kitex.ServiceInfo{
		ServiceName:     serviceName,
		HandlerType:     handlerType,
		Methods:         methods,
		PayloadCodec:    kitex.Thrift,
		KiteXGenVersion: "v0.9.1",
		Extra:           extra,
	}
	return svcInfo
}

func checkfileHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*xuetang.BigFileCheckfileArgs)
	realResult := result.(*xuetang.BigFileCheckfileResult)
	success, err := handler.(xuetang.BigFile).Checkfile(ctx, realArg.FileMd5)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newBigFileCheckfileArgs() interface{} {
	return xuetang.NewBigFileCheckfileArgs()
}

func newBigFileCheckfileResult() interface{} {
	return xuetang.NewBigFileCheckfileResult()
}

func uploadBigFileHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*xuetang.BigFileUploadBigFileArgs)
	realResult := result.(*xuetang.BigFileUploadBigFileResult)
	success, err := handler.(xuetang.BigFile).UploadBigFile(ctx, realArg.Req, realArg.FilePath)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newBigFileUploadBigFileArgs() interface{} {
	return xuetang.NewBigFileUploadBigFileArgs()
}

func newBigFileUploadBigFileResult() interface{} {
	return xuetang.NewBigFileUploadBigFileResult()
}

func getUploadProcessHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*xuetang.BigFileGetUploadProcessArgs)
	realResult := result.(*xuetang.BigFileGetUploadProcessResult)
	success, err := handler.(xuetang.BigFile).GetUploadProcess(ctx, realArg.Filepath, realArg.FileSize)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newBigFileGetUploadProcessArgs() interface{} {
	return xuetang.NewBigFileGetUploadProcessArgs()
}

func newBigFileGetUploadProcessResult() interface{} {
	return xuetang.NewBigFileGetUploadProcessResult()
}

type kClient struct {
	c client.Client
}

func newServiceClient(c client.Client) *kClient {
	return &kClient{
		c: c,
	}
}

func (p *kClient) Checkfile(ctx context.Context, fileMd5 string) (r *xuetang.RestResponse, err error) {
	var _args xuetang.BigFileCheckfileArgs
	_args.FileMd5 = fileMd5
	var _result xuetang.BigFileCheckfileResult
	if err = p.c.Call(ctx, "checkfile", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) UploadBigFile(ctx context.Context, req *xuetang.UploadFileParamsDto, filePath string) (r *xuetang.UploadFileResultDto, err error) {
	var _args xuetang.BigFileUploadBigFileArgs
	_args.Req = req
	_args.FilePath = filePath
	var _result xuetang.BigFileUploadBigFileResult
	if err = p.c.Call(ctx, "UploadBigFile", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) GetUploadProcess(ctx context.Context, filepath string, fileSize float64) (r *xuetang.UploadProcessResult_, err error) {
	var _args xuetang.BigFileGetUploadProcessArgs
	_args.Filepath = filepath
	_args.FileSize = fileSize
	var _result xuetang.BigFileGetUploadProcessResult
	if err = p.c.Call(ctx, "GetUploadProcess", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}
