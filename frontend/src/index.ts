import { ref } from 'vue';
import { createRouter, createWebHistory } from 'vue-router';
import Createaccount from './pages/createaccount.vue';
import Home from './pages/home.vue';
import Login from './pages/login.vue';
import { useAuthStore } from './stores/auth.ts';

export const isRouterReady = ref<boolean>(false);

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: [
    {
      path: '/login',
      name: 'Login',
      component: Login,
      meta: { requiresAuth: false },
    },
    {
      path: '/createAccount',
      name: 'CreateAccount',
      component: Createaccount,
      meta: { requiresAuth: false },
    },
    {
      path: '/home',
      name: 'Home',
      component: Home,
    },
    {
      path: '/',
      redirect: '/home',
    },
    {
      path: '/:pathMatch(.*)*',
      redirect: '/',
    },
  ],
});

router.beforeEach(async (to) => {
  const authStore = useAuthStore();

  if (!authStore.initialized) {
    await authStore.checkAuth();
  }

  isRouterReady.value = true;
  const noAuthRequired = to.meta.requiresAuth === false;

  if (noAuthRequired && authStore.isLoggedIn) {
    return '/home';
  }

  if (!noAuthRequired && !authStore.isLoggedIn) {
    return '/login';
  }
});

export default router;
