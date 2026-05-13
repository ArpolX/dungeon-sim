package engine

import (
	"context"
	logger "dungeon-sim/internal/infrastructure"
	"dungeon-sim/internal/models"
	"testing"

	"go.uber.org/zap"
)

func TestGame_GameSimulation(t *testing.T) {
	logger.Init()
	log := logger.GetLogger()

	type fields struct {
		Config *models.Config
		Player map[int]*models.PlayerState
	}
	type args struct {
		events []models.Event
		l      *zap.SugaredLogger
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{name: "one", fields: fields{
			Config: &models.Config{},
			Player: make(map[int]*models.PlayerState),
		}, args: args{
			events: []models.Event{{EventID: 1}, {EventID: 5}, {EventID: 7}, {EventID: 9}, {EventID: 12}, {EventID: 122}},
			l:      log,
		}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := &Game{
				Config: tt.fields.Config,
				Player: tt.fields.Player,
			}
			g.GameSimulation(context.Background(), tt.args.events, tt.args.l)
		})
	}
}
