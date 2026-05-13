package engine

import (
	"dungeon-sim/internal/formatter"
	"dungeon-sim/internal/models"
	"time"
)

const (
	StateSuccess string = "SUCCESS"
	StateFail    string = "FAIL"
	StateDisqual string = "DISQUAL"
)

func newReport(t time.Time, comment string) models.Report {
	return models.Report{
		Time:    t,
		Comment: comment,
	}
}

func singleReport(t time.Time, comment string) []models.Report {
	return []models.Report{
		newReport(t, comment),
	}
}

func impossibleMove(event models.Event) []models.Report {
	return singleReport(event.Time, formatter.ImpossibleMove(
		event.PlayerID,
		event.EventID,
	),
	)
}
