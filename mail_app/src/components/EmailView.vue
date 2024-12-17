<script setup lang="ts">
  import { defineProps } from 'vue';
  import { useEmailStore } from '@/stores/emails'
  import type { EmailDocument } from '@/types';

  const emailStore = useEmailStore()
  defineProps<{ selectedEmail: EmailDocument | null }>()

</script>

<template>
  <div class="p-4 w-full">
    <!-- Mostrar los detalles del email si hay uno seleccionado -->
    <div v-if="selectedEmail">
      <h2 class="text-xl font-bold">{{ selectedEmail.subject }}</h2>
      <p class="text-gray-500">From: {{ selectedEmail.from }}</p>
      <p class="text-gray-500">To: {{ selectedEmail.to }}</p>
      <div class="mt-4">
        <p>{{ selectedEmail.body }}</p>
      </div>
    </div>
    <!-- Mostrar mensaje si no hay email seleccionado pero hay emails cargados -->
    <div 
      v-else-if="emailStore.emails.length > 0" 
      class="p-4 text-gray-500"
    >
      <p>Please select an email to view its details.</p>
    </div>

    <!-- Mostrar mensaje si no hay emails cargados -->
    <div v-else class="p-4 text-gray-500">
      <p>No emails available. Please load emails.</p>
    </div>
  </div>
</template>

<style scoped>
.text-gray-500 {
  color: #6b7280;
}
</style>
