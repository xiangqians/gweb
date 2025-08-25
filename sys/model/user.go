// @author xiangqian
// @date 2025/07/20 17:05
package model

import "time"

// User 用户
type User struct {
	Id       uint32     `db:"id" json:"id"`             // 主键
	RoleId   uint32     `db:"role_id" json:"roleId"`    // 角色主键
	Name     string     `db:"name" json:"name"`         // 用户名
	Password string     `db:"password" json:"password"` // 密码
	Del      uint8      `db:"del" json:"del"`           // 是否已删除，0-未删除，1-已删除
	AddTime  time.Time  `db:"add_time" json:"addTime"`  // 添加时间
	UpdTime  *time.Time `db:"upd_time" json:"updTime"`  // 修改时间
}

type UserReq struct {
}

type UserRes struct {
}
