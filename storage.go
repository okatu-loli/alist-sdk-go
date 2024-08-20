package alist

import (
	"encoding/json"
	"fmt"
)

// StorageData 包含存储信息的数据结构
type StorageData struct {
	ID              int    `json:"id"`
	MountPath       string `json:"mount_path"`
	Order           int    `json:"order"`
	Driver          string `json:"driver"`
	CacheExpiration int    `json:"cache_expiration"`
	Status          string `json:"status"`
	Addition        string `json:"addition"`
	Remark          string `json:"remark"`
	Modified        string `json:"modified"`
	Disabled        bool   `json:"disabled"`
	EnableSign      bool   `json:"enable_sign"`
	OrderBy         string `json:"order_by"`
	OrderDirection  string `json:"order_direction"`
	ExtractFolder   string `json:"extract_folder"`
	WebProxy        bool   `json:"web_proxy"`
	WebDavPolicy    string `json:"webdav_policy"`
	DownProxyURL    string `json:"down_proxy_url"`
}

// StorageListData 包含存储列表的数据结构
type StorageListData struct {
	Content []StorageData `json:"content"`
	Total   int           `json:"total"`
}

func (c *Client) ListStorages(page, perPage string) (*StorageListData, error) {
	endpoint := fmt.Sprintf("/api/admin/storage/list?page=%s&per_page=%s", page, perPage)
	response, err := c.doRequest("GET", endpoint, nil, "")
	if err != nil {
		return nil, err
	}

	dataBytes, err := json.Marshal(response.Data)
	if err != nil {
		return nil, err
	}

	var storageListData StorageListData
	if err := json.Unmarshal(dataBytes, &storageListData); err != nil {
		return nil, err
	}

	return &storageListData, nil
}

func (c *Client) EnableStorage(id int) (*Response, error) {
	endpoint := fmt.Sprintf("/api/admin/storage/enable?id=%d", id)
	return c.doRequest("POST", endpoint, nil, "")
}

func (c *Client) DisableStorage(id int) (*Response, error) {
	endpoint := fmt.Sprintf("/api/admin/storage/disable?id=%d", id)
	return c.doRequest("POST", endpoint, nil, "")
}

func (c *Client) CreateStorage(storage StorageData) (*Response, error) {
	endpoint := "/api/admin/storage/create"
	return c.doRequest("POST", endpoint, storage, "application/json")
}

func (c *Client) UpdateStorage(storage StorageData) (*Response, error) {
	endpoint := "/api/admin/storage/update"
	return c.doRequest("POST", endpoint, storage, "application/json")
}

func (c *Client) GetStorage(id int) (*StorageData, error) {
	endpoint := fmt.Sprintf("/api/admin/storage/get?id=%d", id)
	response, err := c.doRequest("GET", endpoint, nil, "")
	if err != nil {
		return nil, err
	}

	dataBytes, err := json.Marshal(response.Data)
	if err != nil {
		return nil, err
	}

	var storageData StorageData
	if err := json.Unmarshal(dataBytes, &storageData); err != nil {
		return nil, err
	}

	return &storageData, nil
}

func (c *Client) DeleteStorage(id int) (*Response, error) {
	endpoint := fmt.Sprintf("/api/admin/storage/delete?id=%d", id)
	return c.doRequest("POST", endpoint, nil, "")
}

func (c *Client) ReloadAllStorages() (*Response, error) {
	endpoint := "/api/admin/storage/load_all"
	return c.doRequest("POST", endpoint, nil, "")
}
