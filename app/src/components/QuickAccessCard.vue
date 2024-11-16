<template>
  <button class="quick-access-card" @click="$emit('click')">
    <i :class="['iconfont', icon]" />
    <span>{{ name }}</span>
  </button>
</template>

<script lang="ts" setup>
import { defineProps } from "vue";

defineProps<{
  icon: string;
  name: string;
}>();
</script>

<style scoped>
.quick-access-card {
  display: flex;
  flex-direction: column;
  align-items: center;
  gap: 8px;
  padding: 16px;
  background: var(--color-background);
  border: 1px solid var(--color-border);
  border-radius: 6px;
  cursor: pointer;
  transition: all 0.2s ease;
}

.quick-access-card:hover {
  transform: translateY(-2px);
  box-shadow: 0 2px 8px rgba(0, 0, 0, 0.1);
}

.quick-access-card .iconfont {
  font-size: 24px;
  color: var(--color-primary);
}

.quick-access-card span {
  font-size: 14px;
  color: var(--color-text);
}
</style>
