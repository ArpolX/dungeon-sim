package main

import (
	"context"
	"dungeon-sim/internal/engine"
	logger "dungeon-sim/internal/infrastructure"
	"dungeon-sim/internal/parser"
)

func main() {
	ctx := context.Background()

	logger.Init()
	l := logger.GetLogger()

	config, err := parser.ConfigParser(l)
	if err != nil {
		l.Fatalf("error config parser %v", err)
	}
	events, err := parser.EventsParser(l)
	if err != nil {
		l.Fatalf("error events parser %v", err)
	}

	game := engine.NewGame(config)

	game.GameSimulation(ctx, events, l)
}
