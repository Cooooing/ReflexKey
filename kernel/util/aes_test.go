package util

import (
	"bytes"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"kernel/model/param"
	"strings"
	"testing"
)

// 通用测试数据结构
type testData struct {
	plainTexts map[string]string
	keys       map[string]map[string]string
	deviations map[string]string
}

// 获取通用测试数据
func getTestData() testData {
	// 生成一个不完整的块（非16的倍数）
	plainText := strings.Repeat("A", 40) // 40字节，2个完整块加8字节
	return testData{
		plainTexts: map[string]string{
			"Text":   plainText,
			"Base64": base64.StdEncoding.EncodeToString([]byte(plainText)),
			"Hex":    hex.EncodeToString([]byte(plainText)),
		},
		keys: map[string]map[string]string{
			"16": {
				"Text":   "1234567890123456",
				"Base64": "MTIzNDU2Nzg5MDEyMzQ1Ng==",
				"Hex":    "31323334353637383930313233343536",
			},
			"24": {
				"Text":   "123456789012345678901234",
				"Base64": "MTIzNDU2Nzg5MDEyMzQ1Njc4OTAxMjM0",
				"Hex":    "313233343536373839303132333435363738393031323334",
			},
			"32": {
				"Text":   "12345678901234567890123456789012",
				"Base64": "MTIzNDU2Nzg5MDEyMzQ1Njc4OTAxMjM0NTY3ODkwMTI=",
				"Hex":    "3132333435363738393031323334353637383930313233343536373839303132",
			},
		},
		deviations: map[string]string{
			"Text":   "1234567890123456",
			"Base64": "MTIzNDU2Nzg5MDEyMzQ1Ng==",
			"Hex":    "31323334353637383930313233343536",
		},
	}
}

// 执行加解密测试的通用方法
func runTest(t *testing.T, encryptParam param.AesEncryptParam, decryptParam param.AesDecryptParam) {
	encrypted, err := AesEncrypt(encryptParam)
	if err != nil {
		t.Logf("加密参数: %+v", encryptParam)
		t.Errorf("加密失败: %v", err)
		return
	}
	t.Logf("加密文本: %s", encryptParam.PlainText)
	t.Logf("加密结果: %s", encrypted)

	decryptParam.CipherText = encrypted
	decrypted, err := AesDecrypt(decryptParam)
	if err != nil {
		t.Logf("解密参数: %+v", decryptParam)
		t.Errorf("解密失败: %v", err)
		return
	}
	t.Logf("解密结果: %s", decrypted)

	if decrypted != encryptParam.PlainText {
		t.Errorf("解密结果与原文不匹配\n期望: %s\n实际: %s",
			encryptParam.PlainText, decrypted)
	}
}

// TestStreamModes 测试流加密模式（OFB/CFB/CTR）
func TestStreamModes(t *testing.T) {
	data := getTestData()

	// 流加密模式共同特点：
	// 1. 只使用No填充
	// 2. 需要16字节IV
	// 3. 可以处理任意长度数据
	for _, mode := range []string{"OFB", "CFB", "CTR"} {
		for keyLength, keyMap := range data.keys {
			for _, keyFormat := range aesType["KeyFormat"] {
				for _, plainTextFormat := range aesType["PlainTextFormat"] {
					for _, cipherTextFormat := range aesType["CipherTextFormat"] {
						for _, deviationFormat := range aesType["DeviationFormat"] {
							testName := fmt.Sprintf(
								"%s-%s-%s-%sBytes-%s-%s",
								mode,
								plainTextFormat,
								cipherTextFormat,
								keyLength,
								keyFormat,
								deviationFormat,
							)

							t.Run(testName, func(t *testing.T) {
								encryptParam := param.AesEncryptParam{
									PlainText:        data.plainTexts[plainTextFormat],
									PlainTextFormat:  plainTextFormat,
									OperationMode:    mode,
									Fill:             "No", // 流模式只用No填充
									Key:              keyMap[keyFormat],
									KeyFormat:        keyFormat,
									Deviation:        data.deviations[deviationFormat],
									DeviationFormat:  deviationFormat,
									CipherTextFormat: cipherTextFormat,
								}

								decryptParam := param.AesDecryptParam{
									CipherTextFormat: cipherTextFormat,
									OperationMode:    mode,
									Fill:             "No",
									Key:              keyMap[keyFormat],
									KeyFormat:        keyFormat,
									Deviation:        data.deviations[deviationFormat],
									DeviationFormat:  deviationFormat,
									PlainTextFormat:  plainTextFormat,
								}

								runTest(t, encryptParam, decryptParam)
							})
						}
					}
				}
			}
		}
	}
}

