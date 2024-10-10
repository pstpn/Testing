//go:build unit

package postgres_test

import (
	"context"
	"fmt"
	"time"

	"github.com/ozontech/allure-go/pkg/framework/provider"
	"github.com/ozontech/allure-go/pkg/framework/suite"

	"course/internal/model"
	"course/internal/storage/mocks"
	"course/internal/storage/utils"
)

type CheckpointStorageSuite struct {
	suite.Suite

	checkpointMockStorage mocks.CheckpointStorage
}

func (c *CheckpointStorageSuite) BeforeAll(t provider.T) {
	t.Title("Init checkpoint mock storage")
	c.checkpointMockStorage = *mocks.NewCheckpointStorage(t)
	t.Tags("fixture", "checkpoint", "passage")
}

func (c *CheckpointStorageSuite) Test_CheckpointStorage_CreateCheckpoint(t provider.T) {
	t.Title("[CreateCheckpoint] Create checkpoint test")
	t.Tags("storage", "postgres", "checkpoint")
	t.Parallel()
	t.WithNewStep("Create checkpoint test", func(sCtx provider.StepCtx) {
		sCtx.WithNewParameters("Start", fmt.Sprintf("%s", time.Now()))
		ctx := context.TODO()
		request := utils.CheckpointBuilder{}.
			WithPhoneNumber("123").
			ToCreateCheckpointDTO()
		expCheckpoint := &model.Checkpoint{
			ID:          model.ToCheckpointID(1),
			PhoneNumber: "123",
		}

		c.checkpointMockStorage.
			On("CreateCheckpoint", ctx, request).
			Return(expCheckpoint, nil).
			Once()

		sCtx.WithNewParameters("ctx", ctx, "request", request)

		checkpoint, err := c.checkpointMockStorage.CreateCheckpoint(ctx, request)

		sCtx.Assert().NoError(err)
		sCtx.Assert().NotNil(checkpoint)
		sCtx.Assert().Equal(expCheckpoint, checkpoint)
	})
}

func (c *CheckpointStorageSuite) Test_CheckpointStorage_CreatePassage(t provider.T) {
	t.Title("[CreatePassage] Create passage test")
	t.Tags("storage", "postgres", "checkpoint", "passage")
	t.Parallel()
	t.WithNewStep("Create passage test", func(sCtx provider.StepCtx) {
		sCtx.WithNewParameters("Start", fmt.Sprintf("%s", time.Now()))
		ctx := context.TODO()
		request := utils.CheckpointBuilder{}.
			WithCheckpointID(1).
			WithDocumentID(1).
			WithPassageType(1).
			WithTime(time.Now().UTC()).
			ToCreateDTO()
		expPassage := &model.Passage{
			ID:           model.ToPassageID(1),
			CheckpointID: model.ToCheckpointID(1),
			DocumentID:   model.ToDocumentID(1),
			Type:         model.ToPassageTypeFromInt(1),
			Time:         nil,
		}

		c.checkpointMockStorage.
			On("CreatePassage", ctx, request).
			Return(expPassage, nil).
			Once()

		sCtx.WithNewParameters("ctx", ctx, "request", request)

		passage, err := c.checkpointMockStorage.CreatePassage(ctx, request)

		sCtx.Assert().NoError(err)
		sCtx.Assert().NotNil(passage)
		sCtx.Assert().NotNil(passage.ID)
		sCtx.Assert().Equal(expPassage, passage)
	})
}

func (c *CheckpointStorageSuite) Test_CheckpointStorage_GetCheckpoint(t provider.T) {
	t.Title("[GetCheckpoint] Get checkpoint test")
	t.Tags("storage", "postgres", "checkpoint")
	t.Parallel()
	t.WithNewStep("Get checkpoint test", func(sCtx provider.StepCtx) {
		sCtx.WithNewParameters("Start", fmt.Sprintf("%s", time.Now()))
		ctx := context.TODO()
		request := utils.CheckpointBuilder{}.
			WithCheckpointID(1).
			ToGetCheckpointDTO()
		expCheckpoint := &model.Checkpoint{
			ID:          model.ToCheckpointID(1),
			PhoneNumber: "123",
		}

		c.checkpointMockStorage.
			On("GetCheckpoint", ctx, request).
			Return(expCheckpoint, nil).
			Once()

		sCtx.WithNewParameters("ctx", ctx, "request", request)

		checkpoint, err := c.checkpointMockStorage.GetCheckpoint(ctx, request)

		sCtx.Assert().NoError(err)
		sCtx.Assert().NotNil(checkpoint)
		sCtx.Assert().Equal(expCheckpoint, checkpoint)
	})
}

