// @author xiangqian
// @date 2025/08/03 00:07
package xhttp

import (
	"fmt"
	"gweb/pkg/jwt"
	"gweb/pkg/redis"
	"time"
)

var key = []byte("123456")

// GenToken 生成令牌
func GenToken(userId, roleId uint32) (token string, expiresAt time.Time, err error) {
	expiresAt = time.Now().Add(5 * time.Minute)
	token, err = jwt.Gen(userId, roleId, expiresAt, key)
	return
}

func verifyToken(token string) (*jwt.CustomClaims, error) {
	claims, err := jwt.Verify(token, key)
	if err != nil {
		return nil, err
	}
	claims.RegisteredClaims = nil
	return claims, err
}

// SetSToken 存储令牌
func SetSToken(userId uint32, token string, exp time.Duration) error {
	return redis.Set(fmt.Sprintf("user_%d", userId), token, exp)
}

// DelSToken 删除存储令牌
func DelSToken(userId uint32) error {
	return redis.Del(fmt.Sprintf("user_%d", userId))
}

// 获取存储令牌
func getSToken(userId uint32) (string, error) {
	return redis.Get(fmt.Sprintf("user_%d", userId))
}
