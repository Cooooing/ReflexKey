<template>
  <div class="header">
    <div class="header-left">
      <Breadcrumb :items="breadcrumbs" />
    </div>
    <div class="header-right">
      <ThemeToggle />
    </div>
  </div>
</template>

<script setup lang="ts">
import { computed } from "vue";
import { useRoute } from "vue-router";
import Breadcrumb from "@/components/common/Breadcrumb.vue";
import ThemeToggle from "@/components/common/ThemeToggle.vue";
import type { BreadcrumbItem } from "@/types/common";
import { NAV_ITEMS } from "@/constants/nav";

const route = useRoute();

const breadcrumbs = computed((): BreadcrumbItem[] => {
  const crumbs: BreadcrumbItem[] = [];

  // 获取当前路由的根级别名称
  const rootName = route.meta.root || route.name;

  // 找到对应的导航项
  const currentItem = NAV_ITEMS.find((item) => item.name === rootName);
  if (currentItem) {
    crumbs.push({
      text: currentItem.text,
      path: rootName === "home" ? "/" : `/${rootName}`,
    });
  }

  // 如果有模块标题，添加模块层级
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

<style scoped>
.header {
  height: 48px;
  display: flex;
  align-items: center;
  justify-content: space-between;
  padding: 0 24px;
  background-color: var(--color-surface);
  border-bottom: 1px solid var(--color-border);
}

.header-left,
.header-right {
  display: flex;
  align-items: center;
  gap: 8px;
}
</style>
