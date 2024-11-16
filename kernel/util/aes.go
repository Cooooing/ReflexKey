package util

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"encoding/hex"
	"errors"
	"io"
	"kernel/model/param"
	"slices"
	"strconv"
)

var aesType = map[string][]string{
	"OperationMode":    {"CBC", "ECB", "OFB", "CFB", "CTS", "CTR", "GCM"},
	"Fill":             {"PKCS7", "Zeros", "No"},
	"KeyLength":        {"16", "24", "32"},
	"KeyFormat":        {"Hex", "Base64", "Text"},
	"DeviationFormat":  {"Hex", "Base64", "Text"},
	"PlainTextFormat":  {"Hex", "Base64"},
	"CipherTextFormat": {"Hex", "Base64"},
}

func AesEncrypt(params param.AesEncryptParam) (string, error) {
	var err error
	var res string

	if !slices.Contains(aesType["OperationMode"], params.OperationMode) {
		return "", errors.New("invalid aes operation mode")
	}
	if len(params.PlainText) == 0 {
		return "", errors.New("plainText can not be empty")
	}
	plainText := []byte(params.PlainText)

	var key []byte
	keyFormat := params.KeyFormat
	if !slices.Contains(aesType["KeyFormat"], keyFormat) {
		keyFormat = "Text"
	}
	switch keyFormat {
	case "Text":
		key = []byte(params.Key)
	case "Base64":
		key, err = base64.StdEncoding.DecodeString(params.Deviation)
		if nil != err {
			return "", err
		}
	case "Hex":
		key, err = hex.DecodeString(params.Key)
		if nil != err {
			return "", err
		}
	}
	if !slices.Contains(aesType["KeyLength"], strconv.Itoa(len([]byte(params.Key)))) {
		return "", errors.New("invalid aes key length")
	}

	// ECB 模式不使用初始化向量（IV）
	var deviation []byte
	if params.OperationMode != "ECB" && slices.Contains(aesType["DeviationFormat"], params.DeviationFormat) {
		switch params.DeviationFormat {
		case "Text":
			deviation = []byte(params.Deviation)
		case "Base64":
			deviation, err = base64.StdEncoding.DecodeString(params.Deviation)
			if nil != err {
				return "", err
			}
		case "Hex":
			deviation, err = hex.DecodeString(params.Deviation)
			if nil != err {
				return "", err
			}
		}
		// GCM 模式IV可以有不同的长度，但最常见的长度是12字节96bits（可以简化实现并提高效率）。
		if params.OperationMode == "GCM" && len(deviation) != 12 {
			return "", errors.New("here GCM mode only support 12 bytes deviation")
		}
		// IV 的长度必须是16字节（128bits）
		if len(deviation) != 16 {
			return "", errors.New("invalid aes deviation length")
		}
	} else if params.OperationMode == "ECB" {

	} else {
		return "", errors.New("invalid aes deviation")
	}

	switch params.OperationMode {
	case "CBC":
		if slices.Contains(aesType["Fill"], params.Fill) {
			switch params.Fill {
			case "No":
				if len(params.PlainText)%16 == 0 {
					// CBC 的 无填充模式 密文必须是块大小（16字节128bits） 的整数倍
					return "", errors.New("plainText length must be a multiple of 16")
				}
				res, err = aesCBCEncrypt(plainText, key, deviation, noPadding)
			case "PKCS7":
				res, err = aesCBCEncrypt(plainText, key, deviation, pKCS7Padding)
			case "Zeros":
				res, err = aesCBCEncrypt(plainText, key, deviation, zerosPadding)
			}
		}
	case "ECB":
		if slices.Contains(aesType["Fill"], params.Fill) {
			switch params.Fill {
			case "No":
				res, err = aesECBEncrypt(plainText, key, noPadding)
			case "PKCS7":
				res, err = aesECBEncrypt(plainText, key, pKCS7Padding)
			case "Zeros":
				res, err = aesECBEncrypt(plainText, key, zerosPadding)
			}
		}
	case "OFB":
		if slices.Contains(aesType["Fill"], params.Fill) {
			switch params.Fill {
			case "No":
				res, err = aesOFBEncrypt(plainText, key, deviation, noPadding)
			case "PKCS7":
				res, err = aesOFBEncrypt(plainText, key, deviation, pKCS7Padding)
			case "Zeros":
				res, err = aesOFBEncrypt(plainText, key, deviation, zerosPadding)
			}
		}
	case "CFB":
		if slices.Contains(aesType["Fill"], params.Fill) {
			switch params.Fill {
			case "No":
				res, err = aesCFBEncrypt(plainText, key, deviation, noPadding)
			case "PKCS7":
				res, err = aesCFBEncrypt(plainText, key, deviation, pKCS7Padding)
			case "Zeros":
				res, err = aesCFBEncrypt(plainText, key, deviation, zerosPadding)
			}
		}
	case "CTS":
		if slices.Contains(aesType["Fill"], params.Fill) {
			switch params.Fill {
			case "No":
				res, err = aesCTSEncrypt(plainText, key, deviation, noPadding)
			case "PKCS7":
				res, err = aesCTSEncrypt(plainText, key, deviation, pKCS7Padding)
			case "Zeros":
				res, err = aesCTSEncrypt(plainText, key, deviation, zerosPadding)
			}
		}
	case "CTR":
		if slices.Contains(aesType["Fill"], params.Fill) {
			switch params.Fill {
			case "No":
				res, err = aesCTREncrypt(plainText, key, deviation, noPadding)
			case "PKCS7":
				res, err = aesCTREncrypt(plainText, key, deviation, pKCS7Padding)
			case "Zeros":
				res, err = aesCTREncrypt(plainText, key, deviation, zerosPadding)
			}
		}
	case "GCM":
		if slices.Contains(aesType["Fill"], params.Fill) {
			switch params.Fill {
			case "No":
				res, err = aesGCMEncrypt(plainText, key, deviation)
			case "PKCS7":
				res, err = aesGCMEncrypt(plainText, key, deviation)
			case "Zeros":
				res, err = aesGCMEncrypt(plainText, key, deviation)
			}
		}
	}
	return res, err
}

