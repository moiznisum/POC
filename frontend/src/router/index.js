import { createRouter, createWebHistory } from 'vue-router';
import Home from '../views/Home.vue';
import Booked from '../views/Booked.vue';

const routes = [
    { path: '/', component: Home },
    { path: '/booked', component: Booked },
];

const router = createRouter({
    history: createWebHistory(import.meta.env.BASE_URL),
    routes,
});

export default router;
