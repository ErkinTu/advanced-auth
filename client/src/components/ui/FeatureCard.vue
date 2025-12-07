<script setup lang="ts">
interface Props {
  icon: string
  title: string
  description: string
}

defineProps<Props>()
</script>

<template>
  <div class="card-feature">
    <div class="card-feature-icon">{{ icon }}</div>
    <h3 class="section-subtitle">{{ title }}</h3>
    <p>{{ description }}</p>
  </div>
</template>
