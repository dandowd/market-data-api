package markets

type SymbolInfo struct {
  Symbol string
  Price float64
}

type MarketClient interface {
  GetLatest(tickerSymbol string) (*SymbolInfo, error)
}
