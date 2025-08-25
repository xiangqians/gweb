// @author xiangqian
// @date 2025/07/26 13:07
package redis

import (
	"fmt"
	"testing"
	"time"
)

func TestInit(t *testing.T) {
	cli, pong, err := Init(Config{
		Host:     "localhost",
		Port:     6379,
		Password: "",
		Db:       0,
	})
	if err != nil {
		panic(err)
	}
	defer cli.Close()
	fmt.Println(pong)
}

func TestString(t *testing.T) {
	TestInit(t)

	var key = "string"

	Set(key, "123456", 30*time.Second)
	//Set(key, "123456", KeepTTL)

	value, err := Get(key)
	if err != nil {
		panic(err)
	}
	fmt.Printf("value: %v\n", value)
}

func TestSet(t *testing.T) {
	TestInit(t)

	var key = "test_set"
	//err := SAdd(key, "1", 2, 3, 4, 5, 6, "123")
	//if err != nil {
	//	panic(err)
	//}

	value, err := SGet(key)
	if err != nil {
		panic(err)
	}
	fmt.Printf("value: %v\n", value)

	values, err := SGets(key)
	if err != nil {
		panic(err)
	}
	fmt.Printf("values: %v\n", values)

	count, err := SCount(key)
	if err != nil {
		panic(err)
	}
	fmt.Printf("count: %v\n", count)

	exists, err := SHas(key, 1)
	if err != nil {
		panic(err)
	}
	fmt.Printf("exists: %v\n", exists)

	err = SDel(key, 1)
	if err != nil {
		panic(err)
	}
}

func TestHash(t *testing.T) {
	TestInit(t)
}