// TestBlockModes 测试分组加密模式（CBC/ECB）
func TestBlockModes(t *testing.T) {
	data := getTestData()

	// 分组加密模式共同特点：
	// 1. 支持所有填充模式
	// 2. CBC需要IV，ECB不需要
	// 3. 需要处理块大小对齐
	for _, mode := range []string{"CBC", "ECB"} {
		for _, fill := range aesType["Fill"] {
			for keyLength, keyMap := range data.keys {
				for _, keyFormat := range aesType["KeyFormat"] {
					for _, plainTextFormat := range aesType["PlainTextFormat"] {
						for _, cipherTextFormat := range aesType["CipherTextFormat"] {
							// ECB模式不需要遍历偏移量格式
							deviationFormats := aesType["DeviationFormat"]
							if mode == "ECB" {
								deviationFormats = []string{""}
							}

							for _, deviationFormat := range deviationFormats {
								// 如果是No填充，确保明文长度是16的倍数
								if fill == "No" {
									var textBytes []byte
									var err error
									switch plainTextFormat {
									case "Text":
										textBytes = []byte(data.plainTexts[plainTextFormat])
									case "Base64":
										textBytes, err = base64.StdEncoding.DecodeString(data.plainTexts[plainTextFormat])
									case "Hex":
										textBytes, err = hex.DecodeString(data.plainTexts[plainTextFormat])
									}
									if err != nil || len(textBytes)%16 != 0 {
										continue
									}
								}

								testName := fmt.Sprintf(
									"%s-%s-%s-%s-%sBytes-%s",
									mode,
									fill,
									plainTextFormat,
									cipherTextFormat,
									keyLength,
									keyFormat,
								)
								if mode != "ECB" {
									testName += "-" + deviationFormat
								}

								t.Run(testName, func(t *testing.T) {
									encryptParam := param.AesEncryptParam{
										PlainText:        data.plainTexts[plainTextFormat],
										PlainTextFormat:  plainTextFormat,
										OperationMode:    mode,
										Fill:             fill,
										Key:              keyMap[keyFormat],
										KeyFormat:        keyFormat,
										CipherTextFormat: cipherTextFormat,
									}

									decryptParam := param.AesDecryptParam{
										CipherTextFormat: cipherTextFormat,
										OperationMode:    mode,
										Fill:             fill,
										Key:              keyMap[keyFormat],
										KeyFormat:        keyFormat,
										PlainTextFormat:  plainTextFormat,
									}

									// CBC模式需要设置偏移量
									if mode == "CBC" {
										encryptParam.Deviation = data.deviations[deviationFormat]
										encryptParam.DeviationFormat = deviationFormat
										decryptParam.Deviation = data.deviations[deviationFormat]
										decryptParam.DeviationFormat = deviationFormat
									}

									runTest(t, encryptParam, decryptParam)
								})
							}
						}
					}
				}
			}
		}
	}
}

