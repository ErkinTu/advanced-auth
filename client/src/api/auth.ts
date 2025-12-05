import http from './http';
import type {AuthCredentials, AuthResponse, User} from '../types/auth';


export const registerRequest = (data: AuthCredentials) =>
  http.post<AuthResponse>('/register', data)

export const loginRequest = (data: AuthCredentials) =>
  http.post<AuthResponse>('/login', data)

export const activateRequest = (token: string) =>
  http.get(`/activate/${token}`)

export const getUsersRequest = () => http.get<{users: User[]}>('/users')