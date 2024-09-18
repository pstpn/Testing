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

type checkpointStorageImpl struct {
	*postgres.Postgres
}

func NewCheckpointStorage(db *postgres.Postgres) storage.CheckpointStorage {
	return &checkpointStorageImpl{db}
}

func (c *checkpointStorageImpl) CreateCheckpoint(ctx context.Context, request *dto.CreateCheckpointRequest) (*model.Checkpoint, error) {
	query := c.Builder.
		Insert(checkpointTable).
		Columns(
			phoneNumberField,
		).
		Values(
			request.PhoneNumber,
		).
		Suffix(returningCheckpointColumns())

	sql, args, err := query.ToSql()
	if err != nil {
		return nil, err
	}
	row := c.Pool.QueryRow(ctx, sql, args...)

	return c.rowToCheckpointModel(row)
}

func (c *checkpointStorageImpl) CreatePassage(ctx context.Context, request *dto.CreatePassageRequest) (*model.Passage, error) {
	query := c.Builder.
		Insert(passageTable).
		Columns(
			checkpointIdField,
			documentIdField,
			typeField,
			timeField,
		).
		Values(
			request.CheckpointID,
			request.DocumentID,
			model.ToPassageTypeFromInt(request.Type).String(),
			request.Time,
		).
		Suffix(returningPassageColumns())

	sql, args, err := query.ToSql()
	if err != nil {
		return nil, err
	}
	row := c.Pool.QueryRow(ctx, sql, args...)

	return c.rowToPassageModel(row)
}

func (c *checkpointStorageImpl) GetPassage(ctx context.Context, request *dto.GetPassageRequest) (*model.Passage, error) {
	query := c.Builder.
		Select(
			idField,
			checkpointIdField,
			documentIdField,
			typeField,
			timeField,
		).
		From(passageTable).
		Where(squirrel.Eq{idField: request.PassageID})

	sql, args, err := query.ToSql()
	if err != nil {
		return nil, err
	}
	row := c.Pool.QueryRow(ctx, sql, args...)

	return c.rowToPassageModel(row)
}

func (c *checkpointStorageImpl) GetCheckpoint(ctx context.Context, request *dto.GetCheckpointRequest) (*model.Checkpoint, error) {
	query := c.Builder.
		Select(
			idField,
			phoneNumberField,
		).
		From(checkpointTable).
		Where(squirrel.Eq{idField: request.CheckpointID})

	sql, args, err := query.ToSql()
	if err != nil {
		return nil, err
	}
	row := c.Pool.QueryRow(ctx, sql, args...)

	return c.rowToCheckpointModel(row)
}

func (c *checkpointStorageImpl) ListPassages(ctx context.Context, request *dto.ListPassagesRequest) ([]*model.Passage, error) {
	query := c.Builder.
		Select(
			fullColName(passageTable, idField),
			checkpointIdField,
			documentIdField,
			fullColName(passageTable, typeField),
			timeField,
		).
		From(passageTable).
		Join(on(
			passageTable,
			documentTable,
			documentIdField,
			idField,
		)).
		Where(squirrel.Eq{documentIdField: request.DocumentID})

	sql, args, err := query.ToSql()
	if err != nil {
		return nil, err
	}
	rows, err := c.Pool.Query(ctx, sql, args...)
	if err != nil {
		return nil, err
	}

	var passages []*model.Passage
	for rows.Next() {
		passage, err := c.rowToPassageModel(rows)
		if err != nil {
			return nil, err
		}

		passages = append(passages, passage)
	}

	return passages, nil
}

func (c *checkpointStorageImpl) DeletePassage(ctx context.Context, request *dto.DeletePassageRequest) error {
	query := c.Builder.
		Delete(passageTable).
		Where(squirrel.Eq{idField: request.PassageID})

	sql, args, err := query.ToSql()
	if err != nil {
		return err
	}
	_, err = c.Pool.Exec(ctx, sql, args...)
	if err != nil {
		return err
	}

	return nil
}

func (c *checkpointStorageImpl) DeleteCheckpoint(ctx context.Context, request *dto.DeleteCheckpointRequest) error {
	query := c.Builder.
		Delete(checkpointTable).
		Where(squirrel.Eq{idField: request.CheckpointID})

	sql, args, err := query.ToSql()
	if err != nil {
		return err
	}
	_, err = c.Pool.Exec(ctx, sql, args...)
	if err != nil {
		return err
	}

	return nil
}

func (c *checkpointStorageImpl) rowToPassageModel(row pgx.Row) (*model.Passage, error) {
	var passage model.Passage
	var passageType string
	err := row.Scan(
		&passage.ID,
		&passage.CheckpointID,
		&passage.DocumentID,
		&passageType,
		&passage.Time,
	)
	if err != nil {
		return nil, err
	}
	passage.Type = model.ToPassageTypeFromString(passageType)

	return &passage, nil
}

func (c *checkpointStorageImpl) rowToCheckpointModel(row pgx.Row) (*model.Checkpoint, error) {
	var checkpoint model.Checkpoint
	err := row.Scan(
		&checkpoint.ID,
		&checkpoint.PhoneNumber,
	)
	if err != nil {
		return nil, err
	}

	return &checkpoint, nil
}
