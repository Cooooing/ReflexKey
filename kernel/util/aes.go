package util

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"encoding/hex"
	"errors"
	"fmt"
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
	"PlainTextFormat":  {"Hex", "Base64", "Text"},
	"CipherTextFormat": {"Hex", "Base64"},
}

// 判断是否是流加密模式
func isStreamMode(mode string) bool {
	return mode == "OFB" || mode == "CFB" || mode == "CTR" || mode == "GCM"
}

/*
AesEncrypt 使用AES算法加密数据
所有的文本输入(plainText, key, deviation)默认使用UTF-8编码
当格式为Hex或Base64时，输入应为对应格式的字符串
返回值根据cipherTextFormat参数返回对应格式的字符串
*/
func AesEncrypt(params param.AesEncryptParam) (string, error) {
	// 验证必要数
	if params.OperationMode == "" {
		return "", errors.New("operation mode is required")
	}
	if params.Key == "" {
		return "", errors.New("key is required")
	}
	if params.Fill == "" && params.OperationMode != "GCM" {
		return "", errors.New("fill mode is required")
	}

	var err error
	var res string

	if !slices.Contains(aesType["OperationMode"], params.OperationMode) {
		return "", errors.New("invalid aes operation mode")
	}
	if len(params.PlainText) == 0 {
		return "", errors.New("plainText can not be empty")
	}

	if !slices.Contains(aesType["PlainTextFormat"], params.PlainTextFormat) {
		return "", errors.New("invalid plaintext format")
	}

	var plainTextBytes []byte
	switch params.PlainTextFormat {
	case "Hex":
		plainTextBytes, err = hex.DecodeString(params.PlainText)
		if err != nil {
			return "", fmt.Errorf("failed to decode hex plaintext: %v", err)
		}
	case "Base64":
		plainTextBytes, err = base64.StdEncoding.DecodeString(params.PlainText)
		if err != nil {
			return "", fmt.Errorf("failed to decode base64 plaintext: %v", err)
		}
	default:
		plainTextBytes = []byte(params.PlainText)
	}

	var key []byte
	keyFormat := params.KeyFormat
	if !slices.Contains(aesType["KeyFormat"], keyFormat) {
		keyFormat = "Text"
	}
	switch keyFormat {
	case "Text":
		key = []byte(params.Key)
	case "Base64":
		key, err = base64.StdEncoding.DecodeString(params.Key)
		if err != nil {
			return "", fmt.Errorf("failed to decode base64 key: %v", err)
		}
	case "Hex":
		key, err = hex.DecodeString(params.Key)
		if err != nil {
			return "", fmt.Errorf("failed to decode hex key: %v", err)
		}
	}
	if !slices.Contains(aesType["KeyLength"], strconv.Itoa(len(key))) {
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
			if err != nil {
				return "", fmt.Errorf("failed to decode base64 deviation: %v", err)
			}
		case "Hex":
			deviation, err = hex.DecodeString(params.Deviation)
			if err != nil {
				return "", fmt.Errorf("failed to decode hex deviation: %v", err)
			}
		}
		// GCM 模式IV可以有不同的长度，但最常见的长度12字节96bits（可简化实现并提高效率）。
		if params.OperationMode == "GCM" && len(deviation) != 12 {
			return "", errors.New("here GCM mode only support 12 bytes deviation")
		} else if params.OperationMode != "GCM" && len(deviation) != 16 {
			// 非GCM模式需要16字节IV
			return "", errors.New("invalid aes deviation length")
		}
	} else if params.OperationMode == "ECB" {

	} else {
		return "", errors.New("invalid aes deviation")
	}

	// 流加密模式不需要填充
	if isStreamMode(params.OperationMode) && params.Fill != "No" {
		return "", errors.New("stream modes do not need padding")
	}

	switch params.OperationMode {
	case "CBC":
		if slices.Contains(aesType["Fill"], params.Fill) {
			switch params.Fill {
			case "No":
				if len(params.PlainText)%16 != 0 {
					// CBC 的 无填充模式 密文必须是块大小（16字节128bits） 的整数倍
					return "", errors.New("plainText length must be a multiple of 16")
				}
				res, err = aesCBCEncrypt(plainTextBytes, key, deviation, noPadding)
			case "PKCS7":
				res, err = aesCBCEncrypt(plainTextBytes, key, deviation, pKCS7Padding)
			case "Zeros":
				res, err = aesCBCEncrypt(plainTextBytes, key, deviation, zerosPadding)
			}
		}
	case "ECB":
		if slices.Contains(aesType["Fill"], params.Fill) {
			switch params.Fill {
			case "No":
				res, err = aesECBEncrypt(plainTextBytes, key, noPadding)
			case "PKCS7":
				res, err = aesECBEncrypt(plainTextBytes, key, pKCS7Padding)
			case "Zeros":
				res, err = aesECBEncrypt(plainTextBytes, key, zerosPadding)
			}
		}
	case "OFB":
		if slices.Contains(aesType["Fill"], params.Fill) {
			switch params.Fill {
			case "No":
				res, err = aesOFBEncrypt(plainTextBytes, key, deviation, noPadding)
			case "PKCS7":
				res, err = aesOFBEncrypt(plainTextBytes, key, deviation, pKCS7Padding)
			case "Zeros":
				res, err = aesOFBEncrypt(plainTextBytes, key, deviation, zerosPadding)
			}
		}
	case "CFB":
		if slices.Contains(aesType["Fill"], params.Fill) {
			switch params.Fill {
			case "No":
				res, err = aesCFBEncrypt(plainTextBytes, key, deviation, noPadding)
			case "PKCS7":
				res, err = aesCFBEncrypt(plainTextBytes, key, deviation, pKCS7Padding)
			case "Zeros":
				res, err = aesCFBEncrypt(plainTextBytes, key, deviation, zerosPadding)
			}
		}
	case "CTS":
		if slices.Contains(aesType["Fill"], params.Fill) {
			switch params.Fill {
			case "No":
				res, err = aesCTSEncrypt(plainTextBytes, key, deviation, noPadding)
			case "PKCS7":
				res, err = aesCTSEncrypt(plainTextBytes, key, deviation, pKCS7Padding)
			case "Zeros":
				res, err = aesCTSEncrypt(plainTextBytes, key, deviation, zerosPadding)
			}
		}
	case "CTR":
		if slices.Contains(aesType["Fill"], params.Fill) {
			switch params.Fill {
			case "No":
				res, err = aesCTREncrypt(plainTextBytes, key, deviation, noPadding)
			case "PKCS7":
				res, err = aesCTREncrypt(plainTextBytes, key, deviation, pKCS7Padding)
			case "Zeros":
				res, err = aesCTREncrypt(plainTextBytes, key, deviation, zerosPadding)
			}
		}
	case "GCM":
		res, err = aesGCMEncrypt(plainTextBytes, key, deviation)
	}

	if err != nil {
		return "", err
	}

	if !slices.Contains(aesType["CipherTextFormat"], params.CipherTextFormat) {
		return "", errors.New("invalid ciphertext format")
	}

	switch params.CipherTextFormat {
	case "Hex":
		return hex.EncodeToString([]byte(res)), nil
	case "Base64":
		return base64.StdEncoding.EncodeToString([]byte(res)), nil
	default:
		return res, nil
	}
}

