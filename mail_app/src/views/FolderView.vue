<script setup lang="ts">
  import { ref, computed } from 'vue'
  import Sidebar from '@/components/FoldersSidebar.vue'
  import EmailList from '@/components/EmailList.vue'
  import EmailView from '@/components/EmailView.vue'
  import { useEmailStore } from '@/stores/emails'
  import type { EmailDocument } from '@/types'

  // Obtener el store de correos
  const emailStore = useEmailStore()

  // Email seleccionado
  const selectedEmail = ref<EmailDocument | null>(null)

  // Función para seleccionar un email
  const selectEmail = (email: EmailDocument) => {
    selectedEmail.value = email
  }

  // Comprobación si hay correos
  const hasEmails = computed(() => emailStore.emails.length > 0)
</script>

<template>
  <div class="flex h-full">
    <Sidebar />
    <div class="flex flex-1">
      <EmailList @selectEmail="selectEmail" />
      
      <!-- Solo muestra EmailView si hay correos y uno está seleccionado -->
      <EmailView :selectedEmail="selectedEmail" v-if="selectedEmail" />
      
      <!-- Mensaje si no hay correos disponibles -->
      <div v-else class="p-4 text-gray-500 flex items-center justify-center w-full">
        <p v-if="!hasEmails">No emails available. Please load emails.</p>
        <p v-else>Please select an email to view its details.</p>
      </div>
    </div>
  </div>
</template>
