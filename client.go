package alist

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

// Client 结构体表示一个HTTP客户端，包含基础URL、Token和HTTP客户端实例
type Client struct {
	BaseURL    string
	Token      string
	HTTPClient *http.Client
}

// Response 结构体表示HTTP响应，包含状态码、消息和数据
type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

// NewClient 函数用于创建一个新的Client实例
func NewClient(baseURL, token string) *Client {
	return &Client{
		BaseURL:    baseURL,
		Token:      token,
		HTTPClient: &http.Client{},
	}
}

// doRequest 方法用于执行HTTP请求，并返回响应结果或错误
func (c *Client) doRequest(method, endpoint string, bodyData interface{}, contentType string) (*Response, error) {
	url := c.BaseURL + endpoint

	var body *bytes.Buffer
	if bodyData != nil {
		jsonData, err := json.Marshal(bodyData)
		if err != nil {
			return nil, err
		}
		body = bytes.NewBuffer(jsonData)
	} else {
		body = &bytes.Buffer{}
	}

	req, err := http.NewRequest(method, url, body)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Authorization", c.Token)
	if contentType != "" {
		req.Header.Set("Content-Type", contentType)
	}

	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var response Response
	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		return nil, err
	}
	if response.Code != http.StatusOK {
		return nil, fmt.Errorf("request failed: %v", response.Message)
	}

	return &response, nil
}
