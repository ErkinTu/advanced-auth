import {defineStore} from 'pinia'
import {ref, computed} from 'vue'
import {getCurrentUserRequest, getUsersRequest, loginRequest, logoutRequest, registerRequest} from '@/api/auth'
import type {AuthCredentials, User} from '@/types/auth'
import {router} from '@/router'

export const useAuthStore = defineStore('auth', () => {
  const currentUser = ref<User | null>(null)
  const isLoading = ref(false)
  const isAuthenticated = computed(() => currentUser.value !== null)

  async function register(data: AuthCredentials) {
    isLoading.value = true
    try {
      await registerRequest(data)
      return { success: true }
    } catch (e: any) {
      throw e
    } finally {
      isLoading.value = false
    }
  }

  async function login(data: AuthCredentials) {
    isLoading.value = true
    try {
      await loginRequest(data)
      await loadCurrentUser()
      return { success: true }
    } catch (e: any) {
      throw e
    } finally {
      isLoading.value = false
    }
  }

  async function logout() {
    try {
      await logoutRequest()
    } catch (e) {
      console.error('Logout error:', e)
    } finally {
      currentUser.value = null
      router.push('/login')
    }
  }

  async function loadCurrentUser() {
    isLoading.value = true
    try {
      const res = await getCurrentUserRequest()
      currentUser.value = res.data.user
    } catch (e: any) {
      currentUser.value = null
      if (e.response?.status === 401) {
        console.log('Session expired or invalid')
      }
    } finally {
      isLoading.value = false
    }
  }

  function hasRole(roleName: string): boolean {
    if (!currentUser.value?.roles) return false
    return currentUser.value.roles.some(role => role.name === roleName)
  }

  function hasAnyRole(roleNames: string[]): boolean {
    if (!currentUser.value?.roles) return false
    return currentUser.value.roles.some(role => roleNames.includes(role.name))
  }

  return {
    currentUser,
    isLoading,
    isAuthenticated,
    login,
    register,
    logout,
    loadCurrentUser,
    hasRole,
    hasAnyRole
  }
})

export const useUserStore = defineStore('users', () => {
  const users = ref<User[]>([])
  const isLoading = ref(false)
  const error = ref<string | null>(null)

  async function loadUsers() {
    isLoading.value = true
    error.value = null
    try {
      const res = await getUsersRequest()
      users.value = res.data.users
    } catch (e: any) {
      console.error("Failed to load users", e)
      error.value = e.response?.data?.message || "Failed to load users"
    } finally {
      isLoading.value = false
    }
  }

  return {
    users,
    loadUsers,
    isLoading,
    error
  }
})