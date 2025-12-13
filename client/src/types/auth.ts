export interface Role {
  id: number
  name: string
}

export interface User {
  id: number
  email: string
  is_activated: boolean
  created_at: string
  updated_at: string
  roles?: Role[]
}

export interface AuthCredentials {
  email: string
  password: string
}

export interface AuthResponse {
  message: string
  user?: User
}