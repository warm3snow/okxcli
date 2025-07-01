package cex

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/go-resty/resty/v2"
	"github.com/warm3snow/cexcli/internal/config"
	"github.com/warm3snow/cexcli/logger"
)

// Client represents an CEX API client
type Client struct {
	config     *config.Config
	BaseURL    string
	HTTPClient *resty.Client
}

// NewClient creates a new CEX API client
func NewClient(cfg *config.Config) *Client {
	client := resty.New()
	client.SetTimeout(10 * time.Second)

	if cfg.CEX.API.IsSimulated {
		client.SetHeader("x-simulated-trading", "1")
	}

	return &Client{
		config:     cfg,
		BaseURL:    cfg.CEX.BaseURL,
		HTTPClient: client,
	}
}

// SetBaseURL sets a custom base URL for the API
func (c *Client) SetBaseURL(url string) {
	c.BaseURL = url
}

// SetSimulated sets whether to use simulated trading
func (c *Client) SetSimulated(simulated bool) {
	c.config.CEX.API.IsSimulated = simulated
}

// sign 生成 API 请求签名
func (c *Client) sign(timestamp, method, requestPath, body string) string {
	// 确保时间戳格式正确（ISO8601格式，包含毫秒）
	if !strings.Contains(timestamp, ".") {
		timestamp = strings.Replace(timestamp, "Z", ".000Z", 1)
	}
	message := timestamp + method + requestPath + body
	mac := hmac.New(sha256.New, []byte(c.config.CEX.API.SecretKey))
	mac.Write([]byte(message))
	return base64.StdEncoding.EncodeToString(mac.Sum(nil))
}

// sendRequest 发送请求到 OKX API
func (c *Client) SendRequest(method, requestPath string, body interface{}) ([]byte, error) {
	timestamp := time.Now().UTC().Format("2006-01-02T15:04:05.000Z")

	req := c.HTTPClient.R().
		SetHeaders(map[string]string{
			"OK-ACCESS-KEY":        c.config.CEX.API.APIKey,
			"OK-ACCESS-TIMESTAMP":  timestamp,
			"OK-ACCESS-PASSPHRASE": c.config.CEX.API.Passphrase,
			"Content-Type":         "application/json",
		})

	if body != nil {
		jsonBody, err := json.Marshal(body)
		if err != nil {
			return nil, fmt.Errorf("failed to marshal request body: %w", err)
		}
		req.SetBody(jsonBody)
		req.SetHeader("OK-ACCESS-SIGN", c.sign(timestamp, method, requestPath, string(jsonBody)))
	} else {
		req.SetHeader("OK-ACCESS-SIGN", c.sign(timestamp, method, requestPath, ""))
	}

	logger.Debugf("request headers: %v", req.Header)

	var resp *resty.Response
	var err error

	fullURL := c.config.CEX.BaseURL + requestPath

	switch method {
	case "GET":
		resp, err = req.Get(fullURL)
	case "POST":
		resp, err = req.Post(fullURL)
	case "DELETE":
		resp, err = req.Delete(fullURL)
	case "PUT":
		resp, err = req.Put(fullURL)
	default:
		return nil, fmt.Errorf("unsupported HTTP method: %s", method)
	}

	if err != nil {
		return nil, fmt.Errorf("request failed: %w", err)
	}

	if resp.StatusCode() != http.StatusOK {
		return nil, fmt.Errorf("API request failed with status %d: %s", resp.StatusCode(), resp.String())
	}

	return resp.Body(), nil
}
