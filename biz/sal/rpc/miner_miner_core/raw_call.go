package miner_miner_core

import (
	"context"
	"github.com/cloudwego/kitex/client/callopt"
	"github.com/qcq1/common/env"
	"github.com/qcq1/common/render"
	"github.com/qcq1/rpc_miner_core/kitex_gen/miner_core"
	"miner_api/biz/sal/rpc"

	"github.com/bytedance/gopkg/util/logger"
	"github.com/cloudwego/kitex/client"
	"github.com/qcq1/rpc_miner_core/kitex_gen/miner_core/minercore"
)

var RawCall = NewRawCall()

const (
	ServiceName = "miner-core"
)

type RawCallStruct struct {
	client minercore.Client
}

func NewRawCall() *RawCallStruct {
	r := &RawCallStruct{}
	var err error
	r.client, err = minercore.NewClient(ServiceName, client.WithHostPorts(rpc.Router[ServiceName][env.GetEnv()]))
	if err != nil {
		logger.Errorf("minercore.NewClient failed, err: %v", err)
		panic(err)
	}
	return r
}

func (r *RawCallStruct) QueryJobList(ctx context.Context, req *miner_core.QueryJobListReq, callOptions ...callopt.Option) (resp *miner_core.QueryJobListResp, err error) {
	logger.CtxInfof(ctx, "client.QueryJobList req = %v", render.Render(req))
	resp, err = r.client.QueryJobList(ctx, req, callOptions...)
	if err != nil {
		logger.CtxErrorf(ctx, "client.QueryJobList failed, err: %v", err)
		return
	}
	logger.CtxInfof(ctx, "client.QueryJobList resp = %v", render.Render(resp))
	return
}
