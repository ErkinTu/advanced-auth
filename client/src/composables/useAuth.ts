import type {AuthCredentials} from "@/types/auth.ts";
import {useAuthStore, useUserStore} from "@/store/auth.ts";

export function useAuth() {
  const auth = useAuthStore()

  const login = (data: AuthCredentials) => auth.login(data)
  const register = (data: AuthCredentials) => auth.register(data)

  return {login, register, accessToken: auth.accessToken}
}

export function useUsers() {
  const store = useUserStore()
  return {users: store.users, loadUsers: store.loadUsers}
}