<template>
  <button class="app-button app-button--icon" @click="toggleTheme">
    <i
      :class="[
        currentTheme === 'light' ? 'icon-yueliang' : 'icon-taiyang',
        'iconfont',
      ]"
    />
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
