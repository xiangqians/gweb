// @author xiangqian
// @date 2025/08/02 12:55
package main

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestLoadConfig(t *testing.T) {
	config, err := LoadConfig()
	if err != nil {
		panic(err)
	}
	Print(config)
}

func Print(v any) {
	//fmt.Printf("%+v\n", v)
	data, err := json.MarshalIndent(v, "", "  ")
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s\n", data)
}