func AesDecrypt(params param.AesDecryptParam) (string, error) {
	var err error
	var res []byte

	// 验证操作模式
	if !slices.Contains(aesType["OperationMode"], params.OperationMode) {
		return "", errors.New("invalid aes operation mode")
	}
	if len(params.CipherText) == 0 {
		return "", errors.New("cipherText can not be empty")
	}

	// 处理密钥
	var key []byte
	keyFormat := params.KeyFormat
	if !slices.Contains(aesType["KeyFormat"], keyFormat) {
		keyFormat = "Text"
	}
	switch keyFormat {
	case "Text":
		key = []byte(params.Key)
	case "Base64":
		key, err = base64.StdEncoding.DecodeString(params.Key)
		if nil != err {
			return "", err
		}
	case "Hex":
		key, err = hex.DecodeString(params.Key)
		if nil != err {
			return "", err
		}
	}
	if !slices.Contains(aesType["KeyLength"], strconv.Itoa(len(key))) {
		return "", errors.New("invalid aes key length")
	}

	// 处理偏移量(IV)
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
		if params.OperationMode == "GCM" {
			if len(deviation) != 12 {
				return "", errors.New("here GCM mode only support 12 bytes deviation")
			}
		} else {
			if len(deviation) != 16 {
				return "", errors.New("invalid aes deviation length")
			}
		}
	} else if params.OperationMode == "ECB" {
		// ECB模式不需要IV
	} else {
		return "", errors.New("invalid aes deviation")
	}

	// 处理密文格式
	var cipherTextBytes []byte
	switch params.CipherTextFormat {
	case "Hex":
		cipherTextBytes, err = hex.DecodeString(params.CipherText)
		if err != nil {
			return "", fmt.Errorf("failed to decode hex ciphertext: %v", err)
		}
	case "Base64":
		cipherTextBytes, err = base64.StdEncoding.DecodeString(params.CipherText)
		if err != nil {
			return "", fmt.Errorf("failed to decode base64 ciphertext: %v", err)
		}
	default:
		cipherTextBytes = []byte(params.CipherText)
	}

	// 根据不同模式进行解密
	switch params.OperationMode {
	case "CBC":
		if slices.Contains(aesType["Fill"], params.Fill) {
			switch params.Fill {
			case "No":
				res, err = aesCBCDecrypt(string(cipherTextBytes), key, deviation, noUnpadding)
			case "PKCS7":
				res, err = aesCBCDecrypt(string(cipherTextBytes), key, deviation, pKCS7Unpadding)
			case "Zeros":
				res, err = aesCBCDecrypt(string(cipherTextBytes), key, deviation, zerosUnpadding)
			}
		}
	case "ECB":
		if slices.Contains(aesType["Fill"], params.Fill) {
			switch params.Fill {
			case "No":
				res, err = aesECBDecrypt(string(cipherTextBytes), key, noUnpadding)
			case "PKCS7":
				res, err = aesECBDecrypt(string(cipherTextBytes), key, pKCS7Unpadding)
			case "Zeros":
				res, err = aesECBDecrypt(string(cipherTextBytes), key, zerosUnpadding)
			}
		}
	case "OFB":
		res, err = aesOFBDecrypt(string(cipherTextBytes), key, deviation)
	case "CFB":
		res, err = aesCFBDecrypt(string(cipherTextBytes), key, deviation)
	case "CTS":
		res, err = aesCTSDecrypt(string(cipherTextBytes), key, deviation)
	case "CTR":
		res, err = aesCTRDecrypt(string(cipherTextBytes), key, deviation)
	case "GCM":
		res, err = aesGCMDecrypt(string(cipherTextBytes), key, deviation)
	}

	if err != nil {
		return "", err
	}

	if !slices.Contains(aesType["PlainTextFormat"], params.PlainTextFormat) {
		return "", errors.New("invalid plaintext format")
	}

	// 格式化解密结果
	switch params.PlainTextFormat {
	case "Hex":
		return hex.EncodeToString(res), nil
	case "Base64":
		return base64.StdEncoding.EncodeToString(res), nil
	default:
		return string(res), nil
	}
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
CBC模式通过将每个明文块与前一个密文块进行异或运算后再加密，确保相同的明文块生��不同的密文块。
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

	plaintext := make([]byte, len(data))
	mode := cipher.NewCBCDecrypter(block, iv)
	mode.CryptBlocks(plaintext, data)

	return unpadding(plaintext)
}

