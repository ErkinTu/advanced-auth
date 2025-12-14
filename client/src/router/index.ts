import {createRouter, createWebHistory} from 'vue-router';
import {useAuthStore, useUserStore} from "@/store/auth.ts";


import Layout from '@/layouts/Layout.vue'

const routes = [
  {
    path: '/',
    component: Layout,
    children: [
      {
        path: '',
        name: 'Home',
        component: () => import('../views/Home.vue'),
        meta: { public: true }
      },
      {
        path: 'login',
        name: 'Login',
        component: () => import('../views/Login.vue'),
        meta: { public: true }
      },
      {
        path: 'register',
        name: 'Register',
        component: () => import('../views/Register.vue'),
        meta: { public: true }
      },
      {
        path: 'users',
        name: 'Users',
        component: () => import('../views/Users.vue'),
        meta: { public: true }
      },
      {
        path: 'activation/:token',
        name: 'Activation',
        component: () => import('../views/Activation.vue'),
        meta: { public: true }
      },
      {
        path: 'profile',
        name: 'Profile',
        component: () => import('../views/Profile.vue'),
        meta: { public: true }
      },
      {
        path: 'roles',
        name: 'Roles',
        component: () => import('../views/Roles.vue'),
        meta: { public: true } // requiresAuth: true, role: 'admin'
      }
    ]
  },

  {
    path: '/:pathMatch(.*)*',
    redirect: '/'
  }
]


export const router = createRouter({
  history: createWebHistory(),
  routes,
});

let isSessionInitialized = false

router.beforeEach(async (to, from, next) => {
  const authStore = useAuthStore()

  if (!isSessionInitialized) {
    isSessionInitialized = true
    await authStore.loadCurrentUser()
  }

  if (to.meta.requiresAuth && !authStore.isAuthenticated) {
    return next({ name: 'Login', query: { redirect: to.fullPath } })
  }

  if (authStore.isAuthenticated && (to.name === 'Login' || to.name === 'Register')) {
    return next({ name: 'Profile' })
  }

  if (to.name === 'Users') {
    const userStore = useUserStore()
    if (userStore.users.length === 0) {
      await userStore.loadUsers()
    }
  }

  next()
})