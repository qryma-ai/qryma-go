// Package client contains the Qryma API client implementation.
package client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/qryma-ai/qryma-go/version"
)

// SearchOptions contains the options for a search request
type SearchOptions struct {
	Lang  string `json:"lang"`
	Start int    `json:"start"`
	Safe  bool   `json:"safe"`
	Mode  string `json:"mode"`
}

// ClientConfig contains the configuration for creating a Qryma client
type ClientConfig struct {
	APIKey  string
	BaseURL string
	Timeout time.Duration
}

// QrymaResponse represents the raw API response
type QrymaResponse map[string]interface{}

// QrymaClient is the client for interacting with the Qryma Search API
type QrymaClient struct {
	apiKey     string
	baseURL    string
	timeout    time.Duration
	httpClient *http.Client
	headers    map[string]string
}

// NewQrymaClient creates a new QrymaClient
//
// Parameters:
//   - apiKey: Qryma API key for authentication
//   - baseURL: Base URL for the Qryma API (default: https://search.qryma.com)
//   - timeout: Request timeout in seconds (default: 30)
//
// Returns:
//   - A new QrymaClient instance
//   - An error if the API key is invalid
func NewQrymaClient(apiKey string, opts ...ClientOption) (*QrymaClient, error) {
	if apiKey == "" {
		return nil, fmt.Errorf("API key must be provided")
	}

	client := &QrymaClient{
		apiKey:  apiKey,
		baseURL: "https://search.qryma.com",
		timeout: 30 * time.Second,
		headers: map[string]string{
			"X-Api-Key":    apiKey,
			"Content-Type": "application/json",
			"User-Agent":   fmt.Sprintf("qryma-go/%s", version.Version),
		},
	}

	for _, opt := range opts {
		opt(client)
	}

	client.httpClient = &http.Client{
		Timeout: client.timeout,
	}

	return client, nil
}

// ClientOption is a function type for configuring the client
type ClientOption func(*QrymaClient)

// WithBaseURL sets the base URL for the client
func WithBaseURL(baseURL string) ClientOption {
	return func(c *QrymaClient) {
		if baseURL != "" {
			if len(baseURL) > 0 && baseURL[len(baseURL)-1] == '/' {
				c.baseURL = baseURL[:len(baseURL)-1]
			} else {
				c.baseURL = baseURL
			}
		}
	}
}

// WithTimeout sets the timeout for the client
func WithTimeout(timeout time.Duration) ClientOption {
	return func(c *QrymaClient) {
		if timeout > 0 {
			c.timeout = timeout
		}
	}
}

// Search performs a search using the Qryma API
//
// Parameters:
//   - query: The search query string (required)
//   - options: Search options (optional)
//
// Returns:
//   - The raw API response containing the search results
//   - An error if there's an error with the request or response processing
func (c *QrymaClient) Search(query string, options ...SearchOptions) (QrymaResponse, error) {
	url := c.baseURL + "/api/web"

	var opts SearchOptions
	if len(options) > 0 {
		opts = options[0]
	}

	// 设置默认值：如果 Mode 为空，则默认使用 "snippet"
	mode := opts.Mode
	if mode == "" {
		mode = "snippet"
	}
	// 确保 Mode 值是有效的
	if mode != "snippet" && mode != "fulltext" {
		mode = "snippet"
	}

	payload := map[string]interface{}{
		"query": query,
		"lang":  opts.Lang,
		"start": opts.Start,
		"safe":  opts.Safe,
		"mode":  mode,
	}

	jsonPayload, err := json.Marshal(payload)
	if err != nil {
		return nil, fmt.Errorf("error marshaling request: %w", err)
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonPayload))
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}

	for key, value := range c.headers {
		req.Header.Set(key, value)
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("API request failed: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return nil, fmt.Errorf("API request failed: %d %s - %s", resp.StatusCode, resp.Status, string(body))
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response: %w", err)
	}

	var result QrymaResponse
	err = json.Unmarshal(body, &result)
	if err != nil {
		return nil, fmt.Errorf("error decoding JSON response: %w", err)
	}

	return result, nil
}

// Qryma creates a Qryma client instance
//
// Parameters:
//   - config: Client configuration
//   - config.APIKey: Qryma API key for authentication
//   - config.BaseURL: Base URL for the Qryma API (optional)
//   - config.Timeout: Request timeout (optional)
//
// Returns:
//   - A new QrymaClient instance
//   - An error if the configuration is invalid
//
// Example:
//
//	// To install: go get github.com/qryma-ai/qryma-go
//	client := qryma.Qryma(qryma.ClientConfig{
//		APIKey: "ak-********************",
//	})
//	response, err := client.Search("ces", qryma.SearchOptions{Lang: "zh-CN"})
func Qryma(config ClientConfig) (*QrymaClient, error) {
	var opts []ClientOption
	opts = append(opts, WithBaseURL(config.BaseURL))
	opts = append(opts, WithTimeout(config.Timeout))
	return NewQrymaClient(config.APIKey, opts...)
}
