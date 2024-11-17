<template>
  <div class="tools-container">
    <!-- 工具页面内容 -->
    <template v-if="$route.name !== 'tools'">
      <keep-alive>
        <router-view></router-view>
      </keep-alive>
    </template>

    <template v-else>
      <!-- 常用工具模块 -->
      <div id="common" class="module-section">
        <div class="section-header">
          <i class="iconfont icon-time" />
          <h2 class="section-title">常用工具</h2>
        </div>
        <div class="section-content">
          <div class="tools-grid">
            <div
              v-for="tool in commonTools"
              :key="tool.id"
              class="tool-card"
              @click="handleToolClick(tool)"
            >
              <div class="tool-icon">
                <i :class="['iconfont', tool.icon]" />
              </div>
              <div class="tool-info">
                <h3>{{ tool.name }}</h3>
                <p>{{ tool.description }}</p>
              </div>
            </div>
          </div>
        </div>
      </div>

      <!-- 开发工具模块 -->
      <div id="dev" class="module-section">
        <div class="section-header">
          <i class="iconfont icon-code" />
          <h2 class="section-title">开发工具</h2>
        </div>
        <div class="section-content">
          <div class="tools-grid">
            <div
              v-for="tool in devTools"
              :key="tool.id"
              class="tool-card"
              @click="handleToolClick(tool)"
            >
              <div class="tool-icon">
                <i :class="['iconfont', tool.icon]" />
              </div>
              <div class="tool-info">
                <h3>{{ tool.name }}</h3>
                <p>{{ tool.description }}</p>
              </div>
            </div>
          </div>
        </div>
      </div>
    </template>
  </div>
</template>

<script lang="ts" setup>
import { useRouter } from "vue-router";

const router = useRouter();

interface Tool {
  id: number;
  name: string;
  description: string;
  icon: string;
  route?: string;
  section?: string;
}

const commonTools: Tool[] = [
  {
    id: 1,
    name: "密码生成器",
    description: "生成安全的随机密码",
    icon: "icon-lock",
    route: "password-generator",
    section: "common",
  },
  {
    id: 2,
    name: "文件加密",
    description: "加密重要文件",
    icon: "icon-file",
    route: "file-encrypt",
    section: "common",
  },
];

const devTools: Tool[] = [
  {
    id: 1,
    name: "编码转换",
    description: "各种编码格式转换",
    icon: "icon-code",
    route: "encode-decode",
    section: "dev",
  },
  {
    id: 2,
    name: "JSON 格式化",
    description: "JSON 格式美化与压缩",
    icon: "icon-file-open",
    route: "json-formatter",
    section: "dev",
  },
  {
    id: 3,
    name: "AES 加解密",
    description: "AES 加密与解密工具",
    icon: "icon-lock",
    route: "AES",
    section: "dev",
  },
];

const handleToolClick = (tool: Tool) => {
  if (tool.route) {
    router.push({
      path: `/tools/${tool.route}`,
      query: { section: tool.section },
    });
  }
};
</script>

<style scoped>
.tools-container {
  padding: 16px;
  display: flex;
  flex-direction: column;
  gap: 16px;
}

.module-section {
  background: var(--color-surface);
  border-radius: 8px;
  padding: 16px;
  box-shadow: 0 1px 2px rgba(0, 0, 0, 0.05);
}

.section-header {
  display: flex;
  align-items: center;
  gap: 6px;
  margin-bottom: 16px;
}

.section-header .iconfont {
  font-size: 16px;
  color: var(--color-primary);
}

.section-title {
  font-size: 14px;
  font-weight: 500;
  color: var(--color-text);
  margin: 0;
}

.tools-grid {
  display: grid;
  grid-template-columns: repeat(auto-fill, minmax(240px, 1fr));
  gap: 16px;
}

.tool-card {
  display: flex;
  align-items: center;
  gap: 12px;
  padding: 12px;
  background: var(--color-background);
  border: 1px solid var(--color-border);
  border-radius: 6px;
  cursor: pointer;
  transition: all 0.2s ease;
}

.tool-card:hover {
  transform: translateY(-2px);
  box-shadow: 0 2px 8px rgba(0, 0, 0, 0.1);
}

.tool-icon {
  display: flex;
  align-items: center;
  justify-content: center;
  width: 32px;
  height: 32px;
  border-radius: 6px;
  background: var(--color-primary);
  color: white;
}

.tool-icon .iconfont {
  font-size: 18px;
}

.tool-info {
  flex: 1;
}

.tool-info h3 {
  font-size: 14px;
  font-weight: 500;
  color: var(--color-text);
  margin: 0 0 4px 0;
}

.tool-info p {
  font-size: 12px;
  color: var(--color-text-secondary);
  margin: 0;
}
</style>
