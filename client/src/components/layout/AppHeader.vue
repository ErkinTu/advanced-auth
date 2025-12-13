<script setup lang="ts">
import {useAuth} from '@/composables/useAuth'
import AppNav from './AppNav.vue'
import BaseButton from '@/components/ui/BaseButton.vue'

const {currentUser, isAuthenticated, logout} = useAuth()

const handleLogout = async () => {
  await logout()
}
</script>

<template>
  <header class="flex items-center justify-between px-10 py-4 bg-gray-100/98 fixed w-full top-0 z-10 border-b border-gray-200">
    <div class="flex items-center">
      <router-link to="/" class="text-xl text-cyan-900 font-bold pr-10">
        Auth <span class="text-black">UI</span>
      </router-link>
      <AppNav />
    </div>

    <div class="flex gap-5 font-bold text-sm items-center">
      <template v-if="isAuthenticated && currentUser">
        <span class="text-gray-700 font-normal">{{ currentUser.email }}</span>
        <BaseButton variant="secondary" @click="handleLogout">
          Logout
        </BaseButton>
      </template>

      <template v-else>
        <router-link to="/login">Login</router-link>
        <BaseButton variant="primary" to="/register">
          Register
        </BaseButton>
      </template>
    </div>
  </header>
</template>