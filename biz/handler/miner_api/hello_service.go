// Code generated by hertz generator.

package miner_api

import (
	"context"
	"encoding/json"
	"github.com/qcq1/common/render"
	"miner_api/biz/common/Status"
	miner_api "miner_api/biz/model/miner_api"
	"miner_api/biz/sal/rpc/miner_miner_core"

	"github.com/bytedance/gopkg/util/logger"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	"github.com/qcq1/rpc_miner_core/kitex_gen/miner_core"
)

// HelloMethod .
// @router /hello [GET]
func HelloMethod(ctx context.Context, c *app.RequestContext) {
	NewHelloHandler(ctx, c).Handle()
}

type HelloHandler struct {
	ctx      context.Context
	hertzCtx *app.RequestContext

	respData *miner_api.HelloData
}

func NewHelloHandler(ctx context.Context, hertzCtx *app.RequestContext) *HelloHandler {
	return &HelloHandler{
		ctx:      ctx,
		hertzCtx: hertzCtx,
	}
}

func (h *HelloHandler) Handle() {
	ctx := h.ctx
	var req miner_api.HelloReq
	err := h.hertzCtx.BindAndValidate(&req)
	if err != nil {
		h.hertzCtx.String(consts.StatusBadRequest, err.Error())
		return
	}

	coreResp, err := miner_miner_core.RawCall.GetItem(ctx, &miner_core.GetItemReq{
		Id: 1,
	})
	if err != nil {
		logger.CtxErrorf(ctx, "miner_miner_core.RawCall.GetItem failed, err = %v", err)
		return
	}

	h.respData = h.RpcResp2HttpResp(coreResp)
	h.ReturnResp(Status.Success, err)
}

func (h *HelloHandler) ReturnResp(status *Status.Status, err error) {
	if err != nil {
		logger.CtxErrorf(h.ctx, "Hello failed, err = %v", err)
	}
	resp := new(miner_api.HelloResp)
	resp.Code = status.Code()
	resp.Message = status.Message()
	if status.Code() == Status.Success.Code() && err == nil {
		resp.Data = h.respData
	}
	logger.CtxInfof(h.ctx, "Hello, resp = %v", render.Render(resp))
	h.hertzCtx.JSON(consts.StatusOK, &resp)
}

func (h *HelloHandler) RpcResp2HttpResp(rpcResp *miner_core.GetItemResp) *miner_api.HelloData {
	marshal, err := json.Marshal(rpcResp.Item)
	if err != nil {
		return nil
	}
	return &miner_api.HelloData{
		RespBody: string(marshal),
	}
}
