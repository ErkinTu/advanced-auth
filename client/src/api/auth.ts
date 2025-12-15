import http from './http';
import type {LoginCredentials, AuthResponse, User, Role, AssignRolePayload, RegisterCredentials} from '../types/auth';

export const registerRequest = (data: RegisterCredentials) =>
  http.post<AuthResponse>('/register', data)

export const loginRequest = (data: LoginCredentials) =>
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

export const assignRoleRequest = (data: AssignRolePayload) =>
  http.post('/role/assign', data)