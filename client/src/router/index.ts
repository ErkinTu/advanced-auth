import {createRouter, createWebHistory} from 'vue-router';


const routes = [
  {
    path: '/',
    name: 'Home',
    component: () => import('../views/Home.vue'),
    meta: {public: true}
  },
  {
    path: '/register',
    name: 'Register',
    component: () => import('../views/Register.vue'),
    meta: {public: true} // + guestOnly: true
  },
  {
    path: '/login',
    name: 'Login',
    component: () => import('../views/Login.vue'),
    meta: {public: true} // + guestOnly: true
  },
  {
    path: '/activate/:token',
    name: 'Activation',
    component: () => import('../views/Activation.vue'),
    meta: {public: true}
  },
  {
    path: '/users',
    name: 'Users',
    component: () => import('../views/Users.vue'),
    meta: {public: true} // adminOnly: true
  },
  {
    path: '/:pathMatch(.*)*',
    redirect: '/'
  }
];

export const router = createRouter({
  history: createWebHistory(),
  routes,
});