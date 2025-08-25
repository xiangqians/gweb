// @author xiangqian
// @date 2025/08/03 21:35
package xhttp

import (
	"encoding/json"
	"net/http"
)

// JsonDecode 解析 JSON 请求报文
func JsonDecode(r *http.Request, v any) error {
	decoder := json.NewDecoder(r.Body)

	// 禁止未知字段
	decoder.DisallowUnknownFields()

	return decoder.Decode(v)
}
