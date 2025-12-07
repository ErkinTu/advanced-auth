<script setup lang="ts">
interface Props {
  text: string
}

defineProps<Props>()
</script>

<template>
  <p class="badge">{{ text }}</p>
</template>
