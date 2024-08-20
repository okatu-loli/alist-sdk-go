package alist

import (
	"encoding/json"
	"fmt"
)

// MetaData 包含元信息的数据结构
type MetaData struct {
	ID       int    `json:"id"`
	Path     string `json:"path"`
	Password string `json:"password"`
	PSub     bool   `json:"p_sub"`
	Write    bool   `json:"write"`
	WSub     bool   `json:"w_sub"`
	Hide     string `json:"hide"`
	HSub     bool   `json:"h_sub"`
	Readme   string `json:"readme"`
	RSub     bool   `json:"r_sub"`
}

// MetaListData 包含元信息列表的数据结构
type MetaListData struct {
	Content []MetaData `json:"content"`
	Total   int        `json:"total"`
}

func (c *Client) ListMeta(page, perPage string) (*MetaListData, error) {
	endpoint := fmt.Sprintf("/api/admin/meta/list?page=%s&per_page=%s", page, perPage)
	response, err := c.doRequest("GET", endpoint, nil, "")
	if err != nil {
		return nil, err
	}

	metaListDataBytes, err := json.Marshal(response.Data)
	if err != nil {
		return nil, err
	}

	var metaListData MetaListData
	if err := json.Unmarshal(metaListDataBytes, &metaListData); err != nil {
		return nil, err
	}

	return &metaListData, nil
}

func (c *Client) GetMeta(id string) (*MetaData, error) {
	endpoint := fmt.Sprintf("/api/admin/meta/get?id=%s", id)
	response, err := c.doRequest("GET", endpoint, nil, "")
	if err != nil {
		return nil, err
	}

	metaDataBytes, err := json.Marshal(response.Data)
	if err != nil {
		return nil, err
	}

	var metaData MetaData
	if err := json.Unmarshal(metaDataBytes, &metaData); err != nil {
		return nil, err
	}

	return &metaData, nil
}

func (c *Client) CreateMeta(meta MetaData) (*Response, error) {
	endpoint := "/api/admin/meta/create"
	return c.doRequest("POST", endpoint, meta, "application/json")
}

func (c *Client) UpdateMeta(meta MetaData) (*Response, error) {
	endpoint := "/api/admin/meta/update"
	return c.doRequest("POST", endpoint, meta, "application/json")
}

func (c *Client) DeleteMeta(id string) (*Response, error) {
	endpoint := fmt.Sprintf("/api/admin/meta/delete?id=%s", id)
	return c.doRequest("POST", endpoint, nil, "")
}
