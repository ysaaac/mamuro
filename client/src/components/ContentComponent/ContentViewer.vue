<script setup lang="ts">
import {defineProps} from 'vue'
import {useMailingStore} from "@/stores/mailing";
import type {EmailContentType} from "@/types/MailContent";

interface EmailContentViewer extends EmailContentType {
}

const props = defineProps<EmailContentViewer>()

const mailingStore = useMailingStore()

const contentViewerIcons = {
  star: 'https://www.gstatic.com/images/icons/material/system_gm/1x/star_border_black_20dp.png',
  reply: '//ssl.gstatic.com/ui/v1/icons/mail/gm3/1x/reply_baseline_nv700_20dp.png',
  more: '//ssl.gstatic.com/ui/v1/icons/mail/gm3/1x/more_vert_baseline_nv700_20dp.png'
}
</script>

<template>
  <article class="flex-1 bg-white overflow-y-scroll">
    <!--    <template x-if="postContent === null">-->
    <div class="p-8" v-if="!content">
      <p>Welcome to Mamuro Mailing</p>
      <i class="italic">
        Select the email you want to see in the left side or search it with the top side input
      </i>
    </div>
    <!--    <template x-if="postContent !== null">-->
    <div class="px-8" v-if="content">
      <h1 class="text-xl">{{ subject }}</h1>
      <div>
        <div class="px-[0.25rem] rounded bg-neutral-200 text-gray-800 text-xs w-fit">
          {{ mailingStore.mailSectionType }}
        </div>
      </div>
      <div class="flex flex-row space-x-4 items-center">
        <div class="w-10 h-10">
          <img class="user-avatar" src="https://lh3.googleusercontent.com/a/default-user=s40-p" alt="user-avatar"/>
        </div>
        <div class="flex flex-col justify-between py-2">
          <div><p class="font-medium text-sm"> {{ from }} </p></div>
          <div><p class="text-xs text-gray-700"> {{ to }} </p></div>
        </div>
        <div class="flex-1 flex flex-col justify-between items-end py-2">
          <div class="flex flex-row items-center space-x-2 cursor-pointer">
            <!--            x-text="new Date(postContent.created_utc * 1000).toLocaleString()"-->
            <div><p class="text-gray-500 text-xs">{{ date }}</p></div>
            <div>
              <div class="content-options-icons" :style="{ backgroundImage: `url('${contentViewerIcons.star}')` }"/>
            </div>

            <div>
              <div class="content-options-icons" :style="{ backgroundImage: `url('${contentViewerIcons.reply}')` }"/>
            </div>

            <div>
              <div class="content-options-icons" :style="{ backgroundImage: `url('${contentViewerIcons.more}')` }"/>
            </div>
          </div>
          <div><p class="font-medium text-sm opacity-0">{{ from }}</p></div>
        </div>
      </div>
      <!--      <a-->
      <!--        :href="postContent.url"-->
      <!--        target="_blank"-->
      <!--        class="text-sm text-blue-500 underline pl-14"-->
      <!--        x-text="postContent.url"-->
      <!--      ></a>-->
      <!--      <p-->
      <!--        class="text-sm text-gray-500 pl-14"-->
      <!--        x-text="postContent.selftext"-->
      <!--      ></p>-->

      <div class="ml-14 flex flex-col mt-6">
        <!--        <template x-if="postContentComments.length === 0">-->
        <!--          <p>No comments</p>-->
        <!--        </template>-->

        <!-- Print all comments -->
        <div class="border-l-2 border-gray-300 pb-2">
          <div class="pl-2">
            <p class="text-sm">{{ content }}</p>
          </div>
        </div>
      </div>
    </div>
  </article>
</template>

<style scoped lang="scss">
.user-avatar {
  @apply rounded-full object-cover w-10 h-10 aspect-square;
}

.content-options-icons {
  @apply w-4 h-4 bg-no-repeat bg-contain;
}
</style>