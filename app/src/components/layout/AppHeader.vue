<template>
  <header class="app-header">
    <div class="app-header__container">
      <div class="app-header__logo">
        <img src="@/assets/logo.png" alt="Logo" />
      </div>
      <nav class="app-header__nav">
        <!-- 导航内容 -->
      </nav>
      <div class="app-header__actions">
        <ThemeToggle />
      </div>
    </div>
  </header>
</template>

<script setup lang="ts">
import ThemeToggle from "@/components/common/ThemeToggle.vue";
</script>

<style lang="scss" scoped>
@import "@/assets/styles/base/naming";

@include b(header) {
  position: fixed;
  top: 0;
  left: 0;
  right: 0;
  height: 60px;
  background: var(--color-surface);
  border-bottom: 1px solid var(--color-border);
  z-index: 100;

  @include e(container) {
    max-width: 1200px;
    margin: 0 auto;
    height: 100%;
    display: flex;
    align-items: center;
    justify-content: space-between;
    padding: 0 var(--spacing-lg);
  }

  @include e(logo) {
    height: 32px;

    img {
      height: 100%;
    }
  }

  @include e(nav) {
    display: flex;
    gap: var(--spacing-lg);
  }

  @include e(actions) {
    display: flex;
    align-items: center;
    gap: var(--spacing-md);
  }
}
</style>
