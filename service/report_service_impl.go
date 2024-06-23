package service

import (
	"context"
	"reports/data/request"
	"reports/data/response"
	"reports/helper"
	"reports/model"
	"reports/repository"
	"time"
)

type ReportServiceImpl struct {
	reportRepository repository.ReportRepository
}

func NewReportServiceImpl(reportRepository repository.ReportRepository) ReportService {
	return &ReportServiceImpl{reportRepository: reportRepository}
}

func (r *ReportServiceImpl) Create(ctx context.Context, request request.ReportCreateRequest) error {
	loc, err := time.LoadLocation("Asia/Manila")
	if err != nil {
		helper.ErrorPanic(err)
	}

	now := time.Now().In(loc)

	report := model.Report{
		MonthOf:          request.MonthOf,
		WorkerName:       request.WorkerName,
		AreaOfAssignment: request.AreaOfAssignment,
		NameOfChurch:     request.NameOfChurch,
		WorshipService:   request.WorshipService,
		CreatedAt:        now,
		UpdatedAt:        now,
	}

	report.CalculateAverage()

	err = r.reportRepository.Save(ctx, report)
	if err != nil {
		helper.ErrorPanic(err)
	}

	return err
}

func (r *ReportServiceImpl) Delete(ctx context.Context, reportId int) {
	report, err := r.reportRepository.FindById(ctx, reportId)
	helper.ErrorPanic(err)
	r.reportRepository.Delete(ctx, report.Id)
}

func (r *ReportServiceImpl) FindAll(ctx context.Context) []response.ReportResponse {
	reports := r.reportRepository.FindAll(ctx)

	var reportResp []response.ReportResponse

	for _, value := range reports {
		report := response.ReportResponse{
			Id:                value.Id,
			MonthOf:           value.MonthOf,
			WorkerName:        value.WorkerName,
			AreaOfAssignment:  value.AreaOfAssignment,
			NameOfChurch:      value.NameOfChurch,
			WorshipService:    value.WorshipService,
			AverageAttendance: value.AverageAttendance,
			CreatedAt:         value.CreatedAt,
			UpdatedAt:         value.UpdatedAt,
		}

		reportResp = append(reportResp, report)
	}

	return reportResp
}

func (r *ReportServiceImpl) FindById(ctx context.Context, reportId int) (response.ReportResponse, error) {
	report, err := r.reportRepository.FindById(ctx, reportId)
	helper.ErrorPanic(err)
	return response.ReportResponse(report), nil
}

func (r *ReportServiceImpl) Update(ctx context.Context, request request.ReportUpdateRequest) error {
	report, err := r.reportRepository.FindById(ctx, request.Id)
	helper.ErrorPanic(err)

	report.MonthOf = request.MonthOf
	report.WorkerName = request.WorkerName
	report.AreaOfAssignment = request.AreaOfAssignment
	report.NameOfChurch = request.NameOfChurch

	r.reportRepository.Update(ctx, report)

	return nil
}
