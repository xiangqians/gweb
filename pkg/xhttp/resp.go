// @author xiangqian
// @date 2025/08/02 22:45
package xhttp

import (
	"encoding/json"
	"net/http"
)

// Error 服务器在处理请求时遇到了未预期的错误，导致无法完成请求
func Error(w http.ResponseWriter, msg string) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(Resp{
		Code: "error",
		Msg:  msg,
	})
}

// Ok 服务器已正确处理并返回请求的资源
func Ok(w http.ResponseWriter, data any) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(Resp{
		Code: "ok",
		Msg:  http.StatusText(http.StatusOK),
		Data: data,
	})
}

// Resp 响应
type Resp struct {
	Code string `json:"code"` // 状态码
	Msg  string `json:"msg"`  // 消息
	Data any    `json:"data"` // 数据
}
