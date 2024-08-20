package alist

import (
	"encoding/json"
)

// SettingsData 包含配置的数据结构
type SettingsData struct {
	AllowIndexed            string `json:"allow_indexed"`
	AllowMounted            string `json:"allow_mounted"`
	Announcement            string `json:"announcement"`
	AudioAutoplay           string `json:"audio_autoplay"`
	AudioCover              string `json:"audio_cover"`
	AutoUpdateIndex         string `json:"auto_update_index"`
	DefaultPageSize         string `json:"default_page_size"`
	ExternalPreviews        string `json:"external_previews"`
	Favicon                 string `json:"favicon"`
	FilenameCharMapping     string `json:"filename_char_mapping"`
	ForwardDirectLinkParams string `json:"forward_direct_link_params"`
	HideFiles               string `json:"hide_files"`
	HomeContainer           string `json:"home_container"`
	HomeIcon                string `json:"home_icon"`
	IframePreviews          string `json:"iframe_previews"`
	Logo                    string `json:"logo"`
	MainColor               string `json:"main_color"`
	OcrApi                  string `json:"ocr_api"`
	PackageDownload         string `json:"package_download"`
	PaginationType          string `json:"pagination_type"`
	RobotsTxt               string `json:"robots_txt"`
	SearchIndex             string `json:"search_index"`
	SettingsLayout          string `json:"settings_layout"`
	SiteTitle               string `json:"site_title"`
	SsoLoginEnabled         string `json:"sso_login_enabled"`
	SsoLoginPlatform        string `json:"sso_login_platform"`
	Version                 string `json:"version"`
	VideoAutoplay           string `json:"video_autoplay"`
}

func (c *Client) Ping() (bool, error) {
	endpoint := "/ping"
	response, err := c.doRequest("GET", endpoint, nil, "")
	if err != nil {
		return false, err
	}

	return response.Code == 200, nil
}

func (c *Client) GetSettings() (*SettingsData, error) {
	endpoint := "/api/public/settings"
	response, err := c.doRequest("GET", endpoint, nil, "")
	if err != nil {
		return nil, err
	}

	// Decode the response data into SettingsData
	settingsDataBytes, err := json.Marshal(response.Data)
	if err != nil {
		return nil, err
	}

	var settingsData SettingsData
	if err := json.Unmarshal(settingsDataBytes, &settingsData); err != nil {
		return nil, err
	}

	return &settingsData, nil
}
