package engine

import (
	"context"
	"dungeon-sim/internal/models"
	rep "dungeon-sim/internal/report"
	"fmt"

	"go.uber.org/zap"
)

type Game struct {
	Config *models.Config
	Player map[int]*models.PlayerState
}

func NewGame(cfg *models.Config) *Game {
	return &Game{
		Config: cfg,
		Player: make(map[int]*models.PlayerState),
	}
}

func (g *Game) GameSimulation(ctx context.Context, events []models.Event, l *zap.SugaredLogger) {
	for _, event := range events {
		var report []models.Report

		if _, ok := g.Player[event.PlayerID]; !ok {
			g.Player[event.PlayerID] = &models.PlayerState{
				CurrentHP: 100,
			}
		}

		state := g.Player[event.PlayerID]
		if state.Final == StateDisqual || state.Final == StateFail || state.Final == StateSuccess {
			continue
		}

		switch event.EventID {
		case 1:
			report = g.handleRegister(event)
		case 2:
			report = g.handleEnterDungeon(event)
		case 3:
			report = g.handleKillMonster(event)
		case 4:
			report = g.handleNextFloor(event)
		case 5:
			report = g.handlePreviousFloor(event)
		case 6:
			report = g.handleEnterBossFloor(event)
		case 7:
			report = g.handleKillBoss(event)
		case 8:
			report = g.handleLeaveDungeon(event)
		case 9:
			report = g.handleCannotContinue(event)
		case 10:
			report = g.handleRestoreHealth(event)
		case 11:
			report = g.handleReceiveDamage(event)
		default:
			l.Warnf("No command number %d", event.EventID)
			fmt.Printf("No command number %d", event.EventID)
		}
		rep.OutputLogs(report)
	}

	rep.FinalOutput(g.Player)
}
