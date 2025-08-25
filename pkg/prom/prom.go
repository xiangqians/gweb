// @author xiangqian
// @date 2025/07/26 11:49
package prom

import (
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"net/http"
	"strings"
)

func Handle(config Config) {
	// 暴露指标端点
	var handler = promhttp.Handler()
	http.HandleFunc("/metrics", func(w http.ResponseWriter, r *http.Request) {
		const prefix = "Bearer "
		authorization := strings.TrimSpace(r.Header.Get("Authorization"))
		if authorization == "" || !strings.HasPrefix(authorization, prefix) {
			http.NotFound(w, r)
			return
		}

		token := strings.TrimSpace(strings.TrimPrefix(authorization, prefix))
		if token == "" {
			http.NotFound(w, r)
			return
		}

		if token != config.Token {
			http.NotFound(w, r)
			return
		}

		handler.ServeHTTP(w, r)
	})
}

type Config struct {
	Token string
}
