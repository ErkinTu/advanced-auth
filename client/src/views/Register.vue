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
  <div class="p-4 max-w-sm mx-auto">
    <input v-model="email" class="input" placeholder="Email"/>
    <input v-model="password" type="password" class="input mt-2" placeholder="Password"/>
    <button @click="submit" class="btn mt-4">Register</button>
  </div>
</template>