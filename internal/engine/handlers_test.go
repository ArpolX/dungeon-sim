package engine

import (
	"dungeon-sim/internal/models"
	"reflect"
	"testing"
	"time"
)

func mustParseTime(value string) time.Time {
	t, err := time.Parse("15:04:05", value)
	if err != nil {
		panic(err)
	}

	return t
}

func TestGame_handleRegister(t *testing.T) {
	eventTime := mustParseTime("14:00:00")

	tests := []struct {
		name   string
		player map[int]*models.PlayerState
		event  models.Event
		want   []models.Report
	}{
		{
			name: "successful registration",

			player: map[int]*models.PlayerState{
				1: {
					CurrentHP: 100,
				},
			},

			event: models.Event{
				Time:     eventTime,
				PlayerID: 1,
				EventID:  1,
			},

			want: []models.Report{
				{
					Time:    eventTime,
					Comment: "Player [1] registered",
				},
			},
		},

		{
			name: "duplicate registration",

			player: map[int]*models.PlayerState{
				1: {
					IsRegistered: true,
					CurrentHP:    100,
				},
			},

			event: models.Event{
				Time:     eventTime,
				PlayerID: 1,
				EventID:  1,
			},

			want: []models.Report{
				{
					Time:    eventTime,
					Comment: "Player [1] makes imposible move [1]",
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			g := &Game{
				Player: tt.player,
			}

			got := g.handleRegister(tt.event)

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf(
					"handleRegister() = %v, want %v",
					got,
					tt.want,
				)
			}
		})
	}
}

func TestGame_handleKillMonster(t *testing.T) {
	eventTime := mustParseTime("14:10:00")

	tests := []struct {
		name   string
		player map[int]*models.PlayerState
		event  models.Event
		want   []models.Report
	}{
		{
			name: "successful monster kill",

			player: map[int]*models.PlayerState{
				1: {
					IsRegistered: true,
					IsDungeon:    true,
					CurrentFloor: 1,
					CurrentHP:    100,
				},
			},

			event: models.Event{
				Time:     eventTime,
				PlayerID: 1,
				EventID:  3,
			},

			want: []models.Report{
				{
					Time:    eventTime,
					Comment: "Player [1] killed the monster",
				},
			},
		},

		{
			name: "kill monster outside dungeon",

			player: map[int]*models.PlayerState{
				1: {
					IsRegistered: true,
					IsDungeon:    false,
					CurrentHP:    100,
				},
			},

			event: models.Event{
				Time:     eventTime,
				PlayerID: 1,
				EventID:  3,
			},

			want: []models.Report{
				{
					Time:    eventTime,
					Comment: "Player [1] makes imposible move [3]",
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			g := &Game{
				Config: &models.Config{
					Monsters: 2,
					Floors:   3,
				},
				Player: tt.player,
			}

			got := g.handleKillMonster(tt.event)

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf(
					"handleKillMonster() = %v, want %v",
					got,
					tt.want,
				)
			}
		})
	}
}

func TestGame_handleNextFloor(t *testing.T) {
	eventTime := mustParseTime("14:20:00")

	tests := []struct {
		name   string
		player map[int]*models.PlayerState
		event  models.Event
		want   []models.Report
	}{
		{
			name: "successful next floor",

			player: map[int]*models.PlayerState{
				1: {
					IsRegistered:        true,
					IsDungeon:           true,
					CurrentFloor:        1,
					CurrentMonstersKill: 2,
					CurrentHP:           100,
				},
			},

			event: models.Event{
				Time:     eventTime,
				PlayerID: 1,
				EventID:  4,
			},

			want: []models.Report{
				{
					Time:    eventTime,
					Comment: "Player [1] went to the next floor",
				},
			},
		},

		{
			name: "next floor without enough kills",

			player: map[int]*models.PlayerState{
				1: {
					IsRegistered:        true,
					IsDungeon:           true,
					CurrentFloor:        1,
					CurrentMonstersKill: 1,
					CurrentHP:           100,
				},
			},

			event: models.Event{
				Time:     eventTime,
				PlayerID: 1,
				EventID:  4,
			},

			want: []models.Report{
				{
					Time:    eventTime,
					Comment: "Player [1] makes imposible move [4]",
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			g := &Game{
				Config: &models.Config{
					Monsters: 2,
					Floors:   3,
				},
				Player: tt.player,
			}

			got := g.handleNextFloor(tt.event)

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf(
					"handleNextFloor() = %v, want %v",
					got,
					tt.want,
				)
			}
		})
	}
}

func TestGame_handleRestoreHealth(t *testing.T) {
	eventTime := mustParseTime("14:30:00")

	tests := []struct {
		name   string
		player map[int]*models.PlayerState
		event  models.Event
		wantHP int
	}{
		{
			name: "restore health",

			player: map[int]*models.PlayerState{
				1: {
					IsRegistered: true,
					IsDungeon:    true,
					CurrentHP:    50,
				},
			},

			event: models.Event{
				Time:       eventTime,
				PlayerID:   1,
				EventID:    10,
				ExtraParam: "20",
			},

			wantHP: 70,
		},

		{
			name: "health clamp to 100",

			player: map[int]*models.PlayerState{
				1: {
					IsRegistered: true,
					IsDungeon:    true,
					CurrentHP:    90,
				},
			},

			event: models.Event{
				Time:       eventTime,
				PlayerID:   1,
				EventID:    10,
				ExtraParam: "50",
			},

			wantHP: 100,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			g := &Game{
				Player: tt.player,
			}

			g.handleRestoreHealth(tt.event)

			gotHP := g.Player[1].CurrentHP

			if gotHP != tt.wantHP {
				t.Errorf(
					"CurrentHP = %d, want %d",
					gotHP,
					tt.wantHP,
				)
			}
		})
	}
}

func TestGame_handleReceiveDamage(t *testing.T) {
	eventTime := mustParseTime("14:40:00")

	tests := []struct {
		name      string
		player    map[int]*models.PlayerState
		event     models.Event
		wantHP    int
		wantFinal string
	}{
		{
			name: "receive damage",

			player: map[int]*models.PlayerState{
				1: {
					IsRegistered: true,
					IsDungeon:    true,
					CurrentHP:    100,
				},
			},

			event: models.Event{
				Time:       eventTime,
				PlayerID:   1,
				EventID:    11,
				ExtraParam: "30",
			},

			wantHP: 70,
		},

		{
			name: "player dies",

			player: map[int]*models.PlayerState{
				1: {
					IsRegistered: true,
					IsDungeon:    true,
					CurrentHP:    50,
				},
			},

			event: models.Event{
				Time:       eventTime,
				PlayerID:   1,
				EventID:    11,
				ExtraParam: "100",
			},

			wantHP:    0,
			wantFinal: "FAIL",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			g := &Game{
				Player: tt.player,
			}

			g.handleReceiveDamage(tt.event)

			player := g.Player[1]

			if player.CurrentHP != tt.wantHP {
				t.Errorf(
					"CurrentHP = %d, want %d",
					player.CurrentHP,
					tt.wantHP,
				)
			}

			if tt.wantFinal != "" &&
				player.Final != tt.wantFinal {

				t.Errorf(
					"Final = %s, want %s",
					player.Final,
					tt.wantFinal,
				)
			}
		})
	}
}
