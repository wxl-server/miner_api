package miner_miner_core

import (
	"context"
	"github.com/qcq1/common/env"
	"miner_api/biz/sal/rpc"

	"github.com/bytedance/gopkg/util/logger"
	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/client/callopt"
	"github.com/qcq1/rpc_miner_core/kitex_gen/miner_core"
	"github.com/qcq1/rpc_miner_core/kitex_gen/miner_core/itemservice"
)

var RawCall = NewRawCall()

const (
	ServiceName = "miner-core"
)

type RawCallStruct struct {
	client itemservice.Client
}

func NewRawCall() *RawCallStruct {
	r := &RawCallStruct{}
	var err error
	r.client, err = itemservice.NewClient(ServiceName, client.WithHostPorts(rpc.Router[ServiceName][env.GetEnv()]))
	if err != nil {
		logger.Errorf("itemservice.NewClient failed, err: %v", err)
		panic(err)
	}
	return r
}

func (r *RawCallStruct) GetItem(ctx context.Context, req *miner_core.GetItemReq, callOptions ...callopt.Option) (resp *miner_core.GetItemResp, err error) {
	logger.CtxInfof(ctx, "client.GetItem req: %v", req)
	resp, err = r.client.GetItem(ctx, req, callOptions...)
	if err != nil {
		logger.CtxErrorf(ctx, "client.GetItem failed, err: %v", err)
		return
	}
	logger.CtxInfof(ctx, "client.GetItem resp: %v", resp)
	return
}
