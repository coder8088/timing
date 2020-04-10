package model

import (
	"time"
)

type Timing struct {
	Id              int64      `json:"id"`
	Name            string     `json:"name"`
	DurationSeconds int        `json:"duration_seconds"`
	StartTime       time.Time  `json:"start_time"`
	StopTime        *time.Time `json:"stop_time"`
	Dt              string     `json:"dt"`
}
