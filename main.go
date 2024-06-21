package main

import (
	"log"
	"net/http"
	"reports/config"
	"reports/controller"
	"reports/repository"
	"reports/router"
	"reports/service"

	"reports/helper"
)

func main() {

	loadConfig, err := config.LoadConfig(".")
	if err != nil {
		log.Fatal("cannot load config: ", err)
	}

	// Database
	db := config.ConnectionDB(&loadConfig)

	// Repository
	reportRepository := repository.NewReportRepository(db)

	// Service
	reportService := service.NewReportServiceImpl(reportRepository)

	// Controller
	reportController := controller.NewReportController(reportService)

	router := router.NewRouter(reportController)

	server := &http.Server{
		Addr:    ":8080",
		Handler: router,
	}

	err = server.ListenAndServe()
	helper.ErrorPanic(err)
}
