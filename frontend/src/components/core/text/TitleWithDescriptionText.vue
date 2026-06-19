<script setup lang="ts">
interface Props {
  title: string
  description: string
  titleColor?: string
}

const props = defineProps<Props>();
</script>

<template>
  <div
    class="title-description-text">
    <span
      class="title"
      :style="titleColor ? { color: titleColor } : undefined">
      {{ props.title }}
    </span>

    <span
      class="subtext">
      {{ props.description }}
    </span>
  </div>
</template>

<style lang="less" scoped>
@import '@/styles/theme.less';

.title-description-text {
  display: flex;
  flex-direction: column;
}

.title {
  font-weight: @font-weight-bold;
}

.subtext {
  font-size: @font-size-sm;
  color: @color-text-muted;
}
</style>
