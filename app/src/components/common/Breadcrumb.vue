<template>
  <nav class="app-breadcrumb">
    <ol class="app-breadcrumb__list">
      <li
        v-for="(item, index) in items"
        :key="index"
        class="app-breadcrumb__item"
      >
        <router-link
          v-if="item.path"
          :to="item.path"
          class="app-breadcrumb__link"
        >
          {{ item.name }}
        </router-link>
        <span v-else class="app-breadcrumb__text">{{ item.name }}</span>
      </li>
    </ol>
  </nav>
</template>

<script setup lang="ts">
defineProps<{
  items: Array<{
    name: string;
    path?: string;
  }>;
}>();
</script>
