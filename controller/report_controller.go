package controller

import (
	"net/http"
	"reports/data/request"
	"reports/data/response"
	"reports/helper"
	"reports/service"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ReportController struct {
	reportService service.ReportService
}

func NewReportController(reportService service.ReportService) *ReportController {
	return &ReportController{reportService: reportService}
}

func (controller *ReportController) Create(ctx *gin.Context) {
	reportCreateRequest := request.ReportCreateRequest{}
	err := ctx.ShouldBindJSON(&reportCreateRequest)
	helper.ErrorPanic(err)

	controller.reportService.Create(ctx, reportCreateRequest)
	webResponse := response.Response{
		Code:    http.StatusOK,
		Message: "Success",
		Data:    nil,
	}

	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusOK, webResponse)
}

func (controller *ReportController) FindById(ctx *gin.Context) {
	reportId := ctx.Param("reportId")
	id, err := strconv.Atoi(reportId)
	helper.ErrorPanic(err)

	result := controller.reportService.FindById(ctx, id)
	webResponse := response.Response{
		Code:    http.StatusOK,
		Message: "Success",
		Data:    result,
	}

	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusOK, webResponse)
}

func (controller *ReportController) FindAll(ctx *gin.Context) {
	tagsResponse := controller.reportService.FindAll(ctx)
	webResponse := response.Response{
		Code:    http.StatusOK,
		Message: "Success",
		Data:    tagsResponse,
	}

	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusOK, webResponse)
}

func (controller *ReportController) Delete(ctx *gin.Context) {
	tagId := ctx.Param("reportId")
	id, err := strconv.Atoi(tagId)
	helper.ErrorPanic(err)

	controller.reportService.Delete(ctx, id)

	webResponse := response.Response{
		Code:    http.StatusOK,
		Message: "Success",
		Data:    nil,
	}

	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusOK, webResponse)
}

func (controller *ReportController) Update(ctx *gin.Context) {
	updateTagsRequest := request.ReportUpdateRequest{}
	err := ctx.ShouldBindJSON(&updateTagsRequest)
	helper.ErrorPanic(err)

	tagId := ctx.Param("reportId")
	id, err := strconv.Atoi(tagId)
	helper.ErrorPanic(err)

	updateTagsRequest.Id = id

	controller.reportService.Update(ctx, updateTagsRequest)
	webResponse := response.Response{
		Code:    http.StatusOK,
		Message: "Success",
		Data:    nil,
	}

	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusOK, webResponse)
}
