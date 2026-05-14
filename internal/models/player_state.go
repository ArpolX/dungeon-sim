package models

import "time"

type PlayerState struct {
	IsRegistered   bool
	IsDungeon      bool
	IsBossKilled   bool
	IsPlayerKilled bool

	CurrentFloor int

	CurrentMonstersKill int
	AllMonstersKill     int

	CurrentHP int

	Entered     time.Time
	Exit        time.Time
	LastEventAt time.Time

	CurrentFloorStartedAt time.Time

	BossEnteredAt     time.Time
	BossFightDuration time.Duration

	FloorDurations []time.Duration

	Final string
}
