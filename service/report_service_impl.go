package service

import (
	"context"
	"reports/data/request"
	"reports/data/response"
	"reports/helper"
	"reports/model"
	"reports/repository"
)

type ReportServiceImpl struct {
	reportRepository repository.ReportRepository
}

func NewReportServiceImpl(reportRepository repository.ReportRepository) ReportService {
	return &ReportServiceImpl{reportRepository: reportRepository}
}

func (r *ReportServiceImpl) Create(ctx context.Context, request request.ReportCreateRequest) {
	report := model.Report{
		MonthOf:          request.MonthOf,
		AreaOfAssignment: request.AreaOfAssignment,
		NameOfChurch:     request.NameOfChurch,
	}

	r.reportRepository.Save(ctx, report)
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
			Id:               value.Id,
			MonthOf:          value.MonthOf,
			AreaOfAssignment: value.AreaOfAssignment,
			NameOfChurch:     value.NameOfChurch,
		}

		reportResp = append(reportResp, report)
	}

	return reportResp
}

func (r *ReportServiceImpl) FindById(ctx context.Context, reportId int) response.ReportResponse {
	report, err := r.reportRepository.FindById(ctx, reportId)
	helper.ErrorPanic(err)
	return response.ReportResponse(report)
}

func (r *ReportServiceImpl) Update(ctx context.Context, request request.ReportUpdateRequest) {
	report, err := r.reportRepository.FindById(ctx, request.Id)
	helper.ErrorPanic(err)

	report.MonthOf = request.MonthOf
	report.AreaOfAssignment = request.AreaOfAssignment
	report.NameOfChurch = request.NameOfChurch

	r.reportRepository.Update(ctx, report)
}
