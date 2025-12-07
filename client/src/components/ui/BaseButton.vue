<script setup lang="ts">
import { computed } from 'vue'
import type { RouteLocationRaw } from 'vue-router'

interface Props {
  variant?: 'primary' | 'secondary'
  to?: RouteLocationRaw
  type?: 'button' | 'submit'
}

const props = withDefaults(defineProps<Props>(), {
  variant: 'primary',
  type: 'button'
})

const buttonClass = computed(() => {
  return props.variant === 'primary' ? 'btn-primary' : 'btn-secondary'
})
</script>

<template>
  <router-link v-if="to" :to="to" :class="buttonClass">
    <slot />
  </router-link>
  <button v-else :type="type" :class="buttonClass">
    <slot />
  </button>
</template>
