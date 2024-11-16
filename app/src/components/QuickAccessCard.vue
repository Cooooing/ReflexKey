<template>
  <div class="tool-card" @click="$emit('click')">
    <div class="tool-icon">
      <i :class="['iconfont', icon]" />
    </div>
    <div class="tool-info">
      <h3>{{ name }}</h3>
      <p>{{ description }}</p>
    </div>
  </div>
</template>

<script lang="ts" setup>
import { defineProps } from "vue";

defineProps<{
  icon: string;
  name: string;
  description?: string;
}>();
</script>

<style scoped>
.tool-card {
  display: flex;
  align-items: center;
  gap: 12px;
  padding: 12px;
  background: var(--color-background);
  border: 1px solid var(--color-border);
  border-radius: 6px;
  cursor: pointer;
  transition: all 0.2s ease;
}

.tool-card:hover {
  transform: translateY(-2px);
  box-shadow: 0 2px 8px rgba(0, 0, 0, 0.1);
}

.tool-icon {
  display: flex;
  align-items: center;
  justify-content: center;
  width: 32px;
  height: 32px;
  border-radius: 6px;
  background: var(--color-primary);
  color: white;
}

.tool-icon .iconfont {
  font-size: 18px;
}

.tool-info {
  flex: 1;
}

.tool-info h3 {
  font-size: 14px;
  font-weight: 500;
  color: var(--color-text);
  margin: 0 0 4px 0;
}

.tool-info p {
  font-size: 12px;
  color: var(--color-text-secondary);
  margin: 0;
}
</style>
