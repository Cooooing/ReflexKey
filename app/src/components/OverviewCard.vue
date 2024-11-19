<template>
  <div class="overview-card">
    <div class="overview-icon">
      <i :class="['iconfont', icon]" />
    </div>
    <div class="overview-info">
      <h3>{{ title }}</h3>
      <p class="overview-value">{{ value }}</p>
    </div>
  </div>
</template>

<script lang="ts" setup>

defineProps<{
  title: string;
  value: string | number;
  icon: string;
}>();
</script>

<style scoped>
.overview-card {
  display: flex;
  align-items: center;
  gap: 12px;
  padding: 16px;
  background: var(--color-background);
  border: 1px solid var(--color-border);
  border-radius: 6px;
}

.overview-icon {
  display: flex;
  align-items: center;
  justify-content: center;
  width: 40px;
  height: 40px;
  border-radius: 8px;
  background: var(--color-primary);
  color: white;
}

.overview-icon .iconfont {
  font-size: 20px;
}

.overview-info h3 {
  font-size: 14px;
  color: var(--color-text-secondary);
  margin: 0 0 4px 0;
}

.overview-value {
  font-size: 20px;
  font-weight: 500;
  color: var(--color-text);
  margin: 0;
}
</style>
