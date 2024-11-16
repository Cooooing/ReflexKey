package util

import (
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"kernel/model/param"
	"testing"
)

// TestAesEncrypt 测试所有类型的组合
func TestAesEncrypt(t *testing.T) {
	plainText := "Hello, world!"
	originalIV := "thisisaniv123456"

	for _, operationMode := range aesType["OperationMode"] {
		for _, fill := range aesType["Fill"] {
			//for _, keyLength := range aesType["KeyLength"] {
			for _, keyFormat := range aesType["KeyFormat"] {
				for _, deviationFormat := range aesType["DeviationFormat"] {
					for _, plainTextFormat := range aesType["PlainTextFormat"] {
						// 生成密钥
						key := generateKey("16", keyFormat)
						// 生成IV
						iv := generateIV(originalIV, deviationFormat)
						// 生成明文
						text := generatePlainText(plainText, plainTextFormat)

						params := param.AesEncryptParam{
							PlainText:       text,
							PlainTextFormat: plainTextFormat,
							OperationMode:   operationMode,
							Fill:            fill,
							Key:             key,
							KeyFormat:       keyFormat,
							Deviation:       iv,
							DeviationFormat: deviationFormat,
						}

						encrypted, err := AesEncrypt(params)
						if err != nil {
							t.Errorf("Error: %v", err)
						}
						t.Logf("OperationMode: %s, Fill: %s, KeyLength: %s, KeyFormat: %s, DeviationFormat: %s, PlainTextFormat: %s, Encrypted: %s",
							operationMode, fill, "16", keyFormat, deviationFormat, plainTextFormat, encrypted)
						t.Logf("Encrypted text: %s", encrypted)
						fmt.Println()
						//require.NoError(t, err, "Encryption failed for params: %v", params)
						//assert.NotEmpty(t, encrypted, "Encrypted text should not be empty")
						//}
					}
				}
			}
		}
	}
}

func generateKey(keyLength, keyFormat string) string {
	var rawKey string
	switch keyLength {
	case "16":
		rawKey = "thisisaniv123456" // 示例16字节密钥
	case "24":
		rawKey = "thisisaniv12345612345678" // 示例24字节密钥
	case "32":
		rawKey = "thisisaniv1234561234567812345678" // 示例32字节密钥
	}

	switch keyFormat {
	case "Hex":
		return hex.EncodeToString([]byte(rawKey))
	case "Base64":
		return base64.StdEncoding.EncodeToString([]byte(rawKey))
	default: // "Text" 或其他情况直接返回原密钥
		return rawKey
	}
}
func generateIV(originalIV, deviationFormat string) string {
	switch deviationFormat {
	case "Hex":
		return hex.EncodeToString([]byte(originalIV))
	case "Base64":
		return base64.StdEncoding.EncodeToString([]byte(originalIV))
	default: // "Text" 或其他情况直接返回原IV
		return originalIV
	}
}
func generatePlainText(plainText, plainTextFormat string) string {
	switch plainTextFormat {
	case "Hex":
		// 示例固定16字节明文的Hex表示
		return "4d3c6e5b1d4a9e8fc93476b69a1f7d26"
	case "Base64":
		// 示例固定明文的Base64表示
		return "nH0a3zsy+Fx39Tm0USqK47n5Low1lNBb2TjMWpLT8aI="
	default: // "Text" 或其他情况直接返回原始明文
		return plainText
	}
}
