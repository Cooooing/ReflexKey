import type {
  AxiosInstance,
  AxiosResponse,
  InternalAxiosRequestConfig,
} from "axios";
import axios from "axios";

// 创建 axios 实例
const http: AxiosInstance = axios.create({
  baseURL: import.meta.env.VITE_API_BASE_URL || "http://127.0.0.1:25566",
  timeout: 15000,
  headers: {
    "Content-Type": "application/json",
  },
});

// 请求拦截器
http.interceptors.request.use(
  (config: InternalAxiosRequestConfig) => {
    // 允许跨域携带cookie
    config.withCredentials = true;
    return config;
  },
  (error) => {
    return Promise.reject(error);
  },
);

// 响应拦截器
http.interceptors.response.use(
  (response: AxiosResponse) => {
    const { data } = response;

    // 如果后端返回的不是标准格式，直接返回数据
    if (typeof data === "string") {
      return { data };
    }

    // 处理标准响应格式
    if (data.code === 200) {
      return data;
    }

    // 处理错误，使用后端返回的 message
    return Promise.reject(new Error(data.msg));
  },
  (error) => {
    // 如果是后端返回的错误，使用后端的错误信息
    const message = error.response?.data?.msg;
    return Promise.reject(new Error(message));
  },
);

export { http };
