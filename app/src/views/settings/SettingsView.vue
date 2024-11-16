<template>
  <div class="settings-container">
    <!-- 外观设置模块 -->
    <div class="module-section">
      <div class="section-header">
        <i class="iconfont icon-picture" />
        <h2 class="section-title">外观设置</h2>
      </div>
      <div class="section-content">
        <!-- 主题设置 -->
        <div class="setting-item">
          <div class="setting-label">
            <span class="label-text">主题模式</span>
            <span class="label-desc">选择明亮或暗黑主题</span>
          </div>
          <div class="setting-control">
            <div class="theme-selector">
              <button
                :class="{ active: currentTheme === 'light' }"
                class="theme-option"
                @click="setTheme('light')"
              >
                <i class="iconfont icon-taiyang" />
                <span>明亮</span>
              </button>
              <button
                :class="{ active: currentTheme === 'dark' }"
                class="theme-option"
                @click="setTheme('dark')"
              >
                <i class="iconfont icon-yueliang" />
                <span>暗黑</span>
              </button>
            </div>
          </div>
        </div>

        <!-- 主题色设置 -->
        <div class="setting-item">
          <div class="setting-label">
            <span class="label-text">主题色</span>
            <span class="label-desc">选择您喜欢的主题色调</span>
          </div>
          <div class="setting-control">
            <div class="color-selector">
              <button
                v-for="color in themeColors"
                :key="color.value"
                :class="{ active: currentThemeColor === color.value }"
                :style="{
                  '--theme-color': color.value,
                  '--theme-color-alpha': getThemeColorAlpha(color.value),
                }"
                class="color-option"
                @click="setThemeColor(color.value)"
              >
                <div class="color-circle" />
                <span class="color-name">{{ color.name }}</span>
              </button>
            </div>
          </div>
        </div>
      </div>
    </div>

    <!-- 通用设置模块 -->
    <div class="module-section">
      <div class="section-header">
        <i class="iconfont icon-setting" />
        <h2 class="section-title">通用设置</h2>
      </div>
      <div class="section-content">
        <div class="setting-item">
          <div class="setting-label">
            <span class="label-text">开机自启动</span>
            <span class="label-desc">应用随系统启动时自动运行</span>
          </div>
          <div class="setting-control">
            <button
              :class="{ active: autoStart }"
              class="switch-button"
              @click="toggleAutoStart"
            >
              <span class="switch-slider" />
            </button>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script lang="ts" setup>
import { computed, ref } from "vue";
import { useStore } from "vuex";
import type { Theme } from "@/types/theme";

const store = useStore();

// 使用 Vuex 状态
const currentTheme = computed(() => store.getters["theme/currentTheme"]);
const currentThemeColor = computed(
  () => store.getters["theme/currentThemeColor"]
);

const autoStart = ref(false);

const themeColors = [
  { name: "天青", value: "#40A9FF" },
  { name: "碧绿", value: "#36CFC9" },
  { name: "青柠", value: "#73D13D" },
  { name: "金橙", value: "#FFA940" },
  { name: "玫红", value: "#FF4D4F" },
  { name: "紫罗", value: "#9254DE" },
];

// 修改主题设置方法
const setTheme = (theme: Theme): void => {
  if (theme !== currentTheme.value) {
    store.dispatch("theme/toggleTheme");
  }
};

const setThemeColor = (color: string): void => {
  store.dispatch("theme/setThemeColor", color);
};

// 修改主题判断逻辑
const getThemeColorAlpha = (color: string): string => {
  return currentTheme.value === "dark"
    ? `color-mix(in srgb, ${color} 30%, transparent)`
    : `color-mix(in srgb, ${color} 15%, transparent)`;
};

const toggleAutoStart = (): void => {
  autoStart.value = !autoStart.value;
};
</script>

<style scoped>
.settings-container {
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

.setting-item {
  display: flex;
  align-items: flex-start;
  justify-content: space-between;
  padding: 12px 0;
  border-bottom: 1px solid var(--color-border);
}

.setting-item:last-child {
  border-bottom: none;
  padding-bottom: 0;
}

.setting-item:first-child {
  padding-top: 0;
}

.setting-label {
  flex: 1;
  padding-right: 16px;
}

.label-text {
  display: block;
  font-size: 14px;
  font-weight: 500;
  color: var(--color-text);
  margin-bottom: 4px;
}

.label-desc {
  display: block;
  font-size: 12px;
  color: var(--color-text-secondary);
}

.setting-control {
  flex: 0 0 auto;
  min-width: 200px;
}

/* 主题选择器样式 */
.theme-selector {
  display: flex;
  gap: 8px;
}

.theme-option {
  display: flex;
  align-items: center;
  gap: 6px;
  padding: 6px 12px;
  border: 1px solid var(--color-border);
  border-radius: 6px;
  background: transparent;
  color: var(--color-text);
  cursor: pointer;
  transition: all 0.2s ease;
}

.theme-option:hover {
  background: var(--color-hover);
}

.theme-option.active {
  background: var(--color-primary);
  color: white;
  border-color: var(--color-primary);
}

/* 颜色选择器样式 */
.color-selector {
  display: grid;
  grid-template-columns: repeat(3, 1fr);
  gap: 8px;
}

.color-option {
  display: flex;
  align-items: center;
  gap: 6px;
  padding: 6px 8px;
  border: 1px solid var(--color-border);
  border-radius: 6px;
  background: transparent;
  cursor: pointer;
  transition: all 0.2s ease;
}

.color-option:hover {
  background: var(--theme-color-alpha);
}

.color-option.active {
  background: var(--theme-color-alpha);
  border-color: var(--theme-color);
}

.color-circle {
  width: 14px;
  height: 14px;
  border-radius: 50%;
  background: var(--theme-color);
  border: 1px solid var(--color-border);
}

.color-name {
  font-size: 12px;
  color: var(--color-text);
}

/* 开关按钮样式 */
.switch-button {
  position: relative;
  width: 40px;
  height: 20px;
  border-radius: 10px;
  background: var(--color-border);
  border: none;
  cursor: pointer;
  padding: 0;
  transition: all 0.3s ease;
}

.switch-button.active {
  background: var(--color-primary);
}

.switch-slider {
  position: absolute;
  top: 2px;
  left: 2px;
  width: 16px;
  height: 16px;
  border-radius: 50%;
  background: white;
  transition: all 0.3s ease;
}

.switch-button.active .switch-slider {
  left: 22px;
}
</style>
