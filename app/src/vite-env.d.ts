/// <reference types="vite/client" />

// 声明环境变量的类型
import { DialogApi, LoadingBarApi, MessageApi, NotificationApi } from 'naive-ui'

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

declare namespace NTheme {
  /** 颜色类型 */
  type ColorType = 'primary' | 'info' | 'success' | 'warning' | 'error'
  /** 颜色类型大写值 */
  type ColorTypeCase = 'Primary' | 'Info' | 'Success' | 'Warning' | 'Error'
  /** 颜色场景 */
  type ColorScene = '' | 'suppl' | 'hover' | 'pressed'
  /** 颜色场景大写值 */
  type ColorSceneCase = '' | 'Suppl' | 'Hover' | 'Pressed'
  /** 按钮颜色场景 */
  type ButtonColorScene = '' | 'hover' | 'pressed' | 'focus' | 'disabled'
  /** 按钮颜色场景大写值 */
  type ButtonColorSceneCase = '' | 'Hover' | 'Pressed' | 'Focus' | 'Disabled'
  // 主题配置
  type Config = {
    [key in NTheme.ColorType]: string
  }
}

interface Window {
  $message: MessageApi
  $dialog: DialogApi
  $notification: NotificationApi
  $loadingBar: LoadingBarApi
}
