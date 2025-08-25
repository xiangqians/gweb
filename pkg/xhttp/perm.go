// @author xiangqian
// @date 2025/08/02 22:53
package xhttp

import (
	"database/sql"
	"errors"
	"fmt"
	"gweb/pkg/db"
	"gweb/pkg/redis"
	"log"
)

func GetPermId(method, path string, anon bool) uint32 {
	// 获取权限信息
	var perm struct {
		Id   uint32 `db:"id"`
		Anon uint8  `db:"anon"`
		Del  uint8  `db:"del"`
	}
	err := db.Get(&perm, "SELECT `id`, `anon`, `del` FROM `perm` WHERE `method` = ? AND `path` = ? LIMIT 1", method, path)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return 0
		} else {
			log.Printf("GetPermId(%s, %s, %t) Get error: %v\n", method, path, anon, err)
			return 0
		}
	}
	if perm.Id == 0 {
		return 0
	}

	// 更新权限信息
	var ianon uint8
	if anon {
		ianon = 1
	} else {
		ianon = 0
	}
	if perm.Anon != ianon || perm.Del == 1 {
		_, err = db.Upd("UPDATE `perm` SET `anon` = ?, `del` = 0 WHERE `id` = ?", ianon, perm.Id)
		if err != nil {
			log.Printf("GetPermId(%s, %s, %t) Upd error: %v\n", method, path, anon, err)
		}
	}

	return perm.Id
}

func AddPerm(method, path string, anon bool) (int64, int64) {
	rowsAffected, insertId, err := db.Add("INSERT INTO `perm` (`method`, `path`, `anon`) VALUES (?, ?, ?)", method, path, anon)
	if err != nil {
		log.Printf("AddPerm(%s, %s, %t) error: %v\n", method, path, anon, err)
	}
	return rowsAffected, insertId
}

func hasPermId(roleId, permId uint32) (bool, error) {
	return redis.SHas(fmt.Sprintf("role_%d", roleId), permId)
}
