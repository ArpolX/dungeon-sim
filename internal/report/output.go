package report

import (
	"dungeon-sim/internal/models"
	"dungeon-sim/internal/utils"
	"fmt"
	"sort"
	"time"
)

type exitFinal struct {
	id     int
	report string
}

func OutputLogs(report []models.Report) {
	for _, rep := range report {
		fmt.Printf("%v %v\n", utils.ParseStringTime(rep.Time), rep.Comment)
	}
}

func FinalOutput(state map[int]*models.PlayerState) {
	var final []exitFinal
	for k, v := range state {
		if v.Final == "" {
			v.Final = "FAIL"
		}

		var dungeonTime time.Duration

		if !v.Exit.IsZero() && !v.Entered.IsZero() {
			dungeonTime = v.Exit.Sub(v.Entered)
		}

		rep := fmt.Sprintf("[%s] %d [%v, %v, %v] HP:%d", v.Final, k, formatDuration(dungeonTime), formatDuration(calculateAverageFloorTime(v.FloorDurations)), formatDuration(v.BossFightDuration), v.CurrentHP)

		final = append(final, exitFinal{id: k, report: rep})
	}

	sort.Slice(final, func(i, j int) bool {
		return final[i].id < final[j].id
	})

	for _, f := range final {
		fmt.Println(f.report)
	}
}

func calculateAverageFloorTime(floorDurations []time.Duration) time.Duration {
	if len(floorDurations) == 0 {
		return 0
	}

	var total time.Duration

	for _, d := range floorDurations {
		total += d
	}

	return total / time.Duration(len(floorDurations))
}

func formatDuration(d time.Duration) string {
	h := int(d.Hours())
	m := int(d.Minutes()) % 60
	s := int(d.Seconds()) % 60

	return fmt.Sprintf("%02d:%02d:%02d", h, m, s)
}
