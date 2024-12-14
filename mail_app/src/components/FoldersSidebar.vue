
<script lang="ts" setup>
  import { useEmailStore } from '@/stores/emails';
  import { computed } from 'vue';

  const emailStore = useEmailStore();
  const folders = computed(() => emailStore.folders);
</script>

<template>
  <div class="w-64 bg-gray-100 h-screen p-4">
    <ul>
      <li
        v-for="folder in folders"
        :key="folder"
        class="flex items-center p-2 cursor-pointer hover:bg-gray-200"
        @click="$router.push({ name: 'folder', params: { id: folder } })"
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