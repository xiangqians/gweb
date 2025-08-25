// @author xiangqian
// @date 2025/07/27 13:16
package md5

import (
	"crypto/md5"
	"encoding/hex"
)

// 哈希算法之 MD5
// 生成一个固定长度（16 个字节）的哈希值，通常用于检查数据完整性或生成文件指纹，但由于其碰撞问题，已经不再推荐用于安全场景。

func Hash(data []byte) string {
	sum := md5.Sum(data)
	return hex.EncodeToString(sum[:])
}
