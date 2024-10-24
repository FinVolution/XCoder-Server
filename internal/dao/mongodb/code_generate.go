package mongodb

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"google.golang.org/protobuf/types/known/emptypb"
	"time"
	"xcoder/internal/model/input/codein"
)

func (d *MongodbDao) CodeGenerateInsert(ctx context.Context, in *codein.CodeGenerateInsertMongoRequest) (
	*emptypb.Empty, error) {
	client, err := d.GetClient(ctx, d.DbName)
	if err != nil {
		return nil, err
	}

	currentTime := time.Now()
	year, month, _ := currentTime.Date()
	collectionName := fmt.Sprintf("code_generate_%d_%02d", year, int(month))
	collection := client.Database(d.DbName).Collection(collectionName)
	_, err = collection.InsertOne(ctx, bson.M{
		"insertTime":            currentTime,
		"updateTime":            currentTime,
		"generateUUID":          in.GenerateUUID,
		"codeBeforeCursor":      in.CodeBeforeCursor,
		"codeAfterCursor":       in.CodeAfterCursor,
		"codeBeforeWithContext": in.CodeBeforeWithContext,
		"completionCode":        "",
	})
	if err != nil {
		return nil, err
	}

	return nil, nil
}

func (d *MongodbDao) CodeGenerateUpdate(ctx context.Context, in *codein.CodeGenerateUpdateMongoRequest) (
	*emptypb.Empty, error) {
	client, err := d.GetClient(ctx, d.DbName)
	if err != nil {
		return nil, err
	}

	currentTime := time.Now()
	year, month, _ := currentTime.Date()
	collectionName := fmt.Sprintf("code_generate_%d_%02d", year, int(month))
	collection := client.Database(d.DbName).Collection(collectionName)
	filter := bson.M{"generateUUID": in.GenerateUUID}
	updateFields := bson.M{"$set": bson.M{
		"completionCode":    in.CompletionCode,
		"rawCompletionCode": in.RawCompletionCode,
		"updateTime":        currentTime,
	}}
	_, err = collection.UpdateMany(ctx, filter, updateFields)
	if err != nil {
		return nil, err
	}

	return nil, nil
}
