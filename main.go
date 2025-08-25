// @author xiangqian
// @date 2025/07/20 12:38
package main

import (
	"fmt"
	"gweb/pkg/db"
	"gweb/pkg/xhttp"
	"gweb/pkg/xlog"
	"gweb/sys"
	"log"
	"net/http"
)

// @title        GWeb 系统
// @version      1.0.0
// @description  GWeb 系统
// @host         localhost:58080
func main() {
	err := exec()
	if err != nil {
		log.Fatal(err)
	}
}

func exec() error {
	// [config]
	config, err := LoadConfig()
	if err != nil {
		return err
	}

	// [xlog]
	err = xlog.Init()
	if err != nil {
		return err
	}

	// [db]
	db, err := db.Init(config.Db)
	if err != nil {
		return err
	}
	defer db.Close()

	// [prom]
	//prom.Handle()
	// [sys]
	sys.Handle()

	// [doc]
	xhttp.Doc()

	// [http]
	port := config.Http.Port
	log.Printf("Server starting on port %d ...\n", port)
	return http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
}
