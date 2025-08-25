// @author xiangqian
// @date 2025/07/20 13:05
package gob

import (
	"bytes"
	"encoding/gob"
)

// Serialize GOB 序列化
func Serialize(v any) ([]byte, error) {
	var buf = &bytes.Buffer{}
	encoder := gob.NewEncoder(buf)
	err := encoder.Encode(v)
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

// Deserialize GOB 反序列化
func Deserialize(data []byte, v any) error {
	var buf = bytes.NewBuffer(data)
	decoder := gob.NewDecoder(buf)
	return decoder.Decode(v)
}
