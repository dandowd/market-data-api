package markets

import (
	"stock-price-api/config"

	"github.com/alpacahq/alpaca-trade-api-go/v3/alpaca"
)

type AlpacaClient struct {
	client *alpaca.Client
	config *config.Config
}

func NewAlpacaClient(config *config.Config) *AlpacaClient {
	return &AlpacaClient{
		client: alpaca.NewClient(config.AlpacaConfig),
		config: config,
	}
}
