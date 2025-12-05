import axios from 'axios';


const http = axios.create({
  baseURL: import.meta.env.VITE_API_URL,
  headers: {
    'Content-Type': 'application/json'
  },
  withCredentials: true,
})

// update access token on 401 response
http.interceptors.response.use(
  res => res,
  async err => {
    if (err.response?.status === 401) {
      try {
        const refreshToken = await http.post('/refresh')
        const newToken = refreshToken.data.access_token
        http.defaults.headers.common['Authorization'] = `Bearer ${newToken}`
        err.config.headers['Authorization'] = `Bearer ${newToken}`
        return http(err.config)
      } catch {
        return Promise.reject(err)
      }
    }
    return Promise.reject(err)
  }
)

export default http;