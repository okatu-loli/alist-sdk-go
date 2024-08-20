package alist

import (
	"encoding/json"
	"fmt"
)

// DriverInfo 包含驱动的详细信息结构
type DriverInfo struct {
	Common     []DriverConfig      `json:"common"`
	Additional []DriverConfig      `json:"additional"`
	Config     DriverConfigDetails `json:"config"`
}

// DriverConfig 包含驱动配置的字段信息
type DriverConfig struct {
	Name     string `json:"name"`
	Type     string `json:"type"`
	Default  string `json:"default"`
	Options  string `json:"options"`
	Required bool   `json:"required"`
	Help     string `json:"help"`
}

// DriverConfigDetails 为具体驱动配置的信息
type DriverConfigDetails struct {
	Name        string `json:"name"`
	LocalSort   bool   `json:"local_sort"`
	OnlyLocal   bool   `json:"only_local"`
	OnlyProxy   bool   `json:"only_proxy"`
	NoCache     bool   `json:"no_cache"`
	NoUpload    bool   `json:"no_upload"`
	NeedMs      bool   `json:"need_ms"`
	DefaultRoot string `json:"default_root"`
	Alert       string `json:"alert"`
}

func (c *Client) ListDriverTemplates() (interface{}, error) {
	endpoint := "/api/admin/driver/list"
	response, err := c.doRequest("GET", endpoint, nil, "")
	if err != nil {
		return nil, err
	}

	return response.Data, nil
}

func (c *Client) ListDriverNames() ([]string, error) {
	endpoint := "/api/admin/driver/names"
	response, err := c.doRequest("GET", endpoint, nil, "")
	if err != nil {
		return nil, err
	}

	// 转换数据为期望的类型
	driverNames, ok := response.Data.([]interface{})
	if !ok {
		return nil, fmt.Errorf("unexpected data format for driver names")
	}

	names := make([]string, len(driverNames))
	for i, name := range driverNames {
		names[i], ok = name.(string)
		if !ok {
			return nil, fmt.Errorf("unexpected data type for driver name")
		}
	}

	return names, nil
}

func (c *Client) GetDriverInfo(driverName string) (*DriverInfo, error) {
	endpoint := fmt.Sprintf("/api/admin/driver/info?driver=%s", driverName)
	response, err := c.doRequest("GET", endpoint, nil, "")
	if err != nil {
		return nil, err
	}

	driverInfoBytes, err := json.Marshal(response.Data)
	if err != nil {
		return nil, err
	}

	var driverInfo DriverInfo
	if err := json.Unmarshal(driverInfoBytes, &driverInfo); err != nil {
		return nil, err
	}

	return &driverInfo, nil
}
