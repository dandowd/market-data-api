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
}

func NewPolygonClient(config *config.Config) MarketClient {
	return &polygonClient{
		client: polygon.New(config.PolygonApiKey),
	}
}

func (c *polygonClient) GetLatest(tickerSymbol string) (*SymbolInfo, error) {
	logger := config.GetLogger()

	resp, err := c.client.GetLastTrade(context.Background(), &models.GetLastTradeParams{Ticker: tickerSymbol})
	if err != nil {
		logger.Error("Error getting last trade", zap.Error(err))
		return nil, err
	}

	return &SymbolInfo{
		Symbol: tickerSymbol,
		Price:  resp.Results.Price,
	}, nil
}
