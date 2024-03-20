<script setup lang="ts">
import ContentComponent from '@/components/ContentComponent/ContentComponent.vue'
import {MailSectionType} from "@/enums/MailSection";
</script>

<template>
  <ContentComponent :mailSectionType="MailSectionType.SENT"/>
</template>