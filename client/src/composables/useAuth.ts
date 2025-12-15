import type {LoginCredentials, RegisterCredentials} from "@/types/auth.ts";
import {useAuthStore, useRoleStore, useUserStore} from "@/store/auth.ts";
import {storeToRefs} from "pinia";

export function useAuth() {
  const authStore = useAuthStore()
  const { currentUser, isAuthenticated, isLoading } = storeToRefs(authStore)

  const login = (data: LoginCredentials) => authStore.login(data)
  const register = (data: RegisterCredentials) => authStore.register(data)
  const logout = () => authStore.logout()
  const hasRole = (roleName: string) => authStore.hasRole(roleName)
  const hasAnyRole = (roleNames: string[]) => authStore.hasAnyRole(roleNames)

  return {
    currentUser,
    isAuthenticated,
    isLoading,
    login,
    register,
    logout,
    hasRole,
    hasAnyRole
  }
}

export function useUsers() {
  const store = useUserStore()
  const { users, isLoading, error } = storeToRefs(store)

  return {
    users,
    isLoading,
    error,
    loadUsers: store.loadUsers
  }
}

export function useRoles() {
  const store = useRoleStore()
  const { roles, isLoading, error } = storeToRefs(store)

  return {
    roles,
    isLoading,
    error,
    loadRoles: store.loadRoles,
    assignRole: store.assignRole
  }
}