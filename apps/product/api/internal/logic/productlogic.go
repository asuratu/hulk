package logic

import (
	"context"

	"hulk/apps/product/api/internal/svc"
	"hulk/apps/product/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ProductLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewProductLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ProductLogic {
	return &ProductLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ProductLogic) Product(req *types.Request) (resp *types.Response, err error) {
	// todo: add your logic here and delete this line

	return
}
