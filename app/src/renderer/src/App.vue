<template>
  <div>{{ count }}</div>
</template>

<script setup lang="ts">
import { onMounted, ref } from 'vue'


let count = ref<number>(0)

onMounted(() => {
  setInterval(() => {
    count.value++
    console.log(count.value)
  }, 1000)

})


</script>


