package okx

import (
	"encoding/json"
	"fmt"

	"github.com/warm3snow/okxcli/internal/types"
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

// GetTickers 获取所有产品行情
func (c *Client) GetTickers(instType, uly, instFamily string) ([]types.Ticker, error) {
	url := "/api/v5/market/tickers"
	if instType != "" {
		url += "?instType=" + instType
	}
	if uly != "" {
		url += "?uly=" + uly
	}
	if instFamily != "" {
		url += "?instFamily=" + instFamily
	}
	resp, err := c.SendRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	var result struct {
		Code string         `json:"code"`
		Msg  string         `json:"msg"`
		Data []types.Ticker `json:"data"`
	}
	if err := json.Unmarshal(resp, &result); err != nil {
		return nil, err
	}
	if result.Code != "0" {
		return nil, fmt.Errorf("API error: %s", result.Msg)
	}
	return result.Data, nil
}
