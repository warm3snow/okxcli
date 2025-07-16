package okx

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/warm3snow/okxcli/internal/types"
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

// GetAccountPositions 查询持仓
func (c *Client) GetAccountPositions(params map[string]string) ([]types.AccountPosition, error) {
	url := "/api/v5/account/positions"
	if len(params) > 0 {
		var arr []string
		for k, v := range params {
			arr = append(arr, fmt.Sprintf("%s=%s", k, v))
		}
		url += "?" + strings.Join(arr, "&")
	}
	resp, err := c.SendRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	var result struct {
		Code string                  `json:"code"`
		Msg  string                  `json:"msg"`
		Data []types.AccountPosition `json:"data"`
	}
	if err := json.Unmarshal(resp, &result); err != nil {
		return nil, err
	}
	if result.Code != "0" {
		return nil, fmt.Errorf("API error: %s", result.Msg)
	}
	return result.Data, nil
}

// GetAccountPositionRisk 查询账户持仓风险
func (c *Client) GetAccountPositionRisk(instType string) ([]types.AccountPositionRisk, error) {
	url := "/api/v5/account/account-position-risk"
	if instType != "" {
		url += "?instType=" + instType
	}
	resp, err := c.SendRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	var result struct {
		Code string                      `json:"code"`
		Msg  string                      `json:"msg"`
		Data []types.AccountPositionRisk `json:"data"`
	}
	if err := json.Unmarshal(resp, &result); err != nil {
		return nil, err
	}
	if result.Code != "0" {
		return nil, fmt.Errorf("API error: %s", result.Msg)
	}
	return result.Data, nil
}

// GetAccountConfig 获取账户配置信息
func (c *Client) GetAccountConfig() (*types.AccountConfig, error) {
	url := "/api/v5/account/config"
	resp, err := c.SendRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	var result struct {
		Code string                `json:"code"`
		Msg  string                `json:"msg"`
		Data []types.AccountConfig `json:"data"`
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

// SetLeverage 设置杠杆倍数
func (c *Client) SetLeverage(req *types.SetLeverageRequest) (*types.SetLeverageResponse, error) {
	url := "/api/v5/account/set-leverage"
	resp, err := c.SendRequest("POST", url, req)
	if err != nil {
		return nil, err
	}
	var result struct {
		Code string                      `json:"code"`
		Msg  string                      `json:"msg"`
		Data []types.SetLeverageResponse `json:"data"`
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

// SetPositionMode 设置持仓模式
func (c *Client) SetPositionMode(req *types.SetPositionModeRequest) (*types.SetPositionModeResponse, error) {
	url := "/api/v5/account/set-position-mode"
	resp, err := c.SendRequest("POST", url, req)
	if err != nil {
		return nil, err
	}
	var result struct {
		Code string                          `json:"code"`
		Msg  string                          `json:"msg"`
		Data []types.SetPositionModeResponse `json:"data"`
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
