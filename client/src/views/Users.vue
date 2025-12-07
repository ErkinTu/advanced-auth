<script setup lang="ts">
import {useUsers} from '@/composables/useAuth'
import {onMounted} from 'vue'
import TableContainer from '@/components/ui/TableContainer.vue'
import StatusBadge from '@/components/ui/StatusBadge.vue'


const {users, loadUsers, isLoading, error} = useUsers()

onMounted(loadUsers)
</script>

<template>
  <div class="p-6">
    <h2 class="page-title">System Users</h2>
    <p class="mb-5">Manage user access and activation status</p>

    <TableContainer 
      title="User Directory"
      description="A list of all registered users in the system."
    >
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
        <StatusBadge :active="u.is_activated" />
      </span>
      </div>
    </TableContainer>
  </div>

</template>