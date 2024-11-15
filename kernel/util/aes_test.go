package util

import (
	"encoding/base64"
	"encoding/hex"
	"kernel/model/param"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// TestAesEncrypt 测试所有类型的组合
func TestAesEncrypt(t *testing.T) {
	plainText := "Hello, world!"
	key := "thisisaverysecretkey" // 16字节密钥
	iv := "thisisaniv123456"      // 16字节IV

	// 生成所有可能的组合
	for _, operationMode := range aesType["OperationMode"] {
		for _, fill := range aesType["Fill"] {
			for _, keyLength := range aesType["KeyLength"] {
				for _, keyFormat := range aesType["KeyFormat"] {
					for _, deviationFormat := range aesType["DeviationFormat"] {
						for _, plainTextFormat := range aesType["PlainTextFormat"] {
							// 根据格式转换明文、密钥和偏移量
							var plainTextBytes, keyBytes, ivBytes []byte
							var err error

							switch plainTextFormat {
							case "Hex":
								plainTextBytes, err = hex.DecodeString(hex.EncodeToString([]byte(plainText)))
							case "Base64":
								plainTextBytes, err = base64.StdEncoding.DecodeString(base64.StdEncoding.EncodeToString([]byte(plainText)))
							default:
								plainTextBytes = []byte(plainText)
							}
							require.NoError(t, err)

							switch keyFormat {
							case "Hex":
								keyBytes, err = hex.DecodeString(hex.EncodeToString([]byte(key)))
							case "Base64":
								keyBytes, err = base64.StdEncoding.DecodeString(base64.StdEncoding.EncodeToString([]byte(key)))
							default:
								keyBytes = []byte(key)
							}
							require.NoError(t, err)

							switch deviationFormat {
							case "Hex":
								ivBytes, err = hex.DecodeString(hex.EncodeToString([]byte(iv)))
							case "Base64":
								ivBytes, err = base64.StdEncoding.DecodeString(base64.StdEncoding.EncodeToString([]byte(iv)))
							default:
								ivBytes = []byte(iv)
							}
							require.NoError(t, err)

							// 构建参数
							params := param.AesEncryptParam{
								PlainText:       string(plainTextBytes),
								PlainTextFormat: plainTextFormat,
								OperationMode:   operationMode,
								Fill:            fill,
								Key:             string(keyBytes),
								KeyFormat:       keyFormat,
								Deviation:       string(ivBytes),
								DeviationFormat: deviationFormat,
							}

							// 调用加密函数
							encrypted, err := AesEncrypt(params)
							require.NoError(t, err)
							assert.NotEmpty(t, encrypted, "Encrypted text should not be empty")

							// 打印测试结果（可选）
							t.Logf("OperationMode: %s, Fill: %s, KeyLength: %s, KeyFormat: %s, DeviationFormat: %s, PlainTextFormat: %s, Encrypted: %s",
								operationMode, fill, keyLength, keyFormat, deviationFormat, plainTextFormat, encrypted)
						}
					}
				}
			}
		}
	}
}
