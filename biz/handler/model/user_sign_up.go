package model

import (
	"context"
	"miner_api/biz/common/Status"
	"miner_api/biz/model"
	"miner_api/biz/sal/rpc/common_user_rpc"

	"github.com/bytedance/gopkg/util/logger"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	"github.com/wxl-server/common/gptr"
	"github.com/wxl-server/idl_gen/kitex_gen/common_user"
)

// SignUp .
// @router /user/signup [POST]
func SignUp(ctx context.Context, c *app.RequestContext) {
	NewSignUpHandler(ctx, c).Handle()
}

type SignUpHandler struct {
	ctx      context.Context
	hertzCtx *app.RequestContext

	respData *model.SignUpData
}

func NewSignUpHandler(ctx context.Context, hertzCtx *app.RequestContext) *SignUpHandler {
	return &SignUpHandler{
		ctx:      ctx,
		hertzCtx: hertzCtx,
	}
}

func (h *SignUpHandler) Handle() {
	ctx := h.ctx
	var req model.SignUpReq
	err := h.hertzCtx.BindAndValidate(&req)
	if err != nil {
		h.ReturnResp(Status.RequestParamsInvalid, err)
		return
	}

	rpcResp, err := common_user_rpc.SignUp(ctx, h.HttpReq2RpcReq(&req))
	if err != nil {
		logger.CtxErrorf(ctx, "common_user_rpc.SignUp failed, err = %v", err)
		h.ReturnResp(Status.InternalError, err)
		return
	}

	h.respData = h.RpcResp2HttpResp(rpcResp)
	h.ReturnResp(Status.Success, err)
}

func (h *SignUpHandler) ReturnResp(status *Status.Status, err error) {
	if err != nil {
		logger.CtxErrorf(h.ctx, "sign up failed, err = %v", err)
	}
	resp := new(model.SignUpResp)
	resp.Code = status.Code()
	resp.Message = status.Message()
	if status.Code() == Status.Success.Code() && err == nil {
		resp.Data = h.respData
	}
	h.hertzCtx.JSON(consts.StatusOK, &resp)
}

func (h *SignUpHandler) HttpReq2RpcReq(httpReq *model.SignUpReq) *common_user.SignUpReq {
	return &common_user.SignUpReq{
		Email:    gptr.Indirect(httpReq.Email),
		Password: gptr.Indirect(httpReq.Password),
	}
}

func (h *SignUpHandler) RpcResp2HttpResp(rpcResp *common_user.SignUpResp) *model.SignUpData {
	return &model.SignUpData{
		Id: gptr.Of(rpcResp.Id),
	}
}
