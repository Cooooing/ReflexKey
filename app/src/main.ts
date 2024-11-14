import { createApp } from 'vue'
import App from './App.vue'

import { provideFluentDesignSystem, fluentCard, fluentButton } from '@fluentui/web-components'

provideFluentDesignSystem().register(fluentCard(), fluentButton())

const app = createApp(App)


app.mount('#app')

