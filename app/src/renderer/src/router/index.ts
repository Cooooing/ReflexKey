import { createRouter, createWebHistory } from 'vue-router';
import { logPerformanceMetrics } from '../utils/performance';

const routes = [
  // 你的路由配置
];

const router = createRouter({
  history: createWebHistory(),
  routes,
});

router.afterEach(() => {
  logPerformanceMetrics(); // 每次路由跳转后记录性能数据
});

export default router;
