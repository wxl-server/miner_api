package miner_core_rpc

import (
	"context"

	"github.com/cloudwego/kitex/client/callopt"
	"github.com/cloudwego/kitex/pkg/retry"
	"github.com/wxl-server/common/wxl_cluster"
	"github.com/wxl-server/idl_gen/kitex_gen/miner_core"

	"github.com/bytedance/gopkg/util/logger"
	"github.com/wxl-server/idl_gen/kitex_gen/miner_core/minercore"
)

var client = wxl_cluster.NewClient(minercore.NewClient, "miner_core")

func QueryJobList(ctx context.Context, req *miner_core.QueryJobListReq) (resp *miner_core.QueryJobListResp, err error) {
	// 请求，最大重试2次
	resp, err = client.QueryJobList(ctx, req, callopt.WithRetryPolicy(retry.BuildFailurePolicy(retry.NewFailurePolicy())))
	if err != nil {
		logger.CtxErrorf(ctx, "client.QueryJobList failed, err = %v", err)
		return
	}
	return
}
