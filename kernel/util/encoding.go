package util

import "encoding/base64"

// Base64Encode Base64 编码，便于显示输出
func Base64Encode(data []byte) string {
	return base64.StdEncoding.EncodeToString(data)
}
