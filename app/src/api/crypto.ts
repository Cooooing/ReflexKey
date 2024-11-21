import { http } from "@/utils/request";
import type { AesDecryptParams, AesEncryptParams } from "./types";

export const aesEncrypt = (data: AesEncryptParams) => {
  return http.post<string>("/api/crypto/aesEncrypt", data);
};

export const aesDecrypt = (data: AesDecryptParams) => {
  return http.post<string>("/api/crypto/aesDecrypt", data);
};
