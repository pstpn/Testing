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

type infoCardStorageImpl struct {
	*postgres.Postgres
}

func NewInfoCardStorage(db *postgres.Postgres) storage.InfoCardStorage {
	return &infoCardStorageImpl{db}
}

func (i *infoCardStorageImpl) Create(ctx context.Context, request *dto.CreateInfoCardRequest) (*model.InfoCard, error) {
	query := i.Builder.
		Insert(infoCardTable).
		Columns(
			createdEmployeeIDField,
			isConfirmedField,
			createdDateField,
		).
		Values(
			request.EmployeeID,
			request.IsConfirmed,
			request.CreatedDate,
		).
		Suffix(returningInfoCardColumns())

	sql, args, err := query.ToSql()
	if err != nil {
		return nil, err
	}
	row := i.Pool.QueryRow(ctx, sql, args...)

	return i.rowToModel(row)
}

func (i *infoCardStorageImpl) Validate(ctx context.Context, request *dto.ValidateInfoCardRequest) error {
	query := i.Builder.
		Update(infoCardTable).
		Set(isConfirmedField, request.IsConfirmed).
		Where(squirrel.Eq{idField: request.InfoCardID})

	sql, args, err := query.ToSql()
	if err != nil {
		return err
	}
	_, err = i.Pool.Exec(ctx, sql, args...)
	if err != nil {
		return err
	}

	return nil
}

func (i *infoCardStorageImpl) GetByID(ctx context.Context, request *dto.GetInfoCardByIDRequest) (*model.InfoCard, error) {
	query := i.Builder.
		Select(
			idField,
			createdEmployeeIDField,
			isConfirmedField,
			createdDateField,
		).
		From(infoCardTable).
		Where(squirrel.Eq{idField: request.InfoCardID})

	sql, args, err := query.ToSql()
	if err != nil {
		return nil, err
	}
	row := i.Pool.QueryRow(ctx, sql, args...)

	return i.rowToModel(row)
}

func (i *infoCardStorageImpl) GetByEmployeeID(ctx context.Context, request *dto.GetInfoCardByEmployeeIDRequest) (*model.InfoCard, error) {
	query := i.Builder.
		Select(
			idField,
			createdEmployeeIDField,
			isConfirmedField,
			createdDateField,
		).
		From(infoCardTable).
		Where(squirrel.Eq{createdEmployeeIDField: request.EmployeeID})

	sql, args, err := query.ToSql()
	if err != nil {
		return nil, err
	}
	row := i.Pool.QueryRow(ctx, sql, args...)

	return i.rowToModel(row)
}

func (i *infoCardStorageImpl) List(ctx context.Context, request *dto.ListInfoCardsRequest) ([]*model.FullInfoCard, error) {
	query := i.Builder.
		Select(
			fullColName(infoCardTable, idField),
			createdEmployeeIDField,
			isConfirmedField,
			createdDateField,
			fullNameField,
			phoneNumberField,
			postField,
			dateOfBirthField,
		).
		From(infoCardTable).
		Join(on(
			infoCardTable,
			employeeTable,
			createdEmployeeIDField,
			idField,
		))

	query = request.Pagination.ToSQL(query)

	sql, args, err := query.ToSql()
	if err != nil {
		return nil, err
	}
	rows, err := i.Pool.Query(ctx, sql, args...)
	if err != nil {
		return nil, err
	}

	var fullInfoCards []*model.FullInfoCard
	for rows.Next() {
		fullInfoCard, err := i.rowToFullModel(rows)
		if err != nil {
			return nil, err
		}

		fullInfoCards = append(fullInfoCards, fullInfoCard)
	}

	return fullInfoCards, nil
}

func (i *infoCardStorageImpl) Delete(ctx context.Context, request *dto.DeleteInfoCardRequest) error {
	query := i.Builder.
		Delete(infoCardTable).
		Where(squirrel.Eq{idField: request.InfoCardID})

	sql, args, err := query.ToSql()
	if err != nil {
		return err
	}
	_, err = i.Pool.Exec(ctx, sql, args...)
	if err != nil {
		return err
	}

	return nil
}

func (i *infoCardStorageImpl) rowToModel(row pgx.Row) (*model.InfoCard, error) {
	var infoCard model.InfoCard
	err := row.Scan(
		&infoCard.ID,
		&infoCard.CreatedEmployeeID,
		&infoCard.IsConfirmed,
		&infoCard.CreatedDate,
	)
	if err != nil {
		return nil, err
	}

	return &infoCard, nil
}

func (i *infoCardStorageImpl) rowToFullModel(row pgx.Row) (*model.FullInfoCard, error) {
	var fullInfoCard model.FullInfoCard
	err := row.Scan(
		&fullInfoCard.ID,
		&fullInfoCard.CreatedEmployeeID,
		&fullInfoCard.IsConfirmed,
		&fullInfoCard.CreatedDate,
		&fullInfoCard.FullName,
		&fullInfoCard.PhoneNumber,
		&fullInfoCard.Post,
		&fullInfoCard.DateOfBirth,
	)
	if err != nil {
		return nil, err
	}

	return &fullInfoCard, nil
}
