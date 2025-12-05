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

  async function loadUsers() {
    const res = await getUsersRequest()
    users.value = res.data.users
  }

  return {users, loadUsers}
})