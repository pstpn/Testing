package service_test

import (
	"context"
	"fmt"
	"time"

	"github.com/ozontech/allure-go/pkg/framework/provider"
	"github.com/ozontech/allure-go/pkg/framework/suite"

	"course/internal/model"
	"course/internal/service"
	"course/internal/service/utils"
	"course/internal/storage/mocks"
)

type CheckpointSuite struct {
	suite.Suite
}

func (c *CheckpointSuite) Test_Checkpoint_CreatePassage1(t provider.T) {
	t.Title("[CreatePassage] Incorrect checkpoint ID")
	t.Tags("checkpoint", "passage")
	t.Parallel()
	t.WithNewStep("Incorrect checkpoint ID", func(sCtx provider.StepCtx) {
		ctx := context.TODO()
		request := utils.CheckpointBuilder{}.
			WithCheckpointID(-1).
			WithDocumentID(1).
			WithPassageType(1).
			WithTime(time.Now().UTC()).ToCreateDTO()

		checkpointMockStorage := mocks.NewCheckpointStorage(t)
		checkpointMockStorage.
			On("CreatePassage", ctx, request).
			Return(nil, fmt.Errorf("incorrect checkpoint ID")).
			Once()

		sCtx.WithNewParameters("ctx", ctx, "request", request)

		passage, err := service.NewCheckpointService(utils.NewMockLogger(), checkpointMockStorage).CreatePassage(ctx, request)

		sCtx.Assert().Error(err)
		sCtx.Assert().Nil(passage)
	})
}

func (c *CheckpointSuite) Test_Checkpoint_CreatePassage2(t provider.T) {
	t.Title("[CreatePassage] Success")
	t.Tags("checkpoint", "passage")
	t.Parallel()
	t.WithNewStep("Success", func(sCtx provider.StepCtx) {
		ctx := context.TODO()
		request := utils.CheckpointBuilder{}.
			WithCheckpointID(1).
			WithDocumentID(1).
			WithPassageType(1).
			WithTime(time.Now().UTC()).
			ToCreateDTO()
		tm, _ := time.Parse(time.RFC3339, "2006-01-02T15:04:05Z")
		expPassage := &model.Passage{
			ID:           model.ToPassageID(123),
			CheckpointID: model.ToCheckpointID(1),
			DocumentID:   model.ToDocumentID(1),
			Type:         model.ToPassageTypeFromInt(1),
			Time:         &tm,
		}

		checkpointMockStorage := mocks.NewCheckpointStorage(t)
		checkpointMockStorage.
			On("CreatePassage", ctx, request).
			Return(expPassage, nil).
			Once()

		sCtx.WithNewParameters("ctx", ctx, "request", request)

		passage, err := service.NewCheckpointService(utils.NewMockLogger(), checkpointMockStorage).CreatePassage(ctx, request)

		sCtx.Assert().NoError(err)
		sCtx.Assert().NotNil(passage)
		sCtx.Assert().Equal(expPassage, passage)
	})
}

func (c *CheckpointSuite) Test_Checkpoint_ListPassages1(t provider.T) {
	t.Title("[ListPassages] Incorrect document ID")
	t.Tags("checkpoint", "passage")
	t.Parallel()
	t.WithNewStep("Incorrect document ID", func(sCtx provider.StepCtx) {
		ctx := context.TODO()
		request := utils.CheckpointBuilder{}.
			WithDocumentID(-1).
			ToListDTO()

		checkpointMockStorage := mocks.NewCheckpointStorage(t)
		checkpointMockStorage.
			On("ListPassages", ctx, request).
			Return([]*model.Passage{}, nil).
			Once()

		sCtx.WithNewParameters("ctx", ctx, "request", request)

		passages, err := service.NewCheckpointService(utils.NewMockLogger(), checkpointMockStorage).ListPassages(ctx, request)

		sCtx.Assert().NoError(err)
		sCtx.Assert().Empty(passages)
	})
}

func (c *CheckpointSuite) Test_Checkpoint_ListPassages2(t provider.T) {
	t.Title("[ListPassages] Success")
	t.Tags("checkpoint", "passage")
	t.Parallel()
	t.WithNewStep("Success", func(sCtx provider.StepCtx) {
		ctx := context.TODO()
		request := utils.CheckpointBuilder{}.
			WithDocumentID(1).
			ToListDTO()
		tm, _ := time.Parse(time.RFC3339, "2006-01-02T15:04:05Z")
		expPassages := []*model.Passage{{
			ID:           model.ToPassageID(123),
			CheckpointID: model.ToCheckpointID(1),
			DocumentID:   model.ToDocumentID(1),
			Type:         model.ToPassageTypeFromInt(1),
			Time:         &tm,
		}}

		checkpointMockStorage := mocks.NewCheckpointStorage(t)
		checkpointMockStorage.
			On("ListPassages", ctx, request).
			Return(expPassages, nil).
			Once()

		sCtx.WithNewParameters("ctx", ctx, "request", request)

		passages, err := service.NewCheckpointService(utils.NewMockLogger(), checkpointMockStorage).ListPassages(ctx, request)

		sCtx.Assert().NoError(err)
		sCtx.Assert().NotNil(passages)
		sCtx.Assert().Len(passages, len(expPassages))
		sCtx.Assert().Equal(expPassages[0].ID, passages[0].ID)
		sCtx.Assert().Equal(expPassages[0].CheckpointID, passages[0].CheckpointID)
		sCtx.Assert().Equal(expPassages[0].DocumentID, passages[0].DocumentID)
		sCtx.Assert().Equal(expPassages[0].Type, passages[0].Type)
		sCtx.Assert().Equal(expPassages[0].Time, passages[0].Time)
	})
}
