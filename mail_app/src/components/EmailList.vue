<script setup lang="ts">
  import { ref, onMounted } from 'vue'
  import { useEmailStore } from '@/stores/emails';
  import SearchCard from './SearchCard.vue';
  import type { SearchPayload } from '@/types'

  const emailStore = useEmailStore();
  const itemsPerPage = ref(0); // Número de elementos por página
  const searchQuery = ref(''); // Variable para el campo de búsqueda
  const searchField = ref('from'); // Campo de búsqueda seleccionado
  
  const loadEmails = () => {
    const maxResults = itemsPerPage.value + 1
    itemsPerPage.value++
    emailStore.loadEmails(searchQuery.value, searchField.value, 30 * maxResults);
  };

  const updateSearch = (payload: SearchPayload) => {
    searchQuery.value = payload.query;
    searchField.value = payload.field;
    itemsPerPage.value = 0; // Reiniciar la paginación
    loadEmails(); // Cargar los emails con los nuevos parámetros
  };

  // Cargar los emails al montar el componente
  onMounted(() => {
    loadEmails();
  })
</script>

<template>
  <div class="flex flex-col p-4 w-full">
    <div v-if="emailStore.error" class="text-red-500">
      Error: {{ emailStore.error }}
    </div>

    <SearchCard @updateSearch="updateSearch" />

    <div class="overflow-y-auto h-full max-h-[calc(100vh-200px)]">
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

    <!-- Load more emails -->
    <div class="flex justify-between mt-4 items-center">
      <v-btn @click="loadEmails()" color="primary">
        Load More
      </v-btn>
      <span class="font-semibold">Total Emails: {{ emailStore.totalResults }}</span>
    </div>
  </div>
</template>

<style scoped>
  .hover\:bg-gray-100:hover {
    background-color: #f7f7f7;
  }
  /* Estilo para la barra de desplazamiento */
  .overflow-y-auto {
    scrollbar-width: thin; /* Firefox */
    scrollbar-color: #cbd5e1 #f3f4f6; /* Firefox */
  }

  .overflow-y-auto::-webkit-scrollbar {
    width: 8px; /* Ancho de la barra de desplazamiento */
  }

  .overflow-y-auto::-webkit-scrollbar-track {
    background: #f3f4f6; /* Color de fondo de la pista */
  }

  .overflow-y-auto::-webkit-scrollbar-thumb {
    background-color: #cbd5e1; /* Color de la barra de desplazamiento */
    border-radius: 10px; /* Bordes redondeados */
  }

  .overflow-y-auto::-webkit-scrollbar-thumb:hover {
    background-color: #a0aec0; /* Color de la barra de desplazamiento al pasar el mouse */
  }
</style>