package mongodb

import (
	"context"
	"github.com/gogf/gf/v2/frame/g"
)

var MDao *MongodbDao

func init() {
	// 初始化mongodb
	var err error
	MDao, err = New()
	if err != nil {
		g.Log().Errorf(context.Background(), "mongodb init error: %v", err)
		panic(err)
	}
	// 初始化mongodb完成
	g.Log().Infof(context.Background(), "mongodb init success, MDao: %v", MDao)

	defer MDao.Close()
}
