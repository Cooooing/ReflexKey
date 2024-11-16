<template>
  <button class="theme-toggle" @click="toggleTheme">
    <div class="theme-toggle-icon">
      <i
        :class="[
          currentTheme === 'light' ? 'icon-yueliang' : 'icon-taiyang',
          'animate-spin',
        ]"
        class="iconfont"
      />
    </div>
  </button>
</template>

<script setup lang="ts">
import { computed } from "vue";
import { useStore } from "vuex";
import type { Theme } from "@/types/theme";

const store = useStore();
const currentTheme = computed((): Theme => store.getters["theme/currentTheme"]);
const toggleTheme = () => store.dispatch("theme/toggleTheme");
</script>

<style scoped>
.theme-toggle {
  position: relative;
  width: 40px;
  height: 40px;
  border: none;
  border-radius: 12px;
  background: transparent;
  cursor: pointer;
  padding: 0;
  transition: all 0.3s ease;
}

.theme-toggle:hover {
  background: var(--color-hover);
}

.theme-toggle-icon {
  display: flex;
  align-items: center;
  justify-content: center;
  width: 100%;
  height: 100%;
}

.theme-toggle .iconfont {
  font-size: 20px;
  color: var(--color-text);
  transition: none !important;
}

.theme-toggle:hover .iconfont {
  color: var(--color-primary);
  transform: rotate(180deg);
}
</style>
