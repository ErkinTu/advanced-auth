<script setup lang="ts">

</script>

<template>
  <div class="flex flex-col items-center justify-center text-center p-6">
    <h1 class="text-3xl font-bold mb-4">Welcome</h1>

    <p class="text-gray-600 max-w-md">
      This is a basic public homepage.
      You can register, log in or explore the app.
    </p>

    <div class="flex gap-4 mt-6">
      <router-link
        to="/register"
        class="px-4 py-2 rounded bg-blue-600 text-white hover:bg-blue-700 transition"
      >
        Register
      </router-link>

      <router-link
        to="/login"
        class="px-4 py-2 rounded bg-gray-700 text-white hover:bg-gray-800 transition"
      >
        Login
      </router-link>

      <router-link
        to="/users"
        class="px-4 py-2 rounded bg-green-600 text-white hover:bg-green-700 transition"
      >
        Users
      </router-link>
    </div>
  </div>
</template>
