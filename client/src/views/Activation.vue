<script setup lang="ts">
import {useRoute, useRouter} from 'vue-router'
import {activateRequest} from '@/api/auth'
import {onMounted, ref} from 'vue'
import BaseButton from '@/components/ui/BaseButton.vue'

const route = useRoute()
const router = useRouter()
const message = ref('')
const isLoading = ref(true)
const isSuccess = ref(false)
const error = ref<string | null>(null)

onMounted(async () => {
  try {
    const res = await activateRequest(route.params.token as string)
    message.value = res.data.message || 'Your account has been successfully activated!'
    isSuccess.value = true
  } catch (e: any) {
    error.value = e.response?.data?.message || 'Activation failed. The link may be invalid or expired.'
    isSuccess.value = false
  } finally {
    isLoading.value = false
  }
})

const goToLogin = () => {
  router.push('/login')
}
</script>

<template>
  <div class="flex flex-col items-center justify-center p-6 min-h-[400px]">
    <div class="max-w-md w-full text-center">
      <!-- Loading State -->
      <div v-if="isLoading" class="space-y-4">
        <div class="text-4xl">⏳</div>
        <h2 class="section-title">Activating your account...</h2>
        <p class="text-secondary">Please wait while we verify your email.</p>
      </div>

      <!-- Success State -->
      <div v-else-if="isSuccess" class="space-y-4">
        <div class="text-6xl">✅</div>
        <h2 class="section-title text-green-600">Account Activated!</h2>
        <p class="text-secondary">{{ message }}</p>
        <div class="mt-6">
          <BaseButton variant="primary" @click="goToLogin">
            Go to Login
          </BaseButton>
        </div>
      </div>

      <!-- Error State -->
      <div v-else class="space-y-4">
        <div class="text-6xl">❌</div>
        <h2 class="section-title text-red-600">Activation Failed</h2>
        <p class="text-secondary">{{ error }}</p>
        <div class="mt-6 flex gap-3 justify-center">
          <BaseButton variant="secondary" to="/">
            Go Home
          </BaseButton>
          <BaseButton variant="primary" to="/login">
            Try Login
          </BaseButton>
        </div>
      </div>
    </div>
  </div>
</template>