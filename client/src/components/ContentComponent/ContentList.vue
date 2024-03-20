<script setup lang="ts">

import {useMailingStore} from "@/stores/mailing";
const mailingStore = useMailingStore()

const getMailContent = (mailId: string): any => {
  console.log('Fetching mail content of ', mailId)
}
const savedIconUrl = '//ssl.gstatic.com/ui/v1/icons/mail/gm3/1x/label_baseline_nv700_20dp.png'
</script>

<template>
  <div class="h-full max–h-full flex-1 overflow-y-scroll">
    <!--    :class="{ 'bg-gray-200/80': viewedPosts[item.data.id], 'bg-blue-200': postContent !== null && postContent.id === item.data.id, 'hover:bg-gray-100': postContent !== null && postContent.id !== item.data.id }"-->
    <div
        v-for="(mail, index) in mailingStore.mailList"
        :key="index"
        @click="getMailContent(mail.mailId)"
        class="item-list"
        :class="{ 'bg-gray-200/80': mail.viewed, 'bg-blue-200 hover:bg-gray-100': !mail.viewed}"
    >
      <div>
        <input type="checkbox"/>
      </div>
      <div>
        <div class="w-5 h-5 bg-contain bg-no-repeat" :style="{ backgroundImage: `url('${savedIconUrl}')` }"></div>
      </div>
      <div>
        <div class="w-36 max-w-36">
          <p class="line-clamp-1 text-sm">{{ mail.from }}</p>
        </div>
      </div>
      <div class="flex-1">
        <p class="line-clamp-1 text-sm">{{ mail.subject }}</p>
      </div>
      <div>
        <!--        x-text="(new Date(item.data.created * 1000)).toLocaleString()"-->
        <p class="text-sm">{{ mail.date }}</p>
      </div>
    </div>
  </div>
</template>

<style scoped lang="scss">
.item-list {
  @apply flex flex-row items-center space-x-4 py-2 px-4 cursor-pointer max-w-full border-b border-gray-200;
}
</style>