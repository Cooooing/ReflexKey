import { createApp } from 'vue'
import App from './App.vue'
import { setupStore } from './stores'
async function setupApp() {
  const app = createApp(App)
  setupStore(app)
  app.mount('#app')
}

void setupApp()
