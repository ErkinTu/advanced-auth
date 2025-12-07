<script setup lang="ts">
import { computed } from 'vue'

interface Props {
  active: boolean
  label?: string
}

const props = withDefaults(defineProps<Props>(), {
  label: ''
})

const badgeClass = computed(() => {
  return ['badge-status', props.active ? 'badge-status-active' : 'badge-status-pending']
})

const displayLabel = computed(() => {
  return props.label || (props.active ? 'Active' : 'Pending')
})
</script>

<template>
  <span :class="badgeClass">
    <svg
      v-if="active"
      xmlns="http://www.w3.org/2000/svg"
      class="w-3 h-3"
      fill="none"
      viewBox="0 0 24 24"
      stroke-width="3"
      stroke="currentColor"
    >
      <path stroke-linecap="round" stroke-linejoin="round" d="M4.5 12.75l6 6 9-13.5" />
    </svg>

    <svg
      v-else
      xmlns="http://www.w3.org/2000/svg"
      class="w-3 h-3"
      fill="none"
      viewBox="0 0 24 24"
      stroke-width="3"
      stroke="currentColor"
    >
      <path stroke-linecap="round" stroke-linejoin="round" d="M6 18L18 6M6 6l12 12" />
    </svg>

    {{ displayLabel }}
  </span>
</template>
