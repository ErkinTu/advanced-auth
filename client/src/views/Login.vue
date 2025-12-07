<script setup lang="ts">
import {ref} from 'vue'
import {useAuth} from '@/composables/useAuth'
import BaseInput from '@/components/ui/BaseInput.vue'
import BaseButton from '@/components/ui/BaseButton.vue'


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
      
      <BaseInput
        v-model="email"
        id="email"
        label="Email"
        type="email"
        placeholder="Enter your email"
      />
      
      <BaseInput
        v-model="password"
        id="password"
        label="Password"
        type="password"
        placeholder="Enter your password"
      />
      
      <BaseButton
        @click="submit"
        variant="primary"
        type="submit"
        class="mt-4 w-full"
      >
        Sign in
      </BaseButton>
      
      <p class="mt-2 text-sm">Don't have an account? <router-link to="/register" class="link-primary">Register</router-link></p>
    </div>

  </div>
</template>