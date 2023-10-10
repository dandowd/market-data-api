package markets

import (
	"stock-price-api/config"

	"github.com/alpacahq/alpaca-trade-api-go/v3/marketdata"
	"go.uber.org/zap"
)

type AlpacaClient struct {
	client *marketdata.Client
	logger *zap.Logger
}

// GetLatest implements MarketClient.
func (a *AlpacaClient) GetLatest(tickerSymbol string) (*SymbolInfo, error) {
	tradeInfo, err := a.client.GetTrades(tickerSymbol, marketdata.GetTradesRequest{})
	if err != nil {
		a.logger.Error("Failed to get latest trade info", zap.Error(err))
	}

	if len(tradeInfo) == 0 {
		return nil, nil
	}

	return &SymbolInfo{
		Symbol: tickerSymbol,
		Price:  tradeInfo[0].Price,
	}, nil
}

func NewAlpacaClient() MarketClient {
	logger := config.GetLogger()
	config := config.GetVariables()
	return &AlpacaClient{
		logger: logger,
		client: marketdata.NewClient(marketdata.ClientOpts{
			APIKey:    config.AlpacaApiKey,
			APISecret: config.AlpacaApiSecret,
			BaseURL:   config.AlpacaApiBaseUrl,
		}),
	}
}