func (c *CheckpointStorageSuite) Test_CheckpointStorage_GetPassage(t provider.T) {
	t.Title("[GetPassage] Get passage test")
	t.Tags("storage", "postgres", "checkpoint", "passage")
	t.Parallel()
	t.WithNewStep("Get passage test", func(sCtx provider.StepCtx) {
		sCtx.WithNewParameters("Start", fmt.Sprintf("%s", time.Now()))
		ctx := context.TODO()
		request := utils.CheckpointBuilder{}.
			WithPassageID(1).
			ToGetDTO()

		expPassage := &model.Passage{
			ID:           model.ToPassageID(1),
			CheckpointID: model.ToCheckpointID(1),
			DocumentID:   model.ToDocumentID(1),
			Type:         model.ToPassageTypeFromInt(1),
			Time:         nil,
		}

		c.checkpointMockStorage.
			On("GetPassage", ctx, request).
			Return(expPassage, nil).
			Once()

		sCtx.WithNewParameters("ctx", ctx, "request", request)

		passage, err := c.checkpointMockStorage.GetPassage(ctx, request)

		sCtx.Assert().NoError(err)
		sCtx.Assert().NotNil(passage)
		sCtx.Assert().Equal(expPassage, passage)
	})
}

func (c *CheckpointStorageSuite) Test_CheckpointStorage_ListPassages(t provider.T) {
	t.Title("[ListPassages] List passages test")
	t.Tags("storage", "postgres", "checkpoint", "passage")
	t.Parallel()
	t.WithNewStep("List passages test", func(sCtx provider.StepCtx) {
		sCtx.WithNewParameters("Start", fmt.Sprintf("%s", time.Now()))
		ctx := context.TODO()
		expPassages := []*model.Passage{
			{
				ID:           model.ToPassageID(1),
				CheckpointID: model.ToCheckpointID(1),
				DocumentID:   model.ToDocumentID(1),
				Type:         model.ToPassageTypeFromInt(1),
				Time:         nil,
			},
			{
				ID:           model.ToPassageID(2),
				CheckpointID: model.ToCheckpointID(1),
				DocumentID:   model.ToDocumentID(1),
				Type:         model.ToPassageTypeFromInt(1),
				Time:         nil,
			},
		}
		request := utils.CheckpointBuilder{}.
			WithDocumentID(1).
			ToListDTO()

		c.checkpointMockStorage.
			On("ListPassages", ctx, request).
			Return(expPassages, nil).
			Once()

		sCtx.WithNewParameters("ctx", ctx, "request", request)

		passages, err := c.checkpointMockStorage.ListPassages(ctx, request)

		sCtx.Assert().NoError(err)
		sCtx.Assert().NotNil(passages)
		sCtx.Assert().Len(passages, len(expPassages))
		sCtx.Assert().Equal(passages[0], expPassages[0])
		sCtx.Assert().Equal(passages[1], expPassages[1])
	})
}

func (c *CheckpointStorageSuite) Test_CheckpointStorage_DeleteCheckpoint(t provider.T) {
	t.Title("[DeleteCheckpoint] Delete checkpoint test")
	t.Tags("storage", "postgres", "checkpoint")
	t.Parallel()
	t.WithNewStep("Delete checkpoint test", func(sCtx provider.StepCtx) {
		sCtx.WithNewParameters("Start", fmt.Sprintf("%s", time.Now()))
		ctx := context.TODO()
		request := utils.CheckpointBuilder{}.
			WithCheckpointID(1).
			ToDeleteCheckpointDTO()

		c.checkpointMockStorage.
			On("DeleteCheckpoint", ctx, request).
			Return(nil).
			Once()

		sCtx.WithNewParameters("ctx", ctx, "request", request)

		err := c.checkpointMockStorage.DeleteCheckpoint(ctx, request)

		sCtx.Assert().NoError(err)
	})
}

func (c *CheckpointStorageSuite) Test_CheckpointStorage_DeletePassage(t provider.T) {
	t.Title("[DeletePassage] Delete passage test")
	t.Tags("storage", "postgres", "checkpoint", "passage")
	t.Parallel()
	t.WithNewStep("Delete passage test", func(sCtx provider.StepCtx) {
		sCtx.WithNewParameters("Start", fmt.Sprintf("%s", time.Now()))
		ctx := context.TODO()
		request := utils.CheckpointBuilder{}.
			WithPassageID(1).
			ToDeleteDTO()

		sCtx.WithNewParameters("ctx", ctx, "request", request)

		c.checkpointMockStorage.
			On("DeletePassage", ctx, request).
			Return(nil).
			Once()

		err := c.checkpointMockStorage.DeletePassage(ctx, request)

		sCtx.Assert().NoError(err)
	})
}
