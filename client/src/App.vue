<script setup lang="ts">
import SidebarComponent from '@/components/SidebarComponent/SidebarComponent.vue'
import NavbarComponent from '@/components/NavbarComponent.vue'
import PreListComponent from '@/components/PreListComponent.vue'
import HeaderComponent from '@/components/HeaderComponent/HeaderComponent.vue'
</script>

<template>
  <div class="h-screen flex">
    <SidebarComponent />
    <NavbarComponent />
    <div class="flex-1 h-screen max-h-screen flex flex-col">
      <HeaderComponent />
      <PreListComponent />
      <RouterView />
    </div>
  </div>
</template>

<style scoped lang="scss">
</style>