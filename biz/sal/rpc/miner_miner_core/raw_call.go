package miner_miner_core

import (
	"context"
	"miner_api/biz/sal/rpc"

	"github.com/cloudwego/kitex/client/callopt"
	"github.com/qcq1/common/env"
	"github.com/qcq1/rpc_miner_core/kitex_gen/miner_core"

	"github.com/bytedance/gopkg/util/logger"
	"github.com/cloudwego/kitex/client"
	dns "github.com/kitex-contrib/resolver-dns"
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
	r.client, err = minercore.NewClient(rpc.Router[ServiceName][env.GetEnv()], client.WithResolver(dns.NewDNSResolver()), client.WithMiddleware(rpc.LogMiddleware))
	if err != nil {
		logger.Errorf("minercore.NewClient failed, err: %v", err)
		panic(err)
	}
	return r
}

func (r *RawCallStruct) QueryJobList(ctx context.Context, req *miner_core.QueryJobListReq, callOptions ...callopt.Option) (resp *miner_core.QueryJobListResp, err error) {
	resp, err = r.client.QueryJobList(ctx, req, callOptions...)
	if err != nil {
		logger.CtxErrorf(ctx, "client.QueryJobList failed, err: %v", err)
		return
	}
	return
}
