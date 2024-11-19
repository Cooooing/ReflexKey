<template>
  <div class="app-card" @click="$emit('click')">
    <div class="app-card__icon">
      <i :class="['iconfont', icon]" />
    </div>
    <div class="app-card__content">
      <h3 class="app-card__title">{{ name }}</h3>
      <p class="app-card__description">{{ description }}</p>
    </div>
  </div>
</template>

<script lang="ts" setup>
import {  } from "vue";

defineProps<{
  icon: string;
  name: string;
  description?: string;
}>();
</script>
