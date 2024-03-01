<script setup lang="ts">
import ContentList from '@/components/ContentComponent/ContentList.vue'
import ContentViewer from '@/components/ContentComponent/ContentViewer.vue'
</script>

<template>
  <div class="flex flex-row flex-1 overflow-y-auto">
    <ContentList />
    <ContentViewer />
  </div>
</template>

<style scoped lang="scss">

</style>