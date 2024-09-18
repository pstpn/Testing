package mongodb

import (
	"context"
	"os"
	"testing"

	"github.com/stretchr/testify/require"

	"course/internal/service/dto"
	"course/pkg/storage/mongodb"
)

const connURI = "mongodb://localhost:27017/"

var testMongoDB *mongodb.MongoDB

func TestMain(m *testing.M) {
	testMongoDB, _ = mongodb.New(connURI, "tests", "test")

	os.Exit(m.Run())
}

func Test_photoDataStorageImpl_Save(t *testing.T) {
	photoDataStorage := NewPhotoDataStorage(testMongoDB)

	key, err := photoDataStorage.Save(context.TODO(), &dto.CreatePhotoRequest{
		DocumentID: 1,
		Data:       []byte{'o'},
	})
	require.NoError(t, err)
	require.NotEmpty(t, key)
}

func Test_photoDataStorageImpl_Get(t *testing.T) {
	photoDataStorage := NewPhotoDataStorage(testMongoDB)

	request := &dto.CreatePhotoRequest{
		DocumentID: 1,
		Data:       []byte{'o'},
	}

	key, err := photoDataStorage.Save(context.TODO(), request)
	require.NoError(t, err)
	require.NotEmpty(t, key)

	data, err := photoDataStorage.Get(context.TODO(), key)
	require.NoError(t, err)
	require.NotEmpty(t, data)
	require.Equal(t, data, request.Data)
}

func Test_photoDataStorageImpl_Delete(t *testing.T) {
	photoDataStorage := NewPhotoDataStorage(testMongoDB)

	request := &dto.CreatePhotoRequest{
		DocumentID: 1,
		Data:       []byte{'o'},
	}

	key, err := photoDataStorage.Save(context.TODO(), request)
	require.NoError(t, err)
	require.NotEmpty(t, key)

	err = photoDataStorage.Delete(context.TODO(), key)
	require.NoError(t, err)

	data, err := photoDataStorage.Get(context.TODO(), key)
	require.Error(t, err)
	require.Empty(t, data)
}
