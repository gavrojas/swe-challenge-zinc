<script setup lang="ts">
import router from '@/router';
import { ref } from 'vue'

const registeredUser = ref({ username: '', password: '' })
const registerError = ref('')
const isRegistered = ref(false)

async function register() {
  try {
    const response = await fetch('http://localhost:8080/auth/register', {
      method: 'POST',
      headers: { 'Content-Type': 'application/json' },
      body: JSON.stringify(registeredUser.value),
    })

    if (!response.ok) {
      const errorData = await response.json()
      throw new Error(`Registration failed ${errorData.error}`)
    }

    isRegistered.value = true
    router.push({path: '/login', query: { registered: 'true' } })
  } catch (error) {
    registerError.value = error instanceof Error ? error.message : 'An unexpected error occurred'
  }
}
</script>

<template>
  <div class="p-6 max-w-sm mx-auto">
    <h1 class="text-2xl font-bold mb-4">Register</h1>
    <form @submit.prevent="register">
      <div class="mb-4">
        <label class="block text-gray-700">Username</label>
        <input variant="outlined" v-model="registeredUser.username" class="w-full p-2 border border-gray-300 rounded" type="text" placeholder="Username" />
      </div>
      <div class="mb-4">
        <label class="block text-gray-700">Password</label>
        <input variant="outlined" v-model="registeredUser.password" class="w-full p-2 border border-gray-300 rounded" type="password" placeholder="Password" />
      </div>
      <button class="w-full bg-blue-500 text-white py-2 rounded hover:bg-blue-700">Register</button>
      <p v-if="registerError" class="error text-red-500 mt-2">{{ registerError }}</p>
    </form>
    <p class="mt-4">
      Already have an account? <router-link to="/login" class="text-blue-500">Login here</router-link>
    </p>
  </div>
</template>
