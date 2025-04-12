package miner_integrate_rpc

import (
	"context"
	"github.com/bytedance/gopkg/util/logger"
	"github.com/wxl-server/common/wxl_cluster"
	"github.com/wxl-server/idl_gen/kitex_gen/miner_integrate"
	"github.com/wxl-server/idl_gen/kitex_gen/miner_integrate/minerintegrate"
)

var client = wxl_cluster.NewClient(minerintegrate.NewClient, "miner_integrate")

func UpdateMockProducerQps(ctx context.Context, req *miner_integrate.UpdateMockProducerQpsReq) (resp *miner_integrate.UpdateMockProducerQpsResp, err error) {
	resp, err = client.UpdateMockProducerQps(ctx, req)
	if err != nil {
		logger.CtxErrorf(ctx, "UpdateMockProducerQps failed, err = %v", err)
		return nil, err
	}
	return resp, nil
}
