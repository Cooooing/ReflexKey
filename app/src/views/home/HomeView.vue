<template>
  <div class="page-container">
    <ModuleSection icon="icon-home" title="概览">
      <div class="overview-grid">
        <OverviewCard
          v-for="item in overviewItems"
          :key="item.title"
          :icon="item.icon"
          :title="item.title"
          :value="item.value"
        />
      </div>
    </ModuleSection>

    <ModuleSection icon="icon-history" title="最近使用">
      <div class="recent-list">
        <!-- 最近使用的内容 -->
      </div>
    </ModuleSection>

    <ModuleSection icon="icon-lightning" title="快速入口">
      <div class="quick-access-grid">
        <QuickAccessCard
          v-for="item in quickAccessItems"
          :key="item.name"
          :icon="item.icon"
          :name="item.name"
          @click="item.onClick"
        />
      </div>
    </ModuleSection>
  </div>
</template>

<script lang="ts" setup>
import { ref } from "vue";
import { useRouter } from "vue-router";
import type { OverviewItem } from "@/types";
import ModuleSection from "@/components/ModuleSection.vue";
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
    onClick: () => router.push({ name: "account" }),
  },
  {
    name: "常用工具",
    icon: "icon-tool",
    onClick: () => router.push({ name: "tools" }),
  },
];
</script>

<style scoped>
.overview-grid {
  display: grid;
  grid-template-columns: repeat(auto-fit, minmax(200px, 1fr));
  gap: 16px;
}

.quick-access-grid {
  display: grid;
  grid-template-columns: repeat(auto-fill, minmax(120px, 1fr));
  gap: 16px;
}
</style>
