// @author xiangqian
// @date 2025/07/20 13:06
package gzip

import (
	"bytes"
	"compress/gzip"
	"io"
)

// Compress gzip压缩
func Compress(data []byte) ([]byte, error) {
	var buf bytes.Buffer
	w := gzip.NewWriter(&buf)
	if _, err := w.Write(data); err != nil {
		return nil, err
	}

	// Close() 是写入过程的必要组成部分（写入压缩尾数据）
	// 需要立即检查错误，因为关闭错误意味着压缩数据不完整
	if err := w.Close(); err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

// Decompress gzip解压
func Decompress(data []byte) ([]byte, error) {
	r, err := gzip.NewReader(bytes.NewReader(data))
	if err != nil {
		return nil, err
	}

	// Close() 只是资源清理操作
	defer r.Close()

	return io.ReadAll(r)
}
