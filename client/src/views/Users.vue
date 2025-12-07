<script setup lang="ts">
import {useUsers} from '@/composables/useAuth'
import {onMounted} from 'vue'


const {users, loadUsers, isLoading, error} = useUsers()

onMounted(loadUsers)
</script>

<template>
  <div class="p-6">
    <h2 class="text-2xl font-bold text-black">System Users</h2>
    <p class="mb-5">Manage user access and activation status</p>

    <div class="border border-gray-200 rounded-xl shadow p-6 bg-gray-100/50">
      <h3 class="font-semibold text-lg text-black">User Directory</h3>
      <p class="text-gray-600 mb-6">A list of all registered users in the system.</p>

      <div v-if="isLoading" class="text-gray-500">Loading...</div>
      <div v-else-if="error" class="text-red-500">{{ error }}</div>
      <div v-else-if="users.length === 0" class="text-gray-500">No users found</div>

      <div
        class="grid grid-cols-[3fr_1fr_1.5fr_auto] px-4 py-2 text-sm font-semibold text-gray-600 border-b border-gray-200"
      >
        <span>User Email</span>
        <span>Role</span>
        <span>Joined Date</span>
        <span class="text-right">Status</span>
      </div>

      <div
        v-for="u in users"
        :key="u.id"
        class="grid grid-cols-[3fr_1fr_1.5fr_auto] px-4 py-3 items-center border-b border-gray-100 hover:bg-gray-50 transition"
      >
        <span class="font-medium text-black">
        {{ u.email }}
      </span>

        <span class="text-gray-700">
        Admin
      </span>

        <span class="text-gray-700">
        {{ u.created_at }}
      </span>

        <span class="flex justify-end">
        <span
          :class="[
            'px-3 py-1 text-xs rounded-full font-medium inline-flex items-center gap-2',
            u.is_activated
              ? 'bg-emerald-100 text-emerald-700'
              : 'bg-gray-200 text-gray-700'
          ]"
        >
          <svg
            v-if="u.is_activated"
            xmlns="http://www.w3.org/2000/svg"
            class="w-3 h-3"
            fill="none"
            viewBox="0 0 24 24"
            stroke-width="3"
            stroke="currentColor"
          >
            <path stroke-linecap="round" stroke-linejoin="round"
                  d="M4.5 12.75l6 6 9-13.5" />
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
            <path stroke-linecap="round" stroke-linejoin="round"
                  d="M6 18L18 6M6 6l12 12" />
          </svg>

          {{ u.is_activated ? 'Active' : 'Pending' }}
        </span>
      </span>
      </div>
    </div>
  </div>

</template>