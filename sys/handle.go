// @author xiangqian
// @date 2025/07/20 12:44
package sys

import (
	"gweb/pkg/xhttp"
	"gweb/sys/handler"
	"time"
)

func Handle() {
	var options = xhttp.Options{
		N:       1 << 9, // 512 B
		Timeout: 2 * time.Second,
		Anon:    false,
	}

	// [auth]
	// 获取令牌
	xhttp.JsonHandle("POST /api/v1/sys/user/token", handler.Token, xhttp.Options{N: 1 << 6, Timeout: 2 * time.Second, Anon: true})
	// 撤销令牌
	xhttp.JsonHandle("POST /api/v1/sys/user/revoke", handler.Revoke, options)

	// [user]
	// 新增用户
	xhttp.JsonHandle("POST /api/v1/sys/user", handler.AddUser, options)
	// 修改用户
	xhttp.JsonHandle("PUT /api/v1/sys/user", handler.UpdUser, options)
	// 获取用户
	xhttp.JsonHandle("GET /api/v1/sys/user/{id}", handler.GetUser, options)
	// 获取用户列表
	xhttp.JsonHandle("GET /api/v1/sys/users", handler.GetUsers, options)
	// 删除用户
	xhttp.JsonHandle("DELETE /api/v1/sys/user", handler.DelUser, options)
}
