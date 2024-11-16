import axios, { AxiosRequestConfig } from "axios";
import type { CommonResponse } from "@/types";

const request = axios.create({
  baseURL: import.meta.env.VITE_API_BASE_URL,
  timeout: 10000,
});

request.interceptors.response.use(
  (response) => {
    const data = response.data;
    return {
      ...data,
      success: data.code === 0 || data.code === 200,
    };
  },
  (error) => {
    console.error("请求错误:", error);
    return Promise.reject(error);
  }
);

export const http = {
  async get<T>(
    url: string,
    config?: AxiosRequestConfig
  ): Promise<CommonResponse<T>> {
    return request.get(url, config);
  },

  async post<T>(
    url: string,
    data?: any,
    config?: AxiosRequestConfig
  ): Promise<CommonResponse<T>> {
    return request.post(url, data, config);
  },
};

export default http;
