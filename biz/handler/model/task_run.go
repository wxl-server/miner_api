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
	"github.com/wxl-server/common/gslice"
	"github.com/wxl-server/idl_gen/kitex_gen/miner_core"
)

// RunTask .
// @router /task/run [GET]
func RunTask(ctx context.Context, c *app.RequestContext) {
	NewRunTaskHandlerHandler(ctx, c).Handle()
}

type RunTaskHandler struct {
	ctx      context.Context
	hertzCtx *app.RequestContext

	respData *model.RunTaskData
}

func NewRunTaskHandlerHandler(ctx context.Context, hertzCtx *app.RequestContext) *RunTaskHandler {
	return &RunTaskHandler{
		ctx:      ctx,
		hertzCtx: hertzCtx,
	}
}

func (h *RunTaskHandler) Handle() {
	ctx := h.ctx
	var req model.RunTaskReq
	err := h.hertzCtx.BindAndValidate(&req)
	if err != nil {
		h.ReturnResp(Status.RequestParamsInvalid, err)
		return
	}

	resp, err := miner_core_rpc.RunTask(ctx, h.HttpReq2RpcReq(&req))
	if err != nil {
		logger.CtxErrorf(ctx, "miner_core_rpc.RunTask failed, err = %v", err)
		h.ReturnResp(Status.InternalError, err)
		return
	}

	h.respData = h.RpcResp2HttpResp(resp)
	h.ReturnResp(Status.Success, err)
}

func (h *RunTaskHandler) HttpReq2RpcReq(httpReq *model.RunTaskReq) *miner_core.RunTaskReq {
	return &miner_core.RunTaskReq{
		JobId: gptr.Indirect(httpReq.JobId),
		Rules: gslice.Map(httpReq.Rules, func(v *model.Rule) *miner_core.Rule {
			return &miner_core.Rule{
				Id:           gptr.Indirect(v.Id),
				FactorCode:   gptr.Indirect(v.FactorCode),
				OperatorCode: gptr.Indirect(v.OperatorCode),
				ValueList:    v.ValueList,
			}
		}),
		LogicExpression: gptr.Indirect(httpReq.LogicExpression),
		Limit:           gptr.Indirect(httpReq.Limit),
	}
}
func (h *RunTaskHandler) RpcResp2HttpResp(rpcResp *miner_core.RunTaskResp) *model.RunTaskData {
	return &model.RunTaskData{}
}

func (h *RunTaskHandler) ReturnResp(status *Status.Status, err error) {
	if err != nil {
		logger.CtxErrorf(h.ctx, "RunTask failed, err = %v", err)
	}
	resp := new(model.RunTaskResp)
	resp.Code = status.Code()
	resp.Message = status.Message()
	if status.Code() == Status.Success.Code() && err == nil {
		resp.Data = h.respData
	}
	h.hertzCtx.JSON(consts.StatusOK, &resp)
}
