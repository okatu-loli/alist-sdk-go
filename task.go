package alist

import (
	"encoding/json"
	"fmt"
)

// TaskData 包含任务信息的数据结构
type TaskData struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	State    string `json:"state"`
	Status   string `json:"status"`
	Progress int    `json:"progress"`
	Error    string `json:"error"`
}

// GetTaskInfo 获取特定任务的信息
func (c *Client) GetTaskInfo(taskType, taskID string) ([]TaskData, error) {
	endpoint := fmt.Sprintf("/api/admin/task/%s/info?tid=%s", taskType, taskID)
	return c.requestTasksData(endpoint)
}

// ListCompletedTasks 列出已完成的任务
func (c *Client) ListCompletedTasks(taskType string) ([]TaskData, error) {
	endpoint := fmt.Sprintf("/api/admin/task/%s/done", taskType)
	return c.requestTasksData(endpoint)
}

// ListUndoneTasks 列出未完成的任务
func (c *Client) ListUndoneTasks(taskType string) ([]TaskData, error) {
	endpoint := fmt.Sprintf("/api/admin/task/%s/undone", taskType)
	return c.requestTasksData(endpoint)
}

// DeleteTask 删除指定的任务
func (c *Client) DeleteTask(taskType, taskID string) (*Response, error) {
	endpoint := fmt.Sprintf("/api/admin/task/%s/delete?tid=%s", taskType, taskID)
	return c.doRequest("POST", endpoint, nil, "")
}

// CancelTask 取消指定的任务
func (c *Client) CancelTask(taskType, taskID string) (*Response, error) {
	endpoint := fmt.Sprintf("/api/admin/task/%s/cancel?tid=%s", taskType, taskID)
	return c.doRequest("POST", endpoint, nil, "")
}

// ClearCompletedTasks 清除已完成的任务
func (c *Client) ClearCompletedTasks(taskType string) (*Response, error) {
	endpoint := fmt.Sprintf("/api/admin/task/%s/clear_done", taskType)
	return c.doRequest("POST", endpoint, nil, "")
}

// ClearSucceededTasks 清除已成功的任务
func (c *Client) ClearSucceededTasks(taskType string) (*Response, error) {
	endpoint := fmt.Sprintf("/api/admin/task/%s/clear_succeeded", taskType)
	return c.doRequest("POST", endpoint, nil, "")
}

// RetryTask 重试指定的任务
func (c *Client) RetryTask(taskType, taskID string) (*Response, error) {
	endpoint := fmt.Sprintf("/api/admin/task/%s/retry?tid=%s", taskType, taskID)
	return c.doRequest("POST", endpoint, nil, "")
}

// requestTasksData 发起GET请求以获取任务数据
func (c *Client) requestTasksData(endpoint string) ([]TaskData, error) {
	// 发起请求并获得响应
	response, err := c.doRequest("GET", endpoint, nil, "")
	if err != nil {
		return nil, err
	}

	// 解析返回数据
	dataBytes, err := json.Marshal(response.Data)
	if err != nil {
		return nil, err
	}

	var tasks []TaskData
	if err := json.Unmarshal(dataBytes, &tasks); err != nil {
		return nil, err
	}

	return tasks, nil
}
