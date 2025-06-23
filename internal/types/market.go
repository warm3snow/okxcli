package types

// Ticker represents the ticker information for a trading pair from OKX
type Ticker struct {
	InstType  string `json:"instType"`  // Instrument type
	InstID    string `json:"instId"`    // Instrument ID
	Last      string `json:"last"`      // Last traded price
	LastSz    string `json:"lastSz"`    // Last traded size
	AskPx     string `json:"askPx"`     // Best ask price
	AskSz     string `json:"askSz"`     // Best ask size
	BidPx     string `json:"bidPx"`     // Best bid price
	BidSz     string `json:"bidSz"`     // Best bid size
	Open24h   string `json:"open24h"`   // 24-hour open price
	High24h   string `json:"high24h"`   // 24-hour highest price
	Low24h    string `json:"low24h"`    // 24-hour lowest price
	VolCcy24h string `json:"volCcy24h"` // 24-hour volume in currency
	Vol24h    string `json:"vol24h"`    // 24-hour volume
	Ts        string `json:"ts"`        // Timestamp
	SodUtc0   string `json:"sodUtc0"`   // Open price at UTC 0
	SodUtc8   string `json:"sodUtc8"`   // Open price at UTC+8
}
