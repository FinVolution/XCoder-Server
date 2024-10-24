package main

import (
	"context"
	"errors"
	_ "github.com/gogf/gf/contrib/drivers/mysql/v2"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gctx"
	"xcoder/internal/cmd"
	// 初始化 mongodb
	_ "xcoder/internal/dao/mongodb"
	_ "xcoder/internal/logic"
	_ "xcoder/utility/xapollo"
)

func main() {
	err := connData()
	if err != nil {
		g.Log().Errorf(context.Background(), "mysql conn error: %v", err)
		panic(err)
	}

	cmd.Main.Run(gctx.GetInitCtx())
}

// 检查 database 连接
func connData() error {
	err := g.DB().PingMaster()
	if err != nil {
		return errors.New("database 连接失败")
	}
	return nil
}
