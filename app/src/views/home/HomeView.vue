<template>
  <div class="page-container">
    <!-- 概览模块 -->
    <div class="module-section">
      <div class="section-header">
        <i class="iconfont icon-home" />
        <h2 class="section-title">概览</h2>
      </div>
      <div class="section-content">
        <div class="overview-grid">
          <OverviewCard
            v-for="item in overviewItems"
            :key="item.title"
            :icon="item.icon"
            :title="item.title"
            :value="item.value"
          />
        </div>
      </div>
    </div>

    <!-- 最近使用模块 -->
    <div class="module-section">
      <div class="section-header">
        <i class="iconfont icon-history" />
        <h2 class="section-title">最近使用</h2>
      </div>
      <div class="section-content">
        <div class="recent-list">
          <!-- 最近使用的内容 -->
        </div>
      </div>
    </div>

    <!-- 快速入口模块 -->
    <div class="module-section">
      <div class="section-header">
        <i class="iconfont icon-lightning" />
        <h2 class="section-title">快速入口</h2>
      </div>
      <div class="section-content">
        <div class="quick-access-grid">
          <QuickAccessCard
            v-for="item in quickAccessItems"
            :key="item.name"
            :icon="item.icon"
            :name="item.name"
            :description="item.description"
            @click="item.onClick"
          />
        </div>
      </div>
    </div>
  </div>
</template>

<script lang="ts" setup>
import { ref } from "vue";
import { useRouter } from "vue-router";
import type { OverviewItem } from "@/types";
import OverviewCard from "@/components/OverviewCard.vue";
import QuickAccessCard from "@/components/QuickAccessCard.vue";

const router = useRouter();

const overviewItems = ref<OverviewItem[]>([
  { title: "账号总数", value: 12, icon: "icon-user" },
  { title: "最近使用", value: "2小时前", icon: "icon-time" },
]);

const quickAccessItems = [
  {
    name: "添加账号",
    icon: "icon-add",
    description: "添加新的账号信息",
    onClick: () => router.push({ name: "account" }),
  },
  {
    name: "常用工具",
    icon: "icon-tool",
    description: "查看所有可用工具",
    onClick: () => router.push({ name: "tools" }),
  },
];
</script>

<style scoped>
.page-container {
  padding: 16px;
  display: flex;
  flex-direction: column;
  gap: 16px;
}

.module-section {
  background: var(--color-surface);
  border-radius: 8px;
  padding: 16px;
  box-shadow: 0 1px 2px rgba(0, 0, 0, 0.05);
}

.section-header {
  display: flex;
  align-items: center;
  gap: 6px;
  margin-bottom: 16px;
}

.section-header .iconfont {
  font-size: 16px;
  color: var(--color-primary);
}

.section-title {
  font-size: 14px;
  font-weight: 500;
  color: var(--color-text);
  margin: 0;
}

.overview-grid {
  display: grid;
  grid-template-columns: repeat(auto-fit, minmax(240px, 1fr));
  gap: 16px;
}

.quick-access-grid {
  display: grid;
  grid-template-columns: repeat(auto-fill, minmax(240px, 1fr));
  gap: 16px;
}

.recent-list {
  display: grid;
  gap: 16px;
}
</style>
