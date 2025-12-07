<script setup lang="ts">
import BaseCard from '@/components/ui/BaseCard.vue'
import BaseBadge from '@/components/ui/BaseBadge.vue'
import BaseButton from '@/components/ui/BaseButton.vue'
import FeatureCard from '@/components/ui/FeatureCard.vue'
</script>

<template>
  <div class="flex flex-col items-center justify-center text-center p-6">

    <BaseCard class="px-[var(--spacing-hero-x)] py-[var(--spacing-hero-y)]">
      <BaseBadge text="v1.0 version" />
      <h2 class="hero-title">Header block</h2>
      <p>Description Lorem ipsum dolor sit amet, consectetur adipisicing elit. Delectus, iusto, maiores! Consequatur culpa deleniti dolor doloribus dolorum, eligendi eos facere, inventore iure necessitatibus nemo quasi quod recusandae sint, unde velit!</p>
      <div>
        <BaseButton variant="primary" to="/register" class="mr-5">
          Get Started 🠊
        </BaseButton>
        <BaseButton variant="secondary" to="/users">View Demo Data</BaseButton>
      </div>
    </BaseCard>

    <div class="flex gap-[var(--spacing-2xl)] mt-[var(--spacing-2xl)] text-left">
      <FeatureCard 
        icon="🔐"
        title="Secure by Default"
        description="Enterprise-grade security standards implemented at every layer of the application"
      />
      <FeatureCard 
        icon="🚀"
        title="High Performance"
        description="Optimized for speed with minimal overhead and efficient rendering cycles"
      />
      <FeatureCard 
        icon="👥"
        title="Team Collaboration"
        description="Built-in tools for managing users, roles, and permissions seamlessly"
      />
    </div>

  </div>
</template>
