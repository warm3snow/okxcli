package types

// Balance represents the balance information for a currency
type Balance struct {
	Currency     string `json:"ccy"`       // Currency name, e.g., BTC
	Balance      string `json:"bal"`       // Total balance
	FrozenBal    string `json:"frozenBal"` // Frozen balance
	AvailableBal string `json:"availBal"`  // Available balance
}
