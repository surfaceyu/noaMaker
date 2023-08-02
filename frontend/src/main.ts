import { createApp } from 'vue'
import App from './views/App.vue'
import { router } from './router'
import ElementPlus from 'element-plus'
import 'element-plus/dist/index.css'
import '@wangeditor/editor/dist/css/style.css'

const app = createApp(App).use(ElementPlus).use(router)
app.mount('#app')
