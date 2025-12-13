<script setup lang="ts">
import {ref} from 'vue'
import {useAuth} from '@/composables/useAuth'
import {useRouter, useRoute} from 'vue-router'
import BaseInput from '@/components/ui/BaseInput.vue'
import BaseButton from '@/components/ui/BaseButton.vue'

const router = useRouter()
const route = useRoute()
const {login, isLoading} = useAuth()

const email = ref('')
const password = ref('')
const error = ref<string | null>(null)

async function submit() {
  error.value = null

  try {
    await login({email: email.value, password: password.value})

    const redirectTo = route.query.redirect as string || '/profile'
    router.push(redirectTo)
  } catch (e: any) {
    error.value = e.response?.data?.message || 'Login failed. Please check your credentials.'
  }
}
</script>

<template>
  <div>
    <div class="form-container">
      <h2 class="section-title">Welcome back</h2>
      <p class="form-description">Enter your credentials to access your account</p>

      <div v-if="error" class="mb-4 p-3 bg-red-100 text-red-800 rounded-md text-sm">
        {{ error }}
      </div>

      <BaseInput
        v-model="email"
        id="email"
        label="Email"
        type="email"
        placeholder="Enter your email"
        :disabled="isLoading"
      />

      <BaseInput
        v-model="password"
        id="password"
        label="Password"
        type="password"
        placeholder="Enter your password"
        :disabled="isLoading"
      />

      <BaseButton
        @click="submit"
        variant="primary"
        type="submit"
        class="mt-4 w-full"
        :disabled="isLoading"
      >
        {{ isLoading ? 'Signing in...' : 'Sign in' }}
      </BaseButton>

      <p class="mt-2 text-sm">Don't have an account? <router-link to="/register" class="link-primary">Register</router-link></p>
    </div>
  </div>
</template>