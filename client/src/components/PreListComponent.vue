<script setup lang="ts">
const iconsUrls = {
  box: '//ssl.gstatic.com/ui/v1/icons/mail/gm3/1x/check_box_outline_blank_baseline_nv700_20dp.png',
  arrowDown: '//ssl.gstatic.com/ui/v1/icons/mail/gm3/1x/arrow_drop_down_baseline_nv700_20dp.png',
  boxWithArrow: '//ssl.gstatic.com/ui/v1/icons/mail/gm3/1x/archive_baseline_nv700_20dp.png',
  exclamation: '//ssl.gstatic.com/ui/v1/icons/mail/gm3/1x/report_baseline_nv700_20dp.png',
  delete: '//ssl.gstatic.com/ui/v1/icons/mail/gm3/1x/delete_baseline_nv700_20dp.png',
  message: '//ssl.gstatic.com/ui/v1/icons/mail/gm3/1x/mail_baseline_nv700_20dp.png',
  timer: '//ssl.gstatic.com/ui/v1/icons/mail/gm3/1x/schedule_baseline_nv700_20dp.png',
  checkedCircle: '//ssl.gstatic.com/ui/v1/icons/mail/gm3/1x/add_task_baseline_nv700_20dp.png',
  folderWithArrow: '//ssl.gstatic.com/ui/v1/icons/mail/gm3/1x/drive_file_move_outline_baseline_nv700_20dp.png',
  savedItems: '//ssl.gstatic.com/ui/v1/icons/mail/gm3/1x/label_baseline_nv700_20dp.png',
  threeDots: '//ssl.gstatic.com/ui/v1/icons/mail/gm3/1x/more_vert_baseline_nv700_20dp.png'
}
</script>

<template>
  <div class="w-full flex flex-row items-center p-4">
    <div>
      <div class="pre-list-icon" :style="{ backgroundImage: `url('${iconsUrls.box}')` }"></div>
    </div>
    <div class="pr-4">
      <div class="pre-list-icon" :style="{ backgroundImage: `url('${iconsUrls.arrowDown}')` }"></div>
    </div>
    <div class="pr-4">
      <div class="pre-list-icon" :style="{ backgroundImage: `url('${iconsUrls.boxWithArrow}')` }"></div>
    </div>
    <div class="pr-4">
      <div class="pre-list-icon" :style="{ backgroundImage: `url('${iconsUrls.exclamation}')` }"></div>
    </div>
    <div class="pr-6 border-r border-gray-200">
      <div class="pre-list-icon" :style="{ backgroundImage: `url('${iconsUrls.delete}')` }"></div>
    </div>
    <div class="pl-6 pr-4">
      <div class="pre-list-icon" :style="{ backgroundImage: `url('${iconsUrls.message}')` }"></div>
    </div>
    <div class="pr-4">
      <div class="pre-list-icon" :style="{ backgroundImage: `url('${iconsUrls.timer}')` }"></div>
    </div>
    <div class="pr-6 border-r border-gray-200">
      <div class="pre-list-icon" :style="{ backgroundImage: `url('${iconsUrls.checkedCircle}')` }"></div>
    </div>
    <div class="pl-6 pr-4">
      <div class="pre-list-icon" :style="{ backgroundImage: `url('${iconsUrls.folderWithArrow}')` }"></div>
    </div>
    <div class="pl-4">
      <div class="pre-list-icon" :style="{ backgroundImage: `url('${iconsUrls.savedItems}')` }"></div>
    </div>
    <div class="pl-4">
      <div class="pre-list-icon" :style="{ backgroundImage: `url('${iconsUrls.threeDots}')` }"></div>
    </div>
  </div>
</template>

<style scoped lang="scss">
.pre-list-icon {
  @apply w-5 h-5 bg-contain bg-no-repeat;
}
</style>