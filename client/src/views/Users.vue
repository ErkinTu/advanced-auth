<script setup lang="ts">
import {useUsers} from '@/composables/useAuth'
import {onMounted} from 'vue'


const {users, loadUsers, isLoading, error} = useUsers()

onMounted(loadUsers)
</script>

<template>
  <div class="p-4">
    <div v-if="isLoading" class="text-gray-500">Loading...</div>
    <div v-else-if="error" class="text-red-500">{{ error }}</div>
    <div v-else-if="users.length === 0" class="text-gray-500">No users found</div>
    <div v-else v-for="u in users" :key="u.id" class="border p-2 mb-2 rounded">
      <p>{{ u.email }} — {{ u.is_activated ? 'activated' : 'not activated' }}</p>
    </div>
  </div>
</template>