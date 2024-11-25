<template>
  <n-config-provider :theme="theme" :theme-overrides="themeOverrides">
    <naive-provider>
      <router-view />
    </naive-provider>
  </n-config-provider>
</template>

<script lang="ts" setup>
import NaiveProvider from '@/components/common/naive-provider.vue'
import { useThemeStore } from '@/stores'
import { storeToRefs } from 'pinia'
import { computed } from 'vue'

const themeStore = useThemeStore()
const { theme, themeConfig, themeOverrides, themeColors, modeState } =
  storeToRefs(themeStore)
const colorTypes = computed(
  () => Object.keys(themeConfig.value) as NTheme.ColorType[]
)
</script>

<style></style>