func AesDecrypt(params param.AesDecryptParam) (string, error) {
	var err error
	var res string

	return res, err
}

// generateRandomBytes 生成指定长度的随机字节数组
func generateRandomBytes(n int) ([]byte, error) {
	b := make([]byte, n)
	_, err := io.ReadFull(rand.Reader, b)
	if err != nil {
		return nil, err
	}
	return b, nil
}

// PKCS7Padding 使用PKCS7标准填充数据到块大小的倍数
func pKCS7Padding(src []byte, blockSize int) []byte {
	padding := blockSize - len(src)%blockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(src, padtext...)
}

// PKCS7Unpadding 移除PKCS7填充的数据
func pKCS7Unpadding(src []byte) ([]byte, error) {
	length := len(src)
	unpadding := int(src[length-1])
	if unpadding > length || unpadding == 0 {
		return nil, errors.New("invalid pkcs7 padding")
	}
	return src[:length-unpadding], nil
}

// ZerosPadding 使用零字节填充数据到块大小的倍数
func zerosPadding(src []byte, blockSize int) []byte {
	padding := blockSize - len(src)%blockSize
	padtext := bytes.Repeat([]byte{0}, padding)
	return append(src, padtext...)
}

// ZerosUnpadding 移除零字节填充的数据
func zerosUnpadding(src []byte) ([]byte, error) {
	for i := len(src) - 1; i >= 0; i-- {
		if src[i] != 0 {
			return src[:i+1], nil
		}
	}
	return src, nil
}

// NoPadding 不进行任何填充
func noPadding(src []byte, blockSize int) []byte {
	return src
}

// NoUnpadding 不进行任何填充的移除操作
func noUnpadding(src []byte) ([]byte, error) {
	return src, nil
}

/*
CBC模式通过将每个明文块与前一个密文块进行异或运算后再加密，确保相同的明文块生成不同的密文块。
这种模式提高了安全性，但需要初始化向量（IV），并且加密过程不能并行化。
适用于需要高安全性的场景，如文件加密和网络通信。
*/

// AesCBCEncrypt 使用CBC模式进行AES加密
func aesCBCEncrypt(plainText, key, iv []byte, padding func([]byte, int) []byte) (string, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}
	blockSize := block.BlockSize()
	plainText = padding(plainText, blockSize)
	if len(plainText)%blockSize != 0 {
		return "", errors.New("plaintext is not a multiple of the block size")
	}
	ciphertext := make([]byte, len(plainText))
	mode := cipher.NewCBCEncrypter(block, iv)
	mode.CryptBlocks(ciphertext, plainText)
	return base64.StdEncoding.EncodeToString(ciphertext), nil
}