// TestAuthenticatedMode 测试认证加密模式（GCM）
func TestAuthenticatedMode(t *testing.T) {
	data := getTestData()
	gcmDeviations := map[string]string{
		"Text":   "123456789012",
		"Base64": "MTIzNDU2Nzg5MDEy",
		"Hex":    "313233343536373839303132",
	}

	// GCM模式特点：
	// 1. 不需要填充
	// 2. 需要12字节nonce
	// 3. 提供认证功能
	for keyLength, keyMap := range data.keys {
		for _, keyFormat := range aesType["KeyFormat"] {
			for _, plainTextFormat := range aesType["PlainTextFormat"] {
				for _, cipherTextFormat := range aesType["CipherTextFormat"] {
					for _, deviationFormat := range aesType["DeviationFormat"] {
						testName := fmt.Sprintf(
							"GCM-%s-%s-%sBytes-%s-%s",
							plainTextFormat,
							cipherTextFormat,
							keyLength,
							keyFormat,
							deviationFormat,
						)

						t.Run(testName, func(t *testing.T) {
							// 准备测试数据
							var plainText string
							switch plainTextFormat {
							case "Text":
								plainText = data.plainTexts[plainTextFormat]
							case "Base64":
								plainBytes := []byte(data.plainTexts["Text"])
								plainText = base64.StdEncoding.EncodeToString(plainBytes)
							case "Hex":
								plainBytes := []byte(data.plainTexts["Text"])
								plainText = hex.EncodeToString(plainBytes)
							}

							encryptParam := param.AesEncryptParam{
								PlainText:        plainText,
								PlainTextFormat:  plainTextFormat,
								OperationMode:    "GCM",
								Fill:             "No", // GCM模式使用No填充
								Key:              keyMap[keyFormat],
								KeyFormat:        keyFormat,
								Deviation:        gcmDeviations[deviationFormat],
								DeviationFormat:  deviationFormat,
								CipherTextFormat: cipherTextFormat,
							}

							decryptParam := param.AesDecryptParam{
								CipherTextFormat: cipherTextFormat,
								OperationMode:    "GCM",
								Fill:             "No", // GCM模式使用No填充
								Key:              keyMap[keyFormat],
								KeyFormat:        keyFormat,
								Deviation:        gcmDeviations[deviationFormat],
								DeviationFormat:  deviationFormat,
								PlainTextFormat:  plainTextFormat,
							}

							// 记录原始输入
							t.Logf("原始明文: %s", plainText)
							t.Logf("明文格式: %s", plainTextFormat)

							encrypted, err := AesEncrypt(encryptParam)
							if err != nil {
								t.Logf("加密参数: %+v", encryptParam)
								t.Errorf("加密失败: %v", err)
								return
							}
							t.Logf("加密结果: %s", encrypted)

							decryptParam.CipherText = encrypted
							decrypted, err := AesDecrypt(decryptParam)
							if err != nil {
								t.Logf("解密参数: %+v", decryptParam)
								t.Errorf("解密失败: %v", err)
								return
							}
							t.Logf("解密结果: %s", decrypted)

							if decrypted != plainText {
								t.Errorf("解密结果与原文不匹配\n期望: %s\n实际: %s",
									plainText, decrypted)
							}
						})
					}
				}
			}
		}
	}
}

