package xapollo

import (
	"context"
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/gogf/gf/v2/frame/g"
	"xcoder/internal/model/input/codein"
)

const (
	RetrievalCrossFileContextParamsConfigPath = "xcoder_code_generate_retrieval_cfc_params_conf"
)

var RetrievalCrossFileContextParamsConfigInfo = new(RetrievalCrossFileContextParamsConfig)

type RetrievalCrossFileContextParamsConfig struct {
	Config *codein.RetrievalCrossFileContextParams `json:"config"`
}

func init() {
	viperCfg := NewViperConfig()
	viperCfg.SetConfigName(RetrievalCrossFileContextParamsConfigPath)
	if err := viperCfg.ReadInConfig(); err != nil {
		g.Log().Errorf(
			context.Background(),
			"read config file %s error: %s", RetrievalCrossFileContextParamsConfigPath,
			err.Error(),
		)
	}
	if err := viperCfg.Unmarshal(RetrievalCrossFileContextParamsConfigInfo); err != nil {
		g.Log().Errorf(
			context.Background(),
			"unmarshal config file %s error: %s", RetrievalCrossFileContextParamsConfigPath,
			err.Error(),
		)
		return
	}
	g.Log().Infof(context.Background(), "load config file %s success", RetrievalCrossFileContextParamsConfigPath)

	viperCfg.WatchConfig()
	viperCfg.OnConfigChange(func(in fsnotify.Event) {
		if err := viperCfg.Unmarshal(RetrievalCrossFileContextParamsConfigInfo); err != nil {
			g.Log().Errorf(
				context.Background(),
				"unmarshal config file %s error: %s", RetrievalCrossFileContextParamsConfigPath,
				err.Error(),
			)
			return
		}
		g.Log().Infof(context.Background(), "load config file %s success", RetrievalCrossFileContextParamsConfigPath)
	})
}

func GetRetrievalCrossFileContextParams() (*codein.RetrievalCrossFileContextParams, error) {
	if RetrievalCrossFileContextParamsConfigInfo.Config == nil {
		return nil, fmt.Errorf("模型 cfc 配置信息为空")
	}

	return RetrievalCrossFileContextParamsConfigInfo.Config, nil
}
