<template>
  <div class="app-container">
    <!-- 左侧导航栏 -->
    <div class="sidebar">
      <div class="logo">
        <img alt="Logo" src="./assets/logo.png" />
      </div>
      <nav class="nav-menu">
        <button
          v-for="item in navItems"
          :key="item.name"
          :class="['nav-item', { active: currentRoute === item.name }]"
          @click="navigateTo(item.name)"
        >
          <div class="nav-content">
            <span class="nav-icon">
              <i :class="['iconfont', item.icon]" />
            </span>
            <span class="nav-text">{{ item.text }}</span>
          </div>
        </button>
      </nav>
      <div class="sidebar-footer">
        <button
          :class="['nav-item', { active: currentRoute === 'settings' }]"
          @click="navigateTo('settings')"
        >
          <div class="nav-content">
            <span class="nav-icon">
              <i class="iconfont icon-shezhi" />
            </span>
            <span class="nav-text">设置</span>
          </div>
        </button>
      </div>
    </div>

    <!-- 主内容区 -->
    <main class="main-content">
      <div class="header">
        <div class="header-left">
          <div class="page-navigation">
            <template v-for="(crumb, index) in breadcrumbs" :key="index">
              <span
                :class="{
                  clickable: crumb.path && index < breadcrumbs.length - 1,
                }"
                @click="crumb.path && router.push(crumb.path)"
              >
                {{ crumb.text }}
              </span>
              <span v-if="index < breadcrumbs.length - 1" class="separator"
                >/</span
              >
            </template>
          </div>
        </div>
      </div>
      <div class="content">
        <keep-alive>
          <router-view></router-view>
        </keep-alive>
      </div>
    </main>
  </div>
</template>

<script lang="ts" setup>
import { computed, watch } from "vue";
import { useRoute, useRouter } from "vue-router";
import { useStore } from "vuex";
import type { Theme } from "@/types/theme";
import "@/assets/icons/iconfont.css";
import { NAV_ITEMS, SETTINGS_NAV } from "@/constants/nav";

const store = useStore();
const router = useRouter();
const route = useRoute();

const navItems = NAV_ITEMS;

const currentRoute = computed(() => route.meta.root || route.name);

const navigateTo = (route: string): void => {
  if (route && router.hasRoute(route)) {
    router.push({ name: route }).catch((err) => {
      console.error("Navigation failed:", err);
    });
  } else {
    console.warn(`Route "${route}" does not exist`);
  }
};

// 使用 Vuex 状态
const currentTheme = computed((): Theme => store.getters["theme/currentTheme"]);

// 初始化主题
store.dispatch("theme/applyTheme");

// 监听主题变化
watch(
  currentTheme,
  (newTheme) => {
    console.log("主题已新为：", newTheme);
  },
  { immediate: true }
);

// 计算面包屑
const breadcrumbs = computed(() => {
  const crumbs = [];

  // 获取当前路由的根级别名称
  const rootName = route.meta.root || route.name;

  // 找到对应的导航项
  const currentItem = [...navItems, SETTINGS_NAV].find(
    (item) => item.name === rootName
  );
  if (currentItem) {
    crumbs.push({
      text: currentItem.text,
      path: rootName === "home" ? "/" : `/${rootName}`,
    });
  }

  // 如果有模块题，添加模块层级
  if (route.meta.moduleTitle) {
    crumbs.push({
      text: route.meta.moduleTitle as string,
      path: `${route.meta.parentPath as string}#${
        route.meta.section as string
      }`,
    });
  }

  // 如果是子页面，添加当前页
  if (route.meta.subPage) {
    crumbs.push({
      text: route.meta.subPage as string,
      path: "",
    });
  }

  return crumbs;
});
</script>

<style>
/* 全局主题变量 */
:root {
  /* 主题变量会被 themeManager 动态更新 */
  --color-primary: "";
  --color-secondary: "";
  --color-background: "";
  --color-surface: "";
  --color-text: "";
  --color-text-secondary: "";
  --color-border: "";
  --color-error: "";
  --color-success: "";
  --color-warning: "";
  --color-hover: "";
  --color-focus: "";
  --color-disabled: "";
  --color-backdrop: "";
  --color-divider: "";
  --color-overlay: "";
  --border-radius: "";
  --box-shadow: "";
  --transition: "";
}

/* 只对背景和边框应用过渡效果 */
* {
  transition: background-color 0.2s ease, border-color 0.2s ease;
}

/* 确保背景色应用到整个应用 */
html,
body {
  margin: 0;
  padding: 0;
  height: 100%;
  background-color: var(--color-background);
  color: var(--color-text);
}

/* 布局相关 */
.app-container {
  display: flex;
  height: 100vh;
  background-color: var(--color-background);
}

/* 侧边栏样式 */
.sidebar {
  width: 68px;
  height: 100%;
  display: flex;
  flex-direction: column;
  background-color: var(--color-surface);
  border-right: 1px solid var(--color-border);
  transition: all 0.3s cubic-bezier(0.4, 0, 0.2, 1);
}

.sidebar:hover {
  width: 200px;
}

/* Logo 样式 */
.logo {
  height: 48px;
  display: flex;
  align-items: center;
  justify-content: center;
  border-bottom: 1px solid #e8e8e8;
}

.logo img {
  width: 24px;
  height: 24px;
}

/* 导航菜单样式 */
.nav-menu {
  flex: 1;
  display: flex;
  flex-direction: column;
  padding: 8px 0;
}

/* 导航项样式 */
.nav-item {
  height: 40px;
  width: calc(100% - 16px);
  display: flex;
  align-items: center;
  justify-content: center;
  padding: 0;
  margin: 4px 8px;
  border-radius: 6px;
  cursor: pointer;
  color: var(--color-textSecondary);
  background: transparent;
  border: none;
  transition: all 0.2s;
  position: relative;
}

