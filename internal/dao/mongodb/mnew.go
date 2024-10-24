package mongodb

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"xcoder/utility/xmongo"
)

type MongodbDao struct {
	DbName  string
	mongodb *xmongo.Mongodb
}

func New() (*MongodbDao, error) {
	mongodbConfig, err := xmongo.NewFromFile()
	if err != nil {
		return nil, err
	}

	return &MongodbDao{
		DbName:  mongodbConfig.DBName,
		mongodb: mongodbConfig,
	}, nil
}

func (d *MongodbDao) GetClient(ctx context.Context, dbName string) (*mongo.Client, error) {
	return d.mongodb.GetClient(ctx, dbName)
}

func (d *MongodbDao) GetDBName(ctx context.Context) string {
	return d.mongodb.GetDBName(ctx)
}

func (d *MongodbDao) Close() {
	d.mongodb.Close()
}
