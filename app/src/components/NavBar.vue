<template>
  <div class="nav-bar">
    <router-link to="/">首页</router-link>
    <router-link to="/tools">工具</router-link>
    <router-link to="/settings">设置</router-link>
  </div>
</template>

<style scoped>
.nav-bar {
  width: 200px; /* 确保导航栏有宽度 */
  background-color: var(--color-surface);
  padding: 16px; /* 添加内边距 */
  box-shadow: 2px 0 5px rgba(0, 0, 0, 0.1); /* 添加阴影效果 */
  height: 100vh; /* 确保导航栏高度为视口高度 */
}

.nav-bar a {
  display: block; /* 确保链接占满整个宽度 */
  padding: 8px 16px; /* 添加内边距 */
  color: var(--color-text); /* 使用主题颜色 */
  text-decoration: none; /* 去掉下划线 */
}

.nav-bar a:hover {
  background-color: var(--color-hover); /* 悬停时的背景颜色 */
}
</style>
