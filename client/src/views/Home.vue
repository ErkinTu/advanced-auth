<script setup lang="ts">

</script>

<template>
  <div class="flex flex-col items-center justify-center text-center p-6">

    <div class="card px-[var(--spacing-hero-x)] py-[var(--spacing-hero-y)]">
      <p class="badge">v1.0 version</p>
      <h2 class="hero-title">Header block</h2>
      <p>Description Lorem ipsum dolor sit amet, consectetur adipisicing elit. Delectus, iusto, maiores! Consequatur culpa deleniti dolor doloribus dolorum, eligendi eos facere, inventore iure necessitatibus nemo quasi quod recusandae sint, unde velit!</p>
      <div>
        <router-link to="/register" class="btn-primary mr-5">
          Get Started 🠊
        </router-link>
        <router-link to="/users" class="btn-secondary">View Demo Data</router-link>
      </div>
    </div>

    <div class="flex gap-[var(--spacing-2xl)] mt-[var(--spacing-2xl)] text-left">
      <div class="card-feature">
        <div class="card-feature-icon">🔐</div>
        <h3 class="section-subtitle">Secure by Default</h3>
        <p>Enterprise-grade security standards implemented at every layer of the application</p>
      </div>
      <div class="card-feature">
        <div class="card-feature-icon">🚀</div>
        <h3 class="section-subtitle">High Performance</h3>
        <p>Optimized for speed with minimal overhead and efficient rendering cycles</p>
      </div>
      <div class="card-feature">
        <div class="card-feature-icon">👥</div>
        <h3 class="section-subtitle">Team Collaboration</h3>
        <p>Built-in tools for managing users, roles, and permissions seamlessly</p>
      </div>
    </div>

  </div>
</template>
