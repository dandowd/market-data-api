package markets

type SymbolInfo struct {
	Symbol string  `json:"symbol"`
	Price  float64 `json:"price"`
}

type MarketClient interface {
	GetLatest(tickerSymbol string) (*SymbolInfo, error)
}
