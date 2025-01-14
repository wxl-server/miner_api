package config

import (
	"context"

	"github.com/bytedance/gopkg/util/logger"
	"github.com/qcq1/common/env"
	"github.com/qcq1/common/render"
	"github.com/spf13/viper"
)

type AppConfig struct {
	Nacos NacosConfig `yaml:"nacos"`
}

type NacosConfig struct {
	Host      string `yaml:"host"`
	Port      uint64 `yaml:"port"`
	Namespace string `yaml:"namespace"`
	Username  string `yaml:"username"`
	Password  string `yaml:"password"`
}

func InitAppConfig(ctx context.Context) *AppConfig {
	vip := viper.New()
	vip.SetConfigFile("conf/" + env.GetEnv() + ".yaml")
	err := vip.ReadInConfig()
	if err != nil {
		logger.CtxErrorf(ctx, "[Init] init local config failed, err = %v", err)
		panic(err)
	}
	Config := &AppConfig{}
	err = vip.Unmarshal(Config)
	if err != nil {
		logger.CtxErrorf(ctx, "[Init] unmarshal config failed, err = %v", err)
		panic(err)
	}
	logger.CtxInfof(ctx, "[Init] init local config success, config = %v", render.Render(Config))
	return Config
}
