import { http } from "@/utils/request";
import type { AesEncryptParams, AesDecryptParams } from "./types";
import type { BaseResponse } from "./types";

export const aesEncrypt = (data: AesEncryptParams) => {
  return http.post<string>("/api/crypto/aesEncrypt", data);
};

export const aesDecrypt = (data: AesDecryptParams) => {
  return http.post<string>("/api/crypto/aesDecrypt", data);
};
