package engine

import (
	"dungeon-sim/internal/models"
	"dungeon-sim/internal/utils"
	"time"
)

func (g *Game) canPlay(state *models.PlayerState) bool {
	return state.IsRegistered &&
		!state.IsPlayerKilled
}

func (g *Game) canEnterDungeon(state *models.PlayerState, eventTime time.Time) bool {
	openAt, _ := utils.ParseTime(g.Config.OpenAt)

	closeTime := openAt.Add(
		time.Duration(g.Config.Duration) * time.Hour,
	)

	return state.IsRegistered &&
		!state.IsDungeon &&
		!state.IsPlayerKilled &&
		eventTime.Before(closeTime)
}

func (g *Game) canKillMonster(state *models.PlayerState) bool {
	return g.canPlay(state) &&
		state.IsDungeon &&
		state.CurrentFloor > 0 &&
		state.CurrentFloor <= g.Config.Floors &&
		state.CurrentMonstersKill < g.Config.Monsters
}

func (g *Game) canMoveNextFloor(state *models.PlayerState) bool {
	return g.canPlay(state) &&
		state.IsDungeon &&
		state.CurrentMonstersKill == g.Config.Monsters &&
		state.CurrentFloor < g.Config.Floors
}

func (g *Game) canMovePreviousFloor(state *models.PlayerState) bool {
	return g.canPlay(state) &&
		state.IsDungeon &&
		state.CurrentFloor > 1
}

func (g *Game) canEnterBossFloor(state *models.PlayerState) bool {
	return g.canPlay(state) &&
		state.IsDungeon &&
		state.CurrentFloor == g.Config.Floors &&
		state.CurrentMonstersKill == 0 && !state.IsBossKilled
}

func (g *Game) canKillBoss(state *models.PlayerState) bool {
	return g.canPlay(state) &&
		state.IsDungeon &&
		state.CurrentFloor == g.Config.Floors &&
		!state.IsBossKilled
}

func (g *Game) canRestoreHealth(state *models.PlayerState) bool {
	return g.canPlay(state) && state.IsDungeon
}

func (g *Game) canReceiveDamage(state *models.PlayerState) bool {
	return g.canPlay(state) && state.IsDungeon
}
