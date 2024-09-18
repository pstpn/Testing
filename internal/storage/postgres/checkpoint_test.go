package postgres

import (
	"context"
	"testing"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/stretchr/testify/require"

	"course/internal/model"
	"course/internal/service/dto"
)

func Test_checkpointStorageImpl_CreateCheckpoint(t *testing.T) {
	checkpointStorage := NewCheckpointStorage(testDB)

	request := &dto.CreateCheckpointRequest{
		PhoneNumber: "123432",
	}

	checkpoint, err := checkpointStorage.CreateCheckpoint(context.TODO(), request)
	require.NoError(t, err)
	require.NotEmpty(t, checkpoint)
	require.Equal(t, request.PhoneNumber, checkpoint.PhoneNumber)

	err = checkpointStorage.DeleteCheckpoint(context.TODO(), &dto.DeleteCheckpointRequest{CheckpointID: checkpoint.ID.Int()})
	require.NoError(t, err)
}

func Test_checkpointStorageImpl_CreatePassage(t *testing.T) {
	checkpointStorage := NewCheckpointStorage(testDB)

	tm, _ := time.Parse(time.RFC3339, "2006-01-02T15:04:05Z")

	request := &dto.CreatePassageRequest{
		CheckpointID: ids["checkpointID"],
		DocumentID:   ids["documentID"],
		Type:         0,
		Time:         &tm,
	}

	passage, err := checkpointStorage.CreatePassage(context.TODO(), request)
	require.NoError(t, err)
	require.NotEmpty(t, passage)
	require.Equal(t, model.ToCheckpointID(request.CheckpointID), passage.CheckpointID)
	require.Equal(t, model.ToDocumentID(request.DocumentID), passage.DocumentID)
	require.Equal(t, model.ToPassageTypeFromInt(request.Type), passage.Type)
	require.Equal(t, request.Time, passage.Time)

	err = checkpointStorage.DeletePassage(context.TODO(), &dto.DeletePassageRequest{PassageID: passage.ID.Int()})
	require.NoError(t, err)
}

func Test_checkpointStorageImpl_GetPassage(t *testing.T) {
	checkpointStorage := NewCheckpointStorage(testDB)

	tm, _ := time.Parse(time.RFC3339, "2006-01-02T15:04:05Z")
	passage1, err := checkpointStorage.CreatePassage(context.TODO(), &dto.CreatePassageRequest{
		CheckpointID: ids["checkpointID"],
		DocumentID:   ids["documentID"],
		Type:         0,
		Time:         &tm,
	})
	require.NoError(t, err)
	require.NotEmpty(t, passage1)

	passage2, err := checkpointStorage.GetPassage(context.TODO(), &dto.GetPassageRequest{
		PassageID: passage1.ID.Int(),
	})
	require.NoError(t, err)
	require.NotEmpty(t, passage2)
	require.Equal(t, passage1.ID, passage2.ID)
	require.Equal(t, passage1.CheckpointID, passage2.CheckpointID)
	require.Equal(t, passage1.DocumentID, passage2.DocumentID)
	require.Equal(t, passage1.Type, passage2.Type)
	require.Equal(t, passage1.Time, passage2.Time)

	err = checkpointStorage.DeletePassage(context.TODO(), &dto.DeletePassageRequest{PassageID: passage1.ID.Int()})
	require.NoError(t, err)
}

func Test_checkpointStorageImpl_GetCheckpoint(t *testing.T) {
	checkpointStorage := NewCheckpointStorage(testDB)

	request := &dto.CreateCheckpointRequest{
		PhoneNumber: "123432",
	}

	checkpoint, err := checkpointStorage.CreateCheckpoint(context.TODO(), request)
	require.NoError(t, err)
	require.NotEmpty(t, checkpoint)

	checkpoint2, err := checkpointStorage.GetCheckpoint(context.TODO(), &dto.GetCheckpointRequest{
		CheckpointID: checkpoint.ID.Int(),
	})
	require.NoError(t, err)
	require.NotEmpty(t, checkpoint2)
	require.Equal(t, checkpoint.ID, checkpoint2.ID)
	require.Equal(t, checkpoint.PhoneNumber, checkpoint2.PhoneNumber)

	err = checkpointStorage.DeleteCheckpoint(context.TODO(), &dto.DeleteCheckpointRequest{CheckpointID: checkpoint.ID.Int()})
	require.NoError(t, err)
}

