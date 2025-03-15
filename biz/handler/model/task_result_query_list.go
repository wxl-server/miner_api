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

// QueryTaskResultList .
// @router /task_result/query/list [POST]
func QueryTaskResultList(ctx context.Context, c *app.RequestContext) {
	NewQueryTaskResultListHandler(ctx, c).Handle()
}

type QueryTaskResultListHandler struct {
	ctx      context.Context
	hertzCtx *app.RequestContext

	respData *model.QueryTaskResultListData
}

func NewQueryTaskResultListHandler(ctx context.Context, hertzCtx *app.RequestContext) *QueryTaskResultListHandler {
	return &QueryTaskResultListHandler{
		ctx:      ctx,
		hertzCtx: hertzCtx,
	}
}

func (h *QueryTaskResultListHandler) Handle() {
	ctx := h.ctx
	var req model.QueryTaskResultListReq
	err := h.hertzCtx.BindAndValidate(&req)
	if err != nil {
		h.ReturnResp(Status.RequestParamsInvalid, err)
		return
	}

	coreResp, err := miner_core_rpc.QueryTaskResultList(ctx, h.HttpReq2RpcReq(&req))
	if err != nil {
		logger.CtxErrorf(ctx, "miner_core.RawCall.QueryTaskResultList failed, err = %v", err)
		h.ReturnResp(Status.InternalError, err)
		return
	}

	h.respData = h.RpcResp2HttpResp(coreResp)
	h.ReturnResp(Status.Success, err)
}

func (h *QueryTaskResultListHandler) ReturnResp(status *Status.Status, err error) {
	if err != nil {
		logger.CtxErrorf(h.ctx, "QueryTaskResultList failed, err = %v", err)
	}
	resp := new(model.QueryTaskResultListResp)
	resp.Code = status.Code()
	resp.Message = status.Message()
	if status.Code() == Status.Success.Code() && err == nil {
		resp.Data = h.respData
	}
	h.hertzCtx.JSON(consts.StatusOK, &resp)
}

func (h *QueryTaskResultListHandler) HttpReq2RpcReq(httpReq *model.QueryTaskResultListReq) *miner_core.QueryTaskResultListReq {
	return &miner_core.QueryTaskResultListReq{
		PageNum:  gptr.Indirect(httpReq.PageNum),
		PageSize: gptr.Indirect(httpReq.PageSize),

		TaskId: gptr.Indirect(httpReq.TaskId),
	}
}

func (h *QueryTaskResultListHandler) RpcResp2HttpResp(rpcResp *miner_core.QueryTaskResultListResp) *model.QueryTaskResultListData {
	return &model.QueryTaskResultListData{
		TaskResultList: gslice.Map(rpcResp.ResultList, func(v *miner_core.TaskResult_) *model.TaskResult {
			return &model.TaskResult{
				TaskId:  gptr.Of(v.TaskId),
				EsScore: gptr.Of(v.EsScore),
				Product: &model.Product{
					ProductId:           gptr.Of(v.Product.ProductId),
					ProductName:         gptr.Of(v.Product.ProductName),
					ImageUrls:           v.Product.ImageUrls,
					ShopImageUrl:        gptr.Of(v.Product.ShopImageUrl),
					ShopId:              gptr.Of(v.Product.ShopId),
					Status:              gptr.Of(v.Product.Status),
					ShopName:            gptr.Of(v.Product.ShopName),
					BrandId:             v.Product.BrandId,
					BrandName:           v.Product.BrandName,
					IsBrandAuthorized:   v.Product.IsBrandAuthorized,
					LogoModelBrandId:    v.Product.LogoModelBrandId,
					LogoModelBrandName:  v.Product.LogoModelBrandName,
					ImageModelBrandId:   v.Product.ImageModelBrandId,
					ImageModelBrandName: v.Product.ImageModelBrandName,
					Extra:               v.Product.Extra,
				},
			}
		}),
		Total: gptr.Of(rpcResp.Total),
	}
}
