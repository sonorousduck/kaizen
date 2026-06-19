<script setup lang="ts">
import { computed } from 'vue';

interface Props {
  variant?: 'primary' | 'secondary' | 'ghost'
  type?: 'icon+text' | 'icon-only' | 'text-only'
  disabled?: boolean
  isDestructive?: boolean
}

const props = withDefaults(defineProps<Props>(), {
  variant: 'primary',
  type: 'icon+text',
  disabled: false,
});

const buttonClasses = computed(() => {
  return [
    'duck-button',
    `duck-button-${props.variant}`,
    {
      'duck-button-icon-only': props.type === 'icon-only',
      'duck-button-text-only': props.type === 'text-only',
      'duck-button-destructive-color': props.isDestructive,
    },
  ];
});
</script>

<template>
  <button
    :class="buttonClasses"
    :disabled="disabled">
    <slot />
  </button>
</template>

<style lang="less" scoped>
@import '@/styles/button.less';
</style>
