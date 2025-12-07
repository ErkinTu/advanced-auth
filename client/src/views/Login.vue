<script setup lang="ts">
import {ref} from 'vue'
import {useAuth} from '@/composables/useAuth'


const {login} = useAuth()

const email = ref('')
const password = ref('')

async function submit() {
  await login({email: email.value, password: password.value})
}
</script>


<template>
  <div>

    <div class="form-container">
      <h2 class="section-title">Welcome back</h2>
      <p class="form-description">Enter your credentials to access your account</p>
      <div class="w-full">
        <label for="email" class="form-label">Email</label>
        <input
          v-model="email"
          id="email"
          type="email"
          class="form-input"
          placeholder="Enter your email"
        />
      </div>
      <div class="w-full">
        <label for="password" class="form-label">Password</label>
        <input
          v-model="password"
          id="password"
          type="password"
          class="form-input"
          placeholder="Enter your password"
        />
      </div>
      <button
        @click="submit"
        class="btn-primary mt-4 w-full"
      >
        Sign in
      </button>
      <p class="mt-2 text-sm">Don't have an account? <router-link to="/register" class="link-primary">Register</router-link></p>
    </div>

  </div>
</template>