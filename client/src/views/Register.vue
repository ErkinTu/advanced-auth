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
    <div class="flex flex-col gap-4 min-w-[40%] items-center p-6 mx-auto border-[1px] border-gray-200 rounded-md shadow-md w-min">
      <h2 class="text-2xl text-black font-bold">Create an account</h2>
      <p class="bg-gray-100/70 w-62 rounded-md text-xs text-center">Enter your email below to create your account</p>
      <div class="w-full">
        <label for="email" class="block text-sm font-semibold text-gray-700">Email</label>
        <input
          v-model="email"
          id="email"
          type="email"
          class="w-full p-1 border border-gray-300 rounded-md focus:border-cyan-500 focus:ring focus:ring-cyan-300 focus:ring-opacity-50 text-sm"
          placeholder="Enter your email"
        />
      </div>
      <div class="w-full">
        <label for="password" class="block text-sm font-semibold text-gray-700">Password</label>
        <input
          v-model="password"
          id="password"
          type="password"
          class="w-full p-1 border border-gray-300 rounded-md focus:border-cyan-500 focus:ring focus:ring-cyan-300 focus:ring-opacity-50 text-sm"
          placeholder="Enter your password"
        />
      </div>
      <button
        @click="submit"
        class="bg-cyan-900/80 text-white px-6 py-1 rounded-md font-bold mt-4 w-full hover:bg-cyan-700 transition duration-200"
      >
        Create account
      </button>
      <p class="mt-2 text-sm">Already have an account? <router-link to="/login" class="text-cyan-700 hover:text-cyan-500">Sign in</router-link></p>
    </div>
  </div>
</template>