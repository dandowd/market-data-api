package config

import "os"

type Config struct {
	AlpacaAPIKey    string
	AlpacaAPISecret string
	AlpacaBaseURL   string
}

func NewConfig() *Config {
	return &Config{
		AlpacaAPIKey:    os.Getenv("ALPACA_API_KEY"),
		AlpacaAPISecret: os.Getenv("ALPACA_API_SECRET"),
		AlpacaBaseURL:   os.Getenv("ALPACA_BASE_URL"),
	}
}
