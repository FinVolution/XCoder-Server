package unit_test_test

import (
	"context"
	_ "github.com/gogf/gf/v2/frame/g"
	"github.com/stretchr/testify/mock"
	"xcoder/internal/model/input/chatin"
)

// MockService is a mock implementation of the service.Chat interface
type MockService struct {
	mock.Mock
}

func (m *MockService) SseGenerate(ctx context.Context, req *chatin.ChatSseGenerateReq, resp chan<- *chatin.ChatResult) error {
	args := m.Called(ctx, req, resp)
	return args.Error(0)
}

//func TestControllerV1_SseGenerate(t *testing.T) {
//	mockService := new(MockService)
//	service.Chat = func() service.Chat {
//		return mockService
//	}
//
//	controller := NewV1()
//	ctx := context.Background()
//	req := &v1.ChatSseGenerateReq{}
//
//	t.Run("successful response", func(t *testing.T) {
//		respChan := make(chan *chatin.ChatResult, 1)
//		mockService.On("SseGenerate", ctx, mock.Anything, mock.Anything).Run(func(args mock.Arguments) {
//			resp := args.Get(2).(chan<- *chatin.ChatResult)
//			resp <- &chatin.ChatResult{Data: "success"}
//		}).Return(nil)
//
//		common.GetStreamingChatReq = func(ctx context.Context) *common.StreamingChatReq {
//			return &common.StreamingChatReq{}
//		}
//
//		common.ParseStreamingChatResp = func(ctx context.Context, r *common.StreamingChatReq, resp <-chan *chatin.ChatResult) {
//			result := <-resp
//			assert.Equal(t, "success", result.Data)
//		}
//
//		_, err := controller.SseGenerate(ctx, req)
//		assert.NoError(t, err)
//	})
//
//	t.Run("error in SseGenerate", func(t *testing.T) {
//		respChan := make(chan *chatin.ChatResult, 1)
//		mockService.On("SseGenerate", ctx, mock.Anything, mock.Anything).Return(assert.AnError)
//
//		common.GetStreamingChatReq = func(ctx context.Context) *common.StreamingChatReq {
//			return &common.StreamingChatReq{}
//		}
//
//		common.ParseStreamingChatResp = func(ctx context.Context, r *common.StreamingChatReq, resp <-chan *chatin.ChatResult) {
//			result := <-resp
//			assert.Error(t, result.Error)
//		}
//
//		_, err := controller.SseGenerate(ctx, req)
//		assert.NoError(t, err)
//	})
//
//	t.Run("panic recovery", func(t *testing.T) {
//		respChan := make(chan *chatin.ChatResult, 1)
//		mockService.On("SseGenerate", ctx, mock.Anything, mock.Anything).Run(func(args mock.Arguments) {
//			panic("test panic")
//		}).Return(nil)
//
//		common.GetStreamingChatReq = func(ctx context.Context) *common.StreamingChatReq {
//			return &common.StreamingChatReq{}
//		}
//
//		common.ParseStreamingChatResp = func(ctx context.Context, r *common.StreamingChatReq, resp <-chan *chatin.ChatResult) {
//			select {
//			case <-resp:
//			case <-time.After(1 * time.Second):
//				t.Fatal("timeout waiting for response")
//			}
//		}
//
//		_, err := controller.SseGenerate(ctx, req)
//		assert.NoError(t, err)
//	})
//}
