<script setup lang="ts">
import { useRouter } from 'vue-router'

const iconsUrls = {
  gmailLogo: 'https://ssl.gstatic.com/ui/v1/icons/mail/rfr/logo_gmail_lockup_default_1x_r5.png',
  compose: 'https://www.gstatic.com/images/icons/material/system_gm/1x/create_black_24dp.png',
  labelTag: 'https://www.gstatic.com/images/icons/material/system_gm_filled/1x/label_gm_grey_24dp.png'
}

const labelsList = [
  'random label 1',
  'random label 2',
  'random label 3',
  'random label 4'
]
const router = useRouter()

const isActivePath = (path: string) => {
  return router.currentRoute.value.path === path
}
const disabledAction = () => {
  alert('Temporary disabled option')
}

const mailSections = [
  {
    name: 'Inbox',
    icon: 'https://ssl.gstatic.com/ui/v1/icons/mail/gm3/1x/inbox_fill_baseline_p900_20dp.png',
    path: '/',
    action: () => router.push('/')
  },
  {
    name: 'Sent',
    icon: 'https://ssl.gstatic.com/ui/v1/icons/mail/gm3/1x/send_baseline_nv700_20dp.png',
    path: '/sent',
    action: () => router.push('/sent')
  },
  {
    name: 'Starred',
    icon: 'https://ssl.gstatic.com/ui/v1/icons/mail/gm3/1x/star_baseline_nv700_20dp.png',
    path: '',
    action: disabledAction
  },
  {
    name: 'Draft',
    icon: 'https://ssl.gstatic.com/ui/v1/icons/mail/gm3/1x/draft_baseline_nv700_20dp.png',
    path: '',
    action: disabledAction
  }
]

const addCompose = () => {
  alert('New Compose')
}
</script>

<template>
  <nav class="w-64 bg-gray-50 h-screen py-3 px-6">
    <img :src="iconsUrls.gmailLogo" alt="Gmail Logo" />

    <button @click="addCompose" class="compose-btn">
      <span class="compose-btn__logo" :style="{ backgroundImage: `url('${iconsUrls.compose}')` }" />
      <span class="text-sm pr-4">Compose</span>
    </button>

    <div class="flex flex-col space-y-0">
      <div
        class="flex items-center space-x-2 rounded-full py-2 px-4 cursor-pointer"
        :class="isActivePath(section.path) ? 'bg-blue-200/80 font-bold': 'hover:bg-gray-200/80'"
        v-for="(section, index) in mailSections"
        :key="index"
        @click="section.action"
      >
        <div class="w-4 h-4 rounded-full">
          <img :src="section.icon" :alt="section.name + '_icon'" />
        </div>
        <div class="text-sm font-medium text-gray-700">
          {{ section.name }}
        </div>
      </div>
    </div>

    <div class="pt-8 flex flex-col space-y-0">
      <div class="flex mb-4">
        <div class="flex-1">
          <p class="text-lg">Labels</p>
        </div>
        <div class="hover:bg-gray-300 rounded-full px-2 cursor-pointer">+</div>
      </div>

      <div
        class="flex items-center space-x-4 cursor-not-allowed"
        v-for="(label, index) in labelsList"
        :key="index"
      >
        <div class="label-tag" :style="{ backgroundImage: `url('${iconsUrls.labelTag}')` }"></div>
        <div class="flex-1">
          <p class="text-sm font-light">{{ label }}</p>
        </div>
      </div>
    </div>
  </nav>
</template>

<style scoped>
.compose-btn {
  @apply px-3 py-3 bg-cyan-400/30 rounded-xl my-6 flex flex-row space-x-4 items-center;
}

.compose-btn__logo {
  @apply block w-6 h-6 object-fill bg-no-repeat;
}

.label-tag {
  @apply w-6 h-6 bg-center bg-contain;
}
</style>