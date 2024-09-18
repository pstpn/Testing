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

type photoMetaStorageImpl struct {
	*postgres.Postgres
}

func NewPhotoMetaStorage(db *postgres.Postgres) storage.PhotoMetaStorage {
	return &photoMetaStorageImpl{db}
}

func (p *photoMetaStorageImpl) SaveKey(ctx context.Context, request *dto.CreatePhotoKeyRequest) (*model.PhotoMeta, error) {
	query := p.Builder.
		Insert(photoTable).
		Columns(
			documentIdField,
			keyField,
		).
		Values(
			request.DocumentID,
			request.Key,
		).
		Suffix(returningPhotoMetaColumns())

	sql, args, err := query.ToSql()
	if err != nil {
		return nil, err
	}
	row := p.Pool.QueryRow(ctx, sql, args...)

	return p.rowToModel(row)
}

func (p *photoMetaStorageImpl) GetKey(ctx context.Context, request *dto.GetPhotoRequest) (*model.PhotoMeta, error) {
	query := p.Builder.
		Select(
			idField,
			documentIdField,
			keyField,
		).
		From(photoTable).
		Where(squirrel.Eq{documentIdField: request.DocumentID})

	sql, args, err := query.ToSql()
	if err != nil {
		return nil, err
	}
	row := p.Pool.QueryRow(ctx, sql, args...)

	return p.rowToModel(row)
}

func (p *photoMetaStorageImpl) DeleteKey(ctx context.Context, request *dto.DeletePhotoRequest) error {
	query := p.Builder.
		Delete(photoTable).
		Where(squirrel.Eq{documentIdField: request.DocumentID})

	sql, args, err := query.ToSql()
	if err != nil {
		return err
	}
	_, err = p.Pool.Exec(ctx, sql, args...)
	if err != nil {
		return err
	}

	return nil
}

func (p *photoMetaStorageImpl) rowToModel(row pgx.Row) (*model.PhotoMeta, error) {
	var photoMeta model.PhotoMeta
	err := row.Scan(
		&photoMeta.ID,
		&photoMeta.DocumentID,
		&photoMeta.PhotoKey,
	)
	if err != nil {
		return nil, err
	}

	return &photoMeta, nil
}
