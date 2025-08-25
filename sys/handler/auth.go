// @author xiangqian
// @date 2025/07/20 17:05
package handler

import (
	"golang.org/x/crypto/bcrypt"
	"gweb/pkg/xhttp"
	"gweb/sys/model"
	"gweb/sys/repo"
	"net/http"
	"strings"
	"time"
)

// Token godoc
// @Summary      获取令牌
// @Description  获取令牌
// @Tags         auth
// @Accept       json
// @Param        request body model.TokenReq true "令牌请求信息"
// @Produce      json
// @Success      200 {object} xhttp.Resp{data=model.TokenResp}
// @Router       /api/v1/sys/user/token [post]
func Token(w http.ResponseWriter, r *http.Request) {
	var req model.TokenReq
	if err := xhttp.JsonDecode(r, &req); err != nil {
		xhttp.Error(w, err.Error())
		return
	}

	req.Name = strings.TrimSpace(req.Name)
	req.Password = strings.TrimSpace(req.Password)
	if req.Name == "" || len(req.Name) > 10 || req.Password == "" || len(req.Password) > 20 {
		xhttp.Error(w, "")
		return
	}

	user := repo.GetUserByName(req.Name)
	if user.Id == 0 {
		xhttp.Error(w, "用户名或密码错误")
		return
	}
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password))
	if user.Id == 0 {
		xhttp.Error(w, "用户名或密码错误")
		return
	}

	// 生成令牌
	token, expiresAt, err := xhttp.GenToken(user.Id, user.RoleId)
	if err != nil {
		xhttp.Error(w, err.Error())
		return
	}

	// 存储令牌
	err = xhttp.SetSToken(user.Id, token, expiresAt.Sub(time.Now()))
	if err != nil {
		xhttp.Error(w, err.Error())
		return
	}

	// OK
	xhttp.Ok(w, model.TokenResp{
		AccessToken: token,
		ExpiresAt:   expiresAt,
	})
}

// Revoke godoc
// @Summary      撤销令牌
// @Description  撤销令牌
// @Tags         auth
// @Produce      json
// @Success      200 {object} xhttp.Resp{data=bool}
// @Router       /api/v1/sys/user/revoke [post]
func Revoke(w http.ResponseWriter, r *http.Request) {
	// 获取当前请求上下文中自定义声明信息
	claims := xhttp.Claims(r)
	if claims == nil {
		xhttp.Ok(w, false)
		return
	}

	// 删除存储令牌
	err := xhttp.DelSToken(claims.UserId)
	if err != nil {
		xhttp.Error(w, err.Error())
		return
	}

	// OK
	xhttp.Ok(w, true)
}
