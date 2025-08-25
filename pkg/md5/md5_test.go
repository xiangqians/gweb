// @author xiangqian
// @date 2025/07/27 13:25
package md5

import (
	"fmt"
	"testing"
)

func TestHash(t *testing.T) {
	data := []byte("localhost:58080")
	fmt.Println(Hash(data))
}
