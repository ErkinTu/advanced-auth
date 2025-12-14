<script setup lang="ts">

import TableContainer from "@/components/ui/TableContainer.vue";
import {useRoles, useUsers} from "@/composables/useAuth.ts";
import {onMounted, ref} from "vue";
import type {AssignRolePayload} from "@/types/auth.ts";

const {users, loadUsers, isLoading, error} = useUsers()
const {roles, loadRoles, assignRole} = useRoles()

const assignRolePayload = ref<AssignRolePayload>({
  user_id: '',
  role_name: '',
})

const handleAssignRole = async () => {
  if (assignRolePayload.value.user_id && assignRolePayload.value.role_name) {
    await assignRole(assignRolePayload.value)
  }
}

onMounted(() => {
  loadUsers()
  loadRoles()
})
</script>

<template>
  <div class="p-6">
    <h2 class="page-title">System Roles</h2>
    <p class="mb-5 text-secondary">Manage roles</p>

    <TableContainer
      title="Role Directory"
      description="Assign roles to users to control access."
    >
      <div class="p-4 space-y-4">
        <div class="grid grid-cols-1 md:grid-cols-2 gap-4">
          <!-- User Select -->
          <div>
            <label for="user-select" class="block text-sm font-medium mb-2">
              Select User
            </label>
            <select
              id="user-select"
              v-model="assignRolePayload.user_id"
              class="w-full px-3 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-blue-500"
              :disabled="isLoading"
            >
              <option :value="'0'">Choose a user...</option>
              <option
                v-for="user in users"
                :key="user.id"
                :value="String(user.id)"
              >
                {{ user.email }}
              </option>
            </select>
          </div>

          <!-- Role Select -->
          <div>
            <label for="role-select" class="block text-sm font-medium mb-2">
              Select Role
            </label>
            <select
              id="role-select"
              v-model="assignRolePayload.role_name"
              class="w-full px-3 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-blue-500"
            >
              <option value="">Choose a role...</option>
              <option
                v-for="role in roles"
                :key="role.id"
                :value="role.name"
              >
                {{ role.name }}
              </option>
            </select>
          </div>
        </div>

        <!-- Assign Button -->
        <div class="flex justify-end">
          <button
            @click="handleAssignRole"
            :disabled="!assignRolePayload.user_id || !assignRolePayload.role_name || isLoading"
            class="px-4 py-2 bg-blue-600 text-white rounded-md hover:bg-blue-700 disabled:bg-gray-400 disabled:cursor-not-allowed transition-colors"
          >
            Assign Role
          </button>
        </div>

        <!-- Loading/Error States -->
        <div v-if="isLoading" class="text-center text-gray-500">
          Loading users...
        </div>
        <div v-if="error" class="text-center text-red-500">
          {{ error }}
        </div>
      </div>
    </TableContainer>

  </div>
</template>