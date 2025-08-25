// @author xiangqian
// @date 2025/07/20 18:56
package handler

import "net/http"

// AddUser godoc
// @Summary      新增用户
// @Description  新增用户信息
// @Tags         user
// @Accept       json
// @Param        request body model.User true "新增用户请求信息"
// @Produce      json
// @Success      200 {object} xhttp.Resp{data=bool}
// @Router       /api/v1/sys/user [post]
func AddUser(w http.ResponseWriter, r *http.Request) {
	// 加密密码
	//var password string
	//hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
}

// UpdUser 修改用户
func UpdUser(w http.ResponseWriter, r *http.Request) {
}

// GetUser 获取用户
func GetUser(w http.ResponseWriter, r *http.Request) {
}

// GetUsers 获取用户列表
func GetUsers(w http.ResponseWriter, r *http.Request) {
}

// DelUser 删除用户
func DelUser(w http.ResponseWriter, r *http.Request) {
}
