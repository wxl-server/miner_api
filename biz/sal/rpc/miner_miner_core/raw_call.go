package miner_miner_core

import (
	"context"
	"miner_api/biz/sal/config"
	"miner_api/biz/sal/rpc"

	"github.com/cloudwego/kitex/client/callopt"
	"github.com/cloudwego/kitex/pkg/retry"
	"github.com/nacos-group/nacos-sdk-go/clients"
	"github.com/nacos-group/nacos-sdk-go/vo"
	"github.com/qcq1/rpc_miner_core/kitex_gen/miner_core"

	"github.com/bytedance/gopkg/util/logger"
	"github.com/cloudwego/kitex/client"
	"github.com/kitex-contrib/registry-nacos/resolver"
	"github.com/nacos-group/nacos-sdk-go/common/constant"
	"github.com/qcq1/rpc_miner_core/kitex_gen/miner_core/minercore"
)

var RawCall *RawCallStruct

const (
	ServiceName   = "miner_core"
	MaxRetryTimes = 3
)

type RawCallStruct struct {
	client minercore.Client
	rp     *retry.FailurePolicy
}

func NewRawCall(ctx context.Context, config *config.AppConfig) {
	RawCall = &RawCallStruct{}
	nacosConfig := config.Nacos

	cli, err := clients.NewNamingClient(
		vo.NacosClientParam{
			ClientConfig: &constant.ClientConfig{
				NamespaceId: nacosConfig.Namespace,
				Username:    nacosConfig.Username,
				Password:    nacosConfig.Password,
			},
			ServerConfigs: []constant.ServerConfig{
				*constant.NewServerConfig(nacosConfig.Host, nacosConfig.Port),
			},
		},
	)
	if err != nil {
		logger.Errorf("[Init] clients.NewNamingClient failed, err = %v", err)
		panic(err)
	}
	RawCall.client, err = minercore.NewClient(
		ServiceName,
		client.WithResolver(resolver.NewNacosResolver(cli)),
		client.WithMiddleware(rpc.LogMiddleware),
	)
	if err != nil {
		logger.Errorf("[Init] minercore.NewClient failed, err = %v", err)
		panic(err)
	}
	rp := retry.NewFailurePolicy()
	rp.WithMaxRetryTimes(MaxRetryTimes)
	RawCall.rp = rp
	logger.CtxInfof(ctx, "[Init] init rpc client success, rpc serviceName = %v", ServiceName)
}

func (r *RawCallStruct) QueryJobList(ctx context.Context, req *miner_core.QueryJobListReq) (resp *miner_core.QueryJobListResp, err error) {
	resp, err = r.client.QueryJobList(ctx, req, callopt.WithRetryPolicy(retry.BuildFailurePolicy(r.rp)))
	if err != nil {
		logger.CtxErrorf(ctx, "client.QueryJobList failed, err = %v", err)
		return
	}
	return
}
