<template>
  <div class="module-section">
    <div class="section-header">
      <i :class="['iconfont', icon]" />
      <h2 class="section-title">{{ title }}</h2>
    </div>
    <div class="section-content">
      <slot></slot>
    </div>
  </div>
</template>

<script lang="ts" setup>
import { defineProps } from "vue";

defineProps<{
  icon: string;
  title: string;
}>();
</script>

<style scoped>
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
</style>
