package request

type ReportCreateRequest struct {
	MonthOf           string  `validate:"required,min=1,max=100" json:"month_of"`
	WorkerName        string  `validate:"required,min=1,max=100" json:"worker_name"`
	AreaOfAssignment  string  `validate:"required,min=1,max=100" json:"area_of_assignment"`
	NameOfChurch      string  `validate:"required,min=1,max=100" json:"name_of_church"`
	WorshipService    []int   `validate:"required" json:"worship_service"`
	AverageAttendance float64 `json:"average_attendance"`
}
