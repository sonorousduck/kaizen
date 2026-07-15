<script setup lang="ts">
import type { NavigationItem } from '@/router/navigation';
import { useRouter } from 'vue-router';
import DuckButton from '@/components/core/DuckButton/DuckButton.vue';

interface Props {
  navigationItem: NavigationItem
}

const props = defineProps<Props>();
const router = useRouter();

function onNavigationClicked() {
  router.push(props.navigationItem.to);
}
</script>

<template>
  <button
    class="navigation-button"
    @click="onNavigationClicked">
    {{ props.navigationItem.label }}
  </button>
</template>

<style lang="less" scoped>
@import '@/styles/theme.less';

.navigation-button {
  all: unset;
  font-size: @font-size-md;
  cursor: pointer;
  transition: all 0.2s ease;
  font-weight: @font-weight-bold;
  padding: 12px 8px;
  background-color: @color-background-mute;

  &:hover {
    background-color: @color-background-ghost-hover;
  }
}
</style>
