package postgres

import (
	"context"
	"testing"

	"github.com/jackc/pgx/v5"
	"github.com/stretchr/testify/require"

	"course/internal/model"
	"course/internal/service/dto"
)

func Test_photoMetaStorageImpl_SaveKey(t *testing.T) {
	photoMetaStorage := NewPhotoMetaStorage(testDB)

	request := &dto.CreatePhotoKeyRequest{
		DocumentID: model.ToDocumentID(ids["documentID"]),
		Key:        model.ToPhotoKey("123"),
	}

	photoMeta, err := photoMetaStorage.SaveKey(context.TODO(), request)
	require.NoError(t, err)
	require.NotEmpty(t, photoMeta)
	require.Equal(t, request.DocumentID, photoMeta.DocumentID)
	require.Equal(t, request.Key, photoMeta.PhotoKey)

	err = photoMetaStorage.DeleteKey(context.TODO(), &dto.DeletePhotoRequest{DocumentID: request.DocumentID.Int()})
	require.NoError(t, err)
}

func Test_photoMetaStorageImpl_GetKey(t *testing.T) {
	photoMetaStorage := NewPhotoMetaStorage(testDB)

	photoMeta1, err := photoMetaStorage.SaveKey(context.TODO(), &dto.CreatePhotoKeyRequest{
		DocumentID: model.ToDocumentID(ids["documentID"]),
		Key:        model.ToPhotoKey("123"),
	})
	require.NoError(t, err)
	require.NotEmpty(t, photoMeta1)

	photoMeta2, err := photoMetaStorage.GetKey(context.TODO(), &dto.GetPhotoRequest{DocumentID: photoMeta1.DocumentID.Int()})
	require.NoError(t, err)
	require.NotEmpty(t, photoMeta2)
	require.Equal(t, photoMeta1.ID, photoMeta2.ID)
	require.Equal(t, photoMeta1.DocumentID, photoMeta2.DocumentID)
	require.Equal(t, photoMeta1.PhotoKey, photoMeta2.PhotoKey)

	err = photoMetaStorage.DeleteKey(context.TODO(), &dto.DeletePhotoRequest{DocumentID: photoMeta1.DocumentID.Int()})
	require.NoError(t, err)
}

func Test_photoMetaStorageImpl_DeleteKey(t *testing.T) {
	photoMetaStorage := NewPhotoMetaStorage(testDB)

	photoMeta1, err := photoMetaStorage.SaveKey(context.TODO(), &dto.CreatePhotoKeyRequest{
		DocumentID: model.ToDocumentID(ids["documentID"]),
		Key:        model.ToPhotoKey("123"),
	})
	require.NoError(t, err)
	require.NotEmpty(t, photoMeta1)

	err = photoMetaStorage.DeleteKey(context.TODO(), &dto.DeletePhotoRequest{DocumentID: photoMeta1.DocumentID.Int()})
	require.NoError(t, err)

	photoMeta2, err := photoMetaStorage.GetKey(context.TODO(), &dto.GetPhotoRequest{DocumentID: photoMeta1.DocumentID.Int()})
	require.Error(t, err)
	require.EqualError(t, err, pgx.ErrNoRows.Error())
	require.Empty(t, photoMeta2)

	err = photoMetaStorage.DeleteKey(context.TODO(), &dto.DeletePhotoRequest{DocumentID: photoMeta1.DocumentID.Int()})
	require.NoError(t, err)
}
