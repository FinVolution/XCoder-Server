package mongodb

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"google.golang.org/protobuf/types/known/emptypb"
	"time"
	"xcoder/internal/model/input/chatin"
)

func (d *MongodbDao) ChatMessageInsert(ctx context.Context, in *chatin.ChatMessageInsertMongoRequest) (
	*emptypb.Empty, error) {
	client, err := d.GetClient(ctx, d.DbName)
	if err != nil {
		return nil, err
	}

	currentTime := time.Now()
	year, month, _ := currentTime.Date()
	collectionName := fmt.Sprintf("chat_record_%d_%02d", year, int(month))
	collection := client.Database(d.DbName).Collection(collectionName)
	_, err = collection.InsertOne(ctx, bson.M{
		"insertTime":       currentTime,
		"updateTime":       currentTime,
		"conversationUUID": in.ConversationUUID,
		"selectedCode":     in.SelectedCode,
		"messages":         in.Messages,
	})
	if err != nil {
		return nil, err
	}

	return nil, nil
}

func (d *MongodbDao) ChatMessageUpdate(ctx context.Context, in *chatin.ChatMessageUpdateMongoRequest) (
	*emptypb.Empty, error) {
	client, err := d.GetClient(ctx, d.DbName)
	if err != nil {
		return nil, err
	}

	currentTime := time.Now()
	year, month, _ := currentTime.Date()
	collectionName := fmt.Sprintf("chat_record_%d_%02d", year, int(month))
	collection := client.Database(d.DbName).Collection(collectionName)
	filter := bson.M{"conversationUUID": in.ConversationUUID}
	updateFields := bson.M{"$set": bson.M{
		"response":       in.Response,
		"completionCode": in.CompletionCode,
		"updateTime":     currentTime,
	}}
	_, err = collection.UpdateMany(ctx, filter, updateFields)
	if err != nil {
		return nil, err
	}

	return nil, nil
}
