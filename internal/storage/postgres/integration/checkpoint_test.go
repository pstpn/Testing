//go:build integration

package postgres_test

import (
	"context"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/ozontech/allure-go/pkg/framework/provider"
	"github.com/ozontech/allure-go/pkg/framework/suite"

	"course/internal/model"
	"course/internal/service/dto"
	"course/internal/storage"
	"course/internal/storage/utils"
)

type CheckpointStorageSuite struct {
	suite.Suite

	checkpointStorage storage.CheckpointStorage
	checkpointID      int64
	documentID        int64
}

func (c *CheckpointStorageSuite) Test_IntegrationCheckpointStorage_CreateCheckpoint(t provider.T) {
	t.Title("[INT CreateCheckpoint] Create checkpoint test")
	t.Tags("storage", "postgres", "checkpoint")
	t.Parallel()
	t.WithNewStep("Create checkpoint test", func(sCtx provider.StepCtx) {
		ctx := context.TODO()
		request := utils.CheckpointBuilder{}.
			WithPhoneNumber("123").
			ToCreateCheckpointDTO()

		sCtx.WithNewParameters("ctx", ctx, "request", request)

		checkpoint, err := c.checkpointStorage.CreateCheckpoint(ctx, request)

		sCtx.Assert().NoError(err)
		sCtx.Assert().NotNil(checkpoint)
		sCtx.Assert().NotNil(checkpoint.ID)
		sCtx.Assert().Equal(request.PhoneNumber, checkpoint.PhoneNumber)

		err = c.checkpointStorage.DeleteCheckpoint(context.TODO(), &dto.DeleteCheckpointRequest{CheckpointID: checkpoint.ID.Int()})
		sCtx.Assert().NoError(err)
	})
}

func (c *CheckpointStorageSuite) Test_IntegrationCheckpointStorage_CreatePassage(t provider.T) {
	t.Title("[INT CreatePassage] Create passage test")
	t.Tags("storage", "postgres", "checkpoint", "passage")
	t.Parallel()
	t.WithNewStep("Create passage test", func(sCtx provider.StepCtx) {
		ctx := context.TODO()
		request := utils.CheckpointBuilder{}.
			WithCheckpointID(c.checkpointID).
			WithDocumentID(c.documentID).
			WithPassageType(1).
			WithTime(time.Now().UTC()).
			ToCreateDTO()

		sCtx.WithNewParameters("ctx", ctx, "request", request)

		passage, err := c.checkpointStorage.CreatePassage(ctx, request)

		sCtx.Assert().NoError(err)
		sCtx.Assert().NotNil(passage)
		sCtx.Assert().NotNil(passage.ID)
		sCtx.Assert().Equal(request.CheckpointID, passage.CheckpointID.Int())
		sCtx.Assert().Equal(request.DocumentID, passage.DocumentID.Int())
		sCtx.Assert().Equal(request.Type, passage.Type.Int())
		sCtx.Assert().Equal(request.Time, passage.Time)

		err = c.checkpointStorage.DeletePassage(context.TODO(), &dto.DeletePassageRequest{PassageID: passage.ID.Int()})
		sCtx.Assert().NoError(err)
	})
}

func (c *CheckpointStorageSuite) Test_IntegrationCheckpointStorage_GetCheckpoint(t provider.T) {
	t.Title("[GetCheckpoint] Get checkpoint test")
	t.Tags("storage", "postgres", "checkpoint")
	t.Parallel()
	t.WithNewStep("Get checkpoint test", func(sCtx provider.StepCtx) {
		ctx := context.TODO()
		expCheckpoint, err := c.checkpointStorage.CreateCheckpoint(
			ctx,
			utils.CheckpointBuilder{}.
				WithPhoneNumber("5473").
				ToCreateCheckpointDTO(),
		)
		sCtx.Assert().NoError(err)
		sCtx.Assert().NotNil(expCheckpoint)
		request := utils.CheckpointBuilder{}.
			WithCheckpointID(expCheckpoint.ID.Int()).
			ToGetCheckpointDTO()

		sCtx.WithNewParameters("ctx", ctx, "request", request)

		checkpoint, err := c.checkpointStorage.GetCheckpoint(ctx, request)

		sCtx.Assert().NoError(err)
		sCtx.Assert().NotNil(checkpoint)
		sCtx.Assert().Equal(expCheckpoint.ID, checkpoint.ID)
		sCtx.Assert().Equal(expCheckpoint.PhoneNumber, checkpoint.PhoneNumber)

		err = c.checkpointStorage.DeleteCheckpoint(context.TODO(), &dto.DeleteCheckpointRequest{CheckpointID: checkpoint.ID.Int()})
		sCtx.Assert().NoError(err)
	})
}

func (c *CheckpointStorageSuite) Test_IntegrationCheckpointStorage_GetPassage(t provider.T) {
	t.Title("[GetPassage] Get passage test")
	t.Tags("storage", "postgres", "checkpoint", "passage")
	t.Parallel()
	t.WithNewStep("Get passage test", func(sCtx provider.StepCtx) {
		ctx := context.TODO()
		expPassage, err := c.checkpointStorage.CreatePassage(
			ctx,
			utils.CheckpointBuilder{}.
				WithCheckpointID(c.checkpointID).
				WithDocumentID(c.documentID).
				WithPassageType(1).
				WithTime(time.Now().UTC()).
				ToCreateDTO(),
		)
		sCtx.Assert().NoError(err)
		sCtx.Assert().NotNil(expPassage)
		request := utils.CheckpointBuilder{}.
			WithPassageID(expPassage.ID.Int()).
			ToGetDTO()

		sCtx.WithNewParameters("ctx", ctx, "request", request)

		passage, err := c.checkpointStorage.GetPassage(ctx, request)

		sCtx.Assert().NoError(err)
		sCtx.Assert().NotNil(passage)
		sCtx.Assert().Equal(expPassage.ID, passage.ID)
		sCtx.Assert().Equal(expPassage.CheckpointID, passage.CheckpointID)
		sCtx.Assert().Equal(expPassage.DocumentID, passage.DocumentID)
		sCtx.Assert().Equal(expPassage.Type, passage.Type)
		sCtx.Assert().Equal(expPassage.Time, passage.Time)

		err = c.checkpointStorage.DeletePassage(context.TODO(), &dto.DeletePassageRequest{PassageID: passage.ID.Int()})
		sCtx.Assert().NoError(err)
	})
}

