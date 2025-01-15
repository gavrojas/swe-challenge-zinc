<script setup lang="ts">
  import { ref, computed } from 'vue'
  import Sidebar from '@/components/FoldersSidebar.vue'
  import EmailList from '@/components/EmailList.vue'
  import EmailView from '@/components/EmailView.vue'
  import { useEmailStore } from '@/stores/emails'
  import type { EmailDocument, SearchPayload } from '@/types'

  // Obtener el store de correos
  const emailStore = useEmailStore()

  // Email seleccionado
  const selectedEmail = ref<EmailDocument | null>(null)

  // Función para seleccionar un email
  const selectEmail = (email: EmailDocument) => {
    selectedEmail.value = email
  }

  const clearSelectedEmail = () => {
    selectedEmail.value = null; // Limpiar el email seleccionado
  }

  // Comprobación si hay correos
  const hasEmails = computed(() => emailStore.emails.length > 0)

  const updateSearchWithFolders = (payload: SearchPayload) => {
    emailStore.searchQuery = payload.query;
    emailStore.searchField = payload.field;
    emailStore.itemsPerPage = 0; 
    emailStore.itemsPerPage++;
    emailStore.loadEmails();

    clearSelectedEmail()
  };

</script>

<template>
    <Sidebar @updateSearchWithFolders="updateSearchWithFolders" />
    <div class="flex w-full">
      <EmailList @selectEmail="selectEmail" @clearSelectedEmail="clearSelectedEmail"/>
      
      <!-- Solo muestra EmailView si hay correos y uno está seleccionado -->
      <EmailView :selectedEmail="selectedEmail" v-if="selectedEmail" />
      
      <!-- Mensaje si no hay correos disponibles -->
      <div v-else class="p-4 text-gray-500 w-full">
        <p v-if="!hasEmails">No emails available. Please load emails.</p>
        <p v-else>Please select an email to view its details.</p>
      </div>
    </div>
</template>