/* 导航项左侧标记 */
.nav-item::before {
  content: "";
  position: absolute;
  left: 0;
  top: 8px;
  width: 3px;
  height: 24px;
  background-color: var(--color-primary);
  border-radius: 0 2px 2px 0;
  opacity: 0;
  transform: scaleY(0);
  transition: all 0.2s ease-in-out;
  transform-origin: center;
}

.nav-item.active::before {
  opacity: 1;
  transform: scaleY(1);
}

/* 导航项内容布局 */
.nav-content {
  width: 100%;
  height: 100%;
  display: flex;
  align-items: center;
  justify-content: center;
  padding: 0;
  transition: all 0.3s cubic-bezier(0.4, 0, 0.2, 1);
  position: relative;
}

/* 图标样式 */
.nav-icon {
  width: 20px;
  height: 20px;
  display: flex;
  align-items: center;
  justify-content: center;
  font-size: 16px;
  position: relative;
  transform: translateX(0);
  transition: none !important;
}

/* 文字样式 */
.nav-text {
  opacity: 0;
  max-width: 0;
  width: auto;
  margin-left: 0;
  font-size: 14px;
  white-space: nowrap;
  transition: max-width 0.3s cubic-bezier(0.4, 0, 0.2, 1),
    opacity 0.2s cubic-bezier(0.4, 0, 0.2, 1),
    margin-left 0.3s cubic-bezier(0.4, 0, 0.2, 1);
  overflow: hidden;
}

/* 悬停和激活态 */
.nav-item:hover {
  background-color: var(--color-hover);
  color: var(--color-text);
}

.nav-item.active {
  background-color: var(--color-focus);
  color: var(--color-primary);
}

/* 侧边栏展开状态 */
.sidebar:hover .nav-content {
  justify-content: center;
}

.sidebar:hover .nav-text {
  opacity: 1;
  max-width: 200px;
  margin-left: 8px;
}

/* 底部设置按钮 */
.sidebar-footer {
  padding: 8px 0;
  border-top: 1px solid var(--color-border);
}

.sidebar-footer .nav-item {
  width: calc(100% - 16px);
}

.sidebar-footer .nav-item:hover {
  background-color: var(--color-hover);
  color: var(--color-text);
}

.sidebar-footer .nav-item.active {
  background-color: var(--color-focus);
  color: var(--color-primary);
}

/* 主内容区域 */
.main-content {
  flex: 1;
  display: flex;
  flex-direction: column;
  background-color: var(--color-background);
}

.header {
  height: 48px;
  display: flex;
  align-items: center;
  justify-content: space-between;
  padding: 0 24px;
  background-color: var(--color-surface);
  border-bottom: 1px solid var(--color-border);
}

.header-left {
  display: flex;
  align-items: center;
  gap: 8px;
}

.header h2 {
  margin: 0;
  font-size: 16px;
  font-weight: 500;
  color: var(--color-text);
}

.content {
  flex: 1;
  padding: 24px;
  overflow-y: auto;
  background-color: var(--color-background);
}

/* 滚动条样式 */
::-webkit-scrollbar {
  width: 6px;
  height: 6px;
}

::-webkit-scrollbar-track {
  background: var(--color-surface);
}

::-webkit-scrollbar-thumb {
  background: var(--color-border);
  border-radius: var(--border-radius);
}

::-webkit-scrollbar-thumb:hover {
  background: var(--color-textSecondary);
}

/* 图标相关样式 */
.iconfont {
  font-family: "iconfont" !important;
  font-size: 20px;
  font-style: normal;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  transition: none !important;
}

.nav-icon .iconfont {
  font-size: 16px;
  margin-right: 0;
  display: flex;
  align-items: center;
  justify-content: center;
  transition: none !important;
}

.action-buttons .iconfont {
  font-size: 24px;
  margin: 0 12px;
  cursor: pointer;
}

.action-buttons .iconfont:hover {
  color: #0078d4;
}

/* 更新页面导航样式 */
.page-navigation {
  display: flex;
  align-items: center;
  gap: 8px;
  font-size: 14px;
  color: var(--color-textSecondary);
}

.page-navigation .separator {
  color: var(--color-textSecondary);
  margin: 0;
}

.page-navigation .clickable {
  color: var(--color-primary);
  cursor: pointer;
  transition: all 0.2s;
}

.page-navigation .clickable:hover {
  opacity: 0.8;
  text-decoration: underline;
}

/* 删除旧的标题相关样式 */
.page-navigation h2 {
  display: none;
}

/* 删除旧的包屑样式 */
.breadcrumb {
  display: none;
}

/* 主题切换按钮样式 */
.header-right {
  display: flex;
  align-items: center;
}

.theme-toggle {
  position: relative;
  width: 40px;
  height: 40px;
  border: none;
  border-radius: 12px;
  background: transparent;
  cursor: pointer;
  padding: 0;
  transition: all 0.3s ease;
}

.theme-toggle:hover {
  background: var(--color-hover);
}

.theme-toggle-icon {
  display: flex;
  align-items: center;
  justify-content: center;
  width: 100%;
  height: 100%;
}

.theme-toggle .iconfont {
  font-size: 20px;
  color: var(--color-text);
  transition: none !important;
}

.theme-toggle:hover .iconfont {
  color: var(--color-primary);
  transform: rotate(180deg);
}

.animate-spin {
  transition: none !important;
}

/* 删除旧的主题切换按钮样式 */
.header-right fluent-button {
  display: none;
}
</style>
