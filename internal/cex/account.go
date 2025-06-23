package cex

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/warm3snow/cexcli/internal/types"
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
