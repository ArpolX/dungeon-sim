package engine

import (
	"dungeon-sim/internal/models"
	"dungeon-sim/internal/utils"
	"time"
)

func canPlay(state *models.PlayerState) bool {
	return state.IsRegistered && !state.IsPlayerKilled
}

func canEnterDungeon(state *models.PlayerState, eventTime time.Time, cfg *models.Config) bool {
	openAt, _ := utils.ParseTime(cfg.OpenAt)

	closeTime := openAt.Add(
		time.Duration(cfg.Duration) * time.Hour,
	)

	return state.IsRegistered && !state.IsDungeon && !state.IsPlayerKilled && eventTime.Before(closeTime)
}

func canKillMonster(state *models.PlayerState, cfg *models.Config) bool {
	return canPlay(state) &&
		state.IsDungeon &&
		state.CurrentFloor > 0 &&
		state.CurrentFloor <= cfg.Floors &&
		state.CurrentMonstersKill < cfg.Monsters
}

func canMoveNextFloor(state *models.PlayerState, cfg *models.Config) bool {
	return canPlay(state) &&
		state.IsDungeon &&
		state.CurrentMonstersKill == cfg.Monsters &&
		state.CurrentFloor < cfg.Floors
}

func canMovePreviousFloor(state *models.PlayerState) bool {
	return canPlay(state) &&
		state.IsDungeon &&
		state.CurrentFloor > 1
}

func canEnterBossFloor(state *models.PlayerState, cfg *models.Config) bool {
	return canPlay(state) &&
		state.IsDungeon &&
		state.CurrentFloor == cfg.Floors &&
		state.CurrentMonstersKill == 0 && !state.IsBossKilled
}

func canKillBoss(state *models.PlayerState, cfg *models.Config) bool {
	return canPlay(state) &&
		state.IsDungeon &&
		state.CurrentFloor == cfg.Floors &&
		!state.IsBossKilled
}

func canRestoreHealth(state *models.PlayerState) bool {
	return canPlay(state) && state.IsDungeon
}

func canReceiveDamage(state *models.PlayerState) bool {
	return canPlay(state) && state.IsDungeon
}

func canSuccessPlayer(state *models.PlayerState, cfg *models.Config) bool {
	return canPlay(state) &&
		state.IsDungeon &&
		state.CurrentFloor == cfg.Floors &&
		state.IsBossKilled &&
		state.AllMonstersKill == cfg.Monsters*(cfg.Floors-1)
}
