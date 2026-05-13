package models

import "time"

type Report struct {
	Time    time.Time
	Comment string
}
