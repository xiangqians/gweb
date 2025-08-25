// @author xiangqian
// @date 2025/08/02 22:49
package model

import "time"

// TokenReq 令牌请求信息
type TokenReq struct {
	Name     string `json:"name"`     // 用户名
	Password string `json:"password"` // 密码
}

// TokenResp 令牌响应信息
type TokenResp struct {
	AccessToken string    `json:"accessToken"` // 访问令牌
	ExpiresAt   time.Time `json:"expiresAt"`   // 过期时间
}
