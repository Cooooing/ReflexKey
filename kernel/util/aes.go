package util

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"errors"
	"fmt"
)

// AES 加密模式
const (
	CBC = iota
	ECB
	CFB
	OFB
	CTR
	GCM
)

// 填充方式
const (
	NoPadding = iota
	PKCS7Padding
	ZerosPadding
)

// GenerateKey 用于生成指定长度的 AES 密钥
func GenerateKey(length int) ([]byte, error) {
	// 生成的密钥长度：128bits, 192bits, 256bits
	var key []byte
	if length != 16 && length != 24 && length != 32 {
		return nil, errors.New("invalid key length")
	}
	key = make([]byte, length)
	_, err := rand.Read(key)
	if err != nil {
		return nil, err
	}
	return key, nil
}

// AES 加密函数：支持不同模式、填充方式和密钥长度
func AesEncrypt(plainText []byte, key []byte, mode int, iv []byte, padding int) ([]byte, error) {
	// 对 CBC 和 ECB 模式应用填充
	if mode == CBC || mode == ECB {
		plainText = applyPadding(plainText, padding)
	}

	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	var cipherText []byte
	switch mode {
	case CBC:
		cipherText, err = aesCBCEncrypt(block, plainText, iv)
	case ECB:
		//cipherText, err = aesECBEncrypt(block, plainText)
	case CFB:
		cipherText, err = aesCFBEncrypt(block, plainText, iv)
	case OFB:
		cipherText, err = aesOFBEncrypt(block, plainText, iv)
	case CTR:
		cipherText, err = aesCTREncrypt(block, plainText, iv)
	case GCM:
		cipherText, err = aesGCMEncrypt(block, plainText, iv)
	default:
		err = errors.New("invalid mode")
	}

	if err != nil {
		return nil, err
	}

	return cipherText, nil
}

// AES 解密函数：支持不同模式、填充方式和密钥长度
func AesDecrypt(cipherText []byte, key []byte, mode int, iv []byte, padding int) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	var plainText []byte
	switch mode {
	case CBC:
		plainText, err = aesCBCDecrypt(block, cipherText, iv)
	case ECB:
		//plainText, err = aesECBDecrypt(block, cipherText)
	case CFB:
		plainText, err = aesCFBDecrypt(block, cipherText, iv)
	case OFB:
		plainText, err = aesOFBDecrypt(block, cipherText, iv)
	case CTR:
		plainText, err = aesCTRDecrypt(block, cipherText, iv)
	case GCM:
		plainText, err = aesGCMDecrypt(block, cipherText, iv)
	default:
		err = errors.New("invalid mode")
	}

	if err != nil {
		return nil, err
	}

	// 对 CBC 和 ECB 模式应用填充去除
	if mode == CBC || mode == ECB {
		plainText = removePadding(plainText, padding)
	}

	return plainText, nil
}

// 适用填充方式：PKCS7, Zeros 或 无填充
func applyPadding(data []byte, padding int) []byte {
	switch padding {
	case NoPadding:
		return data
	case PKCS7Padding:
		return pkcs7Padding(data, aes.BlockSize)
	case ZerosPadding:
		return zerosPadding(data, aes.BlockSize)
	default:
		return data
	}
}

// PKCS7 填充
func pkcs7Padding(data []byte, blockSize int) []byte {
	padding := blockSize - len(data)%blockSize
	paddingText := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(data, paddingText...)
}

// Zeros 填充
func zerosPadding(data []byte, blockSize int) []byte {
	padding := blockSize - len(data)%blockSize
	paddingText := make([]byte, padding)
	return append(data, paddingText...)
}

// 去除填充
func removePadding(data []byte, padding int) []byte {
	switch padding {
	case NoPadding:
		return data
	case PKCS7Padding:
		return pkcs7RemovePadding(data)
	case ZerosPadding:
		return zerosRemovePadding(data)
	default:
		return data
	}
}

// PKCS7 去除填充
func pkcs7RemovePadding(data []byte) []byte {
	padding := data[len(data)-1]
	return data[:len(data)-int(padding)]
}

// Zeros 去除填充
func zerosRemovePadding(data []byte) []byte {
	for i := len(data) - 1; i >= 0; i-- {
		if data[i] != 0 {
			return data[:i+1]
		}
	}
	return data
}