/*
ECB模式是最简单的模式，每个明文块独加密，相同的明文块生成相同的密文块。
警告：此模式存在安全隐患，不建议用于敏感数据加密。
仅适用于小数据量且对安全性要求不高的场景。
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
	for bs, be := 0, block.BlockSize(); bs < len(data); bs, be = bs+block.BlockSize(), be+block.BlockSize() {
		block.Decrypt(plaintext[bs:be], data[bs:be])
	}

	return unpadding(plaintext)
}

/*
OFB模式将加密器的输出反馈到下一个块的输入中，生成密钥流，然后与明文块异或生成密文块。
这种模式以转换成流密码，用于实时数据传输，如语音和视频流。
但错误传播问题会导致后续数据受影响。
*/

// AesOFBEncrypt 使用OFB模式进行AES密
func aesOFBEncrypt(plainText, key, iv []byte, padding func([]byte, int) []byte) (string, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}
	// OFB是流模式，不需要填充和块大小检查
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
	stream := cipher.NewOFB(block, iv)
	plaintext := make([]byte, len(data))
	stream.XORKeyStream(plaintext, data)
	return plaintext, nil
}

/*
CFB模式类似于OFB，但使用密文反馈来生成密钥流。
这种模式也适用于实时数据传输，如网络通信。
CFB模式可以理任意长度的数据但错误传播问题会影响后续数据。
*/

