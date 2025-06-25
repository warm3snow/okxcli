package cex

import (
	"encoding/json"
	"fmt"

	"github.com/warm3snow/cexcli/internal/types"
)

// GetInstruments 获取所有可交易产品的信息列表
func (c *Client) GetInstruments(instType, uly, instFamily, instId string) ([]types.Instrument, error) {
	url := "/api/v5/public/instruments"

	if instType != "" {
		url += "?instType=" + instType
	}
	if uly != "" {
		url += "?uly=" + uly
	}
	if instFamily != "" {
		url += "?instFamily=" + instFamily
	}
	if instId != "" {
		url += "?instId=" + instId
	}

	resp, err := c.SendRequest("GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("error making request: %w", err)
	}

	var result struct {
		Code string             `json:"code"`
		Msg  string             `json:"msg"`
		Data []types.Instrument `json:"data"`
	}

	if err := json.Unmarshal(resp, &result); err != nil {
		return nil, fmt.Errorf("error decoding response: %w", err)
	}

	if result.Code != "0" {
		return nil, fmt.Errorf("API error code %s: %s", result.Code, result.Msg)
	}

	return result.Data, nil
}
