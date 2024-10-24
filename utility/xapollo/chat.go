package xapollo

import (
	"context"
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/gogf/gf/v2/frame/g"
	"xcoder/internal/model/input/chatin"
)

const (
	ChatLLMParamsConfigPath = "xcoder_chat_llm_params_conf"
)

var ChatLLMParamsConfigInfo = new(chatin.ChatLLMParamsConfig)

func init() {
	viperCfg := NewViperConfig()
	viperCfg.SetConfigName(ChatLLMParamsConfigPath)
	if err := viperCfg.ReadInConfig(); err != nil {
		g.Log().Errorf(context.Background(), "read config file %s error: %s", ChatLLMParamsConfigPath, err.Error())
	}
	if err := viperCfg.Unmarshal(ChatLLMParamsConfigInfo); err != nil {
		g.Log().Errorf(context.Background(), "unmarshal config file %s error: %s", ChatLLMParamsConfigPath, err.Error())
		return
	}
	g.Log().Infof(context.Background(), "load config file %s success", ChatLLMParamsConfigPath)
	g.Log().Infof(context.Background(), "ChatLLMParamsConfigInfo params: %+v", ChatLLMParamsConfigInfo.Config.LLMParams)

	viperCfg.WatchConfig()
	viperCfg.OnConfigChange(func(in fsnotify.Event) {
		if err := viperCfg.Unmarshal(ChatLLMParamsConfigInfo); err != nil {
			g.Log().Errorf(context.Background(), "unmarshal config file %s error: %s", ChatLLMParamsConfigPath, err.Error())
			return
		}
		g.Log().Infof(context.Background(), "load config file %s success", ChatLLMParamsConfigPath)
	})
}

func GetChatLLMParams() (*chatin.ChatLLMParamsConfig, error) {
	if ChatLLMParamsConfigInfo.Config == nil {
		return nil, fmt.Errorf("ChatLLMParamsConfigInfo.Config is nil")
	}

	return ChatLLMParamsConfigInfo, nil
}
