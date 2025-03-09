package miner_core_rpc

import (
	"context"

	"github.com/cloudwego/kitex/client/callopt"
	"github.com/cloudwego/kitex/pkg/retry"
	"github.com/wxl-server/common/wxl_cluster"
	"github.com/wxl-server/idl_gen/kitex_gen/miner_core"

	"github.com/bytedance/gopkg/util/logger"
	"github.com/wxl-server/idl_gen/kitex_gen/miner_core/minercore"
)

var client = wxl_cluster.NewClient(minercore.NewClient, "miner_core")

func Login(ctx context.Context, req *miner_core.LoginReq) (resp *miner_core.LoginResp, err error) {
	// 请求，最大重试2次
	resp, err = client.Login(ctx, req, callopt.WithRetryPolicy(retry.BuildFailurePolicy(retry.NewFailurePolicy())))
	if err != nil {
		logger.CtxErrorf(ctx, "Login failed, err = %v", err)
		return nil, err
	}
	return resp, nil
}

func SignUp(ctx context.Context, req *miner_core.SignUpReq) (resp *miner_core.SignUpResp, err error) {
	// rpc不保证幂等，不重试
	resp, err = client.SignUp(ctx, req)
	if err != nil {
		logger.CtxErrorf(ctx, "client.SignUp failed, err = %v", err)
		return
	}
	return
}

func QueryUserList(ctx context.Context, req *miner_core.QueryUserListReq) (resp *miner_core.QueryUserListResp, err error) {
	// 请求，最大重试2次
	resp, err = client.QueryUserList(ctx, req, callopt.WithRetryPolicy(retry.BuildFailurePolicy(retry.NewFailurePolicy())))
	if err != nil {
		logger.CtxErrorf(ctx, "client.QueryUserList failed, err = %v", err)
		return
	}
	return
}

func QueryJobList(ctx context.Context, req *miner_core.QueryJobListReq) (resp *miner_core.QueryJobListResp, err error) {
	// 请求，最大重试2次
	resp, err = client.QueryJobList(ctx, req, callopt.WithRetryPolicy(retry.BuildFailurePolicy(retry.NewFailurePolicy())))
	if err != nil {
		logger.CtxErrorf(ctx, "client.QueryJobList failed, err = %v", err)
		return
	}
	return
}

func CreateJob(ctx context.Context, req *miner_core.CreateJobReq) (resp *miner_core.CreateJobResp, err error) {
	// rpc不保证幂等，不重试
	resp, err = client.CreateJob(ctx, req)
	if err != nil {
		logger.CtxErrorf(ctx, "client.CreateJob failed, err = %v", err)
		return
	}
	return
}

func DeleteJob(ctx context.Context, req *miner_core.DeleteJobReq) (resp *miner_core.DeleteJobResp, err error) {
	// 请求，最大重试2次
	resp, err = client.DeleteJob(ctx, req, callopt.WithRetryPolicy(retry.BuildFailurePolicy(retry.NewFailurePolicy())))
	if err != nil {
		logger.CtxErrorf(ctx, "client.DeleteJob failed, err = %v", err)
		return
	}
	return
}

func QueryIndicatorList(ctx context.Context, req *miner_core.QueryIndicatorListReq) (resp *miner_core.QueryIndicatorListResp, err error) {
	// 请求，最大重试2次
	resp, err = client.QueryIndicatorList(ctx, req, callopt.WithRetryPolicy(retry.BuildFailurePolicy(retry.NewFailurePolicy())))
	if err != nil {
		logger.CtxErrorf(ctx, "client.QueryIndicatorList failed, err = %v", err)
		return
	}
	return
}

func QueryTaskList(ctx context.Context, req *miner_core.QueryTaskListReq) (resp *miner_core.QueryTaskListResp, err error) {
	// 请求，最大重试2次
	resp, err = client.QueryTaskList(ctx, req, callopt.WithRetryPolicy(retry.BuildFailurePolicy(retry.NewFailurePolicy())))
	if err != nil {
		logger.CtxErrorf(ctx, "client.QueryTaskList failed, err = %v", err)
		return
	}
	return
}

func RunTask(ctx context.Context, req *miner_core.RunTaskReq) (resp *miner_core.RunTaskResp, err error) {
	// rpc不保证幂等，不重试
	resp, err = client.RunTask(ctx, req)
	if err != nil {
		logger.CtxErrorf(ctx, "client.RunTask failed, err = %v", err)
		return
	}
	return
}

func QueryTaskResultList(ctx context.Context, req *miner_core.QueryTaskResultListReq) (resp *miner_core.QueryTaskResultListResp, err error) {
	// 请求，最大重试2次
	resp, err = client.QueryTaskResultList(ctx, req, callopt.WithRetryPolicy(retry.BuildFailurePolicy(retry.NewFailurePolicy())))
	if err != nil {
		logger.CtxErrorf(ctx, "client.GetTaskResultList failed, err = %v", err)
		return
	}
	return
}
