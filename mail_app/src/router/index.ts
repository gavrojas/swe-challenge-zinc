import { createRouter, createWebHistory } from 'vue-router'
import RegisterView from '@/views/RegisterView.vue'
import LoginView from '@/views/LoginView.vue'
import HomeView from '../views/HomeView.vue'
import FolderView from '@/views/FolderView.vue';
import { useAuthStore } from '@/stores/auth'

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: [
    {
      path: '/register',
      name: 'register',
      component: RegisterView
    },
    {
      path: '/login',
      name: 'login',
      component: LoginView
    },
    {
      path: '/home',
      name: 'home',
      component: HomeView,
      meta: {
        requiresAuth: true
      }
    },
    {
      path: '/folder/:id',
      name: 'folder',
      component: FolderView ,
      meta: {
        requiresAuth: true
      }
    },
    {
      path: '/',
      redirect: '/login'
    }
  ]
})

router.beforeEach((to, from, next) => {
  const authStore = useAuthStore()
  if (to.meta.requiresAuth && !authStore.isLoggedIn) {
    return next({ name: 'login' });
  }
  if (authStore.isLoggedIn && (to.name === 'login' || to.name === 'register')) {
    return next({ name: 'home' });
  }
  next();
})

export default router
