package repository

import (
	"context"
	"reports/model"
)

type ReportRepository interface {
	Save(ctx context.Context, book model.Report)
	Update(ctx context.Context, book model.Report)
	Delete(ctx context.Context, bookId int)
	FindById(ctx context.Context, bookId int) (model.Report, error)
	FindAll(ctx context.Context) []model.Report
}
