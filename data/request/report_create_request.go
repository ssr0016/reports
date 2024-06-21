package request

type ReportCreateRequest struct {
	MonthOf          string `validate:"required min=1,max=100" json:"month_of"`
	AreaOfAssignment string `validate:"required min=1,max=100" json:"area_of_assignment"`
	NameOfChurch     string `validate:"required min=1,max=100" json:"name_of_church"`
}
