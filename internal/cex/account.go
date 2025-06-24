package cex

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/warm3snow/cexcli/internal/types"
)

// GetAccountBalance 获取账户余额
func (c *Client) GetAccountBalance(ccys ...string) ([]types.AccountBalance, error) {
	url := "/api/v5/account/balance"
	if len(ccys) > 0 {
		url += "?ccy=" + strings.Join(ccys, ",")
	}
	resp, err := c.SendRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	var result struct {
		Code string                 `json:"code"`
		Msg  string                 `json:"msg"`
		Data []types.AccountBalance `json:"data"`
	}
	if err := json.Unmarshal(resp, &result); err != nil {
		return nil, err
	}
	if result.Code != "0" {
		return nil, fmt.Errorf("API error: %s", result.Msg)
	}
	return result.Data, nil
}
