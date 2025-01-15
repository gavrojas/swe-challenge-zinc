
<script lang="ts" setup>
  import { useEmailStore } from '@/stores/emails';
  import { computed } from 'vue';

  const emailStore = useEmailStore();
  const folders = computed(() => emailStore.folders);

  const emit = defineEmits(['updateSearchWithFolders']);
</script>

<template>
  <div class="w-64 bg-gray-100 h-fit p-4">
    <ul>
      <li
        v-for="folder in folders"
        :key="folder"
        class="flex items-center p-2 cursor-pointer hover:bg-gray-200"
        @click="emit('updateSearchWithFolders', { query: folder, field: 'folder_path' })"
      >
        <span class="mr-2">📁</span>
        <span>{{ folder }}</span>
      </li>
    </ul>
  </div>
</template>

<style scoped>
.hover\:bg-gray-200:hover {
  background-color: #e5e7eb;
}
</style>