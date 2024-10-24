// =================================================================================
// This is auto-generated by GoFrame CLI tool only once. Fill this file as you wish.
// =================================================================================

package dao

import (
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"xcoder/internal/dao/internal"
	"xcoder/internal/model/entity"
)

// internalPCodeGenRetrievalSnippetMapDao is internal type for wrapping internal DAO implements.
type internalPCodeGenRetrievalSnippetMapDao = *internal.PCodeGenRetrievalSnippetMapDao

// pCodeGenRetrievalSnippetMapDao is the data access object for table p_code_gen_retrieval_snippet_map.
// You can define custom methods on it to extend its functionality as you wish.
type pCodeGenRetrievalSnippetMapDao struct {
	internalPCodeGenRetrievalSnippetMapDao
}

var (
	// PCodeGenRetrievalSnippetMap is globally public accessible object for table p_code_gen_retrieval_snippet_map operations.
	PCodeGenRetrievalSnippetMap = pCodeGenRetrievalSnippetMapDao{
		internal.NewPCodeGenRetrievalSnippetMapDao(),
	}
)

func (dao *pCodeGenRetrievalSnippetMapDao) Create(ctx context.Context, data *entity.PCodeGenRetrievalSnippetMap) error {
	_, err := g.Model(PCodeGenRetrievalSnippetMap.Table()).Data(data).Insert()
	if err != nil {
		return err
	}
	return nil
}

func (dao *pCodeGenRetrievalSnippetMapDao) BatchCreate(ctx context.Context, datas []*entity.PCodeGenRetrievalSnippetMap) error {
	_, err := g.Model(PCodeGenRetrievalSnippetMap.Table()).Data(datas).Insert()
	if err != nil {
		return err
	}
	return nil

}
