<script setup lang="ts">
import {ref} from 'vue'
import {useAuth} from '@/composables/useAuth'
import {useRouter} from 'vue-router'
import BaseInput from '@/components/ui/BaseInput.vue'
import BaseButton from '@/components/ui/BaseButton.vue'

const router = useRouter()
const {register, isLoading} = useAuth()

const email = ref('')
const password = ref('')
const passwordConfirm = ref('')
const error = ref<string | null>(null)
const successMessage = ref<string | null>(null)

async function submit() {
  error.value = null
  successMessage.value = null

  if (!email.value || !password.value || !passwordConfirm.value) {
    error.value = 'All fields are required.'
    return
  }

  if (password.value !== passwordConfirm.value) {
    error.value = 'Passwords do not match.'
    return
  }

  if (password.value.length < 6) {
    error.value = 'Password must be at least 6 characters long.'
    return
  }

  try {
    await register({email: email.value, password: password.value, password_confirm: passwordConfirm.value})
    successMessage.value = 'Registration successful! Please check your email to activate your account.'

    setTimeout(() => {
      router.push('/login')
    }, 3000)
  } catch (e: any) {
    error.value = e.response?.data?.message || 'Registration failed. Please try again.'
  }
}
</script>

<template>
  <div>
    <div class="form-container">
      <h2 class="section-title">Create an account</h2>
      <p class="form-description">Enter your email below to create your account</p>

      <div v-if="successMessage" class="mb-4 p-3 bg-green-100 text-green-800 rounded-md text-sm">
        {{ successMessage }}
      </div>

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

      <BaseInput
        v-model="passwordConfirm"
        id="password_confirm"
        label="Confirm Password"
        type="password"
        placeholder="Confirm your password"
        :disabled="isLoading"
      />

      <BaseButton
        @click="submit"
        variant="primary"
        type="submit"
        class="mt-4 w-full"
        :disabled="isLoading"
      >
        {{ isLoading ? 'Creating account...' : 'Create account' }}
      </BaseButton>

      <p class="mt-2 text-sm">Already have an account? <router-link to="/login" class="link-primary">Sign in</router-link></p>
    </div>
  </div>
</template>