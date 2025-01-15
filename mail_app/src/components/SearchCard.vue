<script setup lang="ts">
  import { ref, watch } from 'vue'

  const emit = defineEmits(['updateSearch']);

  const searchQuery = ref(''); // Variable para el campo de búsqueda
  const searchField = ref('body'); // Campo de búsqueda seleccionado
  const searchFields = ['from', 'to', 'body', 'subject']; // Opciones para el select

  const loadEmails = () => {
    emit('updateSearch', { query: searchQuery.value, field: searchField.value });
  };

  watch(searchField, () => {
    loadEmails(); // Cargar correos cuando cambia el campo de búsqueda del select
  });
</script>
<template>
  <!-- Campo de búsqueda -->
  <v-card class="mb-4">
      <v-card-title>Search</v-card-title>
      <v-card-text>
        <v-row>
          <v-col cols="12" md="4">
            <v-text-field
              v-model="searchQuery"
              label="Search Emails"
              variant="underlined"
              clearable
              @input="loadEmails()"
            ></v-text-field>
          </v-col>
          <v-col cols="12" md="4">
            <v-select
              v-model="searchField"
              :items="searchFields"
              hide-details
              dense
              variant="underlined"
              style="min-width: 100px;" 
              @change="loadEmails()"
            />
          </v-col>
        </v-row>
      </v-card-text>
    </v-card>
</template>