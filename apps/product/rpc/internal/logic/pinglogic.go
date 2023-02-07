package logic

import (
	"context"

	"hulk/apps/product/rpc/internal/svc"
	"hulk/apps/product/rpc/product"

	"github.com/zeromicro/go-zero/core/logx"
)

type PingLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewPingLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PingLogic {
	return &PingLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// Ping is a simple RPC that returns a single response.
func (l *PingLogic) Ping(in *product.Request) (*product.Response, error) {
	// todo: add your logic here and delete this line

	return &product.Response{}, nil
}
