<script setup lang="ts">
import {useUsers} from '@/composables/useAuth'
import {onMounted} from 'vue'


const {users, loadUsers} = useUsers()

onMounted(loadUsers)
</script>

<template>
  <div class="p-4">
    <div v-for="u in users" :key="u.id" class="border p-2 mb-2 rounded">
      <p>{{ u.email }} — {{ u.is_activated ? 'activated' : 'not activated' }}</p>
    </div>
  </div>
</template>