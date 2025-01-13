package rpc

import (
	"miner_api/biz/sal/rpc/miner_miner_core"

	"github.com/qcq1/common/env"
)

var Router = map[string]map[env.Env]string{
	miner_miner_core.ServiceName: {
		env.Prod: "miner-core.miner.svc.cluster.local:8888",
		env.Boe:  "127.0.0.1:8889",
	},
}
