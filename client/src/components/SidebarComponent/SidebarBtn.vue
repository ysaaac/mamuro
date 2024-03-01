<script setup lang="ts">
interface SideBarBtnProps {
  isDisabled: boolean;
  bgImage: string;
}

const props = defineProps<SideBarBtnProps>()
</script>

<template>
  <button :class="[isDisabled ? 'cursor-not-allowed' : 'cursor-pointer']">
    <span class="side-bar-btn" :style="{ backgroundImage: `url('${props.bgImage}')` }"></span>
    <slot class="text-xs"></slot>
  </button>
</template>

<style scoped lang="scss">
.side-bar-btn {
  @apply block w-12 h-8 rounded-full mx-auto bg-center bg-no-repeat ;
}
</style>
