import http from './http';
import type {AuthCredentials, AuthResponse, User, Role} from '../types/auth';

export const registerRequest = (data: AuthCredentials) =>
  http.post<AuthResponse>('/register', data)

export const loginRequest = (data: AuthCredentials) =>
  http.post<AuthResponse>('/login', data)

export const logoutRequest = () =>
  http.post('/logout')

export const activateRequest = (token: string) =>
  http.get(`/activate/${token}`)

export const getUsersRequest = () =>
  http.get<{users: User[]}>('/users')

export const getCurrentUserRequest = () =>
  http.get<{user: User}>('/me')

export const getRolesRequest = () =>
  http.get<{roles: Role[]}>('/roles')

export const assignRoleRequest = (data: { user_id: number; role_name: string }) =>
  http.post('/role/assign', data)