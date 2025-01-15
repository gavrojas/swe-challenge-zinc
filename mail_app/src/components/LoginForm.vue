<script setup lang="ts">
import { ref } from 'vue'
import { useRouter, useRoute } from 'vue-router'
import { useAuthStore } from '@/stores/auth'

const loggedUser = ref({ username: '', password: '' })
const loginError = ref('')
const router = useRouter()
const route = useRoute()
const authStore = useAuthStore()
const showSuccessSnackbar = ref(false);
const successMessage = ref('User registered with success!');
const snackbarTimeout = ref(5000);

if (route.query.registered) {
  showSuccessSnackbar.value = true;
}

async function login() {
  try {
    // const response = await fetch('http://localhost:8080/auth/login', {
    const response = await fetch('https://zinc-gavrojas.serveftp.com/auth/login', {
      method: 'POST',
      headers: { 'Content-Type': 'application/json' },
      body: JSON.stringify(loggedUser.value),
    })
    const data = await response.json()

    if (!response.ok) {
      throw new Error(`${data.error}`)
    }

    authStore.setSession(data.token)
    authStore.username = loggedUser.value.username

    router.push('/home')
  } catch (error) {
    loginError.value = error instanceof Error ? error.message : 'An unexpected error occurred'
  }
}
</script>

<template>
  <div class="p-6 max-w-sm mx-auto">
    <h1 class="text-2xl font-bold mb-4">Login</h1>
    <p v-if="route.query.registered" class="text-green-500 mb-4">Now, you can login!</p>
    <form @submit.prevent="login">
      <div class="mb-4">
        <label class="block text-gray-700">Username</label>
        <input v-model="loggedUser.username" class="w-full p-2 border border-gray-300 rounded" type="text" placeholder="Username" />
      </div>
      <div class="mb-4">
        <label class="block text-gray-700">Password</label>
        <input v-model="loggedUser.password" class="w-full p-2 border border-gray-300 rounded" type="password" placeholder="Password" />
      </div>
      <button class="w-full bg-blue-500 text-white py-2 rounded hover:bg-blue-700">Login</button>
      <p v-if="loginError" class="error text-red-500 mt-2">{{ loginError }}</p>
    </form>
    <p class="mt-4">
      Don't have an account? <router-link to="/register" class="text-blue-500">Register here</router-link>
    </p>

    
    <v-snackbar
      v-model="showSuccessSnackbar"
      :timeout="snackbarTimeout"
      color="success"
      border="left"
      rounded
      class="flex items-end justify-end"
    > 
      {{ successMessage }}
      <template #actions>
        <v-btn
          color="grey-darken-3"
          variant="text"
          @click="showSuccessSnackbar = false"
        >
          Close
        </v-btn>
      </template>
    </v-snackbar>
  </div>
</template>