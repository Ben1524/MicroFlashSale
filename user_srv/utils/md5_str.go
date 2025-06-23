package utils

import (
	"crypto/md5"
	"encoding/hex"
)

func Md5Str(str string) string {
	// 1. 创建一个新的 MD5 哈希对象
	h := md5.New()

	// 2. 写入要哈希的数据
	h.Write([]byte(str))

	// 3. 计算哈希值并返回十六进制字符串
	return hex.EncodeToString(h.Sum(nil))
}
