package model

import (
	"math"
	"time"
)

type Report struct {
	Id                int
	MonthOf           string
	WorkerName        string
	AreaOfAssignment  string
	NameOfChurch      string
	WorshipService    []int
	AverageAttendance float64
	CreatedAt         time.Time
	UpdatedAt         time.Time
}

func (r *Report) CalculateAverage() {
	if len(r.WorshipService) == 0 {
		r.AverageAttendance = 0
		return
	}

	total := 0
	for _, attendance := range r.WorshipService {
		total += attendance
	}

	average := float64(total) / float64(len(r.WorshipService))
	r.AverageAttendance = math.Round(average)
}
