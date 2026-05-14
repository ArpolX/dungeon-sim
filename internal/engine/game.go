package engine

import (
	"context"
	"dungeon-sim/internal/models"
	rep "dungeon-sim/internal/report"

	"go.uber.org/zap"
)

const (
	EventRegister = iota + 1
	EventEnterDungeon
	EventKillMonster
	EventNextFloor
	EventPreviousFloor
	EventEnterBossFloor
	EventKillBoss
	EventLeaveDungeon
	EventCannotContinue
	EventRestoreHealth
	EventReceiveDamage
)

type eventHandler func(event models.Event) []models.Report

type Game struct {
	Config   *models.Config
	Player   map[int]*models.PlayerState
	Handlers map[int]eventHandler
}

func NewGame(cfg *models.Config) *Game {
	game := &Game{
		Config: cfg,
		Player: make(map[int]*models.PlayerState),
	}

	game.Handlers = game.initializeHandlers()

	return game
}

func (g *Game) GameSimulation(ctx context.Context, events []models.Event, l *zap.SugaredLogger) {
	for _, event := range events {
		select {
		case <-ctx.Done():
			l.Warn("context done")
			return
		default:
		}

		var report []models.Report

		state, ok := g.Player[event.PlayerID]
		if !ok {
			state = &models.PlayerState{
				CurrentHP: 100,
			}

			g.Player[event.PlayerID] = state
		}
		if state.Final != "" {
			continue
		}
		if state.IsDungeon {
			state.LastEventAt = event.Time
		}

		if v, ok := g.Handlers[event.EventID]; ok {
			report = v(event)
		} else {
			l.Warnf("No command number %d", event.EventID)
		}

		rep.OutputLogs(report)
	}
	g.finalizeStates()

	rep.FinalOutput(g.Player)
}

func (g *Game) initializeHandlers() map[int]eventHandler {
	return map[int]eventHandler{
		EventRegister:       g.handleRegister,
		EventEnterDungeon:   g.handleEnterDungeon,
		EventKillMonster:    g.handleKillMonster,
		EventNextFloor:      g.handleNextFloor,
		EventPreviousFloor:  g.handlePreviousFloor,
		EventEnterBossFloor: g.handleEnterBossFloor,
		EventKillBoss:       g.handleKillBoss,
		EventLeaveDungeon:   g.handleLeaveDungeon,
		EventCannotContinue: g.handleCannotContinue,
		EventRestoreHealth:  g.handleRestoreHealth,
		EventReceiveDamage:  g.handleReceiveDamage,
	}
}

func (g *Game) finalizeStates() {
	for _, state := range g.Player {
		if state.Final == "" {
			state.Final = "FAIL"
		}

		if state.Exit.IsZero() {
			state.Exit = state.LastEventAt
		}
	}
}
