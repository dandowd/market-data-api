package config

import (
	"os"

	"go.uber.org/zap"
)

type Config struct {
	PolygonApiKey    string
	AlpacaApiKey     string
	AlpacaApiSecret  string
	AlpacaApiBaseUrl string
	Logger           *zap.Logger
}

func NewConfig() *Config {
	logger, nil := zap.NewProduction()
	if nil != nil {
		panic("Failed to create logger")
	}

	return &Config{
		AlpacaApiKey:     os.Getenv("ALPACA_API_KEY"),
		AlpacaApiSecret:  os.Getenv("ALPACA_API_SECRET"),
		AlpacaApiBaseUrl: os.Getenv("ALPACA_API_BASE_URL"),
		PolygonApiKey:    os.Getenv("POLYGON_API_KEY"),
		Logger:           logger,
	}
}