// CBC 加密
func aesCBCEncrypt(block cipher.Block, plaintext []byte, iv []byte) ([]byte, error) {
	if len(plaintext)%aes.BlockSize != 0 {
		return nil, fmt.Errorf("plaintext is not a multiple of block size")
	}
	ciphertext := make([]byte, len(plaintext))
	mode := cipher.NewCBCEncrypter(block, iv)
	mode.CryptBlocks(ciphertext, plaintext)
	return ciphertext, nil
}

// CBC 解密
func aesCBCDecrypt(block cipher.Block, ciphertext []byte, iv []byte) ([]byte, error) {
	if len(ciphertext)%aes.BlockSize != 0 {
		return nil, fmt.Errorf("ciphertext is not a multiple of block size")
	}
	plaintext := make([]byte, len(ciphertext))
	mode := cipher.NewCBCDecrypter(block, iv)
	mode.CryptBlocks(plaintext, ciphertext)
	return plaintext, nil
}

//// ECB 加密
//func aesECBEncrypt(block cipher.Block, plaintext []byte) ([]byte, error) {
//	ciphertext := make([]byte, len(plaintext))
//	mode := cipher.NewECBEncrypter(block)
//	mode.CryptBlocks(ciphertext, plaintext)
//	return ciphertext, nil
//}
//
//// ECB 解密
//func aesECBDecrypt(block cipher.Block, ciphertext []byte) ([]byte, error) {
//	plaintext := make([]byte, len(ciphertext))
//	mode := cipher.NewECBDecrypter(block)
//	mode.CryptBlocks(plaintext, ciphertext)
//	return plaintext, nil
//}

// CFB 加密
func aesCFBEncrypt(block cipher.Block, plaintext []byte, iv []byte) ([]byte, error) {
	stream := cipher.NewCFBEncrypter(block, iv)
	ciphertext := make([]byte, len(plaintext))
	stream.XORKeyStream(ciphertext, plaintext)
	return ciphertext, nil
}

// CFB 解密
func aesCFBDecrypt(block cipher.Block, ciphertext []byte, iv []byte) ([]byte, error) {
	stream := cipher.NewCFBDecrypter(block, iv)
	plaintext := make([]byte, len(ciphertext))
	stream.XORKeyStream(plaintext, ciphertext)
	return plaintext, nil
}

// OFB 加密
func aesOFBEncrypt(block cipher.Block, plaintext []byte, iv []byte) ([]byte, error) {
	stream := cipher.NewOFB(block, iv)
	ciphertext := make([]byte, len(plaintext))
	stream.XORKeyStream(ciphertext, plaintext)
	return ciphertext, nil
}

// OFB 解密
func aesOFBDecrypt(block cipher.Block, ciphertext []byte, iv []byte) ([]byte, error) {
	stream := cipher.NewOFB(block, iv)
	plaintext := make([]byte, len(ciphertext))
	stream.XORKeyStream(plaintext, ciphertext)
	return plaintext, nil
}

// CTR 加密
func aesCTREncrypt(block cipher.Block, plaintext []byte, iv []byte) ([]byte, error) {
	stream := cipher.NewCTR(block, iv)
	ciphertext := make([]byte, len(plaintext))
	stream.XORKeyStream(ciphertext, plaintext)
	return ciphertext, nil
}

// CTR 解密
func aesCTRDecrypt(block cipher.Block, ciphertext []byte, iv []byte) ([]byte, error) {
	stream := cipher.NewCTR(block, iv)
	plaintext := make([]byte, len(ciphertext))
	stream.XORKeyStream(plaintext, ciphertext)
	return plaintext, nil
}

// GCM 加密
func aesGCMEncrypt(block cipher.Block, plaintext []byte, iv []byte) ([]byte, error) {
	nonce := iv
	aead, err := cipher.NewGCM(block)
	if err != nil {
		return nil, err
	}
	ciphertext := aead.Seal(nil, nonce, plaintext, nil)
	return ciphertext, nil
}

// GCM 解密
func aesGCMDecrypt(block cipher.Block, ciphertext []byte, iv []byte) ([]byte, error) {
	nonce := iv
	aead, err := cipher.NewGCM(block)
	if err != nil {
		return nil, err
	}
	plaintext, err := aead.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		return nil, err
	}
	return plaintext, nil
}
