// @author xiangqian
// @date 2025/07/20 16:08
package xhttp

import (
	"context"
	"gweb/pkg/jwt"
	"log"
	"net/http"
	"strings"
	"time"
)

func JsonHandle(pattern string, handler func(http.ResponseWriter, *http.Request), options Options) {
	options.contentType = "application/json"
	handle(pattern, handler, options)
}

func FormDataHandle(pattern string, handler func(http.ResponseWriter, *http.Request), options Options) {
	options.contentType = "multipart/form-data"
	handle(pattern, handler, options)
}

// 路由注册
// pattern  路由模式
// handler  处理器
// options  配置选项
func handle(pattern string, handler func(http.ResponseWriter, *http.Request), options Options) {
	// 存储权限信息（如果不存在的话）
	var permId uint32
	if method, path, ok := strings.Cut(pattern, " "); ok {
		permId = GetPermId(method, path, options.Anon)
		if permId == 0 {
			_, insertId := AddPerm(method, path, options.Anon)
			permId = uint32(insertId)
		}
	}

	http.Handle(pattern, &xhandler{
		permId:  permId,
		options: options,
		handle:  handler,
	})
}

// X处理器
type xhandler struct {
	permId  uint32  // 权限主键
	options Options // 配置选项
	handle  func(http.ResponseWriter, *http.Request)
}

// Options 配置选项
type Options struct {
	contentType string        // 内容类型
	N           int64         // 允许的最大字节数，例如 1 << 10 = 1 KB，1 << 20 = 1 MB
	Timeout     time.Duration // 超时时间
	Anon        bool          // 是否允许匿名访问
}

func (xh *xhandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// 禁止 MIME 嗅探（即使返回 JSON 也可能被浏览器误解析）
	w.Header().Set("X-Content-Type-Options", "nosniff")
	// 禁止嵌入 iframe（保护 API 调试页面）
	w.Header().Set("X-Frame-Options", "DENY")
	// 禁用浏览器猜测的内容类型（防御 Content-Type 绕过攻击）
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	// 禁止所有外部资源加载（API 无需加载 JS/CSS/image 等）
	w.Header().Set("Content-Security-Policy", "default-src 'none'; frame-ancestors 'none'")

	// 限制请求报文大小
	// 防止 DoS 攻击时，恶意客户端发送超大请求报文，如 10GB 的 JSON，占满服务器内存从而导致 OOM 崩溃
	r.Body = http.MaxBytesReader(w, r.Body, xh.options.N)
	// 关闭请求报文
	defer r.Body.Close()

	// 限制解析资源，超出写磁盘
	if err := r.ParseMultipartForm(5 << 20); err != nil {
		if err.Error() == "http: request body too large" {
			Error(w, http.StatusText(http.StatusRequestEntityTooLarge))
		} else {
			Error(w, http.StatusText(http.StatusBadRequest))
		}
		return
	}
	// 清理临时文件
	defer r.MultipartForm.RemoveAll()

	// 缺一不可：
	// 只设 MaxBytesReader 内存可能被海量小文件耗尽
	// 只设 ParseMultipartForm 无法阻止超大单个文件攻击

	// 超时控制
	ctx, cancel := context.WithTimeout(r.Context(), xh.options.Timeout)
	defer cancel()

	// 判断是否支持内容类型
	if r.Header.Get("Content-Type") != xh.options.contentType {
		Error(w, http.StatusText(http.StatusUnsupportedMediaType))
		return
	}

	// 判断内容长度是否超出限制
	if r.ContentLength > xh.options.N {
		Error(w, http.StatusText(http.StatusBadRequest))
		return
	}

	// 是否允许匿名访问
	if xh.options.Anon {
		xh.handle(w, r)
		return
	}

	// 获取令牌
	const prefix = "Bearer "
	authorization := strings.TrimSpace(r.Header.Get("Authorization"))
	if authorization == "" || !strings.HasPrefix(authorization, prefix) {
		Error(w, http.StatusText(http.StatusUnauthorized))
		return
	}
	token := strings.TrimSpace(strings.TrimPrefix(authorization, prefix))
	if token == "" {
		Error(w, http.StatusText(http.StatusUnauthorized))
		return
	}

	// JWT 验证令牌
	claims, err := verifyToken(token)
	if err != nil {
		log.Printf("jwt token verify error: %v\n", err)
		Error(w, http.StatusText(http.StatusUnauthorized))
		return
	}

	// 验证存储令牌
	sToken, err := getSToken(claims.UserId)
	if err != nil {
		log.Printf("getSToken(%d) error: %v\n", claims.UserId, err)
		Error(w, http.StatusText(http.StatusUnauthorized))
		return
	}
	if sToken != token {
		log.Printf("sToken(%s) != token(%s)\n", sToken, token)
		Error(w, http.StatusText(http.StatusUnauthorized))
		return
	}

	// 验证权限
	exists, err := hasPermId(claims.RoleId, xh.permId)
	if err != nil {
		log.Printf("hasPermId(%d, %d) error: %v\n", claims.RoleId, xh.permId, err)
		Error(w, http.StatusText(http.StatusUnauthorized))
		return
	}
	if !exists {
		// 服务器理解客户端的请求，但拒绝授权访问所请求的资源
		Error(w, http.StatusText(http.StatusForbidden))
		return
	}

	// 将自定义声明信息存入当前请求的上下文中
	ctx = context.WithValue(ctx, "claims", claims)
	// 继续处理请求
	xh.handle(w, r.WithContext(ctx))
}

// Claims 获取当前请求上下文中自定义声明信息
func Claims(r *http.Request) *jwt.CustomClaims {
	if claims, ok := r.Context().Value("claims").(*jwt.CustomClaims); ok {
		return claims
	}
	return nil
}
