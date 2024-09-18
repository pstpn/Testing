package postgres

import (
	"context"
	"testing"

	"github.com/jackc/pgx/v5"
	"github.com/stretchr/testify/require"

	"course/internal/model"
	"course/internal/service/dto"
)

func Test_fieldStorageImpl_Create(t *testing.T) {
	fieldStorage := NewFieldStorage(testDB)

	request := &dto.CreateDocumentFieldRequest{
		DocumentID: ids["documentID"],
		Type:       0,
		Value:      "222",
	}

	field, err := fieldStorage.Create(context.TODO(), request)
	require.NoError(t, err)
	require.NotEmpty(t, field)
	require.Equal(t, model.ToDocumentID(request.DocumentID), field.DocumentID)
	require.Equal(t, model.ToFieldTypeFromInt(request.Type), field.Type)
	require.Equal(t, request.Value, field.Value)

	err = fieldStorage.Delete(context.TODO(), &dto.DeleteDocumentFieldRequest{FieldID: field.ID.Int()})
	require.NoError(t, err)
}

func Test_fieldStorageImpl_Get(t *testing.T) {
	fieldStorage := NewFieldStorage(testDB)

	field1, err := fieldStorage.Create(context.TODO(), &dto.CreateDocumentFieldRequest{
		DocumentID: ids["documentID"],
		Type:       0,
		Value:      "222",
	})
	require.NoError(t, err)
	require.NotEmpty(t, field1)

	field2, err := fieldStorage.Get(context.TODO(), &dto.GetDocumentFieldRequest{
		DocumentID: ids["documentID"],
		FieldType:  0,
	})
	require.NoError(t, err)
	require.NotEmpty(t, field2)
	require.Equal(t, field1.ID, field2.ID)
	require.Equal(t, field1.DocumentID, field2.DocumentID)
	require.Equal(t, field1.Type, field2.Type)
	require.Equal(t, field1.Value, field2.Value)

	err = fieldStorage.Delete(context.TODO(), &dto.DeleteDocumentFieldRequest{FieldID: field1.ID.Int()})
	require.NoError(t, err)
}

func Test_fieldStorageImpl_List(t *testing.T) {
	fieldStorage := NewFieldStorage(testDB)

	var fields []*model.Field
	for i := range 2 {

		field, err := fieldStorage.Create(context.TODO(), &dto.CreateDocumentFieldRequest{
			DocumentID: ids["documentID"],
			Type:       int64(i),
			Value:      "222",
		})
		require.NoError(t, err)
		require.NotEmpty(t, field)

		fields = append(fields, field)
	}

	listFields, err := fieldStorage.ListCardFields(context.TODO(), &dto.ListDocumentFieldsRequest{
		DocumentID: ids["documentID"],
	})
	require.NoError(t, err)
	require.NotEmpty(t, listFields)
	require.Equal(t, 2, len(listFields))

	for _, field := range fields {
		err = fieldStorage.Delete(context.TODO(), &dto.DeleteDocumentFieldRequest{FieldID: field.ID.Int()})
		require.NoError(t, err)
	}
}

func Test_fieldStorageImpl_Delete(t *testing.T) {
	fieldStorage := NewFieldStorage(testDB)

	field1, err := fieldStorage.Create(context.TODO(), &dto.CreateDocumentFieldRequest{
		DocumentID: ids["documentID"],
		Type:       0,
		Value:      "222",
	})
	require.NoError(t, err)
	require.NotEmpty(t, field1)

	err = fieldStorage.Delete(context.TODO(), &dto.DeleteDocumentFieldRequest{FieldID: field1.ID.Int()})
	require.NoError(t, err)

	field2, err := fieldStorage.Get(context.TODO(), &dto.GetDocumentFieldRequest{
		DocumentID: ids["documentID"],
		FieldType:  0,
	})
	require.Error(t, err)
	require.EqualError(t, err, pgx.ErrNoRows.Error())
	require.Empty(t, field2)

	err = fieldStorage.Delete(context.TODO(), &dto.DeleteDocumentFieldRequest{FieldID: field1.ID.Int()})
	require.NoError(t, err)
}
