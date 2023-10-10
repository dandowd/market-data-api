package markets

import (
	"context"
	"stock-price-api/config"

	polygon "github.com/polygon-io/client-go/rest"
	"github.com/polygon-io/client-go/rest/models"
	"go.uber.org/zap"
)

type polygonClient struct {
	client *polygon.Client
	config *config.Config
}

func NewPolygonClient(config *config.Config) MarketClient {
	return &polygonClient{
		client: polygon.New(config.PolygonApiKey),
		config: config,
	}
}

func (c *polygonClient) GetLatest(tickerSymbol string) (*SymbolInfo, error) {
	resp, err := c.client.GetLastTrade(context.Background(), &models.GetLastTradeParams{Ticker: tickerSymbol})
	if err != nil {
		c.config.Logger.Error("Error getting last trade", zap.Error(err))
		return nil, err
	}

	return &SymbolInfo{
		Symbol: tickerSymbol,
		Price:  resp.Results.Price,
	}, nil
}
