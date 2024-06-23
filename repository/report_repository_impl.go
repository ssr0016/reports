package repository

import (
	"context"
	"database/sql"
	"encoding/json"
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
			worker_name,
			area_of_assignment,
			name_of_church,
			created_at,
			updated_at,
			worship_service,
			average_attendance
		FROM reports
	`
	result, errQuery := tx.QueryContext(ctx, rawSQL)
	helper.ErrorPanic(errQuery)
	defer result.Close()

	var reports []model.Report

	for result.Next() {
		report := model.Report{}
		var worshipServiceJSON []byte

		err := result.Scan(
			&report.Id,
			&report.MonthOf,
			&report.WorkerName,
			&report.AreaOfAssignment,
			&report.NameOfChurch,
			&report.CreatedAt,
			&report.UpdatedAt,
			&worshipServiceJSON,
			&report.AverageAttendance,
		)
		helper.ErrorPanic(err)

		// Unmarshal worshipServiceJSON into []int
		err = json.Unmarshal(worshipServiceJSON, &report.WorshipService)
		if err != nil {
			helper.ErrorPanic(err)
		}

		reports = append(reports, report)
	}

	return reports
}

// FindById implements BookRepository
func (r *ReportRepositoryImpl) FindById(ctx context.Context, reportId int) (model.Report, error) {
	tx, err := r.Db.Begin()
	helper.ErrorPanic(err)
	defer helper.CommitOrRollback(tx)

	rawSQL := `
		SELECT 
			id,
			month_of,
			worker_name,
			area_of_assignment,
			name_of_church,
			created_at,
			updated_at,
			worship_service,
			average_attendance
		FROM reports
		WHERE 
			id = $1
	`
	result, errQuery := tx.QueryContext(ctx, rawSQL, reportId)
	helper.ErrorPanic(errQuery)
	defer result.Close()

	report := model.Report{}

	if result.Next() {
		var worshipServiceJSON []byte

		err := result.Scan(
			&report.Id,
			&report.MonthOf,
			&report.WorkerName,
			&report.AreaOfAssignment,
			&report.NameOfChurch,
			&report.CreatedAt,
			&report.UpdatedAt,
			&worshipServiceJSON,
			&report.AverageAttendance,
		)
		helper.ErrorPanic(err)

		// Unmarshal worshipServiceJSON into []int
		err = json.Unmarshal(worshipServiceJSON, &report.WorshipService)
		if err != nil {
			helper.ErrorPanic(err)
		}

		return report, nil
	}

	return report, errors.New("report not found")
}

// Save implements BookRepository
func (r *ReportRepositoryImpl) Save(ctx context.Context, report model.Report) error {
	tx, err := r.Db.Begin()
	helper.ErrorPanic(err)
	defer helper.CommitOrRollback(tx)

	worshipServiceJSON, err := json.Marshal(report.WorshipService)
	if err != nil {
		return err
	}

	rawSQL := `
		INSERT INTO reports (
			month_of,
			worker_name,
			area_of_assignment,
			worship_service,
			average_attendance,
			name_of_church,
			created_at,
			updated_at
		) VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
	`

	_, err = tx.ExecContext(ctx, rawSQL,
		report.MonthOf,
		report.WorkerName,
		report.AreaOfAssignment,
		worshipServiceJSON,
		report.AverageAttendance,
		report.NameOfChurch,
		report.CreatedAt,
		report.UpdatedAt,
	)

	return err
}

// Update implements BookRepository
func (r *ReportRepositoryImpl) Update(ctx context.Context, report model.Report) error {
	tx, err := r.Db.Begin()
	helper.ErrorPanic(err)
	defer helper.CommitOrRollback(tx)

	loc, err := time.LoadLocation("Asia/Manila")
	if err != nil {
		helper.ErrorPanic(err)
	}

	now := time.Now().In(loc)

	rawSQL := `
        UPDATE reports
        SET
            month_of = $1,
			worker_name = $2,
            area_of_assignment = $3,
            name_of_church = $4,
            updated_at = $5
        WHERE
            id = $6
    `
	_, err = tx.ExecContext(ctx, rawSQL,
		report.MonthOf,
		report.WorkerName,
		report.AreaOfAssignment,
		report.NameOfChurch,
		now,
		report.Id,
	)

	helper.ErrorPanic(err)

	return nil
}
