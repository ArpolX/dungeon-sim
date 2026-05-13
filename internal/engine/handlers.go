package engine

import (
	"dungeon-sim/internal/formatter"
	"dungeon-sim/internal/models"
	"strconv"
)

func (g *Game) handleRegister(event models.Event) []models.Report {
	state := g.Player[event.PlayerID]

	if state.IsRegistered {
		return impossibleMove(event)
	}

	state.IsRegistered = true

	return singleReport(
		event.Time,
		formatter.Registered(event.PlayerID),
	)
}

func (g *Game) handleEnterDungeon(event models.Event) []models.Report {
	state := g.Player[event.PlayerID]

	if !state.IsRegistered {
		state.Final = StateDisqual

		return singleReport(
			event.Time,
			formatter.Disqualified(event.PlayerID),
		)
	}

	if !g.canEnterDungeon(state, event.Time) {
		return impossibleMove(event)
	}

	state.IsDungeon = true
	state.CurrentFloor = 1
	state.Entered = event.Time
	state.CurrentFloorStartedAt = event.Time

	return singleReport(
		event.Time,
		formatter.EnteredDungeon(event.PlayerID),
	)
}

func (g *Game) handleKillMonster(event models.Event) []models.Report {
	state := g.Player[event.PlayerID]

	if !g.canKillMonster(state) {
		return impossibleMove(event)
	}

	state.CurrentMonstersKill++
	state.AllMonstersKill++

	if state.CurrentMonstersKill == g.Config.Monsters {
		floorDuration := event.Time.Sub(state.CurrentFloorStartedAt)
		state.FloorDurations = append(state.FloorDurations, floorDuration)
	}

	return singleReport(
		event.Time,
		formatter.KilledMonster(event.PlayerID),
	)
}

func (g *Game) handleNextFloor(event models.Event) []models.Report {
	state := g.Player[event.PlayerID]

	if !g.canMoveNextFloor(state) {
		return impossibleMove(event)
	}

	state.CurrentFloor++
	state.CurrentMonstersKill = 0
	state.CurrentFloorStartedAt = event.Time

	return singleReport(
		event.Time,
		formatter.NextFloor(event.PlayerID),
	)
}

func (g *Game) handlePreviousFloor(event models.Event) []models.Report {
	state := g.Player[event.PlayerID]

	if !g.canMovePreviousFloor(state) {
		return impossibleMove(event)
	}

	state.CurrentFloor--

	return singleReport(
		event.Time,
		formatter.PreviousFloor(event.PlayerID),
	)
}

func (g *Game) handleEnterBossFloor(event models.Event) []models.Report {
	state := g.Player[event.PlayerID]

	if !g.canEnterBossFloor(state) {
		return impossibleMove(event)
	}

	state.BossEnteredAt = event.Time

	return singleReport(
		event.Time,
		formatter.EnteredBossFloor(event.PlayerID),
	)
}

func (g *Game) handleKillBoss(event models.Event) []models.Report {
	state := g.Player[event.PlayerID]

	if !g.canKillBoss(state) {
		return impossibleMove(event)
	}

	state.IsBossKilled = true
	state.Final = StateSuccess
	state.BossFightDuration = event.Time.Sub(state.BossEnteredAt)

	return singleReport(
		event.Time,
		formatter.KilledBoss(event.PlayerID),
	)
}

func (g *Game) handleLeaveDungeon(event models.Event) []models.Report {
	state := g.Player[event.PlayerID]

	if !g.canPlay(state) {
		return impossibleMove(event)
	}

	state.Exit = event.Time
	state.IsDungeon = false

	if state.Final == "" {
		state.Final = StateFail
	}

	return singleReport(
		event.Time,
		formatter.LeftDungeon(event.PlayerID),
	)
}

func (g *Game) handleCannotContinue(event models.Event) []models.Report {
	state := g.Player[event.PlayerID]

	if !g.canPlay(state) {
		return impossibleMove(event)
	}

	state.Final = StateDisqual
	state.IsDungeon = false
	state.Exit = event.Time

	return singleReport(
		event.Time,
		formatter.CannotContinue(
			event.PlayerID,
			event.ExtraParam,
		),
	)
}

func (g *Game) handleRestoreHealth(event models.Event) []models.Report {
	state := g.Player[event.PlayerID]

	if !g.canRestoreHealth(state) {
		return impossibleMove(event)
	}

	health, err := strconv.Atoi(event.ExtraParam)
	if err != nil {
		return impossibleMove(event)
	}

	state.CurrentHP += health

	if state.CurrentHP > 100 {
		state.CurrentHP = 100
	}

	return singleReport(
		event.Time,
		formatter.RestoredHealth(event.PlayerID, health),
	)
}

func (g *Game) handleReceiveDamage(event models.Event) []models.Report {
	state := g.Player[event.PlayerID]

	if !g.canReceiveDamage(state) {
		return impossibleMove(event)
	}

	damage, err := strconv.Atoi(event.ExtraParam)
	if err != nil {
		return impossibleMove(event)
	}

	state.CurrentHP -= damage

	reports := []models.Report{
		newReport(
			event.Time,
			formatter.ReceivedDamage(event.PlayerID, damage),
		),
	}

	if state.CurrentHP <= 0 {
		state.CurrentHP = 0
		state.IsPlayerKilled = true
		state.IsDungeon = false
		state.Final = StateFail
		state.Exit = event.Time

		reports = append(reports,
			newReport(
				event.Time,
				formatter.PlayerDead(event.PlayerID),
			),
		)
	}

	return reports
}
