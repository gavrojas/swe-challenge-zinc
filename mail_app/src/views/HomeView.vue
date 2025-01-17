<script setup lang="ts">
  import { ref, computed, onMounted, onBeforeUnmount } from 'vue'
  import Sidebar from '@/components/FoldersSidebar.vue'
  import EmailList from '@/components/EmailList.vue'
  import EmailView from '@/components/EmailView.vue'
  import { useEmailStore } from '@/stores/emails'
  import type { EmailDocument, SearchPayload } from '@/types'

  // Obtener el store de correos
  const emailStore = useEmailStore()

  // Manejo del tamaño de la ventana
  const isMobileOrTablet = ref(false)

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

  const checkWindowSize = () => {
    isMobileOrTablet.value = window.innerWidth <= 768;
  }

  onMounted(() => {
    checkWindowSize()
    window.addEventListener('resize', checkWindowSize)
  })

  onBeforeUnmount(() => {
    window.removeEventListener('resize', checkWindowSize)
  })

</script>

<template>
    <Sidebar v-if="!isMobileOrTablet" @updateSearchWithFolders="updateSearchWithFolders" />
    <div class="flex w-full">
      <!-- Mostrar EmailList en desktop o si no hay emails seleccionados en mobile y tablet-->
      <EmailList v-if="!isMobileOrTablet || (!selectedEmail && isMobileOrTablet)" @selectEmail="selectEmail" @clearSelectedEmail="clearSelectedEmail"/>

      <!-- Solo muestra EmailView si hay correos y uno está seleccionado -->
      <div class="flex w-full">
        <EmailView :selectedEmail="selectedEmail" v-if="selectedEmail" />
        <v-btn class="p-4" v-if="selectedEmail && isMobileOrTablet" color="primary" @click="clearSelectedEmail()" >
          back to emails
        </v-btn> 
      </div>
    
      <!-- Mensaje si no hay correos disponibles -->
      <div v-if="!hasEmails" class="p-4 text-gray-500 w-full">
        <p>No emails available. Please load emails.</p>
      </div>
      <div v-else-if="hasEmails && !selectedEmail && !isMobileOrTablet" class="p-4 text-gray-500 w-full hidden md:block">
        <p>Please select an email to view its details.</p>
      </div>
    </div>
</template>