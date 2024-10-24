package xapollo

import (
	"github.com/spf13/viper"
	"xcoder/utility/xutils"
)

var viperInstance *viper.Viper

func NewViperConfig() *viper.Viper {
	viperInstance = viper.New()
	viperInstance.SetConfigType("json")
	cfgBaseDir := xutils.GetEnv("CFG_BASE_DIR", "/app/config")
	viperInstance.AddConfigPath(cfgBaseDir)
	return viperInstance
}
