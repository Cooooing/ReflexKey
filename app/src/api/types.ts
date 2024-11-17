// 分页请求参数接口
export interface PaginationParams {
  page: number;
  pageSize: number;

  [key: string]: any;
}

// 分页响应数据接口
export interface PaginationResponse<T> {
  list: T[];
  total: number;
  page: number;
  pageSize: number;
}

// 基础响应接口
export interface BaseResponse<T> {
  code: number;
  data: T;
  message: string;
  success: boolean;
}

// 加密参数接口
export interface AesEncryptParams {
  plainText: string; // 明文
  plainTextFormat: string; // 明文格式
  operationMode: string; // 加密模式
  fill: string; // 填充模式
  key: string; // 密钥
  keyFormat: string; // 密钥格式
  deviation: string; // 偏移量
  deviationFormat: string; // 偏移量格式
  cipherTextFormat: string; // 密文输出格式
}

export interface AesDecryptParams {
  cipherText: string; // 密文
  cipherTextFormat: string; // 密文格式
  operationMode: string; // 解密模式
  fill: string; // 填充模式
  key: string; // 密钥
  keyFormat: string; // 密钥格式
  deviation: string; // 偏移量
  deviationFormat: string; // 偏移量格式
  plainTextFormat: string; // 明文格式
}
