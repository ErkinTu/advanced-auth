<script setup lang="ts">
import AppNav from './AppNav.vue'
import BaseButton from '@/components/ui/BaseButton.vue'
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
      <router-link to="/login">Login</router-link>
      <BaseButton variant="primary" to="/register">
        Register
      </BaseButton>
    </div>
  </header>
</template>
