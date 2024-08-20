package alist

import (
	"bytes"
	"net/http"
)

func (c *Client) Mkdir(path string) (*Response, error) {
	endpoint := "/api/fs/mkdir"
	bodyData := map[string]string{
		"path": path,
	}

	return c.doRequest("POST", endpoint, bodyData, "application/json")
}

func (c *Client) Rename(name, path string) (*Response, error) {
	endpoint := "/api/fs/rename"
	bodyData := map[string]string{
		"name": name,
		"path": path,
	}

	return c.doRequest("POST", endpoint, bodyData, "application/json")
}

func (c *Client) UploadFile(formData []byte, filePath string) (*Response, error) {
	endpoint := "/api/fs/form"
	req, err := http.NewRequest("PUT", c.BaseURL+endpoint, bytes.NewBuffer(formData))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "multipart/form-data")
	req.Header.Set("File-Path", filePath)
	req.Header.Set("Authorization", c.Token)

	return c.doRequest(req.Method, endpoint, formData, "multipart/form-data")
}

func (c *Client) List(path string) (*Response, error) {
	endpoint := "/api/fs/list"
	bodyData := map[string]interface{}{
		"path":     path,
		"password": "",
		"page":     1,
		"per_page": 0,
		"refresh":  false,
	}

	return c.doRequest("POST", endpoint, bodyData, "application/json")
}

func (c *Client) GetFileInfo(path, password string, page, perPage int, refresh bool) (*Response, error) {
	endpoint := "/api/fs/get"
	bodyData := map[string]interface{}{
		"path":     path,
		"password": password,
		"page":     page,
		"per_page": perPage,
		"refresh":  refresh,
	}

	return c.doRequest("POST", endpoint, bodyData, "application/json")
}

func (c *Client) SearchFiles(parent, keywords string, scope, page, perPage int, password string) (*Response, error) {
	endpoint := "/api/fs/search"
	bodyData := map[string]interface{}{
		"parent":   parent,
		"keywords": keywords,
		"scope":    scope,
		"page":     page,
		"per_page": perPage,
		"password": password,
	}

	return c.doRequest("POST", endpoint, bodyData, "application/json")
}

func (c *Client) GetDirectories(path, password string, forceRoot bool) (*Response, error) {
	endpoint := "/api/fs/dirs"
	bodyData := map[string]interface{}{
		"path":       path,
		"password":   password,
		"force_root": forceRoot,
	}

	return c.doRequest("POST", endpoint, bodyData, "application/json")
}

func (c *Client) BatchRename(srcDir string, renameObjects []map[string]string) (*Response, error) {
	endpoint := "/api/fs/batch_rename"
	bodyData := map[string]interface{}{
		"src_dir":        srcDir,
		"rename_objects": renameObjects,
	}

	return c.doRequest("POST", endpoint, bodyData, "application/json")
}

func (c *Client) RegexRename(srcDir, srcNameRegex, newNameRegex string) (*Response, error) {
	endpoint := "/api/fs/regex_rename"
	bodyData := map[string]string{
		"src_dir":        srcDir,
		"src_name_regex": srcNameRegex,
		"new_name_regex": newNameRegex,
	}

	return c.doRequest("POST", endpoint, bodyData, "application/json")
}

func (c *Client) MoveFile(srcDir, dstDir string, names []string) (*Response, error) {
	endpoint := "/api/fs/move"
	bodyData := map[string]interface{}{
		"src_dir": srcDir,
		"dst_dir": dstDir,
		"names":   names,
	}

	return c.doRequest("POST", endpoint, bodyData, "application/json")
}

func (c *Client) RecursiveMove(srcDir, dstDir string) (*Response, error) {
	endpoint := "/api/fs/recursive_move"
	bodyData := map[string]string{
		"src_dir": srcDir,
		"dst_dir": dstDir,
	}

	return c.doRequest("POST", endpoint, bodyData, "application/json")
}

func (c *Client) CopyFile(srcDir, dstDir string, names []string) (*Response, error) {
	endpoint := "/api/fs/copy"
	bodyData := map[string]interface{}{
		"src_dir": srcDir,
		"dst_dir": dstDir,
		"names":   names,
	}

	return c.doRequest("POST", endpoint, bodyData, "application/json")
}

func (c *Client) RemoveFiles(dir string, names []string) (*Response, error) {
	endpoint := "/api/fs/remove"
	bodyData := map[string]interface{}{
		"dir":   dir,
		"names": names,
	}

	return c.doRequest("POST", endpoint, bodyData, "application/json")
}

func (c *Client) RemoveEmptyDirectory(srcDir string) (*Response, error) {
	endpoint := "/api/fs/remove_empty_directory"
	bodyData := map[string]string{
		"src_dir": srcDir,
	}

	return c.doRequest("POST", endpoint, bodyData, "application/json")
}

func (c *Client) StreamUpload(filePath string, data []byte) (*Response, error) {
	endpoint := "/api/fs/put"
	req, err := http.NewRequest("PUT", c.BaseURL+endpoint, bytes.NewBuffer(data))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/octet-stream")
	req.Header.Set("File-Path", filePath)
	req.Header.Set("Authorization", c.Token)

	return c.doRequest(req.Method, endpoint, data, "application/octet-stream")
}

func (c *Client) AddOfflineDownload(path string, urls []string, tool, deletePolicy string) (*Response, error) {
	endpoint := "/api/fs/add_offline_download"
	bodyData := map[string]interface{}{
		"path":          path,
		"urls":          urls,
		"tool":          tool,
		"delete_policy": deletePolicy,
	}

	return c.doRequest("POST", endpoint, bodyData, "application/json")
}
