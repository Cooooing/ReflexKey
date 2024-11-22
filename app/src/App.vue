<template>
  <n-config-provider>
    <n-message-provider>
      <n-button @click="click">222</n-button>
      <n-icon size="40" :component="Archive" color="blue" />
    </n-message-provider>
  </n-config-provider>
</template>
<script setup lang="ts">
import { Archive } from '@vicons/ionicons5'
import { useMessage } from 'naive-ui'
import { onMounted } from 'vue'

const messageRef = ref(null)

const click = () => {
  if (messageRef.value) {
    messageRef.value.info('点击了')
  }
}

// 在 onMounted 中初始化 message
onMounted(() => {
  messageRef.value = useMessage()
})
</script>

<style></style>
