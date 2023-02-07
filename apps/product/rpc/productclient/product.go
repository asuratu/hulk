// Code generated by goctl. DO NOT EDIT.
// Source: product.proto

package productclient

import (
	"context"

	"hulk/apps/product/rpc/product"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	Request  = product.Request
	Response = product.Response

	Product interface {
		Ping(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error)
	}

	defaultProduct struct {
		cli zrpc.Client
	}
)

func NewProduct(cli zrpc.Client) Product {
	return &defaultProduct{
		cli: cli,
	}
}

func (m *defaultProduct) Ping(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error) {
	client := product.NewProductClient(m.cli.Conn())
	return client.Ping(ctx, in, opts...)
}
