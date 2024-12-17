<script setup lang="ts">
  import { onMounted } from 'vue'
  import { useEmailStore } from '@/stores/emails';

  const emailStore = useEmailStore();
  // Cargar los emails al montar el componente
  onMounted(() => {
    emailStore.loadEmails('mjones7@txu.com', 'from') // Reemplaza 'user@example.com' por el correo del usuario
  })
</script>

<template>
  <div class="flex flex-col p-4 w-full">
    <div v-if="emailStore.error" class="text-red-500">
      Error: {{ emailStore.error }}
    </div>
    
    <div
      v-for="email in emailStore.emails"
      :key="email.message_id"
      class="p-2 border-b cursor-pointer hover:bg-gray-100"
      @click="$emit('selectEmail', email)"
    >
      <p class="font-bold">{{ email.subject }}</p>
      <p>{{ email.from  }}</p>
      <p class="text-sm text-gray-500">{{ email.date }}</p>
    </div>

    <div v-if="emailStore.emails.length === 0" class="text-gray-500">
      No emails found.
    </div>
  </div>
</template>

<style scoped>
.hover\:bg-gray-100:hover {
  background-color: #f7f7f7;
}
</style>