package service

import (
	"context"
	"github.com/2u35-Wooje/gorder-v2/stock/app"
)

// 初始化app 传入main.go的app

func NewApplication(ctx context.Context) app.Application {
	return app.Application{}
}
