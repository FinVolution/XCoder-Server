package common

import (
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/util/gconv"
	"xcoder/core/llms"
	"xcoder/internal/dao"
	"xcoder/internal/dao/mongodb"
	"xcoder/internal/model/entity"
	"xcoder/internal/model/input/codein"
	"xcoder/utility/xconcurrent"
	"xcoder/utility/xcontext"
)

func getModelNameAndVersionToDB(ctx context.Context, llmModelDefaultParams *codein.CodeGenerateLLMParams) (string, string) {
	return llmModelDefaultParams.ModelName, llmModelDefaultParams.ModelVersion
}

func CodeGenerateRecordInsertDB(ctx context.Context, in *codein.CodeGenerateRecordInsertDBRequest) error {
	modelNameToDB, modelVersionToDB := getModelNameAndVersionToDB(ctx, in.LLMParams)
	dataToMysql := &entity.PCodeGenRecords{
		Isactive:         1,
		DeleteToken:      "NA",
		Inserttime:       gtime.Now(),
		Updatetime:       gtime.Now(),
		GenerateUuid:     in.Input.GenerateUUID,
		GenerateType:     in.GenerateType,
		IsSingleLine:     gconv.Int(in.IsSingleLine),
		CreateUser:       in.Input.CreateUser,
		GitRepo:          in.Input.GitRepo,
		GitBranch:        in.Input.GitBranch,
		CodePath:         in.Input.CodePath,
		CodeLanguage:     in.Input.CodeLanguage,
		CodeTotalLines:   in.Input.CodeTotalLines,
		CrossfileCtxNums: in.CfcNums,
		IdeInfo:          in.Input.IDEInfo,
		StartCursorIdx:   gconv.Uint(in.Input.StartCursorIdx),
		PrefixCodeTokens: gconv.Uint(llms.CountTokens(modelNameToDB, in.CodeBeforeCursor)),
		SuffixCodeTokens: gconv.Uint(llms.CountTokens(modelNameToDB, in.CodeAfterCursor)),
		ModelName:        modelNameToDB,
		ModelVersion:     modelVersionToDB,
	}

	dataToMongoDB := &codein.CodeGenerateInsertMongoRequest{
		GenerateUUID:          in.Input.GenerateUUID,
		CodeBeforeCursor:      in.CodeBeforeCursor,
		CodeAfterCursor:       in.CodeAfterCursor,
		CodeBeforeWithContext: in.CodeBeforeWithContext,
	}

	p := xconcurrent.NewBase("codeGenerateRecordInsertDB")
	p.Compute(xcontext.WithProtect(ctx), func(ctx context.Context) error {
		g.Log().Infof(ctx, "AsyncCodeGenerateInsertDB start to insert to mysql")
		// 请求字段保存到 mysql
		err := dao.PCodeGenRecords.Create(ctx, dataToMysql)
		if err != nil {
			g.Log().Errorf(ctx, "Mysql CodeGenerateInsert failed: %v", err)
			return err
		}

		// 检索记录保存到 mysql
		if len(in.CodeGenSnippetModels) > 0 {
			err := dao.PCodeGenRetrievalSnippetMap.BatchCreate(ctx, in.CodeGenSnippetModels)
			if err != nil {
				g.Log().Errorf(ctx, "Mysql CodeGenerateRetrievalSnippetMapBatchCreate failed: %v", err)
				return err
			}
			g.Log().Infof(ctx, "Mysql CodeGenerateRetrievalSnippetMapBatchCreate success, "+
				"generateUUID: %s", in.Input.GenerateUUID)
		}

		g.Log().Infof(ctx, "AsyncCodeGenerateInsertDB start to insert to mongodb")
		// 保存代码到 mongodb
		_, err = mongodb.MDao.CodeGenerateInsert(ctx, dataToMongoDB)
		if err != nil {
			g.Log().Errorf(ctx, "Mongodb CodeGenerateInsert failed: %v", err)
			return err
		}

		g.Log().Infof(ctx, "AsyncCodeGenerateInsertDB finished")
		return nil
	})

	return nil
}
