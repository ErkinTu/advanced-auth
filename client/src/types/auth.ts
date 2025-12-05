export interface User {
  id: number
  email: string
  is_activated: boolean
}


export interface AuthCredentials {
  email: string
  password: string
}


export interface AuthResponse {
  message: string
  token: string // access token
}


export interface ApiResponse<T> {
  message: string
  data: T
}