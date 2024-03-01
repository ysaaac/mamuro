<script setup lang="ts">
import { ref, withDefaults } from 'vue'
import SidebarBtn from '@/components/SidebarComponent/SidebarBtn.vue'

interface NavigationProperties {
  showSideBar: boolean
}

const props = withDefaults(defineProps<NavigationProperties>(), {
  showSideBar: true
})

const showSideBar = ref(props.showSideBar)
const toggleSideBar = () => (showSideBar.value = !showSideBar.value)

const sidebarBtns = [
  {
    name: 'Mail',
    disabled: false,
    img: 'https://ssl.gstatic.com/ui/v1/icons/mail/gm3/1x/mail_fill_baseline_p900_20dp.png'
  },
  {
    name: 'Chat',
    disabled: true,
    img: 'https://ssl.gstatic.com/ui/v1/icons/mail/gm3/1x/chat_bubble_baseline_nv700_20dp.png'
  },
  {
    name: 'Spaces',
    disabled: true,
    img: 'https://ssl.gstatic.com/ui/v1/icons/mail/gm3/1x/gmail_groups_baseline_nv700_20dp.png'
  },
  {
    name: 'Meet',
    disabled: true,
    img: 'https://ssl.gstatic.com/ui/v1/icons/mail/gm3/1x/meet_baseline_nv700_20dp.png'
  }
]
</script>

<template>
  <div class="h-screen flex">
    <aside class="w-16 h-full bg-[#eaf1fb] flex flex-col items-center space-y-4 py-2">
      <button @click="toggleSideBar" class="toggle-btn">
        <svg focusable="false" viewBox="0 0 24 24" class="w-6 h-6">
          <path d="M3 18h18v-2H3v2zm0-5h18v-2H3v2zm0-7v2h18V6H3z"></path>
        </svg>
      </button>

      <SidebarBtn
        v-for="(btn, index) in sidebarBtns"
        :key="index"
        :isDisabled="btn.disabled"
        :bgImage="btn.img"
      >
        <p class="text-xs">{{ btn.name }}</p>
      </SidebarBtn>

    </aside>
  </div>
</template>

<style scoped lang="scss">
.side-bar-btn {
  @apply w-12 h-8 rounded-full mx-auto bg-center bg-no-repeat cursor-not-allowed;
}

.toggle-btn {
  @apply flex items-center justify-center hover:bg-[#dae1eb] rounded-full w-12 h-12;
}

</style>