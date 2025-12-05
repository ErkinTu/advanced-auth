import {defineStore} from 'pinia'
import {ref} from 'vue'
import {getUsersRequest, loginRequest, registerRequest} from '@/api/auth'
import type {AuthCredentials, User} from '@/types/auth'


export const useAuthStore = defineStore('auth', () => {
  const accessToken = ref<string>('')

  async function register(data: AuthCredentials) {
    await registerRequest(data)
  }

  async function login(data: AuthCredentials) {
    const res = await loginRequest(data)
    accessToken.value = res.data.token
  }

  return {accessToken, login, register}
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
      error.value = "Failed to load users"
    } finally {
      isLoading.value = false
    }
  }

  return {users, loadUsers, isLoading, error}
})