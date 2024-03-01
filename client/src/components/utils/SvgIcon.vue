<script setup lang="ts">
import { ref, defineProps } from 'vue'

const { svg } = defineProps<{ svg: string }>()
const svgContent = ref<string | null>(null)

fetch(svg)
  .then(response => response.text())
  .then(data => {
    svgContent.value = data
  })
  .catch(error => {
    console.error('Error loading SVG:', error)
  })
</script>

<template>
  <div v-html="svgContent"></div>
</template>

<style scoped lang="scss">

</style>