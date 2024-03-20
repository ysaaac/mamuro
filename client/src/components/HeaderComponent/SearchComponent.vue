<script setup lang="ts">
import { ref } from 'vue'
import SvgIcon from '@/components/utils/SvgIcon.vue'
import searchIcon from '@/assets/images/search_icon.svg'
import filterIcon from '@/assets/images/search_filter_icon.svg'
import { useMailingStore } from '@/stores/mailing'

const mailingStore = useMailingStore()

const searchInput = ref<HTMLInputElement | null>(null)
const focusInput = () => {
  if (searchInput.value) {
    searchInput.value.focus()
  }
}

const handleEnter = (event: KeyboardEvent) => {
  if (event.key === 'Enter') {
    searchBy()
  }
}

const searchBy = () => {
  if (searchInput.value) {
    const inputValue = searchInput.value.value
    mailingStore.searchMailBy(inputValue)
  }
}
</script>

<template>
  <div class="search-container" @click="focusInput">
    <SvgIcon :svg="searchIcon" class="cursor-pointer" @click="() => searchBy()" />

    <input ref="searchInput" type="text" placeholder="Search in mail" class="search-input" @keypress="handleEnter" />

    <SvgIcon :svg="filterIcon" class="cursor-not-allowed" />
  </div>
</template>

<style scoped lang="scss">

.search-container {
  @apply w-[48%] h-12 px-4 border-0 rounded-xl bg-blue-100/50 my-2 flex items-center space-x-4;
}

.search-input {
  @apply flex-1 bg-transparent outline-none placeholder-gray-600;
}
</style>