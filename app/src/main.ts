import { createApp } from 'vue'
import App from './App.vue'
import naive from 'naive-ui'

async function setupApp() {
  const app = createApp(App)

  app.use(naive)
  app.mount('#app')
}

void setupApp()
