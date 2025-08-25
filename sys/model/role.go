// @author xiangqian
// @date 2025/08/01 22:47
package model

import "time"

// Role 角色
type Role struct {
	Id      uint32     `json:"id"`      // 主键
	Name    string     `json:"name"`    // 名称
	Code    string     `json:"code"`    // 标识码
	Del     uint8      `json:"del"`     // 是否已删除，0-未删除，1-已删除
	AddTime time.Time  `json:"addTime"` // 添加时间
	UpdTime *time.Time `json:"updTime"` // 修改时间
}

// RolePerm 角色-权限
type RolePerm struct {
	RoleId  uint32    `json:"roleId"`  // 角色主键
	PermId  uint32    `json:"permId"`  // 权限主键
	AddTime time.Time `json:"addTime"` // 添加时间
}
