<script setup lang="ts">
import AppHeader from '@/components/layout/AppHeader.vue'
import AppFooter from '@/components/layout/AppFooter.vue'
</script>

<template>
  <div class="min-h-screen flex flex-col text-gray-600">

    <AppHeader />

    <main class="flex-1 p-4 mt-15">
      <RouterView v-slot="{ Component }">
        <transition
          mode="out-in"
          enter-active-class="animate-bounce-in-up"
          leave-active-class="animate-fade-out"
        >
          <component :is="Component" />
        </transition>
      </RouterView>
    </main>

    <AppFooter />

  </div>
</template>