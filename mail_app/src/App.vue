<template>
  <div class="flex flex-col h-screen w-full">
    <header class="flex items-center justify-between p-4 bg-gray-800 text-white w-full">
      <img alt="Vue logo" class="logo" src="@/assets/logo.svg" width="50" height="50" />
      <p class="hidden md:block text-center text-lg font-bold">MAIL APP</p>
      <nav>
        <RouterLink to="/login" class="px-2">Login</RouterLink>
        <RouterLink to="/register" class="px-2">Register</RouterLink>
      </nav>
    </header>
    <div class="flex flex-1 w-full">
      <router-view class="flex-1" />
    </div>
  
  </div>
</template>
