package parser

import (
	"bufio"
	"context"
	"dungeon-sim/internal/models"
	"dungeon-sim/internal/utils"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"go.uber.org/zap"
)

func EventsParser(ctx context.Context, l *zap.SugaredLogger) ([]models.Event, error) {
	cwd, err := os.Getwd()
	if err != nil {
		l.Errorf("error os.Getwd %v", err)
		return nil, fmt.Errorf("error os.Getwd %w", err)
	}

	path := filepath.Join(cwd, "data", "events.txt")
	file, err := os.Open(path)
	if err != nil {
		l.Errorf("error os.Open %v", err)
		return nil, fmt.Errorf("error os.Open %w", err)
	}
	defer file.Close()

	var events []models.Event

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		text := scanner.Text()

		text = strings.TrimSpace(text)
		e := strings.Split(text, " ")

		if len(e) == 0 {
			break
		}

		var event models.Event

		str := e[0]
		str = strings.ReplaceAll(str, "[", "")
		str = strings.ReplaceAll(str, "]", "")
		event.Time, err = utils.ParseTime(str)
		if err != nil {
			l.Errorf("error parse time %v", err)
			return nil, fmt.Errorf("error parse time %w", err)
		}

		event.PlayerID, err = strconv.Atoi(e[1])
		if err != nil {
			l.Errorf("error string->int %v", err)
			return nil, fmt.Errorf("error string->int %w", err)
		}

		event.EventID, err = strconv.Atoi(e[2])
		if err != nil {
			l.Errorf("error string->int %v", err)
			return nil, fmt.Errorf("error string->int %w", err)
		}

		if len(e) > 3 {
			event.ExtraParam = strings.Join(e[3:], " ")
		}

		events = append(events, event)
	}

	return events, nil
}