// AesCBCDecrypt 使用CBC模式进行AES解密
func aesCBCDecrypt(cipherText string, key, iv []byte, unpadding func([]byte) ([]byte, error)) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	data, err := base64.StdEncoding.DecodeString(cipherText)
	if err != nil {
		return nil, err
	}
	if len(data)%aes.BlockSize != 0 {
		return nil, errors.New("ciphertext is not a multiple of the block size")
	}
	mode := cipher.NewCBCDecrypter(block, iv)
	mode.CryptBlocks(data, data)
	return unpadding(data)
}

/*
ECB模式是最简单的模式，每个明文块独立加密，相同的明文块生成相同的密文块。
实现简单，但安全性较低，容易暴露数据模式。
适用于小数据量且对安全性要求不高的场景，如简单的数据加密。
*/

// AesECBEncrypt 使用ECB模式进行AES加密
func aesECBEncrypt(plainText, key []byte, padding func([]byte, int) []byte) (string, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}
	blockSize := block.BlockSize()
	plainText = padding(plainText, blockSize)
	ciphertext := make([]byte, len(plainText))
	for bs, be := 0, blockSize; bs < len(plainText); bs, be = bs+blockSize, be+blockSize {
		block.Encrypt(ciphertext[bs:be], plainText[bs:be])
	}
	return base64.StdEncoding.EncodeToString(ciphertext), nil
}

// AesECBDecrypt 使用ECB模式进行AES解密
func aesECBDecrypt(cipherText string, key []byte, unpadding func([]byte) ([]byte, error)) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	data, err := base64.StdEncoding.DecodeString(cipherText)
	if err != nil {
		return nil, err
	}
	if len(data)%aes.BlockSize != 0 {
		return nil, errors.New("ciphertext is not a multiple of the block size")
	}
	plaintext := make([]byte, len(data))
	blockSize := block.BlockSize()
	for bs, be := 0, blockSize; bs < len(data); bs, be = bs+blockSize, be+blockSize {
		block.Decrypt(plaintext[bs:be], data[bs:be])
	}
	return unpadding(plaintext)
}

/*
OFB模式将加密器的输出反馈到下一个块的输入中，生成密钥流，然后与明文块异或生成密文块。
这种模式可以转换成流密码，适用于实时数据传输，如语音和视频流。
但错误传播问题会导致后续数据受影响。
*/

// AesOFBEncrypt 使用OFB模式进行AES加密
func aesOFBEncrypt(plainText, key, iv []byte, padding func([]byte, int) []byte) (string, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}
	blockSize := block.BlockSize()
	plainText = padding(plainText, blockSize)
	if len(plainText)%blockSize != 0 {
		return "", errors.New("plaintext is not a multiple of the block size")
	}
	stream := cipher.NewOFB(block, iv)
	ciphertext := make([]byte, len(plainText))
	stream.XORKeyStream(ciphertext, plainText)
	return base64.StdEncoding.EncodeToString(ciphertext), nil
}

// AesOFBDecrypt 使用OFB模式进行AES解密
func aesOFBDecrypt(cipherText string, key, iv []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	data, err := base64.StdEncoding.DecodeString(cipherText)
	if err != nil {
		return nil, err
	}
	if len(data)%aes.BlockSize != 0 {
		return nil, errors.New("ciphertext is not a multiple of the block size")
	}
	stream := cipher.NewOFB(block, iv)
	plaintext := make([]byte, len(data))
	stream.XORKeyStream(plaintext, data)
	return plaintext, nil
}

/*
CFB模式类似于OFB，但使用密文反馈来生成密钥流。
这种模式也适用于实时数据传输，如网络通信。
CFB模式可以处理任意长度的数据，但错误传播问题会影响后续数据。
*/

