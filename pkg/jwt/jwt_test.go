// @author xiangqian
// @date 2025/07/26 10:16
package jwt

import (
	"encoding/json"
	"fmt"
	"testing"
	"time"
)

var key = []byte("123456")

func TestGen(t *testing.T) {
	var userId uint32 = 1
	var roleId uint32 = 2
	var expiresAt = time.Now().Add(1 * time.Minute)
	token, err := Gen(userId, roleId, expiresAt, key)
	if err != nil {
		panic(err)
	}
	fmt.Println(len(token), token)
}

func TestVerify(t *testing.T) {
	var token = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVc2VySWQiOjEsIlJvbGVJZCI6MiwiaXNzIjoiZ3dlYiIsInN1YiI6IjEiLCJhdWQiOlsid2ViIiwibW9iaWxlIl0sImV4cCI6MTc1NDU0MTU5MywibmJmIjoxNzU0NTQxNTMzLCJpYXQiOjE3NTQ1NDE1MzMsImp0aSI6IkdIWi11SC1RVmh0ZF9lRldKQUFZWWcifQ.4XqrGmuI7sZ5pu581R9uNuz5vTfYRMfxjBtijuilnBo"
	claims, err := Verify(token, key)
	if err != nil {
		panic(err)
	}

	//fmt.Printf("%+v\n", claims)
	data, err := json.MarshalIndent(claims, "", "  ")
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s\n", data)
}
