// Code generated by hertz generator.

package model

import (
	"context"
	"miner_api/biz/common/Status"
	"miner_api/biz/sal/rpc/miner_core_rpc"

	model "miner_api/biz/model"

	"github.com/bytedance/gopkg/util/logger"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	"github.com/wxl-server/common/gptr"
	"github.com/wxl-server/idl_gen/kitex_gen/miner_core"
)

// Login .
// @router /user/login [POST]
func Login(ctx context.Context, c *app.RequestContext) {
	NewLoginHandler(ctx, c).Handle()
}

type LoginHandler struct {
	ctx      context.Context
	hertzCtx *app.RequestContext

	respData *model.LoginData
}

func NewLoginHandler(ctx context.Context, hertzCtx *app.RequestContext) *LoginHandler {
	return &LoginHandler{
		ctx:      ctx,
		hertzCtx: hertzCtx,
	}
}

func (h *LoginHandler) Handle() {
	ctx := h.ctx
	var req model.LoginReq
	err := h.hertzCtx.BindAndValidate(&req)
	if err != nil {
		h.ReturnResp(Status.RequestParamsInvalid, err)
		return
	}

	rpcResp, err := miner_core_rpc.Login(ctx, h.HttpReq2RpcReq(&req))
	if err != nil {
		logger.CtxErrorf(ctx, "miner_core_rpc.Login failed, err = %v", err)
		h.ReturnResp(Status.InternalError, err)
		return
	}

	h.respData = h.RpcResp2HttpResp(rpcResp)
	h.ReturnResp(Status.Success, err)
}

func (h *LoginHandler) ReturnResp(status *Status.Status, err error) {
	if err != nil {
		logger.CtxErrorf(h.ctx, "login failed, err = %v", err)
	}
	resp := new(model.LoginResp)
	resp.Code = status.Code()
	resp.Message = status.Message()
	if status.Code() == Status.Success.Code() && err == nil {
		resp.Data = h.respData
	}
	h.hertzCtx.JSON(consts.StatusOK, &resp)
}

func (h *LoginHandler) HttpReq2RpcReq(httpReq *model.LoginReq) *miner_core.LoginReq {
	return &miner_core.LoginReq{
		Email:    gptr.Indirect(httpReq.Email),
		Password: gptr.Indirect(httpReq.Password),
	}
}

func (h *LoginHandler) RpcResp2HttpResp(rpcResp *miner_core.LoginResp) *model.LoginData {
	return &model.LoginData{
		Token: gptr.Of(rpcResp.Token),
	}
}
