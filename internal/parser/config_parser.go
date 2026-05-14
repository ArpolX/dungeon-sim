package parser

import (
	"dungeon-sim/internal/models"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"

	"go.uber.org/zap"
)

func ConfigParser(l *zap.SugaredLogger) (*models.Config, error) {
	cwd, err := os.Getwd()
	if err != nil {
		l.Errorf("error os.Getwd %v", err)
		return nil, fmt.Errorf("error os.Getwd %w", err)
	}

	path := filepath.Join(cwd, "data", "config.json")
	file, err := os.Open(path)
	if err != nil {
		l.Errorf("error os.Open %v", err)
		return nil, fmt.Errorf("error os.Open %w", err)
	}
	defer file.Close()

	var config models.Config
	enc := json.NewDecoder(file)

	if err := enc.Decode(&config); err != nil {
		l.Errorf("error encode struct %v", err)
		return nil, fmt.Errorf("error encode struct %w", err)
	}

	return &config, nil
}
