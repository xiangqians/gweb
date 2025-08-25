// @author xiangqian
// @date 2025/08/03 14:40
package repo

import "gweb/pkg/db"

func DelPerm(id uint32) {
	db.Upd("UPDATE `perm` SET `del` = 1 WHERE `id` = ?", id)
}
