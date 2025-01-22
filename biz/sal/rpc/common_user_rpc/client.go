package common_user_rpc

import (
	"context"

	"github.com/bytedance/gopkg/util/logger"
	"github.com/wxl-server/common/wxl_cluster"
	"github.com/wxl-server/idl_gen/kitex_gen/common_user"
	"github.com/wxl-server/idl_gen/kitex_gen/common_user/commonuser"
)

var client = wxl_cluster.NewClient(commonuser.NewClient, "common_user")

func SignUp(ctx context.Context, req *common_user.SignUpReq) (resp *common_user.SignUpResp, err error) {
	// rpc不保证幂等，不重试
	resp, err = client.SignUp(ctx, req)
	if err != nil {
		logger.CtxErrorf(ctx, "client.SignUp failed, err = %v", err)
		return
	}
	return
}

func UpdatePassword(ctx context.Context, req *common_user.UpdatePasswordReq) (resp *common_user.UpdatePasswordResp, err error) {
	// rpc不保证幂等，不重试
	resp, err = client.UpdatePassword(ctx, req)
	if err != nil {
		logger.CtxErrorf(ctx, "client.UpdatePassword failed, err = %v", err)
		return
	}
	return
}
