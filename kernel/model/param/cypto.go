package param

// swagger:model
type AesEncryptParam struct {
	PlainText        string `json:"plainText"`        // 明文
	PlainTextFormat  string `json:"plainTextFormat"`  // 明文格式
	OperationMode    string `json:"operationMode"`    // 加密模式
	Fill             string `json:"fill"`             // 填充模式
	Key              string `json:"key"`              // 密钥
	KeyFormat        string `json:"keyFormat"`        // 密钥格式
	Deviation        string `json:"deviation"`        // 偏移量
	DeviationFormat  string `json:"deviationFormat"`  // 偏移量格式
	CipherTextFormat string `json:"cipherTextFormat"` // 密文输出格式
}

// swagger:model
type AesDecryptParam struct {
	CipherText       string `json:"cipherText"`       // 密文
	CipherTextFormat string `json:"cipherTextFormat"` // 密文格式
	OperationMode    string `json:"operationMode"`    // 解密模式
	Fill             string `json:"fill"`             // 填充模式
	Key              string `json:"key"`              // 密钥
	KeyFormat        string `json:"keyFormat"`        // 密钥格式
	Deviation        string `json:"deviation"`        // 偏移量
	DeviationFormat  string `json:"deviationFormat"`  // 偏移量格式
	PlainTextFormat  string `json:"plainTextFormat"`  // 明文格式
}

// swagger:model
type BcryptHashParam struct {
	Password string `json:"password"` // 明文
	Cost     int    `json:"cost"`     // 哈希成本
}

// swagger:model
type BcryptCompareParam struct {
	HashedPassword string `json:"hashedPassword"` // 密文
	Password       string `json:"password"`       // 明文
}

type BcryptCostParam struct {
	HashedPassword string `json:"hashedPassword"` // 密文
}