func (c *CheckpointStorageSuite) Test_IntegrationCheckpointStorage_ListPassages(t provider.T) {
	t.Title("[ListPassages] List passages test")
	t.Tags("storage", "postgres", "checkpoint", "passage")
	t.Parallel()
	t.WithNewStep("List passages test", func(sCtx provider.StepCtx) {
		ctx := context.TODO()
		var expPassages []*model.Passage
		for range 3 {
			expPassage, err := c.checkpointStorage.CreatePassage(
				ctx,
				utils.CheckpointBuilder{}.
					WithCheckpointID(c.checkpointID).
					WithDocumentID(c.documentID).
					WithPassageType(1).
					WithTime(time.Now().UTC()).
					ToCreateDTO(),
			)
			sCtx.Assert().NoError(err)
			sCtx.Assert().NotNil(expPassage)

			expPassages = append(expPassages, expPassage)
		}
		request := utils.CheckpointBuilder{}.
			WithDocumentID(c.documentID).
			ToListDTO()

		sCtx.WithNewParameters("ctx", ctx, "request", request)

		passages, err := c.checkpointStorage.ListPassages(ctx, request)

		sCtx.Assert().NoError(err)
		sCtx.Assert().NotNil(passages)
		sCtx.Assert().GreaterOrEqual(len(passages), len(expPassages))

		for _, passage := range expPassages {
			err = c.checkpointStorage.DeletePassage(context.TODO(), &dto.DeletePassageRequest{PassageID: passage.ID.Int()})
			sCtx.Assert().NoError(err)
		}
	})
}

func (c *CheckpointStorageSuite) Test_IntegrationCheckpointStorage_DeleteCheckpoint(t provider.T) {
	t.Title("[DeleteCheckpoint] Delete checkpoint test")
	t.Tags("storage", "postgres", "checkpoint")
	t.Parallel()
	t.WithNewStep("Delete checkpoint test", func(sCtx provider.StepCtx) {
		ctx := context.TODO()
		expCheckpoint, err := c.checkpointStorage.CreateCheckpoint(
			ctx,
			utils.CheckpointBuilder{}.
				WithPhoneNumber("123543").
				ToCreateCheckpointDTO(),
		)
		sCtx.Assert().NoError(err)
		sCtx.Assert().NotNil(expCheckpoint)
		request := utils.CheckpointBuilder{}.
			WithCheckpointID(expCheckpoint.ID.Int()).
			ToDeleteCheckpointDTO()

		sCtx.WithNewParameters("ctx", ctx, "request", request)

		err = c.checkpointStorage.DeleteCheckpoint(ctx, request)

		sCtx.Assert().NoError(err)

		checkpoint, err := c.checkpointStorage.GetCheckpoint(
			ctx,
			utils.CheckpointBuilder{}.
				WithCheckpointID(expCheckpoint.ID.Int()).
				ToGetCheckpointDTO(),
		)

		sCtx.Assert().Error(err)
		sCtx.Assert().EqualError(err, pgx.ErrNoRows.Error())
		sCtx.Assert().Nil(checkpoint)

		err = c.checkpointStorage.DeleteCheckpoint(ctx, request)

		sCtx.Assert().NoError(err)
	})
}

func (c *CheckpointStorageSuite) Test_IntegrationCheckpointStorage_DeletePassage(t provider.T) {
	t.Title("[DeletePassage] Delete passage test")
	t.Tags("storage", "postgres", "checkpoint", "passage")
	t.Parallel()
	t.WithNewStep("Delete passage test", func(sCtx provider.StepCtx) {
		ctx := context.TODO()
		expPassage, err := c.checkpointStorage.CreatePassage(
			ctx,
			utils.CheckpointBuilder{}.
				WithCheckpointID(c.checkpointID).
				WithDocumentID(c.documentID).
				WithPassageType(1).
				WithTime(time.Now().UTC()).
				ToCreateDTO(),
		)
		sCtx.Assert().NoError(err)
		sCtx.Assert().NotNil(expPassage)
		request := utils.CheckpointBuilder{}.
			WithPassageID(expPassage.ID.Int()).
			ToDeleteDTO()

		sCtx.WithNewParameters("ctx", ctx, "request", request)

		err = c.checkpointStorage.DeletePassage(ctx, request)

		sCtx.Assert().NoError(err)

		passage, err := c.checkpointStorage.GetPassage(
			ctx,
			utils.CheckpointBuilder{}.
				WithPassageID(expPassage.ID.Int()).
				ToGetDTO(),
		)

		sCtx.Assert().Error(err)
		sCtx.Assert().EqualError(err, pgx.ErrNoRows.Error())
		sCtx.Assert().Nil(passage)

		err = c.checkpointStorage.DeletePassage(ctx, request)

		sCtx.Assert().NoError(err)
	})
}