// TestCTSMode 测试CTS模式
func TestCTSMode(t *testing.T) {
	data := getTestData()

	// CTS模式特点：
	// 1. 需要至少两个块的数据
	// 2. 最后一个块可以不完整
	// 3. 需要16字节IV
	// 4. 不需要填充
	for keyLength, keyMap := range data.keys {
		for _, keyFormat := range aesType["KeyFormat"] {
			for _, plainTextFormat := range aesType["PlainTextFormat"] {
				for _, cipherTextFormat := range aesType["CipherTextFormat"] {
					for _, deviationFormat := range aesType["DeviationFormat"] {
						testName := fmt.Sprintf(
							"CTS-%s-%s-%sBytes-%s-%s",
							plainTextFormat,
							cipherTextFormat,
							keyLength,
							keyFormat,
							deviationFormat,
						)

						t.Run(testName, func(t *testing.T) {
							// 准备测试数据 - 确保数据长度大于32字节（两个块）
							rawPlainText := strings.Repeat("A", 48) // 48字节，3个完整块
							var testPlainText string
							switch plainTextFormat {
							case "Text":
								testPlainText = rawPlainText
							case "Base64":
								testPlainText = base64.StdEncoding.EncodeToString([]byte(rawPlainText))
							case "Hex":
								testPlainText = hex.EncodeToString([]byte(rawPlainText))
							}

							encryptParam := param.AesEncryptParam{
								PlainText:        testPlainText,
								PlainTextFormat:  plainTextFormat,
								OperationMode:    "CTS",
								Fill:             "No", // CTS模式不需要填充
								Key:              keyMap[keyFormat],
								KeyFormat:        keyFormat,
								Deviation:        data.deviations[deviationFormat],
								DeviationFormat:  deviationFormat,
								CipherTextFormat: cipherTextFormat,
							}

							decryptParam := param.AesDecryptParam{
								CipherTextFormat: cipherTextFormat,
								OperationMode:    "CTS",
								Fill:             "No", // CTS模式不需要填充
								Key:              keyMap[keyFormat],
								KeyFormat:        keyFormat,
								Deviation:        data.deviations[deviationFormat],
								DeviationFormat:  deviationFormat,
								PlainTextFormat:  plainTextFormat,
							}

							encrypted, err := AesEncrypt(encryptParam)
							if err != nil {
								t.Logf("加密参数: %+v", encryptParam)
								t.Errorf("加密失败: %v", err)
								return
							}
							t.Logf("加密参数: %+v", encryptParam)
							t.Logf("加密结果: %s", encrypted)

							decryptParam.CipherText = encrypted
							decrypted, err := AesDecrypt(decryptParam)
							if err != nil {
								t.Logf("解密参数: %+v", decryptParam)
								t.Errorf("解密失败: %v", err)
								return
							}
							t.Logf("解密结果: %s", decrypted)

							// 根据不同格式比较解密结果
							var expectedBytes, actualBytes []byte
							switch plainTextFormat {
							case "Text":
								expectedBytes = []byte(rawPlainText)
								actualBytes = []byte(decrypted)
							case "Base64":
								expectedBytes, err = base64.StdEncoding.DecodeString(testPlainText)
								if err != nil {
									t.Errorf("解码期望的Base64失败: %v", err)
									return
								}
								actualBytes, err = base64.StdEncoding.DecodeString(decrypted)
								if err != nil {
									t.Errorf("解码实际的Base64失败: %v", err)
									return
								}
							case "Hex":
								expectedBytes, err = hex.DecodeString(testPlainText)
								if err != nil {
									t.Errorf("解码期望的Hex失败: %v", err)
									return
								}
								actualBytes, err = hex.DecodeString(decrypted)
								if err != nil {
									t.Errorf("解码实际的Hex失败: %v", err)
									return
								}
							}

							// 添加更多的调试信息
							t.Logf("期望的字节: %x", expectedBytes)
							t.Logf("实际的字节: %x", actualBytes)
							t.Logf("期望的长度: %d", len(expectedBytes))
							t.Logf("实际的长度: %d", len(actualBytes))

							if !bytes.Equal(expectedBytes, actualBytes) {
								t.Errorf("解密结果与原文不匹配\n期望: %x\n实际: %x",
									expectedBytes, actualBytes)
							}
						})
					}
				}
			}
		}
	}
}

