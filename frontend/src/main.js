import { createApp } from 'vue'
import App from './App.vue'
import router from './router';
import store from './store';
import "./style.css";
import 'vue-cal/dist/vuecal.css';

createApp(App)
    .use(router)
    .use(store)
    .mount('#app')
