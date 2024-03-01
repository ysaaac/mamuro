<script setup lang="ts">

import SearchComponent from '@/components/HeaderComponent/SearchComponent.vue'
import UserInfoComponent from '@/components/HeaderComponent/UserInfoComponent.vue'
</script>

<template>
  <header class="w-full flex flex-row items-center pr-4">
    <SearchComponent />
    <UserInfoComponent />
  </header>
</template>