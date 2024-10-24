package xapollo

import (
	"context"
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/gogf/gf/v2/frame/g"
	"xcoder/internal/model/input/codein"
)

const (
	CodeGenerateLLMParamsConfigPath = "xcoder_code_generate_llm_params_conf"
)

var CodeGenerateLLMParamsConfigInfo = new(CodeGenerateLLMParamsConfig)

type CodeGenerateLLMParamsConfig struct {
	Config        map[string]*codein.CodeGenerateLLMParams `json:"config"`
	SelectedModel string                                   `json:"selectedModel"`
}

func init() {
	viperCfgC := NewViperConfig()
	viperCfgC.SetConfigName(CodeGenerateLLMParamsConfigPath)
	if err := viperCfgC.ReadInConfig(); err != nil {
		g.Log().Errorf(context.Background(), "read config file %s error: %s", CodeGenerateLLMParamsConfigPath, err.Error())
	}
	if err := viperCfgC.Unmarshal(CodeGenerateLLMParamsConfigInfo); err != nil {
		g.Log().Errorf(context.Background(), "unmarshal config file %s error: %s", CodeGenerateLLMParamsConfigPath, err.Error())
		return
	}
	g.Log().Infof(context.Background(), "load config file %s success", CodeGenerateLLMParamsConfigPath)

	viperCfgC.WatchConfig()
	viperCfgC.OnConfigChange(func(in fsnotify.Event) {
		if err := viperCfgC.Unmarshal(CodeGenerateLLMParamsConfigInfo); err != nil {
			g.Log().Errorf(context.Background(), "unmarshal config file %s error: %s", CodeGenerateLLMParamsConfigPath, err.Error())
			return
		}
		g.Log().Infof(context.Background(), "load config file %s success", CodeGenerateLLMParamsConfigPath)
	})
}

func GetSelectedModel() string {
	return CodeGenerateLLMParamsConfigInfo.SelectedModel
}

func GetCodeGenerateLLmParams() (*codein.CodeGenerateLLMParams, error) {
	selectedModel := GetSelectedModel()
	if selectedModel == "" {
		return nil, fmt.Errorf("必须选择一个模型使用")
	}
	if CodeGenerateLLMParamsConfigInfo.Config == nil {
		return nil, fmt.Errorf("模型配置信息为空")
	}

	if params, ok := CodeGenerateLLMParamsConfigInfo.Config[selectedModel]; ok {
		return params, nil
	}

	return nil, fmt.Errorf("模型: %s 配置信息不存在", CodeGenerateLLMParamsConfigInfo.SelectedModel)
}
