package alist

import (
	"crypto/sha256"
	"encoding/hex"
)

// PostLogin 发送登录请求，包含用户名、密码和OTP码
func (c *Client) PostLogin(username, password, otpCode string) (*Response, error) {
	endpoint := "/api/auth/login"
	bodyData := map[string]string{
		"username": username,
		"password": password,
		"otp_code": otpCode,
	}

	return c.doRequest("POST", endpoint, bodyData, "application/json")
}

// PostLoginHash 发送带有哈希密码的登录请求，包含用户名、哈希密码和OTP码
func (c *Client) PostLoginHash(username, password, otpCode string) (*Response, error) {
	endpoint := "/api/auth/login/hash"
	hashedPassword := computeHashedPassword(password)
	bodyData := map[string]string{
		"username": username,
		"password": hashedPassword,
		"otp_code": otpCode,
	}

	return c.doRequest("POST", endpoint, bodyData, "application/json")
}

// Generate2FA 生成两步验证码
func (c *Client) Generate2FA() (*Response, error) {
	endpoint := "/api/auth/2fa/generate"
	return c.doRequest("POST", endpoint, nil, "application/json")
}

// Verify2FA 验证两步验证码
func (c *Client) Verify2FA(code, secret string) (*Response, error) {
	endpoint := "/api/auth/2fa/verify"
	bodyData := map[string]string{
		"code":   code,
		"secret": secret,
	}

	return c.doRequest("POST", endpoint, bodyData, "application/json")
}

// GetUserInfo 获取用户信息
func (c *Client) GetUserInfo() (*Response, error) {
	endpoint := "/api/me"
	return c.doRequest("GET", endpoint, nil, "")
}

// computeHashedPassword 计算密码的哈希值
func computeHashedPassword(password string) string {
	hasher := sha256.New()
	hasher.Write([]byte(password + "-https://github.com/alist-org/alist"))
	return hex.EncodeToString(hasher.Sum(nil))
}
