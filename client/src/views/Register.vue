<script setup lang="ts">
import {ref} from 'vue'
import {useAuth} from '@/composables/useAuth'
import BaseInput from '@/components/ui/BaseInput.vue'
import BaseButton from '@/components/ui/BaseButton.vue'


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
        Create account
      </BaseButton>
      
      <p class="mt-2 text-sm">Already have an account? <router-link to="/login" class="link-primary">Sign in</router-link></p>
    </div>
  </div>
</template>