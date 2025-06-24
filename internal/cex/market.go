package cex

import (
	"encoding/json"
	"fmt"

	"github.com/warm3snow/cexcli/internal/types"
)

// GetTicker fetches ticker information for a given symbol
func (c *Client) GetTicker(symbol string) (*types.Ticker, error) {
	url := fmt.Sprintf("/api/v5/market/ticker?instId=%s", symbol)

	// Set headers
	resp, err := c.SendRequest("GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("error making request: %w", err)
	}

	var result struct {
		Code string         `json:"code"`
		Msg  string         `json:"msg"`
		Data []types.Ticker `json:"data"`
	}

	if err := json.Unmarshal(resp, &result); err != nil {
		return nil, fmt.Errorf("error decoding response: %w", err)
	}

	if result.Code != "0" {
		return nil, fmt.Errorf("API error: %s", result.Msg)
	}

	if len(result.Data) == 0 {
		return nil, fmt.Errorf("no data returned for symbol %s", symbol)
	}

	return &result.Data[0], nil
}
