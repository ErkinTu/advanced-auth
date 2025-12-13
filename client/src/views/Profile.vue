<script setup lang="ts">
import {useAuth} from '@/composables/useAuth'
import TableContainer from "@/components/ui/TableContainer.vue"
import StatusBadge from '@/components/ui/StatusBadge.vue'

const {currentUser, isLoading} = useAuth()
</script>

<template>
  <div class="p-6">
    <h2 class="page-title">My Profile</h2>
    <p class="mb-5">Your account information and settings</p>

    <div v-if="isLoading" class="text-muted">Loading...</div>

    <TableContainer
      v-else-if="currentUser"
      :title="currentUser.email"
      description="User information and activation status."
    >
      <div class="space-y-3">
        <div class="flex items-center gap-2">
          <span class="text-secondary">ID:</span>
          <strong class="text-[var(--color-text-primary)]">{{ currentUser.id }}</strong>
        </div>

        <div class="flex items-center gap-2">
          <span class="text-secondary">Status:</span>
          <StatusBadge :active="currentUser.is_activated" />
        </div>

        <div v-if="currentUser.roles && currentUser.roles.length > 0">
          <span class="text-secondary">Roles:</span>
          <div class="flex gap-2 mt-1 flex-wrap">
            <span
              v-for="role in currentUser.roles"
              :key="role.id"
              class="px-2 py-1 bg-cyan-100 text-cyan-800 rounded-md text-xs font-medium"
            >
              {{ role.name }}
            </span>
          </div>
        </div>

        <div class="flex items-center gap-2">
          <span class="text-secondary">Created:</span>
          <span class="text-sm">{{ new Date(currentUser.created_at).toLocaleDateString() }}</span>
        </div>
      </div>
    </TableContainer>

    <div v-else class="text-error">
      Unable to load profile information
    </div>
  </div>
</template>