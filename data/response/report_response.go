package response

import "time"

type ReportResponse struct {
	Id               int       `json:"id"`
	MonthOf          string    `json:"month_of"`
	AreaOfAssignment string    `json:"area_of_assignment"`
	NameOfChurch     string    `json:"name_of_church"`
	CreatedAt        time.Time `json:"created_at"`
	UpdatedAt        time.Time `json:"updated_at"`
}
