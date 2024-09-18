package postgres

import (
	"context"
	"github.com/jackc/pgx/v5"
	"github.com/stretchr/testify/require"
	"testing"

	"course/internal/model"
	"course/internal/service/dto"
)

func Test_documentStorageImpl_Create(t *testing.T) {
	documentStorage := NewDocumentStorage(testDB)

	request := &dto.CreateDocumentRequest{
		SerialNumber: "123",
		InfoCardID:   ids["infoCardID"],
		DocumentType: 1,
	}

	document, err := documentStorage.Create(context.TODO(), request)
	require.NoError(t, err)
	require.NotEmpty(t, document)
	require.Equal(t, model.ToInfoCardID(request.InfoCardID), document.InfoCardID)
	require.Equal(t, model.ToDocumentTypeFromInt(request.DocumentType), document.Type)

	err = documentStorage.Delete(context.TODO(), &dto.DeleteDocumentRequest{DocumentID: document.ID.Int()})
	require.NoError(t, err)
}

func Test_documentStorageImpl_GetByID(t *testing.T) {
	documentStorage := NewDocumentStorage(testDB)

	document1, err := documentStorage.Create(context.TODO(), &dto.CreateDocumentRequest{
		SerialNumber: "123",
		InfoCardID:   ids["infoCardID"],
		DocumentType: 1,
	})
	require.NoError(t, err)
	require.NotEmpty(t, document1)

	document2, err := documentStorage.GetByID(context.TODO(), &dto.GetDocumentByIDRequest{
		DocumentID: document1.ID.Int(),
	})
	require.NoError(t, err)
	require.NotEmpty(t, document2)
	require.Equal(t, document1.ID, document2.ID)
	require.Equal(t, document1.InfoCardID, document2.InfoCardID)
	require.Equal(t, document1.Type, document2.Type)

	err = documentStorage.Delete(context.TODO(), &dto.DeleteDocumentRequest{DocumentID: document1.ID.Int()})
	require.NoError(t, err)
}

func Test_documentStorageImpl_List(t *testing.T) {
	documentStorage := NewDocumentStorage(testDB)

	var documents []*model.Document
	for range 10 {
		document, err := documentStorage.Create(context.TODO(), &dto.CreateDocumentRequest{
			SerialNumber: "123",
			InfoCardID:   ids["infoCardID"],
			DocumentType: 1,
		})
		require.NoError(t, err)
		require.NotEmpty(t, document)

		documents = append(documents, document)
	}

	listDocuments, err := documentStorage.List(context.TODO(), &dto.ListEmployeeDocumentsRequest{
		EmployeeID: ids["employeeID"],
	})
	require.NoError(t, err)
	require.NotEmpty(t, listDocuments)
	require.Equal(t, 11, len(listDocuments))

	for _, document := range documents {
		err = documentStorage.Delete(context.TODO(), &dto.DeleteDocumentRequest{DocumentID: document.ID.Int()})
		require.NoError(t, err)
	}
}

func Test_documentStorageImpl_Delete(t *testing.T) {
	documentStorage := NewDocumentStorage(testDB)

	document1, err := documentStorage.Create(context.TODO(), &dto.CreateDocumentRequest{
		SerialNumber: "123",
		InfoCardID:   ids["infoCardID"],
		DocumentType: 1,
	})
	require.NoError(t, err)
	require.NotEmpty(t, document1)

	err = documentStorage.Delete(context.TODO(), &dto.DeleteDocumentRequest{DocumentID: document1.ID.Int()})
	require.NoError(t, err)

	document2, err := documentStorage.GetByID(context.TODO(), &dto.GetDocumentByIDRequest{
		DocumentID: document1.ID.Int(),
	})
	require.Error(t, err)
	require.EqualError(t, err, pgx.ErrNoRows.Error())
	require.Empty(t, document2)

	err = documentStorage.Delete(context.TODO(), &dto.DeleteDocumentRequest{DocumentID: document1.ID.Int()})
	require.NoError(t, err)
}
