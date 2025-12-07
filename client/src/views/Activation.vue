<script setup lang="ts">
import {useRoute} from 'vue-router'
import {activateRequest} from '@/api/auth'
import {onMounted, ref} from 'vue'


const route = useRoute()
const message = ref('')


onMounted(async () => {
  const res = await activateRequest(route.params.token as string)
  message.value = res.data.message
})
</script>


<template>
  <div class="page-title p-6 text-center">{{ message }}</div>
</template>