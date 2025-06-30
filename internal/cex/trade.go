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

	if instType != "" {
		url += "?instType=" + instType
	}
	if instId != "" {
		url += "&instId=" + instId
	}
	if ordType != "" {
		url += "&ordType=" + ordType
	}
	if state != "" {
		url += "&state=" + state
	}
	if after != "" {
		url += "&after=" + after
	}
	if before != "" {
		url += "&before=" + before
	}
	if limit != "" {
		url += "&limit=" + limit
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

// GetOrder 获取订单信息
func (c *Client) GetOrder(instId, ordId, clOrdId string) ([]types.PendingOrdersResponse, error) {
	url := "/api/v5/trade/order"
	if instId != "" {
		url += "?instId=" + instId
	}
	if ordId != "" {
		url += "&ordId=" + ordId
	}
	if clOrdId != "" {
		url += "&clOrdId=" + clOrdId
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

// AmendOrder 修改订单
func (c *Client) AmendOrder(req *types.AmendOrderRequest) ([]types.AmendOrderResponse, error) {
	url := "/api/v5/trade/amend-order"

	resp, err := c.SendRequest("POST", url, req)
	if err != nil {
		return nil, fmt.Errorf("error making request: %w", err)
	}

	var result struct {
		Code string                     `json:"code"`
		Msg  string                     `json:"msg"`
		Data []types.AmendOrderResponse `json:"data"`
	}

	if err := json.Unmarshal(resp, &result); err != nil {
		return nil, fmt.Errorf("error decoding response: %w", err)
	}

	if result.Code != "0" {
		return nil, fmt.Errorf("API error code %s: %s", result.Code, result.Msg)
	}

	// 检查具体修改订单的返回状态
	for _, amendResult := range result.Data {
		if amendResult.SCode != "0" {
			return nil, fmt.Errorf("order amendment failed for ordId '%s': %s (sCode: %s)",
				amendResult.OrdId, amendResult.SMsg, amendResult.SCode)
		}
	}

	return result.Data, nil
}
