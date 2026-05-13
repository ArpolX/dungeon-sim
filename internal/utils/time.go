package utils

import (
	"fmt"
	"time"
)

func ParseTime(t string) (time.Time, error) {
	Time, err := time.Parse("15:04:05", t)
	if err != nil {
		return time.Now(), fmt.Errorf("error parse time %w", err)
	}

	return Time, nil
}

func ParseStringTime(t time.Time) string {
	h := t.Hour()
	m := t.Minute()
	s := t.Second()

	return fmt.Sprintf("[%02d:%02d:%02d]", h, m, s)
}
