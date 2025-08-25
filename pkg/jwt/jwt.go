// @author xiangqian
// @date 2025/07/26 09:46
package jwt

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"time"
)

// Gen 生成令牌
// userId    用户主键
// roleId    角色主键
// expiresAt 过期时间
// key       密钥
func Gen(userId, roleId uint32, expiresAt time.Time, key []byte) (string, error) {
	// 生成 16 字节（128 位）的随机数
	buf := make([]byte, 16)
	if _, err := rand.Read(buf); err != nil {
		return "", fmt.Errorf("failed to generate random JTI: %w", err)
	}
	id := base64.RawURLEncoding.EncodeToString(buf)

	// 主题
	var subject = fmt.Sprintf("%d", userId)
	// 接收方
	var audience = []string{"web", "mobile"}

	// 自定义声明信息
	claims := &CustomClaims{
		UserId: userId,
		RoleId: roleId,
		RegisteredClaims: &jwt.RegisteredClaims{
			ID:        id,                             // ID
			Issuer:    "gweb",                         // 签发者
			Subject:   subject,                        // 主题
			Audience:  audience,                       // 接收方
			IssuedAt:  jwt.NewNumericDate(time.Now()), // 签发时间
			NotBefore: jwt.NewNumericDate(time.Now()), // 生效时间
			ExpiresAt: jwt.NewNumericDate(expiresAt),  // 过期时间
		},
	}

	// 创建令牌对象，指定签名算法
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	// 使用密钥签名令牌
	return token.SignedString(key)
}

// Verify 验证令牌
// token 令牌
// key   密钥
func Verify(token string, key []byte) (*CustomClaims, error) {
	// 解析令牌
	var claims = &CustomClaims{}
	tok, err := jwt.ParseWithClaims(token, claims, func(tok *jwt.Token) (any, error) {
		// 验证签名算法是否正确
		if _, ok := tok.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", tok.Header["alg"])
		}
		return key, nil
	})
	if err != nil {
		return nil, err
	}

	// 验证令牌是否有效
	if !tok.Valid {
		return nil, fmt.Errorf("invalid token")
	}

	return claims, nil
}

// CustomClaims 自定义声明信息
type CustomClaims struct {
	UserId uint32 // 用户主键
	RoleId uint32 // 角色主键
	// 内嵌标准 Claims（包含标准字段如过期时间等）
	*jwt.RegisteredClaims
}
