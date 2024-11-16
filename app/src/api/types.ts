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