func Test_checkpointStorageImpl_ListPassages(t *testing.T) {
	checkpointStorage := NewCheckpointStorage(testDB)

	tm, _ := time.Parse(time.RFC3339, "2006-01-02T15:04:05Z")

	var passages []*model.Passage
	for range 10 {
		passage, err := checkpointStorage.CreatePassage(context.TODO(), &dto.CreatePassageRequest{
			CheckpointID: ids["checkpointID"],
			DocumentID:   ids["documentID"],
			Type:         0,
			Time:         &tm,
		})
		require.NoError(t, err)
		require.NotEmpty(t, passage)

		passages = append(passages, passage)
	}

	listPassages, err := checkpointStorage.ListPassages(context.TODO(), &dto.ListPassagesRequest{
		InfoCardID: ids["infoCardID"],
	})
	require.NoError(t, err)
	require.NotEmpty(t, listPassages)
	require.Equal(t, 10, len(listPassages))

	for _, passage := range passages {
		err = checkpointStorage.DeletePassage(context.TODO(), &dto.DeletePassageRequest{PassageID: passage.ID.Int()})
		require.NoError(t, err)
	}
}

func Test_checkpointStorageImpl_DeletePassage(t *testing.T) {
	checkpointStorage := NewCheckpointStorage(testDB)

	tm, _ := time.Parse(time.RFC3339, "2006-01-02T15:04:05Z")
	passage1, err := checkpointStorage.CreatePassage(context.TODO(), &dto.CreatePassageRequest{
		CheckpointID: ids["checkpointID"],
		DocumentID:   ids["documentID"],
		Type:         0,
		Time:         &tm,
	})
	require.NoError(t, err)
	require.NotEmpty(t, passage1)

	err = checkpointStorage.DeletePassage(context.TODO(), &dto.DeletePassageRequest{PassageID: passage1.ID.Int()})
	require.NoError(t, err)

	passage2, err := checkpointStorage.GetPassage(context.TODO(), &dto.GetPassageRequest{
		PassageID: passage1.ID.Int(),
	})
	require.Error(t, err)
	require.EqualError(t, err, pgx.ErrNoRows.Error())
	require.Empty(t, passage2)

	err = checkpointStorage.DeletePassage(context.TODO(), &dto.DeletePassageRequest{PassageID: passage1.ID.Int()})
	require.NoError(t, err)
}

func Test_checkpointStorageImpl_DeleteCheckpoint(t *testing.T) {
	checkpointStorage := NewCheckpointStorage(testDB)

	request := &dto.CreateCheckpointRequest{
		PhoneNumber: "123432",
	}

	checkpoint, err := checkpointStorage.CreateCheckpoint(context.TODO(), request)
	require.NoError(t, err)
	require.NotEmpty(t, checkpoint)

	err = checkpointStorage.DeleteCheckpoint(context.TODO(), &dto.DeleteCheckpointRequest{CheckpointID: checkpoint.ID.Int()})
	require.NoError(t, err)

	checkpoint2, err := checkpointStorage.GetCheckpoint(context.TODO(), &dto.GetCheckpointRequest{
		CheckpointID: checkpoint.ID.Int(),
	})
	require.Error(t, err)
	require.EqualError(t, err, pgx.ErrNoRows.Error())
	require.Empty(t, checkpoint2)

	err = checkpointStorage.DeleteCheckpoint(context.TODO(), &dto.DeleteCheckpointRequest{CheckpointID: checkpoint.ID.Int()})
	require.NoError(t, err)
}
