package config

import "os"

type Config struct {
	PolygonApiKey string
}

func NewConfig() *Config {
	return &Config{
		PolygonApiKey:    os.Getenv("POLYGON_API_KEY"),
	}
}
