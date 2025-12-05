import type {AuthCredentials} from "@/types/auth.ts";
import {useAuthStore, useUserStore} from "@/store/auth.ts";
import {storeToRefs} from "pinia";

export function useAuth() {
  const auth = useAuthStore()

  const login = (data: AuthCredentials) => auth.login(data)
  const register = (data: AuthCredentials) => auth.register(data)

  return {login, register, accessToken: auth.accessToken}
}

export function useUsers() {
  const store = useUserStore()
  const { users, isLoading, error } = storeToRefs(store)

  return {users: store.users, loadUsers: store.loadUsers, isLoading: store.isLoading, error: store.error}
}