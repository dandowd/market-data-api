package config

import (
	"os"

	"go.uber.org/zap"
)

type Config struct {
	PolygonApiKey string
	Logger *zap.Logger
}

func NewConfig() *Config {
	logger, nil := zap.NewProduction()
	if nil != nil {
		panic("Failed to create logger")
	}
	
	return &Config{
		PolygonApiKey:    os.Getenv("POLYGON_API_KEY"),
		Logger: logger,
	}
}
