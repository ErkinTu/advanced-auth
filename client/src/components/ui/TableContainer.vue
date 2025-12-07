<script setup lang="ts">
interface Props {
  title: string
  description: string
}

defineProps<Props>()
</script>

<template>
  <div class="table-container">
    <h3 class="section-subtitle">{{ title }}</h3>
    <p class="text-secondary mb-6">{{ description }}</p>
    <slot />
  </div>
</template>
