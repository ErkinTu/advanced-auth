<script setup lang="ts">
import {ref} from "vue";
import TableContainer from "@/components/ui/TableContainer.vue";

const user = ref({
  id: 1,
  email: 'erkin@gmail.com',
  is_activated: true,
  role: 'admin'
})
</script>

<template>
  <div class="p-6">
    <TableContainer
      :title="user.email"
      description="User information and activation status."
    >
      <p>Role <strong>{{user.role}}</strong></p>
      <p>Status <strong>{{user.is_activated ? 'activated' : 'no activated'}}</strong></p>
    </TableContainer>
  </div>
</template>