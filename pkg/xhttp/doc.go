// @author xiangqian
// @date 2025/07/26 22:08
package xhttp

import (
	"encoding/json"
	"log"
	"os"
)

func Doc() {
	// 读取文件内容
	data, err := os.ReadFile("api.json")
	if err != nil {
		log.Printf("Error reading api.json: %v\n", err)
		return
	}

	// 解析 JSON 到 map
	var m map[string]interface{}
	err = json.Unmarshal(data, &m)
	if err != nil {
		log.Printf("Error unmarshalling api.json: %v\n", err)
		return
	}

	log.Println(m)
}
