package postgres

import (
	"context"

	"github.com/Masterminds/squirrel"
	"github.com/jackc/pgx/v5"

	"course/internal/model"
	"course/internal/service/dto"
	"course/internal/storage"
	"course/pkg/storage/postgres"
)

type fieldStorageImpl struct {
	*postgres.Postgres
}

func NewFieldStorage(db *postgres.Postgres) storage.FieldStorage {
	return &fieldStorageImpl{db}
}

func (f *fieldStorageImpl) Create(ctx context.Context, request *dto.CreateDocumentFieldRequest) (*model.Field, error) {
	query := f.Builder.
		Insert(fieldTable).
		Columns(
			documentIdField,
			typeField,
			valueField,
		).
		Values(
			request.DocumentID,
			model.ToFieldTypeFromInt(request.Type).String(),
			request.Value,
		).
		Suffix(returningFieldColumns())

	sql, args, err := query.ToSql()
	if err != nil {
		return nil, err
	}
	row := f.Pool.QueryRow(ctx, sql, args...)

	return f.rowToModel(row)
}

func (f *fieldStorageImpl) Get(ctx context.Context, request *dto.GetDocumentFieldRequest) (*model.Field, error) {
	query := f.Builder.
		Select(
			idField,
			documentIdField,
			typeField,
			valueField,
		).
		From(fieldTable).
		Where(
			squirrel.And{
				squirrel.Eq{documentIdField: request.DocumentID},
				squirrel.Eq{typeField: model.ToFieldTypeFromInt(request.FieldType).String()},
			},
		)

	sql, args, err := query.ToSql()
	if err != nil {
		return nil, err
	}
	row := f.Pool.QueryRow(ctx, sql, args...)

	return f.rowToModel(row)
}

func (f *fieldStorageImpl) ListCardFields(ctx context.Context, request *dto.ListDocumentFieldsRequest) ([]*model.Field, error) {
	query := f.Builder.
		Select(
			idField,
			documentIdField,
			typeField,
			valueField,
		).
		From(fieldTable).
		Where(squirrel.Eq{documentIdField: request.DocumentID})

	sql, args, err := query.ToSql()
	if err != nil {
		return nil, err
	}
	rows, err := f.Pool.Query(ctx, sql, args...)
	if err != nil {
		return nil, err
	}

	var fields []*model.Field
	for rows.Next() {
		field, err := f.rowToModel(rows)
		if err != nil {
			return nil, err
		}

		fields = append(fields, field)
	}

	return fields, nil
}

func (f *fieldStorageImpl) Delete(ctx context.Context, request *dto.DeleteDocumentFieldRequest) error {
	query := f.Builder.
		Delete(fieldTable).
		Where(squirrel.Eq{idField: request.FieldID})

	sql, args, err := query.ToSql()
	if err != nil {
		return err
	}
	_, err = f.Pool.Exec(ctx, sql, args...)
	if err != nil {
		return err
	}

	return nil
}

func (f *fieldStorageImpl) rowToModel(row pgx.Row) (*model.Field, error) {
	var field model.Field
	var fieldType string
	err := row.Scan(
		&field.ID,
		&field.DocumentID,
		&fieldType,
		&field.Value,
	)
	if err != nil {
		return nil, err
	}
	field.Type = model.ToFieldTypeFromString(fieldType)

	return &field, nil
}
