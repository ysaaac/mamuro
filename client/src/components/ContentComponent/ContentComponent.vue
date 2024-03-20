<script setup lang="ts">
import ContentList from '@/components/ContentComponent/ContentList.vue'
import ContentViewer from '@/components/ContentComponent/ContentViewer.vue'
import { defineProps, onMounted, ref } from 'vue'
import { MailSectionType } from '@/enums/MailSection'
import { useMailingStore } from '@/stores/mailing'
import { fetchData } from '@/api'

interface ContentComponent {
  mailSectionType: MailSectionType,
}

const props = defineProps<ContentComponent>()
const mailingStore = useMailingStore()
mailingStore.mailSectionType = props.mailSectionType

const emailContent = ref<HTMLElement | null>(null)

const fetchMailsList = async () => {
  try {
    const endpointUrl = `/${mailingStore.mailSectionType.toLowerCase()}`
    const response = await fetchData(endpointUrl)
    mailingStore.setMailListData(response)
  } catch (error) {
    console.error('Error fetching data:', error)
  }
}

onMounted(fetchMailsList)
</script>

<template>
  <div class="flex flex-row flex-1 overflow-y-auto email-content-wrapper" ref="emailContent">
    <ContentList />
    <ContentViewer />
  </div>
</template>

<style scoped lang="scss">
</style>
