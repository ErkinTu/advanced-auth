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

export interface LoginCredentials {
  email: string
  password: string
}

export interface RegisterCredentials {
  email: string
  password: string
  password_confirm: string
}

export interface AuthResponse {
  message: string
  user?: User
}

export interface AssignRolePayload {
  user_id: string
  role_name: string
}