// AesCFBEncrypt 使用CFB模式进行AES加密
func aesCFBEncrypt(plainText, key, iv []byte, padding func([]byte, int) []byte) (string, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}
	// CFB是流模式，不需要填充和块大小检查
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
CTS模式通过“窃取”最后一个完整块的一部分来处理不完整块，适用于需要处理任意长度数据的景，如文件加密。
*/

// AesCTSEncrypt 使用CTS模式进行AES加密
func aesCTSEncrypt(plainText, key, iv []byte, padding func([]byte, int) []byte) (string, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}

	if len(plainText) < aes.BlockSize {
		return "", errors.New("plaintext too short")
	}

	// CTS模式需要至少两个块
	if len(plainText) < 2*aes.BlockSize {
		return "", errors.New("CTS mode needs at least two blocks of plaintext")
	}

	// 实现CBC-CTS (Cipher Text Stealing)
	ciphertext := make([]byte, len(plainText))

	// 1. 先用CBC模式加密除最后两个块之外的所有块
	mode := cipher.NewCBCEncrypter(block, iv)
	if len(plainText) > 2*aes.BlockSize {
		mode.CryptBlocks(ciphertext[:len(plainText)-2*aes.BlockSize],
			plainText[:len(plainText)-2*aes.BlockSize])
	}

	// 2. 加密倒数第二个块
	penultimateBlock := make([]byte, aes.BlockSize)
	copy(penultimateBlock, plainText[len(plainText)-2*aes.BlockSize:len(plainText)-aes.BlockSize])
	mode.CryptBlocks(ciphertext[len(plainText)-2*aes.BlockSize:len(plainText)-aes.BlockSize],
		penultimateBlock)

	// 3. 加密最后一个块
	lastBlock := make([]byte, aes.BlockSize)
	copy(lastBlock, plainText[len(plainText)-aes.BlockSize:])

	// 使用倒数第二个密文块作为IV
	lastIV := ciphertext[len(plainText)-2*aes.BlockSize : len(plainText)-aes.BlockSize]
	lastMode := cipher.NewCBCEncrypter(block, lastIV)
	lastMode.CryptBlocks(ciphertext[len(plainText)-aes.BlockSize:], lastBlock)

	return base64.StdEncoding.EncodeToString(ciphertext), nil
}

// AesCTSDecrypt 使用CTS模式进行AES解密
func aesCTSDecrypt(cipherText string, key, iv []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	data, err := base64.StdEncoding.DecodeString(cipherText)
	if err != nil {
		return nil, err
	}

	if len(data) < 2*aes.BlockSize {
		return nil, errors.New("CTS mode needs at least two blocks of ciphertext")
	}

	// 实现CBC-CTS解密
	plaintext := make([]byte, len(data))

	// 1. 先用CBC模式解密除最后两个块之外的所有块
	mode := cipher.NewCBCDecrypter(block, iv)
	if len(data) > 2*aes.BlockSize {
		mode.CryptBlocks(plaintext[:len(data)-2*aes.BlockSize],
			data[:len(data)-2*aes.BlockSize])
	}

	// 2. 解密倒数第二个块
	penultimateBlock := make([]byte, aes.BlockSize)
	copy(penultimateBlock, data[len(data)-2*aes.BlockSize:len(data)-aes.BlockSize])
	mode.CryptBlocks(plaintext[len(data)-2*aes.BlockSize:len(data)-aes.BlockSize],
		penultimateBlock)

	// 3. 解密最后一个块
	lastBlock := make([]byte, aes.BlockSize)
	copy(lastBlock, data[len(data)-aes.BlockSize:])

	// 使用倒数第二个密文块作为IV
	lastIV := data[len(data)-2*aes.BlockSize : len(data)-aes.BlockSize]
	lastMode := cipher.NewCBCDecrypter(block, lastIV)
	lastMode.CryptBlocks(plaintext[len(data)-aes.BlockSize:], lastBlock)

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
	// CTR是流模式，不需要填充和块大小检查
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
GCM模式是CTR模式的扩展，增加了认证功能，可以同时提供数据加密和整性验证。
这种模式适用于需要高安全性和数据完整性的场景，如TLS协议和安全通信。
*/

// AesGCMEncrypt 使用GCM模式进行AES加密
func aesGCMEncrypt(plainText, key, nonce []byte) (string, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}

	aesgcm, err := cipher.NewGCM(block)
	if err != nil {
		return "", err
	}

	// 可以添加可选的认证数据(aad)
	// ciphertext := aesgcm.Seal(nil, nonce, plainText, aad)
	ciphertext := aesgcm.Seal(nil, nonce, plainText, nil)
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
