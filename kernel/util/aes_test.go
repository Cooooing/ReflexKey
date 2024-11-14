package util

import (
	"fmt"
	"log"
	"testing"
)

// 测试 AES 加解密
func TestAes(t *testing.T) {
	// 设置密钥和 IV
	key := []byte("1234567890123456") // 16 字节密钥
	iv := []byte("1234567890123456")  // 16 字节 IV
	plainText := []byte("This is a test message.")

	// 测试不同模式与填充方式
	modes := []struct {
		mode        string
		paddingMode PaddingMode
		encrypt     func([]byte, []byte, []byte, PaddingMode) ([]byte, error)
		decrypt     func([]byte, []byte, []byte, PaddingMode) ([]byte, error)
	}{
		{"CBC", PKCS7, aesCBCEncrypt, aesCBCDecrypt},
		{"ECB", PKCS7, aesECBEncrypt, aesECBDecrypt},
	}

	for _, test := range modes {
		t.Run(fmt.Sprintf("Mode: %s, Padding: %d", test.mode, test.paddingMode), func(t *testing.T) {
			// 加密
			cipherText, err := test.encrypt(plainText, key, iv, test.paddingMode)
			if err != nil {
				t.Fatalf("Encryption failed: %v", err)
			}
			fmt.Printf("%s Encrypted: %s\n", test.mode, bytesToBase64(cipherText))

			// 解密
			decPlainText, err := test.decrypt(cipherText, key, iv, test.paddingMode)
			if err != nil {
				t.Fatalf("Decryption failed: %v", err)
			}
			fmt.Printf("%s Decrypted: %s\n", test.mode, string(decPlainText))

			// 验证是否解密正确
			if string(decPlainText) != string(plainText) {
				t.Fatalf("Decrypted text doesn't match original")
			}
		})
	}
}

func main() {
	// 运行测试
	err := testing.RunTests(func(pat, str string) (bool, error) {
		return true, nil
	}, []testing.InternalTest{
		{
			Name: "TestAes",
			F:    TestAes,
		},
	})
	if err != nil {
		log.Fatalf("Test failed: %v", err)
	}
}
