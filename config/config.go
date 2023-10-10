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
}

var configVars *Config
var logger *zap.Logger

func Init() {
	configVars = &Config{
		AlpacaApiKey:     os.Getenv("ALPACA_API_KEY"),
		AlpacaApiSecret:  os.Getenv("ALPACA_API_SECRET"),
		AlpacaApiBaseUrl: os.Getenv("ALPACA_API_BASE_URL"),
		PolygonApiKey:    os.Getenv("POLYGON_API_KEY"),
	}

	zapLogger, nil := zap.NewProduction()
	if nil != nil {
		panic("Failed to create logger")
	}

	logger = zapLogger
}

func GetVariables() *Config {
	return configVars
}

func GetLogger() *zap.Logger {
	return logger
}
