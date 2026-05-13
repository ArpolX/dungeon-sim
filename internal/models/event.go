package models

import "time"

type Event struct {
	Time       time.Time
	PlayerID   int
	EventID    int
	ExtraParam string
}
