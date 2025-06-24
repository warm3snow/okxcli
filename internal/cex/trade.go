package cex

import (
	"encoding/json"
	"fmt"

	"github.com/warm3snow/cexcli/internal/types"
)

// PlaceOrder 执行下单操作
func (c *Client) PlaceOrder(req *types.PlaceOrderRequest) ([]types.PlaceOrderResponse, error) {
	url := "/api/v5/trade/order"

	resp, err := c.SendRequest("POST", url, req)
	if err != nil {
		return nil, fmt.Errorf("error making request: %w", err)
	}

	var result struct {
		Code string                     `json:"code"`
		Msg  string                     `json:"msg"`
		Data []types.PlaceOrderResponse `json:"data"`
	}

	if err := json.Unmarshal(resp, &result); err != nil {
		return nil, fmt.Errorf("error decoding response: %w", err)
	}

	if result.Code != "0" {
		return nil, fmt.Errorf("API error code %s: %s", result.Code, result.Msg)
	}

	// 检查具体订单的返回状态
	for _, orderResult := range result.Data {
		if orderResult.SCode != "0" {
			return nil, fmt.Errorf("order placement failed for clOrdId '%s': %s (sCode: %s)",
				orderResult.ClOrdID, orderResult.SMsg, orderResult.SCode)
		}
	}

	return result.Data, nil
}

// CancelOrder 执行撤单操作
func (c *Client) CancelOrder(req *types.CancelOrderRequest) ([]types.CancelOrderResponse, error) {
	url := "/api/v5/trade/cancel-order"

	resp, err := c.SendRequest("POST", url, req)
	if err != nil {
		return nil, fmt.Errorf("error making request: %w", err)
	}

	var result struct {
		Code string                      `json:"code"`
		Msg  string                      `json:"msg"`
		Data []types.CancelOrderResponse `json:"data"`
	}

	if err := json.Unmarshal(resp, &result); err != nil {
		return nil, fmt.Errorf("error decoding response: %w", err)
	}

	if result.Code != "0" {
		return nil, fmt.Errorf("API error code %s: %s", result.Code, result.Msg)
	}

	// 检查具体撤单的返回状态
	for _, cancelResult := range result.Data {
		if cancelResult.SCode != "0" {
			return nil, fmt.Errorf("order cancellation failed for ordId '%s': %s (sCode: %s)",
				cancelResult.OrdID, cancelResult.SMsg, cancelResult.SCode)
		}
	}

	return result.Data, nil
}

// GetPendingOrders 获取未完成订单列表
func (c *Client) GetPendingOrders(instType, instId, ordType, state, after, before, limit string) ([]types.PendingOrdersResponse, error) {
	url := "/api/v5/trade/orders-pending"

	params := make(map[string]string)
	if instType != "" {
		params["instType"] = instType
	}
	if instId != "" {
		params["instId"] = instId
	}
	if ordType != "" {
		params["ordType"] = ordType
	}
	if state != "" {
		params["state"] = state
	}
	if after != "" {
		params["after"] = after
	}
	if before != "" {
		params["before"] = before
	}
	if limit != "" {
		params["limit"] = limit
	}

	resp, err := c.SendRequest("GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("error making request: %w", err)
	}

	var result struct {
		Code string                        `json:"code"`
		Msg  string                        `json:"msg"`
		Data []types.PendingOrdersResponse `json:"data"`
	}

	if err := json.Unmarshal(resp, &result); err != nil {
		return nil, fmt.Errorf("error decoding response: %w", err)
	}

	if result.Code != "0" {
		return nil, fmt.Errorf("API error code %s: %s", result.Code, result.Msg)
	}

	return result.Data, nil
}