// TestAesEncryptErrors 测试错误情况
func TestAesEncryptErrors(t *testing.T) {
	testCases := []struct {
		name        string
		params      param.AesEncryptParam
		expectedErr string
	}{
		// 1. 参数验证错误
		{
			name: "Error-EmptyOperationMode",
			params: param.AesEncryptParam{
				PlainText:       "test",
				PlainTextFormat: "Text",
				Fill:            "PKCS7",
				Key:             "1234567890123456",
			},
			expectedErr: "operation mode is required",
		},
		{
			name: "Error-EmptyKey",
			params: param.AesEncryptParam{
				PlainText:       "test",
				PlainTextFormat: "Text",
				OperationMode:   "CBC",
				Fill:            "PKCS7",
			},
			expectedErr: "key is required",
		},
		{
			name: "Error-EmptyFill",
			params: param.AesEncryptParam{
				PlainText:       "test",
				PlainTextFormat: "Text",
				OperationMode:   "CBC",
				Key:             "1234567890123456",
				KeyFormat:       "Text",
				Deviation:       "1234567890123456",
				DeviationFormat: "Text",
			},
			expectedErr: "fill mode is required",
		},
		{
			name: "Error-EmptyPlainText",
			params: param.AesEncryptParam{
				PlainText:       "",
				PlainTextFormat: "Text",
				OperationMode:   "CBC",
				Fill:            "PKCS7",
				Key:             "1234567890123456",
			},
			expectedErr: "plainText can not be empty",
		},

		// 2. 格式验证错误
		{
			name: "Error-InvalidOperationMode",
			params: param.AesEncryptParam{
				PlainText:       "test",
				PlainTextFormat: "Text",
				OperationMode:   "INVALID",
				Fill:            "PKCS7",
				Key:             "1234567890123456",
			},
			expectedErr: "invalid aes operation mode",
		},
		{
			name: "Error-InvalidPlainTextFormat",
			params: param.AesEncryptParam{
				PlainText:       "test",
				PlainTextFormat: "INVALID",
				OperationMode:   "CBC",
				Fill:            "PKCS7",
				Key:             "1234567890123456",
			},
			expectedErr: "invalid plaintext format",
		},
		{
			name: "Error-InvalidCipherTextFormat",
			params: param.AesEncryptParam{
				PlainText:        "test",
				PlainTextFormat:  "Text",
				OperationMode:    "CBC",
				Fill:             "PKCS7",
				Key:              "1234567890123456",
				KeyFormat:        "Text",
				Deviation:        "1234567890123456",
				DeviationFormat:  "Text",
				CipherTextFormat: "INVALID",
			},
			expectedErr: "invalid ciphertext format",
		},

		// 3. 密钥相关错误
		{
			name: "Error-InvalidKeyLength-TooShort",
			params: param.AesEncryptParam{
				PlainText:       "test",
				PlainTextFormat: "Text",
				OperationMode:   "CBC",
				Fill:            "PKCS7",
				Key:             "123",
				KeyFormat:       "Text",
			},
			expectedErr: "invalid aes key length",
		},
		{
			name: "Error-InvalidKeyLength-Wrong",
			params: param.AesEncryptParam{
				PlainText:       "test",
				PlainTextFormat: "Text",
				OperationMode:   "CBC",
				Fill:            "PKCS7",
				Key:             "12345678901234567", // 17字节
				KeyFormat:       "Text",
			},
			expectedErr: "invalid aes key length",
		},
		{
			name: "Error-InvalidKeyFormat-Base64",
			params: param.AesEncryptParam{
				PlainText:       "test",
				PlainTextFormat: "Text",
				OperationMode:   "CBC",
				Fill:            "PKCS7",
				Key:             "not-base64-string",
				KeyFormat:       "Base64",
			},
			expectedErr: "failed to decode base64 key: illegal base64 data at input byte 3",
		},
		{
			name: "Error-InvalidKeyFormat-Hex",
			params: param.AesEncryptParam{
				PlainText:       "test",
				PlainTextFormat: "Text",
				OperationMode:   "CBC",
				Fill:            "PKCS7",
				Key:             "not-hex-string",
				KeyFormat:       "Hex",
			},
			expectedErr: "failed to decode hex key: encoding/hex: invalid byte: U+006E 'n'",
		},

		// 4. 偏移量相关错误
		{
			name: "Error-InvalidDeviationLength-CBC",
			params: param.AesEncryptParam{
				PlainText:       "test",
				PlainTextFormat: "Text",
				OperationMode:   "CBC",
				Fill:            "PKCS7",
				Key:             "1234567890123456",
				KeyFormat:       "Text",
				Deviation:       "123",
				DeviationFormat: "Text",
			},
			expectedErr: "invalid aes deviation length",
		},
		{
			name: "Error-InvalidDeviationLength-GCM",
			params: param.AesEncryptParam{
				PlainText:       "test",
				PlainTextFormat: "Text",
				OperationMode:   "GCM",
				Key:             "1234567890123456",
				KeyFormat:       "Text",
				Deviation:       "1234567890123456", // 16字节，GCM需要12字节
				DeviationFormat: "Text",
			},
			expectedErr: "here GCM mode only support 12 bytes deviation",
		},

		// 5. 填充模式错误
		{
			name: "Error-StreamModeWithPadding-CFB",
			params: param.AesEncryptParam{
				PlainText:       "test",
				PlainTextFormat: "Text",
				OperationMode:   "CFB",
				Fill:            "PKCS7",
				Key:             "1234567890123456",
				KeyFormat:       "Text",
				Deviation:       "1234567890123456",
				DeviationFormat: "Text",
			},
			expectedErr: "stream modes do not need padding",
		},
		{
			name: "Error-NoFillWithInvalidLength",
			params: param.AesEncryptParam{
				PlainText:       "test",
				PlainTextFormat: "Text",
				OperationMode:   "CBC",
				Fill:            "No",
				Key:             "1234567890123456",
				KeyFormat:       "Text",
				Deviation:       "1234567890123456",
				DeviationFormat: "Text",
			},
			expectedErr: "plainText length must be a multiple of 16",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			_, err := AesEncrypt(tc.params)
			if err == nil {
				t.Error("期望出错但成功了")
				return
			}
			if err.Error() != tc.expectedErr {
				t.Errorf("错误信息不匹配\n期望: %s\n实际: %s", tc.expectedErr, err.Error())
			}
		})
	}
}
