import { createApp } from 'vue'
import App from './App.vue'
import router from './router'
import store from './store'
import ElementPlus from 'element-plus'
import 'element-plus/dist/index.css'
import axios from "axios";
import auth from './auth/auth'
const app = createApp(App);
// 提供全局对象
app.provide('axios', axios);
app.provide('auth', auth);
app.use(ElementPlus).use(store).use(router).mount('#app')
