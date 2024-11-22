/// <reference types="vite/client" />

// 声明环境变量的类型
interface ImportMetaEnv {
  VITE_APP_TITLE: string
  VITE_APP_API_BASE_URL: string
  // 更多环境变量...
}

// 声明静态资源模块
declare module '*.png' {
  const value: string
  export default value
}

declare module '*.svg' {
  const value: string
  export default value
}

// 声明 Vue 组件
declare module '*.vue' {
  import type { DefineComponent } from 'vue'
  const component: DefineComponent<object, object, unknown>
  export default component
}
