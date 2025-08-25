// @author xiangqian
// @date 2025/07/20 22:05
package db

import (
	"testing"
)

func TestInit(t *testing.T) {
	db, err := Init(Config{
		Driver:     "mysql",
		DataSource: "root:root@tcp(localhost:3306)/web?parseTime=true",
	})
	if err != nil {
		panic(err)
	}
	defer db.Close()
}
