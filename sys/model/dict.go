// @author xiangqian
// @date 2025/08/01 22:41
package model

import "time"

// Dict 字典
type Dict struct {
	Id      uint32     `json:"id"`      // 主键
	Type    string     `json:"type"`    // 类型
	Name    string     `json:"name"`    // 名称
	Value   string     `json:"value"`   // 值
	Sort    uint8      `json:"sort"`    // 排序
	AddTime time.Time  `json:"addTime"` // 添加时间
	UpdTime *time.Time `json:"updTime"` // 修改时间
}
