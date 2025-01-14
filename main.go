// Code generated by hertz generator.

package main

import (
	"context"
	"miner_api/biz/sal/config"
	"miner_api/biz/sal/rpc/miner_miner_core"

	"github.com/bytedance/gopkg/util/logger"
	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/qcq1/common/render"
	"go.uber.org/dig"
)

var (
	initCtx   = context.Background()
	container = dig.New()
)

func main() {
	initContainer()

	h := server.Default()

	register(h)
	h.Spin()
}

func initContainer() {
	// context
	{
		mustProvide(func() context.Context { return initCtx })
	}

	// config
	{
		mustProvide(config.InitAppConfig)
	}

	// rpc
	{
		mustInvoke(miner_miner_core.NewRawCall)
	}
}

func mustProvide(constructor interface{}, opts ...dig.ProvideOption) {
	if err := container.Provide(constructor); err != nil {
		logger.Errorf("container provide failed, err = %v, constructor = %v", err, render.Render(constructor))
		panic(err)
	}
}

func mustInvoke(function interface{}, opts ...dig.InvokeOption) {
	if err := container.Invoke(function); err != nil {
		logger.Errorf("container invoke failed, err = %v, function = %v", err, render.Render(function))
		panic(err)
	}
}
