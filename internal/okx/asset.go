package okx

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/warm3snow/okxcli/internal/types"
)

// GetBalances fetches the account balances for all currencies or specified currencies
func (c *Client) GetBalances(currencies ...string) ([]types.Balance, error) {
	url := "/api/v5/asset/balances"
	if len(currencies) > 0 {
		// Join currencies with comma, max 20 currencies allowed
		if len(currencies) > 20 {
			currencies = currencies[:20]
		}
		url += fmt.Sprintf("?ccy=%s", strings.Join(currencies, ","))
	}

	resp, err := c.SendRequest("GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("error making request: %w", err)
	}

	var result struct {
		Code string          `json:"code"`
		Msg  string          `json:"msg"`
		Data []types.Balance `json:"data"`
	}

	if err := json.Unmarshal(resp, &result); err != nil {
		return nil, fmt.Errorf("error decoding response: %w", err)
	}

	if result.Code != "0" {
		return nil, fmt.Errorf("API error: %s", result.Msg)
	}

	return result.Data, nil
}

// GetAssetValuation 获取账户资产估值
func (c *Client) GetAssetValuation(ccy string) (*types.AssetValuation, error) {
	url := "/api/v5/asset/asset-valuation"
	if ccy != "" {
		url += "?ccy=" + ccy
	}

	resp, err := c.SendRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	var result struct {
		Code string                 `json:"code"`
		Msg  string                 `json:"msg"`
		Data []types.AssetValuation `json:"data"`
	}
	if err := json.Unmarshal(resp, &result); err != nil {
		return nil, err
	}
	if result.Code != "0" {
		return nil, fmt.Errorf("API error: %s", result.Msg)
	}
	if len(result.Data) == 0 {
		return nil, fmt.Errorf("no data returned")
	}
	return &result.Data[0], nil
}
