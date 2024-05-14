// Code generated by Kitex v0.9.1. DO NOT EDIT.

package bigfile

import (
	"context"
	client "github.com/cloudwego/kitex/client"
	callopt "github.com/cloudwego/kitex/client/callopt"
	xuetang "xuetang/kitex_gen/xuetang"
)

// Client is designed to provide IDL-compatible methods with call-option parameter for kitex framework.
type Client interface {
	Checkfile(ctx context.Context, fileMd5 string, callOptions ...callopt.Option) (r *xuetang.RestResponse, err error)
	UploadBigFile(ctx context.Context, req *xuetang.UploadFileParamsDto, filePath string, callOptions ...callopt.Option) (r *xuetang.UploadFileResultDto, err error)
	GetUploadProcess(ctx context.Context, filepath string, fileSize float64, callOptions ...callopt.Option) (r *xuetang.UploadProcessResult_, err error)
}

// NewClient creates a client for the service defined in IDL.
func NewClient(destService string, opts ...client.Option) (Client, error) {
	var options []client.Option
	options = append(options, client.WithDestService(destService))

	options = append(options, opts...)

	kc, err := client.NewClient(serviceInfoForClient(), options...)
	if err != nil {
		return nil, err
	}
	return &kBigFileClient{
		kClient: newServiceClient(kc),
	}, nil
}

// MustNewClient creates a client for the service defined in IDL. It panics if any error occurs.
func MustNewClient(destService string, opts ...client.Option) Client {
	kc, err := NewClient(destService, opts...)
	if err != nil {
		panic(err)
	}
	return kc
}

type kBigFileClient struct {
	*kClient
}

func (p *kBigFileClient) Checkfile(ctx context.Context, fileMd5 string, callOptions ...callopt.Option) (r *xuetang.RestResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.Checkfile(ctx, fileMd5)
}

func (p *kBigFileClient) UploadBigFile(ctx context.Context, req *xuetang.UploadFileParamsDto, filePath string, callOptions ...callopt.Option) (r *xuetang.UploadFileResultDto, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.UploadBigFile(ctx, req, filePath)
}

func (p *kBigFileClient) GetUploadProcess(ctx context.Context, filepath string, fileSize float64, callOptions ...callopt.Option) (r *xuetang.UploadProcessResult_, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.GetUploadProcess(ctx, filepath, fileSize)
}
