<script setup lang="ts">
import type { NavigationItem } from '@/router/navigation';
import { computed } from 'vue';
import { useRoute, useRouter } from 'vue-router';

interface Props {
  navigationItem: NavigationItem
}

const props = defineProps<Props>();
const router = useRouter();
const route = useRoute();

const isActive = computed(() => route.path === props.navigationItem.to);

function onNavigationClicked() {
  router.push(props.navigationItem.to);
}
</script>

<template>
  <button
    class="bottom-navigation-button"
    :class="{ 'is-active': isActive }"
    @click="onNavigationClicked">
    <v-icon
      :name="props.navigationItem.icon" />
    {{ props.navigationItem.label }}
  </button>
</template>

<style lang="less" scoped>
@import '@/styles/theme.less';

.bottom-navigation-button {
  all: unset;
  font-size: @font-size-md;
  cursor: pointer;
  transition: all 0.2s ease;
  font-weight: @font-weight-bold;
  padding: 8px;
  background-color: @color-background-mute;
  gap: 2px;
  display: flex;
  flex-grow: 1;
  flex-direction: column;
  align-items: center;
  border-top: 2px solid transparent;
  box-sizing: border-box;
  min-width: 0;
  min-height: 0;

  &:hover {
    background-color: @color-background-ghost-hover;
  }

  &.is-active {
    border-top-color: @color-primary;
  }
}
</style>
