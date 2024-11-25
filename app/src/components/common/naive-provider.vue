<template>
  <n-loading-bar-provider>
    <n-dialog-provider>
      <n-notification-provider>
        <n-message-provider>
          <slot></slot>
          <naive-provider-content />
        </n-message-provider>
      </n-notification-provider>
    </n-dialog-provider>
  </n-loading-bar-provider>
</template>

<script setup lang="ts">
import { computed, defineComponent, h } from 'vue'
import {
  type ConfigProviderProps,
  createDiscreteApi,
  useDialog,
  useLoadingBar,
  useMessage,
  useNotification
} from 'naive-ui'
import { useThemeStore } from '@/stores'
import { storeToRefs } from 'pinia'

defineOptions({ name: 'NaiveProvider' })

// 挂载naive组件的方法至window, 以便在路由钩子函数和请求函数里面调用
function registerNaiveTools() {
  const themeStore = useThemeStore()
  const { theme, themeOverrides } = storeToRefs(themeStore)
  const configProviderProps = computed<ConfigProviderProps>(() => {
    return {
      theme: theme.value,
      themeOverrides: themeOverrides.value
    }
  })
  const { message, dialog, notification, loadingBar } = createDiscreteApi(
    ['message', 'dialog', 'notification', 'loadingBar'],
    { configProviderProps }
  )
  window.$loadingBar = loadingBar
  window.$dialog = dialog
  window.$message = message
  window.$notification = notification
}

const NaiveProviderContent = defineComponent({
  name: 'NaiveProviderContent',
  setup() {
    registerNaiveTools()
  },
  render() {
    return h('div')
  }
})
</script>
<style scoped></style>
