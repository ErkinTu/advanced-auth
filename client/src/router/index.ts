import {createRouter, createWebHistory} from 'vue-router';
import {useUserStore} from "@/store/auth.ts";


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
      },
      {
        path: 'profile',
        name: 'Profile',
        component: () => import('../views/Profile.vue')
      }
    ]
  },

  // 404 fallback
  {
    path: '/:pathMatch(.*)*',
    redirect: '/'
  }
]


export const router = createRouter({
  history: createWebHistory(),
  routes,
});

router.beforeEach(async (to, from, next) => {
  if (to.name === 'Users') {
    const store = useUserStore()
    if (store.users.length === 0) {
      await store.loadUsers()
    }
  }
  next()
})