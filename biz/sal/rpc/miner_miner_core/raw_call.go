package miner_miner_core

import (
	"context"

	"github.com/bytedance/gopkg/util/logger"
	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/client/callopt"
	"github.com/qcq1/rpc_miner_core/kitex_gen/miner_core"
	"github.com/qcq1/rpc_miner_core/kitex_gen/miner_core/itemservice"
)

var RawCall = NewRawCall()

const (
	K8SServiceName = "miner-core"
	K8SHostPort    = "miner-core.miner.svc.cluster.local:8888"
)

type RawCallStruct struct {
	client itemservice.Client
}

func NewRawCall() *RawCallStruct {
	r := &RawCallStruct{}
	var err error
	r.client, err = itemservice.NewClient(K8SServiceName, client.WithHostPorts(K8SHostPort))
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
