<template>
  <div class="page-navigation">
    <template v-for="(crumb, index) in items" :key="index">
      <span
        :class="{ clickable: crumb.path && index < items.length - 1 }"
        @click="handleClick(crumb)"
      >
        {{ crumb.text }}
      </span>
      <span v-if="index < items.length - 1" class="separator">/</span>
    </template>
  </div>
</template>

<script setup lang="ts">
import { useRouter } from "vue-router";
import type { BreadcrumbItem } from "@/types/common";

defineProps<{
  items: BreadcrumbItem[];
}>();

const router = useRouter();

const handleClick = (crumb: BreadcrumbItem): void => {
  if (crumb.path) {
    router.push(crumb.path);
  }
};
</script>

<style scoped>
.page-navigation {
  display: flex;
  align-items: center;
  gap: 8px;
  font-size: 14px;
  color: var(--color-text-secondary);
}

.separator {
  color: var(--color-text-secondary);
}

.clickable {
  color: var(--color-primary);
  cursor: pointer;
  transition: opacity 0.2s;
}

.clickable:hover {
  opacity: 0.8;
  text-decoration: underline;
}
</style>
