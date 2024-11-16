import type { RouteRecordRaw } from "vue-router";
import { createRouter, createWebHashHistory } from "vue-router";

const routes: RouteRecordRaw[] = [
  {
    path: "/",
    name: "home",
    component: () => import("../views/home/HomeView.vue"),
    meta: {
      breadcrumb: "主页",
    },
  },
  {
    path: "/account",
    name: "account",
    component: () => import("../views/account/AccountView.vue"),
    meta: {
      breadcrumb: "账号",
    },
  },
  {
    path: "/account/:id",
    name: "account-detail",
    component: () => import("../views/account/AccountDetailView.vue"),
    meta: {
      breadcrumb: "账号",
      subPage: "详情",
      parentPath: "/account",
    },
  },
  {
    path: "/tools",
    name: "tools",
    component: () => import("../views/tools/ToolsView.vue"),
    meta: {
      breadcrumb: "工具",
      root: "tools",
    },
    children: [
      {
        path: "json-formatter",
        name: "json-formatter",
        component: () => import("../views/tools/JsonFormatter.vue"),
        meta: {
          root: "tools",
          breadcrumb: "工具",
          moduleTitle: "开发工具",
          subPage: "JSON 格式化",
          parentPath: "/tools",
          section: "dev",
        },
      },
    ],
  },
  {
    path: "/settings",
    name: "settings",
    component: () => import("../views/settings/SettingsView.vue"),
    meta: {
      breadcrumb: "设置",
    },
  },
];

const router = createRouter({
  history: createWebHashHistory(),
  routes,
});

export default router;
