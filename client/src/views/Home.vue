<script setup lang="ts">

</script>

<template>
  <div class="flex flex-col items-center justify-center text-center p-6">

    <div class="bg-gray-200 rounded-md shadow-md px-30 py-20 flex flex-col items-center justify-center gap-5">
      <p class="text-[8px] bg-gray-100 rounded-sm w-15 font-bold">v1.0 version</p>
      <h2 class="text-5xl font-bold">Header block</h2>
      <p>Description Lorem ipsum dolor sit amet, consectetur adipisicing elit. Delectus, iusto, maiores! Consequatur culpa deleniti dolor doloribus dolorum, eligendi eos facere, inventore iure necessitatibus nemo quasi quod recusandae sint, unde velit!</p>
      <div>
        <router-link to="/register" class="bg-cyan-900/80 text-white px-6 py-2 rounded-md font-bold mr-5">
          Get Started 🠊
        </router-link>
        <router-link to="/users" class="bg-white border-[1px] px-6 py-2 rounded-md">View Demo Data</router-link>
      </div>
    </div>

    <div class="flex gap-10 mt-10 text-left">
      <div class="flex flex-col gap-2 shadow-xs border-[2px] border-gray-100 rounded-md p-5">
        <div class="bg-gray-200 w-7 h-7 rounded-md text-center">🔐</div>
        <h3 class="text-black font-bold">Secure by Default</h3>
        <p>Enterprise-grade security standards implemented at every layer of the application</p>
      </div>
      <div class="flex flex-col gap-2 shadow-xs border-[2px] border-gray-100 rounded-md p-5">
        <div class="bg-gray-200 w-7 h-7 rounded-md text-center">🚀</div>
        <h3 class="text-black font-bold">High Performance</h3>
        <p>Optimized for speed with minimal overhead and efficient rendering cycles</p>
      </div>
      <div class="flex flex-col gap-2 shadow-xs border-[2px] border-gray-100 rounded-md p-5">
        <div class="bg-gray-200 w-7 h-7 rounded-md text-center">👥</div>
        <h3 class="text-black font-bold">Team Collaboration</h3>
        <p>Built-in tools for managing users, roles, and permissions seamlessly</p>
      </div>
    </div>

  </div>
</template>
