<script setup lang="ts">

</script>

<template>
  <div class="min-h-screen flex flex-col">

    <!-- Header -->
    <header class="p-4 border-b bg-white">
      <h1 class="text-xl font-bold">My App (header)</h1>
    </header>

    <!-- Main -->
    <main class="flex-1 p-4">
      <router-view/>
    </main>

    <!-- Footer -->
    <footer class="p-4 border-t bg-gray-100 text-center text-sm">
      © 2025 My App (footer)
    </footer>

  </div>
</template>