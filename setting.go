package alist

import (
	"encoding/json"
	"fmt"
)

// SettingData 包含每个设置项的数据结构
type SettingData struct {
	Key     string `json:"key"`
	Value   string `json:"value"`
	Help    string `json:"help"`
	Type    string `json:"type"`
	Options string `json:"options"`
	Group   int    `json:"group"`
	Flag    int    `json:"flag"`
}

func (c *Client) ListSettings(groups, group string) ([]SettingData, error) {
	endpoint := fmt.Sprintf("/api/admin/setting/list?groups=%s&group=%s", groups, group)
	response, err := c.doRequest("GET", endpoint, nil, "")
	if err != nil {
		return nil, err
	}

	dataBytes, err := json.Marshal(response.Data)
	if err != nil {
		return nil, err
	}

	var settings []SettingData
	if err := json.Unmarshal(dataBytes, &settings); err != nil {
		return nil, err
	}

	return settings, nil
}

func (c *Client) GetSetting(key string) (*SettingData, error) {
	endpoint := fmt.Sprintf("/api/admin/setting/get?key=%s", key)
	response, err := c.doRequest("GET", endpoint, nil, "")
	if err != nil {
		return nil, err
	}

	dataBytes, err := json.Marshal(response.Data)
	if err != nil {
		return nil, err
	}

	var setting SettingData
	if err := json.Unmarshal(dataBytes, &setting); err != nil {
		return nil, err
	}

	return &setting, nil
}

func (c *Client) SaveSettings(settings []SettingData) (*Response, error) {
	endpoint := "/api/admin/setting/save"
	return c.doRequest("POST", endpoint, settings, "application/json")
}

func (c *Client) DeleteSetting(key string) (*Response, error) {
	endpoint := fmt.Sprintf("/api/admin/setting/delete?key=%s", key)
	return c.doRequest("POST", endpoint, nil, "")
}

func (c *Client) ResetToken() (string, error) {
	endpoint := "/api/admin/setting/reset_token"
	response, err := c.doRequest("POST", endpoint, nil, "")
	if err != nil {
		return "", err
	}

	if token, ok := response.Data.(string); ok {
		return token, nil
	}

	return "", fmt.Errorf("failed to reset token: %v", response.Message)
}

func (c *Client) SetAria2(uri, secret string) (string, error) {
	endpoint := "/api/admin/setting/set_aria2"
	bodyData := map[string]string{"uri": uri, "secret": secret}
	response, err := c.doRequest("POST", endpoint, bodyData, "application/json")
	if err != nil {
		return "", err
	}

	if version, ok := response.Data.(string); ok {
		return version, nil
	}

	return "", fmt.Errorf("failed to set Aria2: %v", response.Message)
}

func (c *Client) SetQBit(url, seedtime string) (string, error) {
	endpoint := "/api/admin/setting/set_qbit"
	bodyData := map[string]string{"url": url, "seedtime": seedtime}
	response, err := c.doRequest("POST", endpoint, bodyData, "application/json")
	if err != nil {
		return "", err
	}

	if version, ok := response.Data.(string); ok {
		return version, nil
	}

	return "", fmt.Errorf("failed to set qBittorrent: %v", response.Message)
}
