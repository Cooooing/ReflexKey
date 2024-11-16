import type { RouteRecordRaw } from "vue-router";

// 定义路由元数据的接口
export interface RouteMeta {
  root?: string;
  moduleTitle?: string;
  subPage?: string;
  parentPath?: string;
  section?: string;
  breadcrumb: string;
}

// 扩展路由记录类型
export type AppRouteRecordRaw = Omit<RouteRecordRaw, "meta" | "children"> & {
  meta: RouteMeta;
  children?: AppRouteRecordRaw[];
};
