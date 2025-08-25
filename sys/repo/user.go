// @author xiangqian
// @date 2025/07/20 18:51
package repo

import (
	"gweb/pkg/db"
	"gweb/sys/model"
	"log"
)

func AddUser() {
	//if err != nil {
	//	log.Printf("AddPerm(%+v) error: %v\n", perm, err)
	//}
}

func UpdUser() {
}

func GetUser() {
}

func GetUserByName(name string) model.User {
	var user model.User
	err := db.Get(&user, "SELECT `id`, `role_id`, `name`, `password` FROM `user` WHERE `name` = ? AND `del` = 0 LIMIT 1", name)
	if err != nil {
		log.Printf("GetUserByName(%s) error: %v\n", name, err)
	}
	return user
}

func DelUser() {
}

//public <T, R> LazyList<R> lazyList(Class<R> type, LazyList lazyList, SqlBuilder<T> sql) {
//    int rows = lazyList.getRows();
//    int offset = lazyList.getOffset() * rows;
//    // 预加载 1 行
//    rows += 1;
//    sql.limit(offset, rows);
//    List<R> data = list(type, sql);
//
//    boolean hasNext = false;
//    if (CollectionUtils.size(data) > rows) {
//        hasNext = true;
//        data = data.subList(0, rows);
//    }
//    lazyList.setData(data);
//    lazyList.setHasNext(hasNext);
//    return lazyList;
//}
