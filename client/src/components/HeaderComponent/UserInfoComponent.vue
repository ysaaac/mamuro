<script setup lang="ts">
import SvgIcon from '@/components/utils/SvgIcon.vue'
import userInfoMoreApps from '@/assets/images/user_info_more_apps.svg'
import userInfoSettings from '@/assets/images/user_info_settings.svg'
import userInfoMoreInfo from '@/assets/images/user_info_more_info.svg'
import userInfoArrowDown from '@/assets/images/user_info_arrow_down.svg'

const iconsUrls = {
  gsuite: 'https://www.google.com/u/1/ac/images/logo.gif?uid=105939349603080396778&service=google_gsuite',
  me: 'https://i.stack.imgur.com/34AD2.jpg'
}


// original href for anchor tag was https://github.com/niturobert/gmailkit/
</script>

<template>
  <div class="flex-1 flex items-center justify-end space-x-4 cursor-not-allowed">
    <div>
      <div class="border border-gray-300 pl-4 pr-2 py-2 rounded-full flex items-center space-x-2">
        <div>
          <div class="rounded-full w-3 h-3 bg-green-700"></div>
        </div>
        <div>
          <p class="text-sm text-gray-500 font-medium">Active</p>
        </div>
        <SvgIcon :svg="userInfoArrowDown" />
      </div>
    </div>

    <SvgIcon :svg="userInfoMoreInfo" class="cursor-not-allowed" />
    <SvgIcon :svg="userInfoSettings" class="cursor-not-allowed" />
    <SvgIcon :svg="userInfoMoreApps" class="cursor-not-allowed" />

    <a href="https://github.com/ysaaac/mamuro" target="_blank" class="user-info-btn">
      <img :src="iconsUrls.gsuite" class="h-6" alt="gsuite" />
      <img :src="iconsUrls.me" class="h-8 w-8 object-cover rounded-full" alt="me" />
    </a>
  </div>
</template>

<style scoped lang="scss">

.user-info-btn {
  @apply border border-gray-200 p-2 flex space-x-2 items-center rounded-md cursor-pointer hover:bg-gray-100;
}
</style>