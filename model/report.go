package model

import "time"

type Report struct {
	Id               int
	MonthOf          string
	AreaOfAssignment string
	NameOfChurch     string
	CreatedAt        time.Time
	UpdatedAt        time.Time
}
