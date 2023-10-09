package markets

import (
	"context"
	"stock-price-api/config"

	polygon "github.com/polygon-io/client-go/rest"
	"github.com/polygon-io/client-go/rest/models"
)

type polygonClient struct {
  client *polygon.Client
}

func NewPolygonClient(config *config.Config) *polygonClient {
	return &polygonClient{
		client: polygon.New(config.PolygonApiKey),
	}
}

func (c *polygonClient) GetLatestPrice(tickerSymbol string) (*SymbolInfo, error) {
	resp, err := c.client.GetLastTrade(context.Background(), &models.GetLastTradeParams{ Ticker: tickerSymbol })
	if err != nil {
		// Add logging here
		return nil, err
	}

	return &SymbolInfo{
		Symbol: tickerSymbol,
		Price: resp.Results.Price,
	}, nil
}
