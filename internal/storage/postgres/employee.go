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

type employeeStorageImpl struct {
	*postgres.Postgres
}

func NewEmployeeStorage(db *postgres.Postgres) storage.EmployeeStorage {
	return &employeeStorageImpl{db}
}

func (e *employeeStorageImpl) Register(ctx context.Context, request *dto.RegisterEmployeeRequest) (*model.Employee, error) {
	query := e.Builder.
		Insert(employeeTable).
		Columns(
			phoneNumberField,
			fullNameField,
			companyIdField,
			postField,
			passwordField,
			refreshTokenField,
			tokenExpiredAtField,
			dateOfBirthField,
		).
		Values(
			request.PhoneNumber,
			request.FullName,
			request.CompanyID,
			model.ToPostTypeFromInt(request.Post).String(),
			request.Password,
			request.RefreshToken,
			request.TokenExpiredAt,
			request.DateOfBirth,
		).
		Suffix(returningEmployeeColumns())

	sql, args, err := query.ToSql()
	if err != nil {
		return nil, err
	}
	row := e.Pool.QueryRow(ctx, sql, args...)

	return e.rowToModel(row)
}

func (e *employeeStorageImpl) UpdateRefreshToken(ctx context.Context, request *dto.UpdateToken) error {
	query := e.Builder.
		Update(employeeTable).
		Set(refreshTokenField, request.RefreshToken).
		Set(tokenExpiredAtField, request.TokenExpiredAt).
		Where(squirrel.Eq{idField: request.EmployeeID})

	sql, args, err := query.ToSql()
	if err != nil {
		return err
	}
	_, err = e.Pool.Exec(ctx, sql, args...)
	if err != nil {
		return err
	}

	return nil
}

func (e *employeeStorageImpl) GetByPhone(ctx context.Context, request *dto.GetEmployeeRequest) (*model.Employee, error) {
	query := e.Builder.
		Select(
			idField,
			phoneNumberField,
			fullNameField,
			companyIdField,
			postField,
			passwordField,
			refreshTokenField,
			tokenExpiredAtField,
			dateOfBirthField,
		).
		From(employeeTable).
		Where(squirrel.Eq{phoneNumberField: request.PhoneNumber})

	sql, args, err := query.ToSql()
	if err != nil {
		return nil, err
	}
	row := e.Pool.QueryRow(ctx, sql, args...)

	return e.rowToModel(row)
}

func (e *employeeStorageImpl) GetByInfoCardID(ctx context.Context, request *dto.GetEmployeeByInfoCardIDRequest) (*model.Employee, error) {
	query := e.Builder.
		Select(
			fullColName(employeeTable, idField),
			phoneNumberField,
			fullNameField,
			companyIdField,
			postField,
			passwordField,
			refreshTokenField,
			tokenExpiredAtField,
			dateOfBirthField,
		).
		From(employeeTable).
		Join(on(
			employeeTable,
			infoCardTable,
			idField,
			createdEmployeeIDField,
		)).
		Where(squirrel.Eq{fullColName(infoCardTable, idField): request.InfoCardID})

	sql, args, err := query.ToSql()
	if err != nil {
		return nil, err
	}
	row := e.Pool.QueryRow(ctx, sql, args...)

	return e.rowToModel(row)
}

func (e *employeeStorageImpl) Delete(ctx context.Context, request *dto.DeleteEmployeeRequest) error {
	query := e.Builder.
		Delete(employeeTable).
		Where(squirrel.Eq{idField: request.EmployeeID})

	sql, args, err := query.ToSql()
	if err != nil {
		return err
	}
	_, err = e.Pool.Exec(ctx, sql, args...)
	if err != nil {
		return err
	}

	return nil
}

func (e *employeeStorageImpl) rowToModel(row pgx.Row) (*model.Employee, error) {
	var employee model.Employee
	var post string

	err := row.Scan(
		&employee.ID,
		&employee.PhoneNumber,
		&employee.FullName,
		&employee.CompanyID,
		&post,
		&employee.Password,
		&employee.RefreshToken,
		&employee.TokenExpiredAt,
		&employee.DateOfBirth,
	)
	if err != nil {
		return nil, err
	}

	employee.Post = model.ToPostTypeFromString(post)

	return &employee, nil
}
