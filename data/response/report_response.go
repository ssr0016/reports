package response

import "time"

type ReportResponse struct {
	Id                int       `json:"id"`
	MonthOf           string    `json:"month_of"`
	WorkerName        string    `json:"worker_name"`
	AreaOfAssignment  string    `json:"area_of_assignment"`
	NameOfChurch      string    `json:"name_of_church"`
	WorshipService    []int     `json:"worship_service"`
	AverageAttendance float64   `json:"average_attendance"`
	CreatedAt         time.Time `json:"created_at"`
	UpdatedAt         time.Time `json:"updated_at"`
}
