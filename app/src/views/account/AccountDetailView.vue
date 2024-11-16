<template>
  <div class="account-detail-container">
    <div class="account-detail-header">
      <h2>账号详情</h2>
    </div>
    <div class="account-detail-content">
      <!-- 账号详情内容将在这里实现 -->
    </div>
  </div>
</template>

<script lang="ts" setup>
// 账号详情相关逻辑将在这里实现
</script>

<style scoped>
.account-detail-container {
  padding: 20px;
}

.account-detail-header {
  margin-bottom: 24px;
}

.account-detail-header h2 {
  font-size: 20px;
  font-weight: 500;
  color: var(--color-text);
  margin: 0;
}

.account-detail-content {
  background: var(--color-surface);
  border-radius: 8px;
  padding: 20px;
  box-shadow: 0 1px 2px rgba(0, 0, 0, 0.05);
}
</style>
