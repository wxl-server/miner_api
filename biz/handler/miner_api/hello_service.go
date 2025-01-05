// Code generated by hertz generator.

package miner_api

import (
	"context"
	"encoding/json"
	"github.com/bytedance/gopkg/util/logger"
	"github.com/qcq1/rpc_miner_core/kitex_gen/miner_core"
	"github.com/qcq1/rpc_miner_core/kitex_gen/miner_core/itemservice"
	"log"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	"github.com/cloudwego/kitex/client"
	dns "github.com/kitex-contrib/resolver-dns"
	miner_api "miner_api/biz/model/miner_api"
)

// HelloMethod .
// @router /hello [GET]
func HelloMethod(ctx context.Context, c *app.RequestContext) {
	var err error
	var req miner_api.HelloReq
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	cli, err := itemservice.NewClient("miner-core.miner", client.WithResolver(dns.NewDNSResolver()))
	if err != nil {
		log.Fatal(err)
	}

	resp1, err := cli.GetItem(ctx, &miner_core.GetItemReq{
		Id: 1,
	})
	if err != nil {
		logger.CtxErrorf(ctx, "client.GetItem failed, err: %v", err)
		return
	}

	marshal, err := json.Marshal(resp1)
	if err != nil {
		logger.CtxErrorf(ctx, "json.Marshal failed, err: %v", err)
		return
	}
	resp := &miner_api.HelloResp{
		RespBody: string(marshal),
	}

	c.JSON(consts.StatusOK, resp)
}
