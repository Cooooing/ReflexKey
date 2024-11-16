<template>
  <div class="account-container">
    <!-- 账号管理模块 -->
    <div class="module-section">
      <div class="section-header">
        <i class="iconfont icon-lock" />
        <h2 class="section-title">密码管理</h2>
      </div>
      <div class="section-content">
        <button class="add-button">
          <i class="iconfont icon-add" />
          <span>添加账号</span>
        </button>
        <!-- 账号列表将在这里实现 -->
      </div>
    </div>

    <!-- 安全设置模块 -->
    <div class="module-section">
      <div class="section-header">
        <i class="iconfont icon-setting" />
        <h2 class="section-title">安全设置</h2>
      </div>
      <div class="section-content">
        <!-- 安全设置内容 -->
      </div>
    </div>

    <!-- 可以继续添加其他模块 -->
  </div>
</template>

<script lang="ts" setup>
// 账号管理相关逻辑将在这里实现
</script>

<style scoped>
.account-container {
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

.section-content {
  width: 100%;
}

.add-button {
  display: flex;
  align-items: center;
  gap: 6px;
  padding: 8px 16px;
  border: none;
  border-radius: 6px;
  background: var(--color-primary);
  color: white;
  cursor: pointer;
  transition: all 0.2s ease;
}

.add-button:hover {
  opacity: 0.9;
}

.add-button .iconfont {
  font-size: 14px;
}
</style>
