package repository

import (
	"context"
	"database/sql"
	"errors"
	"reports/helper"
	"reports/model"
	"time"
)

type ReportRepositoryImpl struct {
	Db *sql.DB
}

func NewReportRepository(Db *sql.DB) ReportRepository {
	return &ReportRepositoryImpl{Db: Db}
}

// Delete implements BookRepository
func (r *ReportRepositoryImpl) Delete(ctx context.Context, reportId int) {
	tx, err := r.Db.Begin()
	helper.ErrorPanic(err)
	defer helper.CommitOrRollback(tx)

	rawSQL := `
		DELETE FROM reports
		WHERE
			 id = $1
	`
	_, err = tx.ExecContext(ctx, rawSQL, reportId)
	helper.ErrorPanic(err)
}

// FindAll implements BookRepository
func (r *ReportRepositoryImpl) FindAll(ctx context.Context) []model.Report {
	tx, err := r.Db.Begin()
	helper.ErrorPanic(err)
	defer helper.CommitOrRollback(tx)

	rawSQL := `
		SELECT 
			id,
			month_of,
			area_of_assignment,
			name_of_church
		FROM reports
	`
	result, errQuery := tx.QueryContext(ctx, rawSQL)
	helper.ErrorPanic(errQuery)
	defer result.Close()

	var reports []model.Report

	for result.Next() {
		report := model.Report{}
		err := result.Scan(
			&report.Id,
			&report.MonthOf,
			&report.AreaOfAssignment,
			&report.NameOfChurch,
		)
		helper.ErrorPanic(err)

		reports = append(reports, report)
	}

	return reports
}

// FindById implements BookRepository
func (r *ReportRepositoryImpl) FindById(ctx context.Context, bookId int) (model.Report, error) {
	tx, err := r.Db.Begin()
	helper.ErrorPanic(err)
	defer helper.CommitOrRollback(tx)

	rawSQL := `
		SELECT 
			id,
			month_of,
			area_of_assignment,
			name_of_church
		FROM reports
		WHERE 
			id = $1
	`
	result, errQuery := tx.QueryContext(ctx, rawSQL, bookId)
	helper.ErrorPanic(errQuery)
	defer result.Close()

	report := model.Report{}

	if result.Next() {
		err := result.Scan(
			&report.Id,
			&report.MonthOf,
			&report.AreaOfAssignment,
			&report.NameOfChurch,
		)
		helper.ErrorPanic(err)
		return report, nil
	} else {
		return report, errors.New("report id not found")
	}
}

// Save implements BookRepository
func (r *ReportRepositoryImpl) Save(ctx context.Context, book model.Report) {

	tx, err := r.Db.Begin()
	helper.ErrorPanic(err)
	defer helper.CommitOrRollback(tx)

	rawSQL := `
		INSERT INTO reports
		(month_of, area_of_assignment, name_of_church)
		VALUES
		($1, $2, $3, $4, $5)
	`
	now := time.Now().Format(time.RFC3339Nano)
	parsedTime, err := time.Parse(time.RFC3339Nano, now)
	if err != nil {
		helper.ErrorPanic(err)
	}
	_, err = tx.ExecContext(ctx, rawSQL, model.Report{
		MonthOf:          book.MonthOf,
		AreaOfAssignment: book.AreaOfAssignment,
		NameOfChurch:     book.NameOfChurch,
		CreatedAt:        parsedTime,
		UpdatedAt:        parsedTime,
	})

	helper.ErrorPanic(err)
}

// Update implements BookRepository
func (r *ReportRepositoryImpl) Update(ctx context.Context, report model.Report) {

	tx, err := r.Db.Begin()
	helper.ErrorPanic(err)
	defer helper.CommitOrRollback(tx)

	rawSQL := `
        UPDATE reports
        SET
            month_of = $1,
            area_of_assignment = $2,
            name_of_church = $3,
            updated_at = $4
        WHERE
            id = $5
    `
	now := time.Now().Format(time.RFC3339Nano)
	parsedTime, err := time.Parse(time.RFC3339Nano, now)
	if err != nil {
		helper.ErrorPanic(err)
	}
	_, err = tx.ExecContext(ctx, rawSQL,
		report.MonthOf,
		report.AreaOfAssignment,
		report.NameOfChurch,
		parsedTime,
		report.Id,
	)
	helper.ErrorPanic(err)
}
