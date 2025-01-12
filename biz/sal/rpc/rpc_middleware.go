package rpc

import (
	"context"
	"github.com/bytedance/gopkg/util/logger"
	"github.com/cloudwego/kitex/pkg/endpoint"
	"github.com/cloudwego/kitex/pkg/utils"
	"github.com/qcq1/common/render"
)

func LogMiddleware(next endpoint.Endpoint) endpoint.Endpoint {
	return func(ctx context.Context, request, response interface{}) error {
		if arg, ok := request.(utils.KitexArgs); ok {
			if req := arg.GetFirstArgument(); req != nil {
				logger.CtxInfof(ctx, "Send rpc request = %v", render.Render(req))
			}
		}
		err := next(ctx, request, response)
		if result, ok := response.(utils.KitexResult); ok {
			if resp := result.GetResult(); resp != nil {
				logger.CtxInfof(ctx, "Get rpc response = %v", render.Render(resp))
			}
		}
		return err
	}
}
