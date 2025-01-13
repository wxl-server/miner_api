package miner_miner_core

import (
	"context"
	"miner_api/biz/sal/rpc"

	"github.com/cloudwego/kitex/client/callopt"
	"github.com/cloudwego/kitex/pkg/retry"
	"github.com/qcq1/common/k8s_dns"
	"github.com/qcq1/rpc_miner_core/kitex_gen/miner_core"

	"github.com/bytedance/gopkg/util/logger"
	"github.com/cloudwego/kitex/client"
	dns "github.com/kitex-contrib/resolver-dns"
	"github.com/qcq1/rpc_miner_core/kitex_gen/miner_core/minercore"
)

var RawCall = NewRawCall()

const (
	ServiceName    = "miner.miner-core"
	BoeServiceName = "127.0.0.1:8889"
	MaxRetryTimes  = 3
)

type RawCallStruct struct {
	client minercore.Client
	rp     *retry.FailurePolicy
}

func NewRawCall() *RawCallStruct {
	r := &RawCallStruct{}
	var err error
	r.client, err = minercore.NewClient(
		k8s_dns.BuildK8sDestServiceName(
			k8s_dns.WithDestServiceName(ServiceName),
			k8s_dns.WithBoeDestServiceName(BoeServiceName),
		),
		client.WithResolver(dns.NewDNSResolver()),
		client.WithMiddleware(rpc.LogMiddleware),
	)
	if err != nil {
		logger.Errorf("minercore.NewClient failed, err: %v", err)
		panic(err)
	}
	rp := retry.NewFailurePolicy()
	rp.WithMaxRetryTimes(MaxRetryTimes)
	r.rp = rp
	return r
}

func (r *RawCallStruct) QueryJobList(ctx context.Context, req *miner_core.QueryJobListReq) (resp *miner_core.QueryJobListResp, err error) {
	resp, err = r.client.QueryJobList(ctx, req, callopt.WithRetryPolicy(retry.BuildFailurePolicy(r.rp)))
	if err != nil {
		logger.CtxErrorf(ctx, "client.QueryJobList failed, err: %v", err)
		return
	}
	return
}
