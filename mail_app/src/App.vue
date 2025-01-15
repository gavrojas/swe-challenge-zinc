<script setup lang="ts">
  import { useAuthStore } from '@/stores/auth'
  import { watchEffect } from 'vue';

  const authStore = useAuthStore()

  watchEffect(() => {
    authStore.init();
  });

  const logout = (e: Event) => {
    e.preventDefault();
    authStore.clearSession(false);
  }
</script>

<template>
  <div class="flex flex-col min-h-screen w-full">
    <header class="flex items-center justify-between p-4 bg-gray-800 text-white w-full">
      <img alt="Vue logo" class="logo" src="@/assets/logo.svg" width="50" height="50" />
      <p class="hidden md:block text-center text-lg font-bold">MAIL APP</p>
      <nav>
        <RouterLink v-if="!authStore.isLoggedIn" to="/login" class="px-2">Login</RouterLink>
        <RouterLink v-if="!authStore.isLoggedIn" to="/register" class="px-2">Register</RouterLink>
        <button v-if="authStore.isLoggedIn" @click.prevent="logout" class="px-2 bg-transparent text-white border-0 cursor-pointer hover:underline">Logout</button>
      </nav>
    </header>
    <main
      class="flex flex-1 p-4">
      <img v-if="!authStore.isLoggedIn" src="@/assets/login.svg" alt="Emails" class="hidden md:block w-1/2 h-full object-cover" />
      <div class="flex flex-1 w-full">
        <router-view class="flex-1" />
      </div>
    </main>
    <footer class="bg-gray-800 text-white p-2">
      <div class="flex justify-center space-x-4">
        <a href="https://www.linkedin.com/in/gavrojas-dev" target="_blank" class="hover:underline">
          <i class="fab fa-linkedin"></i>
          @gavrojas-dev</a>
        <a href="https://github.com/gavrojas">
          <i class="fab fa-github"></i>
          @gavrojas</a>
      </div>
      <p class="text-center mt-2">Developed by 'Gabriela Rojas' with 🩵, ☕ and good music.</p>
    </footer>
  </div>
</template>
