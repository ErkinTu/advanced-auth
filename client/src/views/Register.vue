<script setup lang="ts">
import {ref} from 'vue'
import {useAuth} from '@/composables/useAuth'


const {register} = useAuth()

const email = ref('')
const password = ref('')

async function submit() {
  await register({email: email.value, password: password.value})
}
</script>


<template>
  <div>
    <div class="form-container">
      <h2 class="section-title">Create an account</h2>
      <p class="form-description">Enter your email below to create your account</p>
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
        Create account
      </button>
      <p class="mt-2 text-sm">Already have an account? <router-link to="/login" class="link-primary">Sign in</router-link></p>
    </div>
  </div>
</template>