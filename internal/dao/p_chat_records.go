// =================================================================================
// This is auto-generated by GoFrame CLI tool only once. Fill this file as you wish.
// =================================================================================

package dao

import (
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/gconv"
	"xcoder/internal/dao/internal"
	"xcoder/internal/model/entity"
)

// internalPChatRecordsDao is internal type for wrapping internal DAO implements.
type internalPChatRecordsDao = *internal.PChatRecordsDao

// pChatRecordsDao is the data access object for table p_chat_records.
// You can define custom methods on it to extend its functionality as you wish.
type pChatRecordsDao struct {
	internalPChatRecordsDao
}

var (
	// PChatRecords is globally public accessible object for table p_chat_records operations.
	PChatRecords = pChatRecordsDao{
		internal.NewPChatRecordsDao(),
	}
)

func (dao *pChatRecordsDao) Create(ctx context.Context, data *entity.PChatRecords) error {
	_, err := g.Model(PChatRecords.Table()).Data(data).Insert()
	if err != nil {
		return err
	}
	return nil
}

func (dao *pChatRecordsDao) GetOneByConversationUuid(ctx context.Context, conversationUuid string) (
	result *entity.PChatRecords, err error) {
	data, err := g.Model(PChatRecords.Table()).Where(map[string]interface{}{
		"isactive": 1, "conversation_uuid": conversationUuid,
	}).One()
	if err != nil {
		return nil, err
	}

	err = gconv.Struct(data, &result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (dao *pChatRecordsDao) UpdateByConversationUuid(ctx context.Context, conversationUuid string, updated map[string]interface{}) error {
	_, err := g.Model(PChatRecords.Table()).Where(map[string]interface{}{
		"isactive": 1, "conversation_uuid": conversationUuid,
	}).Data(updated).Update()
	if err != nil {
		return err
	}

	return nil
}