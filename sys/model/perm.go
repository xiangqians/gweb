// @author xiangqian
// @date 2025/07/26 23:13
package model

import "time"

// Perm 权限
type Perm struct {
	Id      uint32     `db:"id" json:"id"`            // 主键
	Tag     string     `db:"tag" json:"tag"`          // 标签
	Name    string     `db:"name" json:"name"`        // 名称
	Desc    string     `db:"desc" json:"desc"`        // 描述
	Method  string     `db:"method" json:"method"`    // 方法
	Path    string     `db:"path" json:"path"`        // 路径
	Anon    uint8      `db:"anon" json:"anon"`        // 是否允许匿名访问，0-不允许，1-允许
	Del     uint8      `db:"del" json:"del"`          // 是否已删除，0-未删除，1-已删除
	AddTime time.Time  `db:"add_time" json:"addTime"` // 添加时间
	UpdTime *time.Time `db:"upd_time" json:"updTime"` // 修改时间
}
