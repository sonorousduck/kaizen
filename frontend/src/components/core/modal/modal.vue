<script setup lang="ts">
import { useResponsive } from '@/stores/responsive';

interface Props {
  title: string
}

const props = defineProps<Props>();
const responsive = useResponsive();

const dialogVisible = defineModel<boolean>('dialogVisible');
</script>

<template>
  <el-dialog
    v-model="dialogVisible"
    :title="props.title"
    :fullscreen="responsive.isMobile"
    :transition="responsive.isMobile ? 'dialog-slide' : 'dialog-fade'">
    <slot />
  </el-dialog>
</template>

<style lang="less" scoped>
@import '@/styles/theme.less';
</style>

<style lang="less">
.dialog-slide-enter-active,
.dialog-slide-leave-active,
.dialog-slide-enter-active .el-dialog,
.dialog-slide-leave-active .el-dialog {
  transition: all 0.3s cubic-bezier(0.25, 0.46, 0.45, 0.94);
}

.dialog-slide-enter-from,
.dialog-slide-leave-to {
  opacity: 0;
}

.dialog-slide-enter-from .el-dialog,
.dialog-slide-leave-to .el-dialog {
  transform: translateY(600px);
  opacity: 0;
}
</style>
