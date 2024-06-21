package service

import (
	"context"
	"reports/data/request"
	"reports/data/response"
)

type ReportService interface {
	Create(ctx context.Context, request request.ReportCreateRequest)
	Update(ctx context.Context, request request.ReportUpdateRequest)
	Delete(ctx context.Context, bookId int)
	FindById(ctx context.Context, bookId int) response.ReportResponse
	FindAll(ctx context.Context) []response.ReportResponse
}
