<script setup lang="ts">
import {useUsers} from '@/composables/useAuth'
import {onMounted} from 'vue'


const {users, loadUsers, isLoading, error} = useUsers()

onMounted(loadUsers)
</script>

<template>
  <div class="p-6">
    <h2 class="page-title">System Users</h2>
    <p class="mb-5">Manage user access and activation status</p>

    <div class="table-container">
      <h3 class="section-subtitle">User Directory</h3>
      <p class="text-secondary mb-6">A list of all registered users in the system.</p>

      <div v-if="isLoading" class="text-muted">Loading...</div>
      <div v-else-if="error" class="text-error">{{ error }}</div>
      <div v-else-if="users.length === 0" class="text-muted">No users found</div>

      <div class="table-header">
        <span>User Email</span>
        <span>Role</span>
        <span>Joined Date</span>
        <span class="text-right">Status</span>
      </div>

      <div
        v-for="u in users"
        :key="u.id"
        class="table-row"
      >
        <span class="font-medium text-[var(--color-text-primary)]">
        {{ u.email }}
      </span>

        <span class="text-secondary">
        Admin
      </span>

        <span class="text-secondary">
        {{ u.created_at }}
      </span>

        <span class="flex justify-end">
        <span
          :class="[
            'badge-status',
            u.is_activated
              ? 'badge-status-active'
              : 'badge-status-pending'
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