package alist

import (
	"encoding/json"
	"fmt"
)

// UserData 包含用户信息的数据结构
type UserData struct {
	ID         int    `json:"id"`
	Username   string `json:"username"`
	Password   string `json:"password"`
	BasePath   string `json:"base_path"`
	Role       int    `json:"role"`
	Disabled   bool   `json:"disabled"`
	Permission int    `json:"permission"`
	SsoID      string `json:"sso_id"`
}

// UsersListData 包含所有用户信息的数据结构
type UsersListData struct {
	Content []UserData `json:"content"`
	Total   int        `json:"total"`
}

func (c *Client) ListAllUsers() (*UsersListData, error) {
	endpoint := "/api/admin/user/list"
	response, err := c.doRequest("GET", endpoint, nil, "")
	if err != nil {
		return nil, err
	}

	// Decode the response data
	usersListDataBytes, err := json.Marshal(response.Data)
	if err != nil {
		return nil, err
	}

	var usersListData UsersListData
	if err := json.Unmarshal(usersListDataBytes, &usersListData); err != nil {
		return nil, err
	}

	return &usersListData, nil
}

func (c *Client) GetUser(id string) (*UserData, error) {
	endpoint := fmt.Sprintf("/api/admin/user/get?id=%s", id)
	response, err := c.doRequest("GET", endpoint, nil, "")
	if err != nil {
		return nil, err
	}

	userDataBytes, err := json.Marshal(response.Data)
	if err != nil {
		return nil, err
	}

	var userData UserData
	if err := json.Unmarshal(userDataBytes, &userData); err != nil {
		return nil, err
	}

	return &userData, nil
}

func (c *Client) CreateUser(user UserData) (*Response, error) {
	endpoint := "/api/admin/user/create"
	return c.doRequest("POST", endpoint, user, "application/json")
}

func (c *Client) UpdateUser(user UserData) (*Response, error) {
	endpoint := "/api/admin/user/update"
	return c.doRequest("POST", endpoint, user, "application/json")
}

func (c *Client) CancelUser2FA(id string) (*Response, error) {
	endpoint := fmt.Sprintf("/api/admin/user/cancel_2fa?id=%s", id)
	return c.doRequest("POST", endpoint, nil, "")
}

func (c *Client) DeleteUser(id string) (*Response, error) {
	endpoint := fmt.Sprintf("/api/admin/user/delete?id=%s", id)
	return c.doRequest("POST", endpoint, nil, "")
}

func (c *Client) DeleteUserCache(username string) (*Response, error) {
	endpoint := fmt.Sprintf("/api/admin/user/del_cache?username=%s", username)
	return c.doRequest("POST", endpoint, nil, "")
}