// AesCFBEncrypt 使用CFB模式进行AES加密
func aesCFBEncrypt(plainText, key, iv []byte, padding func([]byte, int) []byte) (string, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}
	blockSize := block.BlockSize()
	plainText = padding(plainText, blockSize)
	if len(plainText)%blockSize != 0 {
		return "", errors.New("plaintext is not a multiple of the block size")
	}
	stream := cipher.NewCFBEncrypter(block, iv)
	ciphertext := make([]byte, len(plainText))
	stream.XORKeyStream(ciphertext, plainText)
	return base64.StdEncoding.EncodeToString(ciphertext), nil
}

// AesCFBDecrypt 使用CFB模式进行AES解密
func aesCFBDecrypt(cipherText string, key, iv []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	data, err := base64.StdEncoding.DecodeString(cipherText)
	if err != nil {
		return nil, err
	}
	stream := cipher.NewCFBDecrypter(block, iv)
	plaintext := make([]byte, len(data))
	stream.XORKeyStream(plaintext, data)
	return plaintext, nil
}

/*
CTS模式是对CBC模式的改进，解决了CBC模式中最后一个块不足一个完整块的问题。
CTS模式通过“窃取”最后一个完整块的一部分来处理不完整块，适用于需要处理任意长度数据的场景，如文件加密。
*/

// AesCTSEncrypt 使用CTR模式进行AES加密
func aesCTSEncrypt(plainText, key, iv []byte, padding func([]byte, int) []byte) (string, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}
	blockSize := block.BlockSize()
	plainText = padding(plainText, blockSize)
	if len(plainText)%blockSize != 0 {
		return "", errors.New("plaintext is not a multiple of the block size")
	}
	stream := cipher.NewCTR(block, iv)
	ciphertext := make([]byte, len(plainText))
	stream.XORKeyStream(ciphertext, plainText)
	return base64.StdEncoding.EncodeToString(ciphertext), nil
}

// AesCTSDecrypt 使用CTR模式进行AES解密
func aesCTSDecrypt(cipherText string, key, iv []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	data, err := base64.StdEncoding.DecodeString(cipherText)
	if err != nil {
		return nil, err
	}
	stream := cipher.NewCTR(block, iv)
	plaintext := make([]byte, len(data))
	stream.XORKeyStream(plaintext, data)
	return plaintext, nil
}

/*
CTR模式使用一个计数器作为加密器的输入，每次加密时计数器增加，生成的密钥流与明文块异或生成密文块。
这种模式可以并行处理，速度快，适用于大量数据的加密，如硬盘加密和大数据处理。
*/

// AesCTREncrypt 使用CTR模式进行AES加密
func aesCTREncrypt(plainText, key, iv []byte, padding func([]byte, int) []byte) (string, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}
	blockSize := block.BlockSize()
	plainText = padding(plainText, blockSize)
	if len(plainText)%blockSize != 0 {
		return "", errors.New("plaintext is not a multiple of the block size")
	}
	stream := cipher.NewCTR(block, iv)
	ciphertext := make([]byte, len(plainText))
	stream.XORKeyStream(ciphertext, plainText)
	return base64.StdEncoding.EncodeToString(ciphertext), nil
}

// AesCTRDecrypt 使用CTR模式进行AES解密
func aesCTRDecrypt(cipherText string, key, iv []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	data, err := base64.StdEncoding.DecodeString(cipherText)
	if err != nil {
		return nil, err
	}
	stream := cipher.NewCTR(block, iv)
	plaintext := make([]byte, len(data))
	stream.XORKeyStream(plaintext, data)
	return plaintext, nil
}

/*
GCM模式是CTR模式的扩展，增加了认证功能，可以同时提供数据加密和完整性验证。
这种模式适用于需要高安全性和数据完整性的场景，如TLS协议和安全通信。
*/

// AesGCMEncrypt 使用GCM模式进行AES加密
func aesGCMEncrypt(plainText, key, nonce []byte) (string, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}
	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return "", err
	}
	ciphertext := gcm.Seal(nil, nonce, plainText, nil)
	return base64.StdEncoding.EncodeToString(ciphertext), nil
}

// AesGCMDecrypt 使用GCM模式进行AES解密
func aesGCMDecrypt(cipherText string, key, nonce []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return nil, err
	}
	data, err := base64.StdEncoding.DecodeString(cipherText)
	if err != nil {
		return nil, err
	}
	plaintext, err := gcm.Open(nil, nonce, data, nil)
	if err != nil {
		return nil, err
	}
	return plaintext, nil
}
