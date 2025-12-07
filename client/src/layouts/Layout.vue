<script setup lang="ts">

</script>

<template>
  <div class="min-h-screen flex flex-col text-gray-600">

    <!-- Header -->
    <header class="flex items-center justify-between px-10 py-4 bg-gray-100/98 fixed w-full top-0 z-10">
      <div class="flex items-center">
        <router-link to="/" class="text-xl text-cyan-900 font-bold pr-10">Auth <span class="text-black">UI</span></router-link>
        <div class="flex gap-2 text-sm=">
          <router-link to="/">Home</router-link>
          <router-link to="/users">Users</router-link>
        </div>
      </div>
      <div class="flex gap-5 font-bold text-sm items-center">
        <router-link to="/login">Login</router-link>
        <router-link class="bg-cyan-900/80 text-white px-2 py-1 rounded-md" to="/register">
          Register
        </router-link>
      </div>
    </header>

    <!-- Main -->
    <main class="flex-1 p-4 mt-15">
      <router-view/>
    </main>

    <!-- Footer -->
    <footer class="p-4 bg-gray-100 text-center text-sm">
      © 2025 My App (footer)
    </footer>

  </div>
</template>